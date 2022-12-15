package utils

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	scattertoken "github.com/gtfintechlab/scatter-protocol/node/contracts"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var contractAddress common.Address = common.HexToAddress(os.Getenv("ETHEREUM_GOERLI_CONTRACT_ADDRESS"))
var client, _ = ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))
var contract, _ = scattertoken.NewScattertoken(contractAddress, client)

func getOwnStake() *big.Int {
	stake, err := contract.GetOwnStake(nil)

	if err != nil {
		log.Fatal(err)
	}
	return stake
}
