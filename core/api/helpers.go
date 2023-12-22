package api

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gtfintechlab/scatter-protocol/core/protocol"
	scattertoken "github.com/gtfintechlab/scatter-protocol/core/protocol/scatter-token"
)

func transferToken(privateKey string, amount int64, recipient string, ethereumNode string, contractAddress string) {
	auth := getTransactor(privateKey, ethereumNode)
	var ethereumClient, _ = ethclient.Dial(ethereumNode)
	contract, _ := scattertoken.NewScattertoken(common.HexToAddress(contractAddress), ethereumClient)
	contract.Transfer(auth, common.HexToAddress(recipient), big.NewInt(amount))
}

func getTransactor(ownerPrivateKey string, ethereumNode string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(ownerPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, protocol.CHAIN)
	auth.Value = big.NewInt(0)
	var ethereumClient, _ = ethclient.Dial(ethereumNode)
	gas, _ := ethereumClient.SuggestGasPrice(context.Background())
	auth.GasPrice = (*big.Int)(gas)

	return auth
}
