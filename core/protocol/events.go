package protocol

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"github.com/gtfintechlab/scatter-protocol/core/networking"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"

	scatterprotocol "github.com/gtfintechlab/scatter-protocol/core/protocol/scatter-protocol"
	"github.com/gtfintechlab/scatter-protocol/core/utils"
	"golang.org/x/exp/slices"
)

func TrainingEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(utils.SCATTER_PROTOCOL_CONTRACT)},
		Topics:    [][]common.Hash{{contractABI.Events["TrainingInitialized"].ID}},
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
			log.Fatal(err)
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
		Addresses: []common.Address{common.HexToAddress(utils.SCATTER_PROTOCOL_CONTRACT)},
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
			log.Fatal(err)
		case event := <-logs:
			if GetRoleByAddress(node, *node.BlockchainAddress) != utils.PEER_REQUESTOR {
				return
			}

			eventUnpacked := utils.EvaluationRequestEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "RequestForEvaluationSet", event.Data)
			if common.HexToAddress(*node.BlockchainAddress) == eventUnpacked.Requestor {
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
		Addresses: []common.Address{common.HexToAddress(utils.SCATTER_PROTOCOL_CONTRACT)},
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
			log.Fatal(err)
		case event := <-logs:
			if GetRoleByAddress(node, *node.BlockchainAddress) != utils.PEER_VALIDATOR {
				return
			}
			eventUnpacked := utils.ModelReadyToValidateEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "ModelReadyToValidate", event.Data)
			if IsValidatorForRequestorAndTopic(node, eventUnpacked.Requestor.String(), eventUnpacked.TopicName) {
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
		Addresses: []common.Address{common.HexToAddress(utils.SCATTER_PROTOCOL_CONTRACT)},
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
			log.Fatal(err)
		case event := <-logs:
			eventUnpacked := utils.DebugEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "DebugEvent", event.Data)
			log.Printf("%s%s%s", utils.Yellow, eventUnpacked.Message, utils.Reset)
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
			score := big.NewInt(75)
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
		SubmitEvaluationScore(node, requestorAddress, topicName, trainer, score)
	}
}