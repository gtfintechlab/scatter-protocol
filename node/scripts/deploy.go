package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	scattertoken "github.com/gtfintechlab/scatter-protocol/node/contracts"
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

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, CHAIN)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	fmt.Println(fromAddress)
	address, transaction, _, err := scattertoken.DeployScattertoken(auth, client, big.NewInt(int64(100000000000*math.Pow(10, 18))), big.NewInt(int64(100)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract Address: " + address.Hex())
	fmt.Println("Transaction Hash: " + transaction.Hash().Hex())
}
