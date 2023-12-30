package protocol

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	scatterlogs "github.com/gtfintechlab/scatter-protocol/core/logs"
	"github.com/gtfintechlab/scatter-protocol/core/networking"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	scatterprotocol "github.com/gtfintechlab/scatter-protocol/core/protocol/scatter-protocol"
	"github.com/gtfintechlab/scatter-protocol/core/utils"
	"golang.org/x/exp/slices"
)

func TrainingEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["TrainingInitialized"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	node.Subscriptions.TrainingInitialized = &subscription
	defer subscription.Unsubscribe()

	for {
		select {
		case err := <-subscription.Err():
			log.Printf("Subscription Closed %s", err)
			return
		case event := <-logs:
			eventUnpacked := utils.TrainingInitializedEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "TrainingInitialized", event.Data)
			if GetRoleByAddress(node, *node.BlockchainAddress) != utils.PEER_TRAINER {
				continue
			}

			_, subMapExists := (*node.TrainingLock)[eventUnpacked.Requestor.String()]
			if subMapExists && (*node.TrainingLock)[eventUnpacked.Requestor.String()][eventUnpacked.TopicName] {
				continue
			}

			if subMapExists {
				(*node.TrainingLock)[eventUnpacked.Requestor.String()][eventUnpacked.TopicName] = true
			} else {
				(*node.TrainingLock)[eventUnpacked.Requestor.String()] = map[string]bool{
					eventUnpacked.TopicName: true,
				}
			}

			// Run training procedure only if the node's address is a topic
			trainersForTopic := GetAllTrainersByAddressAndTopic(node, eventUnpacked.Requestor.String(), eventUnpacked.TopicName)
			if *node.LogMode {
				log.Println(utils.Green + "Training has been initialized" + utils.Reset)
			}
			if slices.Contains(trainersForTopic, strings.ToLower(*node.BlockchainAddress)) {
				node.JobQueue.Enqueue(utils.Job{
					ID:       uuid.New().String(),
					Function: TrainingHandler,
					Args:     []interface{}{node, eventUnpacked.Requestor.String(), eventUnpacked.TopicName},
				})
			}
		}
	}
}

func EvaluationRequestListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["RequestForEvaluationSet"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case err := <-subscription.Err():
			log.Printf("Subscription Closed %s", err)
			return
		case event := <-logs:
			if GetRoleByAddress(node, *node.BlockchainAddress) != utils.PEER_REQUESTOR {
				return
			}

			eventUnpacked := utils.EvaluationRequestEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "RequestForEvaluationSet", event.Data)
			if common.HexToAddress(*node.BlockchainAddress) == eventUnpacked.Requestor {
				if *node.LogMode {
					log.Println(utils.Green + "Evaluation set has been requested" + utils.Reset)
				}

				// Requestor sends the evaluation request to smart contract
				node.JobQueue.Enqueue(utils.Job{
					ID:       uuid.New().String(),
					Function: EvaluationRequestHandler,
					Args:     []interface{}{node, eventUnpacked.TopicName},
				})
			}
		}
	}
}

func ModelValidationListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["ModelReadyToValidate"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case err := <-subscription.Err():
			log.Printf("Subscription Closed %s", err)
			return
		case event := <-logs:
			if GetRoleByAddress(node, *node.BlockchainAddress) != utils.PEER_VALIDATOR {
				return
			}
			eventUnpacked := utils.ModelReadyToValidateEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "ModelReadyToValidate", event.Data)
			if IsValidatorForRequestorAndTopic(node, eventUnpacked.Requestor.String(), eventUnpacked.TopicName) {
				if *node.LogMode {
					log.Println(utils.Green + "Model is ready to be validated" + utils.Reset)
				}
				node.JobQueue.Enqueue(utils.Job{
					ID:       uuid.New().String(),
					Function: ModelValidationHandler,
					Args:     []interface{}{node, eventUnpacked.Requestor.String(), eventUnpacked.TopicName},
				})
			}
		}
	}
}

func DebugEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["DebugEvent"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case err := <-subscription.Err():
			log.Printf("Subscription Closed %s", err)
			return
		case event := <-logs:
			eventUnpacked := utils.DebugEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "DebugEvent", event.Data)
			log.Printf("%s%s%s", utils.Yellow, eventUnpacked.Message, utils.Reset)
		}
	}
}

func JobCompleteEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["JobComplete"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case err := <-subscription.Err():
			log.Printf("Subscription Closed %s", err)
			return
		case event := <-logs:
			eventUnpacked := utils.JobCompleteEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "JobComplete", event.Data)
			if !(*node.LogMode) {
				return
			}
			if *node.LogMode {

				if node.PeerType == utils.PEER_TRAINER || node.PeerType == utils.PEER_VALIDATOR {
					log.Println(utils.Green + "Training is Complete: Log is being written to database" + utils.Reset)
					balance, _ := new(big.Float).SetInt(GetScatterTokenBalance(node)).Float64()
					timestamp := float64(time.Now().UnixMilli())
					scatterlogs.CreateLogEvent(utils.LOG_EVENT_TOKEN_BALANCE, timestamp, balance, node)
				}

				log.Println(utils.Green + "Training is Complete: Token supply is being updated" + utils.Reset)
				UpdateTokenSupply(node)

				lotteryBalance, _ := new(big.Float).SetInt(GetLotteryBalance(node)).Float64()
				timestamp := float64(time.Now().UnixMilli())
				scatterlogs.CreateLogEvent(utils.LOTTERY_BALANCE, timestamp, lotteryBalance, node)

			}

		}
	}
}

func TrainingHandler(node *utils.PeerNode, requestorAddress string, topicName string) {
	if *node.DummyLoad {
		basePath, _ := os.Getwd()
		modelPath := fmt.Sprintf("%s/training/dummy/model.pth", basePath)
		PublishModel(node, modelPath, requestorAddress, topicName)
		return
	}

	dockerSetup()
	ipfsCid := GetTrainingJobFromAddressAndTopic(node, requestorAddress, topicName)
	dataPath := peerDatabase.GetDatapathFromAddressAndIpfs(node, requestorAddress, ipfsCid)
	downloadTrainingJob(ipfsCid, requestorAddress)
	buildTrainingImage(requestorAddress, ipfsCid, dataPath)
	runTrainingContainer(requestorAddress, ipfsCid)
	submitModel(node, requestorAddress, ipfsCid, topicName)

}

func EvaluationRequestHandler(node *utils.PeerNode, topicName string) {
	evaluationJobDataPath := peerDatabase.GetEvaluationJobDataFromAddressAndTopic(node, *node.BlockchainAddress, topicName)
	zippedJobBytes, _ := networking.ZipFolder(evaluationJobDataPath)
	zippedPath := fmt.Sprintf("%s/evaluation.zip", evaluationJobDataPath)
	networking.WriteBytesToFile(
		zippedPath,
		zippedJobBytes.Bytes(),
	)

	PublishEvaluationData(node, zippedPath, topicName)
	os.RemoveAll(zippedPath)
}

func ModelValidationHandler(node *utils.PeerNode, requestorAddress string, topicName string) {

	if *node.DummyLoad {
		trainers := GetAllTrainersByAddressAndTopic(node, requestorAddress, topicName)
		for _, trainer := range trainers {
			score := big.NewInt(utils.GetRandomNumber(40, 100))
			SubmitEvaluationScore(node, requestorAddress, topicName, trainer, score)
		}
		return
	}

	evaluationJobCid := GetEvaluationJobFromAddressAndTopic(node, requestorAddress, topicName)
	evaluationDataCid := GetEvaluationDataFromAddressAndTopic(node, requestorAddress, topicName)

	downloadEvaluationJob(requestorAddress, evaluationJobCid)
	downloadEvaluationJobData(requestorAddress, evaluationJobCid, evaluationDataCid)

	trainers := GetAllTrainersByAddressAndTopic(node, requestorAddress, topicName)
	for _, trainer := range trainers {
		downloadTrainerModel(node, requestorAddress, topicName, evaluationJobCid, trainer)
		buildEvaluationImage(requestorAddress, evaluationJobCid)
		score := runEvaluationContainer(requestorAddress, evaluationJobCid)
		fmt.Println("SCORE:", score)
		SubmitEvaluationScore(node, requestorAddress, topicName, trainer, score)
	}
}

func UpdateTokenSupply(node *utils.PeerNode) {
	client, _ := utils.DbConnect()
	workspaceId, _ := primitive.ObjectIDFromHex(*node.WorkspaceId)
	filter := bson.D{
		{Key: "workspaceId", Value: workspaceId},
		{Key: "blockchainAddress", Value: *node.BlockchainAddress},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	tokenSupply := GetScatterTokenBalance(node).Int64()
	update := bson.D{{"$set", bson.D{{"tokenSupply", tokenSupply}}}}
	var updatedDocument bson.M
	client.Collection("protocolnodes").FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedDocument)
}
