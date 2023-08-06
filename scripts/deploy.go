package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	evaluationtoken "github.com/gtfintechlab/scatter-protocol/token/evaluation"
	scattertoken "github.com/gtfintechlab/scatter-protocol/token/scatter"
	trainingtoken "github.com/gtfintechlab/scatter-protocol/token/training"

	"github.com/joho/godotenv"
)

const (
	GOERLI = 5
)

const (
	SCATTER    = "scatter"
	EVALUATION = "evaluation"
	TRAINING   = "training"
)

var CHAIN = big.NewInt(GOERLI)

func main() {
	var tokenType string

	flag.StringVar(&tokenType, "type", "", "The smart contract you want to deploy")
	flag.Parse()

	switch tokenType {
	case SCATTER:
		deployScatterToken()
	case EVALUATION:
		deployEvaluationJobToken()
	case TRAINING:
		deployTrainingJobToken()
	}
}

func deployScatterToken() {
	godotenv.Load(".env")
	client, err := ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, CHAIN)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	address, transaction, _, err := scattertoken.DeployScattertoken(auth, client, big.NewInt(int64(100000000000)), big.NewInt(int64(100)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract Address: " + address.Hex())
	fmt.Println("Transaction Hash: " + transaction.Hash().Hex())

}

func deployEvaluationJobToken() {
	godotenv.Load(".env")
	client, err := ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, CHAIN)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	address, transaction, _, err := evaluationtoken.DeployEvaluationtoken(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract Address: " + address.Hex())
	fmt.Println("Transaction Hash: " + transaction.Hash().Hex())

}

func deployTrainingJobToken() {
	godotenv.Load(".env")
	client, err := ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, CHAIN)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	address, transaction, _, err := trainingtoken.DeployTrainingtoken(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract Address: " + address.Hex())
	fmt.Println("Transaction Hash: " + transaction.Hash().Hex())

}
