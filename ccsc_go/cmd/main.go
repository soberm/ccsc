package main

import (
	"ccsc_go/pkg/ccsc"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/big"
	"os"
	"os/signal"
	"time"
)

func main() {

	configFile := flag.String("c", "./configs/config.json", "filename of the config file")
	flag.Parse()

	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config: %v", err)
	}

	var config ccsc.Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unmarshal config into struct: %v", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sig
		cancel()
	}()

	crossChainCaller, err := ccsc.NewCrossChainCaller(config.CallerConfig)
	if err != nil {
		log.Fatalf("new cross chain caller: %v", err)
	}

	err = crossChainCaller.Deposit(ctx)
	if err != nil {
		log.Fatalf("deposit: %v", err)
	}

	go func() {
		if err := crossChainCaller.Run(ctx); err != nil {
			log.Fatalf("caller run: %v", err)
		}
	}()

	ethClient, err := ethclient.Dial(config.EthereumHost)
	if err != nil {
		log.Fatalf("dial eth client: %v", err)
	}

	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatalf("chain ID: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatalf("private key hex to ecdsa: %v", err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("keyed transactor with chain id: %v", err)
	}

	clientContract, err := ccsc.NewClientContract(common.HexToAddress(config.ClientContractAddress), ethClient)
	if err != nil {
		log.Fatalf("new client contract: %v", err)
	}

	f, err := os.OpenFile(config.CallDurationMeasurementsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	durationsCSV := csv.NewWriter(f)

	f, err = os.OpenFile(config.CostsMeasurementsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	costsMeasurementsCSV := csv.NewWriter(f)

	timeBefore := time.Now()

	log.Infof("Calling store function of client contract")
	tx, err := clientContract.Store(transactOpts, big.NewInt(42))
	if err != nil {
		log.Fatalf("store: %v", err)
	}
	receipt, err := bind.WaitMined(ctx, ethClient, tx)
	if err != nil {
		log.Fatalf("wait mined: %v", err)
	}

	sink := make(chan *ccsc.ClientContractStored)
	defer close(sink)

	sub, err := clientContract.WatchStored(
		&bind.WatchOpts{
			Context: ctx,
		},
		sink,
		nil,
	)
	if err != nil {
		log.Fatalf("watch stored: %v", err)
	}

	for {
		select {
		case event := <-sink:
			timeAfter := time.Now()
			duration := timeAfter.Sub(timeBefore).Seconds()

			err = costsMeasurementsCSV.Write([]string{event.Id.String(), fmt.Sprint(receipt.GasUsed)})
			if err != nil {
				log.Fatalf("write duration to csv: %v", err)
			}
			costsMeasurementsCSV.Flush()

			timeBefore = timeAfter
			err = durationsCSV.Write([]string{event.Id.String(), fmt.Sprint(duration)})
			if err != nil {
				log.Fatalf("write duration to csv: %v", err)
			}
			durationsCSV.Flush()

			log.Infof("Calling store function of client contract")
			tx, err = clientContract.Store(transactOpts, big.NewInt(42))
			if err != nil {
				log.Fatalf("store: %v", err)
			}

			receipt, err = bind.WaitMined(ctx, ethClient, tx)
			if err != nil {
				log.Fatalf("wait mined: %v", err)
			}
		case err = <-sub.Err():
			log.Fatalf("sub: %v", err)
		}
	}
}
