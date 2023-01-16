package main

import (
	"context"
	"fmt"
	"log"

	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	scattertoken "github.com/gtfintechlab/scatter-protocol/token"
)

const (
	GOERLI = 5
)

var CHAIN = big.NewInt(GOERLI)

func main() {
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
