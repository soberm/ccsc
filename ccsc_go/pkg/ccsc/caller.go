package ccsc

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"time"
)

type crossChainCaller struct {
	srcClient        *ethclient.Client
	targetClient     *ethclient.Client
	srcChainID       *big.Int
	targetChainID    *big.Int
	proxyContract    *ProxyContract
	serverContract   *ServerContract
	ethRelayContract *Testimonium
	bscRelayContract *BSCRelayContract
	privateKey       *ecdsa.PrivateKey
	measurementFile  *os.File
}

func NewCrossChainCaller(config CallerConfig) (*crossChainCaller, error) {

	srcClient, err := ethclient.Dial(config.SourceHost)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}

	targetClient, err := ethclient.Dial(config.TargetHost)
	if err != nil {
		return nil, fmt.Errorf("dial bsc client: %v", err)
	}

	proxyContract, err := NewProxyContract(common.HexToAddress(config.ProxyContractAddress), srcClient)
	if err != nil {
		return nil, fmt.Errorf("new proxy contract: %w", err)
	}

	serverContract, err := NewServerContract(common.HexToAddress(config.ServerContractAddress), targetClient)
	if err != nil {
		return nil, fmt.Errorf("new server contract: %w", err)
	}

	srcRelayContract, err := NewTestimonium(common.HexToAddress(config.SrcRelayContractAddress), targetClient)
	if err != nil {
		return nil, fmt.Errorf("new src relay contract: %w", err)
	}

	targetRelayContract, err := NewBSCRelayContract(common.HexToAddress(config.TargetRelayContractAddress), srcClient)
	if err != nil {
		return nil, fmt.Errorf("new target relay contract: %w", err)
	}

	srcChainID, err := srcClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get src chain id: %w", err)
	}

	targetChainID, err := targetClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get target chain id: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("private key hex to ecdsa: %v", err)
	}

	f, err := os.OpenFile(config.CallerMeasurementFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open caller measurement file: %w", err)
	}

	return &crossChainCaller{
		srcClient:        srcClient,
		targetClient:     targetClient,
		srcChainID:       srcChainID,
		targetChainID:    targetChainID,
		proxyContract:    proxyContract,
		serverContract:   serverContract,
		ethRelayContract: srcRelayContract,
		bscRelayContract: targetRelayContract,
		privateKey:       privateKey,
		measurementFile:  f,
	}, nil
}

func (c *crossChainCaller) Deposit(ctx context.Context) error {
	minDeposit, err := c.proxyContract.MINDEPOSITPERCALL(nil)
	if err != nil {
		return fmt.Errorf("min deposit per call: %w", err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.srcChainID)
	if err != nil {
		return fmt.Errorf("keyed transactor with chain id: %w", err)
	}
	transactOpts.Value = minDeposit

	tx, err := c.proxyContract.Deposit(transactOpts)
	if err != nil {
		return fmt.Errorf("deposit: %w", err)
	}

	_, err = bind.WaitMined(ctx, c.srcClient, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}
	return nil
}

func (c *crossChainCaller) Withdraw(ctx context.Context) error {
	transactOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.srcChainID)
	if err != nil {
		return fmt.Errorf("keyed transactor with chain id: %w", err)
	}

	tx, err := c.proxyContract.Withdraw(transactOpts)
	if err != nil {
		return fmt.Errorf("withdraw: %w", err)
	}

	_, err = bind.WaitMined(ctx, c.srcClient, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}

	return nil
}

func (c *crossChainCaller) Run(ctx context.Context) error {
	sink := make(chan *ProxyContractCrossChainCallInitiated)
	defer close(sink)

	sub, err := c.proxyContract.WatchCrossChainCallInitiated(
		&bind.WatchOpts{
			Context: ctx,
		},
		sink,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			go func() {
				if err := c.handleCrossChainCallInitiated(ctx, event); err != nil {
					log.Errorf("handle cross chain call initiated: %v", err)
				}
			}()
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *crossChainCaller) handleCrossChainCallInitiated(ctx context.Context, event *ProxyContractCrossChainCallInitiated) error {
	log.Infof("Handle cross chain call initiated %d", event.CallID.Int64())

	measurement := make([]string, 6)
	measurement[0] = event.CallID.String()

	transactOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.srcChainID)
	if err != nil {
		return fmt.Errorf("keyed transactor with chain id: %w", err)
	}

	req := ProtocolCrossChainRequest{
		Id:          event.CallID,
		Origin:      event.Origin,
		CallBackTo:  event.CallBackTo,
		CallData:    event.CallData,
		Gas:         event.Gas,
		CallBack:    event.CallBackName,
		CallBackGas: event.CallBackGas,
		Timeout:     event.Timeout,
	}

	log.Infof("Prepared cross chain call %d", event.CallID.Int64())
	tx, err := c.proxyContract.PrepareCrossChainCall(transactOpts, req)
	if err != nil {
		return fmt.Errorf("prepare cross chain call: %w", err)
	}
	log.Infof("Submit prepare cross chain call tx %s", tx.Hash())

	receipt, err := bind.WaitMined(ctx, c.srcClient, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}
	log.Infof("Mined prepare cross chain call tx %s", tx.Hash())

	measurement[1] = fmt.Sprint(receipt.GasUsed)

	block, err := c.srcClient.BlockByNumber(ctx, receipt.BlockNumber)
	if err != nil {
		return fmt.Errorf("block by number: %w", err)
	}

	receipts, err := ReceiptsByTransactions(ctx, c.srcClient, block.Body().Transactions)
	if err != nil {
		return fmt.Errorf("receipts by transactions: %w", err)
	}

	header := new(bytes.Buffer)
	err = rlp.Encode(header, block.Header())
	if err != nil {
		return fmt.Errorf("rlp encode block header: %w", err)
	}

	merkleProofTransaction, err := MerkleProof(block.Transactions(), receipt.TransactionIndex)
	if err != nil {
		return fmt.Errorf("merkle proof for transaction: %w", err)
	}

	merkleProofReceipt, err := MerkleProof(receipts, receipt.TransactionIndex)
	if err != nil {
		return fmt.Errorf("merkle proof for receipt: %w", err)
	}

	transactOpts, err = bind.NewKeyedTransactorWithChainID(c.privateKey, c.targetChainID)
	if err != nil {
		return fmt.Errorf("keyed transactor with chain id: %w", err)
	}
	transactOpts.Value = big.NewInt(2)

	log.Infof("Wait until block %d is included in ETH relay", receipt.BlockNumber.Int64())
	timeBefore := time.Now()
	err = c.WaitConfirmedETHRelay(ctx, receipt.BlockHash)
	if err != nil {
		return fmt.Errorf("wait confirmed eth relay: %w", err)
	}
	timeAfter := time.Now()
	measurement[4] = fmt.Sprint(timeAfter.Sub(timeBefore).Seconds())

	log.Infof("Execute the cross chain call %d", event.CallID.Int64())
	tx, err = c.serverContract.ExecuteCrossChainCall(transactOpts, ProtocolProof{
		RlpHeader:              header.Bytes(),
		RlpEncodedTx:           merkleProofTransaction.value,
		RlpEncodedReceipt:      merkleProofReceipt.value,
		Path:                   merkleProofTransaction.path,
		RlpEncodedTxNodes:      merkleProofTransaction.nodes,
		RlpEncodedReceiptNodes: merkleProofReceipt.nodes,
	})
	if err != nil {
		return fmt.Errorf("execute cross chain call: %w", err)
	}
	log.Infof("Submit execute cross chain call tx %s", tx.Hash())

	receipt, err = bind.WaitMined(ctx, c.targetClient, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}
	log.Infof("Mined execute cross chain call tx %s", tx.Hash())
	measurement[2] = fmt.Sprint(receipt.GasUsed)

	block, err = c.targetClient.BlockByNumber(ctx, receipt.BlockNumber)
	if err != nil {
		return fmt.Errorf("block by number: %w", err)
	}

	receipts, err = ReceiptsByTransactions(ctx, c.targetClient, block.Body().Transactions)
	if err != nil {
		return fmt.Errorf("receipts by transactions: %w", err)
	}

	header = new(bytes.Buffer)
	err = rlp.Encode(header, block.Header())
	if err != nil {
		return fmt.Errorf("rlp encode block header: %w", err)
	}

	merkleProofTransaction, err = MerkleProof(block.Transactions(), receipt.TransactionIndex)
	if err != nil {
		return fmt.Errorf("merkle proof for transaction: %w", err)
	}

	merkleProofReceipt, err = MerkleProof(receipts, receipt.TransactionIndex)
	if err != nil {
		return fmt.Errorf("merkle proof for receipt: %w", err)
	}

	transactOpts, err = bind.NewKeyedTransactorWithChainID(c.privateKey, c.srcChainID)
	if err != nil {
		return fmt.Errorf("keyed transactor with chain id: %w", err)
	}

	log.Infof("Wait until block %d is included in BSC relay", receipt.BlockNumber.Int64())
	timeBefore = time.Now()
	err = c.WaitConfirmedBSCRelay(ctx, receipt.BlockHash)
	if err != nil {
		return fmt.Errorf("wait confirmed bsc relay: %w", err)
	}
	timeAfter = time.Now()
	measurement[5] = fmt.Sprint(timeAfter.Sub(timeBefore).Seconds())

	log.Infof("Acknowledge cross chain call %d", event.CallID.Int64())
	tx, err = c.proxyContract.AcknowledgeCrossChainCall(transactOpts, ProtocolProof{
		RlpHeader:              header.Bytes(),
		RlpEncodedTx:           merkleProofTransaction.value,
		RlpEncodedReceipt:      merkleProofReceipt.value,
		Path:                   merkleProofTransaction.path,
		RlpEncodedTxNodes:      merkleProofTransaction.nodes,
		RlpEncodedReceiptNodes: merkleProofReceipt.nodes,
	}, req)
	if err != nil {
		return fmt.Errorf("ack cross chain call: %w", err)
	}
	log.Infof("Submit ack cross chain call tx %s", tx.Hash())

	receipt, err = bind.WaitMined(ctx, c.srcClient, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}
	measurement[3] = fmt.Sprint(receipt.GasUsed)

	costsCSV := csv.NewWriter(c.measurementFile)
	err = costsCSV.Write(measurement)
	if err != nil {
		return fmt.Errorf("write measurement: %w", err)
	}
	costsCSV.Flush()

	return nil
}

func (c *crossChainCaller) WaitConfirmedETHRelay(ctx context.Context, blockHash common.Hash) error {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}

		headerToConfirm, err := c.ethRelayContract.GetHeader(nil, blockHash)
		if err != nil {
			return fmt.Errorf("get header: %w", err)
		}

		if headerToConfirm.BlockNumber.Int64() == 0 {
			continue
		}

		longestChainEndpointHash, err := c.ethRelayContract.GetLongestChainEndpoint(nil)
		if err != nil {
			return fmt.Errorf("get longest chain endpoint: %w", err)
		}

		longestChainEndpoint, err := c.ethRelayContract.GetHeader(nil, longestChainEndpointHash)
		if err != nil {
			return fmt.Errorf("get header: %w", err)
		}

		if (headerToConfirm.BlockNumber.Int64() + 7) > longestChainEndpoint.BlockNumber.Int64() {
			continue
		}

		i := 0
		currentHeader := longestChainEndpoint
		for headerToConfirm.BlockNumber.Int64() <= currentHeader.BlockNumber.Int64() {
			if bytes.Equal(headerToConfirm.Hash[:], currentHeader.Hash[:]) {
				return nil
			}

			header, err := c.srcClient.BlockByHash(ctx, currentHeader.Hash)
			if err != nil {
				return fmt.Errorf("block by hash: %w", err)
			}

			currentHeader, err = c.ethRelayContract.GetHeader(nil, header.ParentHash())
			if err != nil {
				return fmt.Errorf("get header: %w", err)
			}
			i++
		}
	}
}

func (c *crossChainCaller) WaitConfirmedBSCRelay(ctx context.Context, blockHash common.Hash) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		isBlockFinal, err := c.bscRelayContract.IsBlockFinal(nil, blockHash)
		if err != nil {
			return fmt.Errorf("is block final: %w", err)
		}
		if isBlockFinal {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}
