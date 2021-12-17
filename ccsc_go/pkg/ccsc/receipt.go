package ccsc

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ReceiptsByTransactions(ctx context.Context, c *ethclient.Client, transactions []*types.Transaction) (types.Receipts, error) {
	receipts := make([]*types.Receipt, len(transactions))
	for i := 0; i < len(receipts); i++ {
		receipt, err := c.TransactionReceipt(ctx, transactions[i].Hash())
		if err != nil {
			return nil, err
		}
		receipts[i] = receipt
	}
	return receipts, nil
}
