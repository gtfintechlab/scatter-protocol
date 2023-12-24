package main

import (
	"context"
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
	evaluationtoken "github.com/gtfintechlab/scatter-protocol/core/protocol/evaluation"
	modelToken "github.com/gtfintechlab/scatter-protocol/core/protocol/model"
	reputationmanager "github.com/gtfintechlab/scatter-protocol/core/protocol/reputation"
	scatterprotocol "github.com/gtfintechlab/scatter-protocol/core/protocol/scatter-protocol"
	scattertoken "github.com/gtfintechlab/scatter-protocol/core/protocol/scatter-token"
	trainingtoken "github.com/gtfintechlab/scatter-protocol/core/protocol/training"
	votemanager "github.com/gtfintechlab/scatter-protocol/core/protocol/vote-manager"
	"github.com/gtfintechlab/scatter-protocol/core/utils"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/joho/godotenv"
)

const (
	PROTOCOL   = "protocol"
	MODEL      = "model"
	TOKEN      = "token"
	EVALUATION = "evaluation"
	TRAINING   = "training"
	REPUTATION = "reputation"
	ALL        = "all"
)

var _ = godotenv.Load(".env")
var CHAIN_ID, _ = strconv.Atoi(getChainId())
var CHAIN = big.NewInt(int64(CHAIN_ID))
var client = getClient()

var DEPLOY_PAUSE time.Duration = 3

func getChainId() string {
	if os.Getenv("CHAIN_ID") == "" {
		return "31337"
	}

	return os.Getenv("CHAIN_ID")
}

func getEthNode() string {
	if os.Getenv("ETHEREUM_NODE") == "" {
		return "ws://127.0.0.1:8545"
	}

	return os.Getenv("ETHEREUM_NODE")
}

func getClient() *ethclient.Client {
	var client, err = ethclient.Dial(getEthNode())
	if err != nil {
		log.Fatal("An error occurred when connecting to client:", err)
	}

	return client
}

func main() {
	var contractName string
	var privateKey string
	flag.StringVar(&contractName, "contract", "all",
		"Name of contract to deploy")
	flag.StringVar(&privateKey, "pkey", os.Getenv("PRIVATE_KEY"),
		"Private key to deploy protocol with")
	flag.Parse()

	if len(privateKey) == 0 {
		privateKey = os.Getenv("PRIVATE_KEY")
	}

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
	case REPUTATION:
		deployReputationManagerContract(privateKey)
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

	reputationManagerAddress, reputationManagerTransaction, reputationManagerInstance, err := reputationmanager.DeployReputationmanager(auth, client)
	if err != nil {
		log.Fatal("Failed to deploy reptuation manager", err.Error())
	}
	time.Sleep(time.Second * DEPLOY_PAUSE)

	voteManagerAddress, voteManagerTransaction, voteManagerInstance, err := votemanager.DeployVotemanager(auth, client)
	if err != nil {
		log.Fatal("Failed to deploy reptuation manager", err.Error())
	}
	time.Sleep(time.Second * DEPLOY_PAUSE)

	scatterProtocolAddress, scatterProtocolTransaction, scatterProtocolInstance, err := scatterprotocol.DeployScatterprotocol(
		auth, client,
		common.HexToAddress(trainingAddress.Hash().Hex()),
		common.HexToAddress(evaluationAddress.Hash().Hex()),
		common.HexToAddress(scatterTokenAddress.Hash().Hex()),
		common.HexToAddress(modelAddress.Hash().Hex()),
		common.HexToAddress(reputationManagerAddress.Hash().Hex()),
		common.HexToAddress(voteManagerAddress.Hash().Hex()),
	)
	if err != nil {
		log.Fatal("Failed to deploy scatter protocol", err.Error())
	}
	time.Sleep(time.Second * 10)

	trainingInstance.SetScatterContractAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	evaluationInstance.SetScatterContractAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	scatterTokenInstance.SetScatterProtocolAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	modelInstance.SetScatterContractAddress(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	reputationManagerInstance.SetScatterProtocolContract(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))
	voteManagerInstance.SetScatterProtocolContract(auth, common.HexToAddress(scatterProtocolAddress.Hash().Hex()))

	scatterTokenInstance.SetEvaluationJobTokenContract(auth, common.HexToAddress(evaluationAddress.Hash().Hex()))
	scatterTokenInstance.SetReputationManagerContract(auth, common.HexToAddress(reputationManagerAddress.Hash().Hex()))

	reputationManagerInstance.SetEvaluationJobContract(auth, common.HexToAddress(evaluationAddress.Hash().Hex()))
	reputationManagerInstance.SetModelTokenContract(auth, common.HexToAddress(modelAddress.Hash().Hex()))
	reputationManagerInstance.SetVoteManagerContract(auth, common.HexToAddress(voteManagerAddress.Hash().Hex()))

	evaluationInstance.SetVoteManagerContract(auth, common.HexToAddress(voteManagerAddress.Hash().Hex()))

	voteManagerInstance.SetEvaluationJobTokenContract(auth, common.HexToAddress(evaluationAddress.Hash().Hex()))

	scatterProtocolInstance.InitRequestorNode(auth)

	log.Println("Transaction Info:")
	log.Printf("Training Token: %s\n", trainingTransaction.Hash())
	log.Printf("Model Token: %s\n", modelTransaction.Hash())
	log.Printf("Evaluation Token: %s\n", evaluationTransaction.Hash())
	log.Printf("Scatter Token: %s\n", scatterTokenTransaction.Hash())
	log.Printf("Reputation Manager: %s\n", reputationManagerTransaction.Hash())
	log.Printf("Vote Manager: %s\n", voteManagerTransaction.Hash())
	log.Printf("Scatter Protocol: %s\n", scatterProtocolTransaction.Hash())
	log.Println("===============")

	log.Println("Contract Info:")
	log.Printf("Training Token: %s\n", trainingAddress.Hex())
	log.Printf("Model Token: %s\n", modelAddress.Hex())
	log.Printf("Evaluation Token: %s\n", evaluationAddress.Hex())
	log.Printf("Scatter Token: %s\n", scatterTokenAddress.Hex())
	log.Printf("Reputation Manager: %s\n", reputationManagerAddress.Hex())
	log.Printf("Vote Manager: %s\n", voteManagerAddress.Hex())
	log.Printf("Scatter Protocol: %s\n", scatterProtocolAddress.Hex())

	contractInfo := utils.Contracts{
		ScatterProtocolContract:   scatterProtocolAddress.Hex(),
		ScatterTokenContract:      scatterTokenAddress.Hex(),
		TrainingTokenContract:     trainingAddress.Hex(),
		EvaluationTokenContract:   evaluationAddress.Hex(),
		ModelTokenContract:        modelAddress.Hex(),
		ReputationManagerContract: reputationManagerAddress.Hex(),
		VoteManagerContract:       voteManagerAddress.Hex(),
	}

	UpdateContractInfo(contractInfo)
}

func deployScatterProtocol(privateKey string) {
	auth := GetTransactor(privateKey)

	address, transaction, _, err := scatterprotocol.DeployScatterprotocol(
		auth, client,
		common.HexToAddress(GetContractInfo().TrainingTokenContract),
		common.HexToAddress(GetContractInfo().EvaluationTokenContract),
		common.HexToAddress(GetContractInfo().ScatterTokenContract),
		common.HexToAddress(GetContractInfo().ModelTokenContract),
		common.HexToAddress(GetContractInfo().ReputationManagerContract),
		common.HexToAddress(GetContractInfo().VoteManagerContract),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contract Address: " + address.Hex())
	log.Println("Transaction Hash: " + transaction.Hash().Hex())
}

func UpdateContractInfo(contracts utils.Contracts) {
	client, _ := utils.DbConnect()
	contractsCollection := client.Collection("contracts")
	contractsCollection.DeleteMany(context.Background(), bson.D{})
	contractsCollection.InsertOne(context.Background(), contracts)

}

func GetContractInfo() utils.Contracts {
	client, _ := utils.DbConnect()
	contractsCollection := client.Collection("contracts")
	singleResult := contractsCollection.FindOne(context.Background(), bson.D{})
	var contracts utils.Contracts
	if err := singleResult.Decode(&contracts); err != nil {
		return utils.Contracts{}
	}
	return contracts
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

func deployReputationManagerContract(privateKey string) {
	auth := GetTransactor(privateKey)

	address, transaction, _, err := reputationmanager.DeployReputationmanager(auth, client)
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
	gas, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	auth.GasPrice = (*big.Int)(gas)

	return auth
}
