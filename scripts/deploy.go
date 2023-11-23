package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	evaluationtoken "github.com/gtfintechlab/scatter-protocol/protocol/evaluation"
	modelToken "github.com/gtfintechlab/scatter-protocol/protocol/model"
	scatterprotocol "github.com/gtfintechlab/scatter-protocol/protocol/scatter-protocol"
	scattertoken "github.com/gtfintechlab/scatter-protocol/protocol/scatter-token"
	trainingtoken "github.com/gtfintechlab/scatter-protocol/protocol/training"
	"github.com/gtfintechlab/scatter-protocol/utils"

	"github.com/joho/godotenv"
)

const (
	PROTOCOL   = "protocol"
	MODEL      = "model"
	TOKEN      = "token"
	EVALUATION = "evaluation"
	TRAINING   = "training"
	ALL        = "all"
)

var _ = godotenv.Load(".env")
var CHAIN_ID, _ = strconv.Atoi(os.Getenv("CHAIN_ID"))
var CHAIN = big.NewInt(int64(CHAIN_ID))
var client, _ = ethclient.Dial(os.Getenv("ETHEREUM_NODE"))

var DEPLOY_PAUSE time.Duration = 3

func main() {
	var contractName string
	var privateKey string
	flag.StringVar(&contractName, "contract", "all",
		"Name of contract to deploy")
	flag.StringVar(&privateKey, "pkey", os.Getenv("PRIVATE_KEY"),
		"Private key to deploy protocol with")
	flag.Parse()

	DeployContracts(contractName, privateKey)
}

func DeployContracts(contract string, privateKey string) {
	switch contract {
	case PROTOCOL:
		deployScatterProtocol(privateKey)
	case EVALUATION:
		deployEvaluationJobToken(privateKey)
	case TRAINING:
		deployTrainingJobToken(privateKey)
	case TOKEN:
		deployScatterToken(privateKey)
	case MODEL:
		deployModelToken(privateKey)
	case ALL:
		deployAllContracts(privateKey)
	}
}

func deployAllContracts(privateKey string) {
	auth := GetTransactor(privateKey)

	trainingAddress, trainingTransaction, trainingInstance, err := trainingtoken.DeployTrainingtoken(auth, client)
	if err != nil {
		log.Fatal("Failed to deploy training token", err.Error())
	}
	time.Sleep(time.Second * DEPLOY_PAUSE)

	evaluationAddress, evaluationTransaction, evaluationInstance, err := evaluationtoken.DeployEvaluationtoken(auth, client)
	if err != nil {
		log.Fatal("Failed to deploy evaluation token", err.Error())
	}
	time.Sleep(time.Second * DEPLOY_PAUSE)

	modelAddress, modelTransaction, modelInstance, err := modelToken.DeployModeltoken(auth, client)
	if err != nil {
		log.Fatal("Failed to deploy model token", err.Error())
	}
	time.Sleep(time.Second * DEPLOY_PAUSE)

	scatterTokenAddress, scatterTokenTransaction, scatterTokenInstance, err := scattertoken.DeployScattertoken(auth, client, big.NewInt(int64(100000000000)))
	if err != nil {
		log.Fatal("Failed to deploy scatter token", err.Error())
	}
	time.Sleep(time.Second * DEPLOY_PAUSE)

	scatterProtocolAddress, scatterProtocolTransaction, scatterInstance, err := scatterprotocol.DeployScatterprotocol(
		auth, client,
		common.HexToAddress(trainingAddress.Hash().Hex()),
		common.HexToAddress(evaluationAddress.Hash().Hex()),
		common.HexToAddress(scatterTokenAddress.Hash().Hex()),
		common.HexToAddress(modelAddress.Hash().Hex()),
	)
	if err != nil {
		log.Fatal("Failed to deploy scatter protocol", err.Error())
	}
	time.Sleep(time.Second * 10)

	trainingInstance.SetScatterContractAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	evaluationInstance.SetScatterContractAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	scatterTokenInstance.SetScatterProtocolAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	modelInstance.SetScatterContractAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	scatterInstance.InitRequestorNode(auth)

	log.Println("Transaction Info:")
	log.Printf("Training Token: %s\n", trainingTransaction.Hash())
	log.Printf("Model Token: %s\n", modelTransaction.Hash())
	log.Printf("Evaluation Token: %s\n", evaluationTransaction.Hash())
	log.Printf("Scatter Token: %s\n", scatterTokenTransaction.Hash())
	log.Printf("Scatter Protocol: %s\n", scatterProtocolTransaction.Hash())
	log.Println("===============")

	log.Println("Contract Info:")
	log.Printf("Training Token: %s\n", trainingAddress.Hex())
	log.Printf("Model Token: %s\n", modelAddress.Hex())
	log.Printf("Evaluation Token: %s\n", evaluationAddress.Hex())
	log.Printf("Scatter Token: %s\n", scatterTokenAddress.Hex())
	log.Printf("Scatter Protocol: %s\n", scatterProtocolAddress.Hex())

	contractInfo := map[string]string{
		"SCATTER_PROTOCOL_CONTRACT": scatterProtocolAddress.Hex(),
		"SCATTER_TOKEN_CONTRACT":    scatterTokenAddress.Hex(),
		"TRAINING_TOKEN_CONTRACT":   trainingAddress.Hex(),
		"EVALUATION_TOKEN_CONTRACT": evaluationAddress.Hex(),
		"MODEL_TOKEN_CONTRACT":      modelAddress.Hex(),
	}

	jsonData, _ := json.MarshalIndent(contractInfo, "", "  ")
	os.WriteFile("utils/contracts.json", jsonData, 0644)
}

func deployScatterProtocol(privateKey string) {
	auth := GetTransactor(privateKey)

	address, transaction, _, err := scatterprotocol.DeployScatterprotocol(
		auth, client,
		common.HexToAddress(utils.TRAINING_TOKEN_CONTRACT),
		common.HexToAddress(utils.EVALUATION_TOKEN_CONTRACT),
		common.HexToAddress(utils.SCATTER_TOKEN_CONTRACT),
		common.HexToAddress(utils.MODEL_TOKEN_CONTRACT))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract Address: " + address.Hex())
	log.Println("Transaction Hash: " + transaction.Hash().Hex())

}

func deployScatterToken(privateKey string) {
	auth := GetTransactor(privateKey)

	address, transaction, _, err := scattertoken.DeployScattertoken(
		auth, client, big.NewInt(int64(1_000_000_000_000)))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract Address: " + address.Hex())
	log.Println("Transaction Hash: " + transaction.Hash().Hex())

}

func deployEvaluationJobToken(privateKey string) {
	auth := GetTransactor(privateKey)

	address, transaction, _, err := evaluationtoken.DeployEvaluationtoken(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract Address: " + address.Hex())
	log.Println("Transaction Hash: " + transaction.Hash().Hex())

}

func deployModelToken(privateKey string) {
	auth := GetTransactor(privateKey)

	address, transaction, _, err := modelToken.DeployModeltoken(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract Address: " + address.Hex())
	log.Println("Transaction Hash: " + transaction.Hash().Hex())

}

func deployTrainingJobToken(privateKey string) {
	auth := GetTransactor(privateKey)
	address, transaction, _, err := trainingtoken.DeployTrainingtoken(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract Address: " + address.Hex())
	log.Println("Transaction Hash: " + transaction.Hash().Hex())

}

func GetTransactor(privateKey string) *bind.TransactOpts {
	privateKeyObject, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKeyObject, CHAIN)
	auth.Value = big.NewInt(0)
	gas, _ := client.SuggestGasPrice(context.Background())
	auth.GasPrice = (*big.Int)(gas)

	return auth
}
