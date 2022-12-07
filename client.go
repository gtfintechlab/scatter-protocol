package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	client, _ := ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))

	fmt.Println(client)
	fmt.Println(os.Getenv("ETHEREUM_NODE"))

}
