package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	client, _ := ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))
	context := context.Background()
	transaction, pending, _ := client.TransactionByHash(context, common.HexToHash("0xe9fbda3ac83da492cd4bd863e4849552150bb26a19b1b980165ad5bb1c3e4c58"))

	if pending {
		fmt.Println("Transaction is Pending")
	} else {
		fmt.Println("Transaction is not Pending")
		fmt.Println(transaction)
	}
}
