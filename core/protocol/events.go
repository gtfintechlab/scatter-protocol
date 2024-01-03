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

	if node.Subscriptions.TrainingInitialized != nil {
		(*node.Subscriptions.TrainingInitialized).Unsubscribe()
	}
	node.Subscriptions.TrainingInitialized = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
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
	if node.Subscriptions.RequestForEvaluationSet != nil {
		(*node.Subscriptions.RequestForEvaluationSet).Unsubscribe()
	}
	node.Subscriptions.RequestForEvaluationSet = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
			return
		case event := <-logs:
			if GetRoleByAddress(node, *node.BlockchainAddress) != utils.PEER_REQUESTOR {
				continue
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
	if node.Subscriptions.ModelReadyToValidate != nil {
		(*node.Subscriptions.ModelReadyToValidate).Unsubscribe()
	}
	node.Subscriptions.ModelReadyToValidate = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
			return
		case event := <-logs:
			if GetRoleByAddress(node, *node.BlockchainAddress) != utils.PEER_VALIDATOR {
				continue
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
	if node.Subscriptions.DebugEvent != nil {
		(*node.Subscriptions.DebugEvent).Unsubscribe()
	}
	node.Subscriptions.DebugEvent = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
			return
		case event := <-logs:
			eventUnpacked := utils.DebugEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "DebugEvent", event.Data)
			log.Printf("%s%s%s", utils.Yellow, eventUnpacked.Message, utils.Reset)
		}
	}
}

func RequestForChallengesEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["RequestForChallenges"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	if node.Subscriptions.RequestForChallenges != nil {
		(*node.Subscriptions.RequestForChallenges).Unsubscribe()
	}
	node.Subscriptions.RequestForChallenges = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
			return
		case event := <-logs:
			eventUnpacked := utils.RequestForChallengesEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "RequestForChallenges", event.Data)
			node.JobQueue.Enqueue(utils.Job{
				ID:       uuid.New().String(),
				Function: RequestForChallengesHandler,
				Args:     []interface{}{node, eventUnpacked},
			})

		}
	}
}

func ChallengeStartedEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["ChallengeStarted"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	if node.Subscriptions.ChallengeStarted != nil {
		(*node.Subscriptions.ChallengeStarted).Unsubscribe()
	}
	node.Subscriptions.ChallengeStarted = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
			return
		case event := <-logs:
			eventUnpacked := utils.ChallengeStartedEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "ChallengeStarted", event.Data)
			node.JobQueue.Enqueue(utils.Job{
				ID:       uuid.New().String(),
				Function: ChallengeStartedHandler,
				Args:     []interface{}{node, eventUnpacked},
			})

		}
	}
}

func RecordLogsEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(GetContractInfo().ScatterProtocolContract)},
		Topics:    [][]common.Hash{{contractABI.Events["RecordLogs"].ID}},
	}

	logs := make(chan types.Log)
	subscription, err := ethereumClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	if node.Subscriptions.RecordLogs != nil {
		(*node.Subscriptions.RecordLogs).Unsubscribe()
	}
	node.Subscriptions.RecordLogs = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
			return
		case event := <-logs:
			eventUnpacked := utils.RecordLogsEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "RecordLogs", event.Data)
			RecordLogsHandler(node, eventUnpacked)
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
	if node.Subscriptions.JobComplete != nil {
		(*node.Subscriptions.JobComplete).Unsubscribe()
	}
	node.Subscriptions.JobComplete = &subscription
	// defer subscription.Unsubscribe()

	for {
		select {
		case <-subscription.Err():
			return
		case event := <-logs:
			eventUnpacked := utils.JobCompleteEvent{}
			contractABI.UnpackIntoInterface(&eventUnpacked, "JobComplete", event.Data)
			node.JobQueue.Enqueue(utils.Job{
				ID:       uuid.New().String(),
				Function: JobCompleteHandler,
				Args:     []interface{}{node, eventUnpacked},
			})

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
			var score *big.Int
			if scatterlogs.IsNodeMalicious(node, trainer) {
				// If validator is malicious, submit good score for bad trainer
				if scatterlogs.IsNodeMalicious(node, *node.BlockchainAddress) {
					score = big.NewInt(utils.GetRandomNumber(40, 100))

				} else {
					score = big.NewInt(utils.GetRandomNumber(0, 39))
				}
			} else {
				// If validator is malicious, submit bad score for good trainer
				if scatterlogs.IsNodeMalicious(node, *node.BlockchainAddress) {
					score = big.NewInt(utils.GetRandomNumber(30, 39))

				} else {
					score = big.NewInt(utils.GetRandomNumber(80, 100))
				}
			}

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

func UpdateTokenSupply(node *utils.PeerNode) {
	client, err := utils.DbConnect()
	for client == nil || err != nil {
		client, err = utils.DbConnect()
	}
	defer client.Client().Disconnect(context.Background())
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

func RequestForChallengesHandler(node *utils.PeerNode, eventUnpacked utils.RequestForChallengesEvent) {
	_, subMapExists := (*node.ChallengeLock)[eventUnpacked.RequestorAddress.String()]
	if subMapExists && (*node.ChallengeLock)[eventUnpacked.RequestorAddress.String()][eventUnpacked.TopicName] {
		return
	}

	if subMapExists {
		(*node.ChallengeLock)[eventUnpacked.RequestorAddress.String()][eventUnpacked.TopicName] = true
	} else {
		(*node.ChallengeLock)[eventUnpacked.RequestorAddress.String()] = map[string]bool{
			eventUnpacked.TopicName: true,
		}
	}

	if GetRoleByAddress(node, *node.BlockchainAddress) == utils.PEER_CHALLENGER {
		trainers := GetAllTrainersByAddressAndTopic(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName)
		validators := GetAllValidatorsByAddressAndTopic(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName)
		combined := append(trainers, validators...)

		if *node.LogMode {
			lotteryBalance, _ := new(big.Float).SetInt(GetLotteryBalance(node)).Float64()
			blockNum := GetBlockNumber()
			scatterlogs.CreateLogEvent(utils.LOTTERY_BALANCE, blockNum, lotteryBalance, node)

			log.Println(utils.Green + "Now Challenging Nodes" + utils.Reset)
			for _, nodeAddress := range combined {
				if scatterlogs.IsNodeMalicious(node, nodeAddress) {
					ChallengeNode(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName, nodeAddress, true)
				} else {
					ChallengeNode(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName, nodeAddress, false)
				}
			}
		}

		ChallengeComplete(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName)
	}
}

func ChallengeStartedHandler(node *utils.PeerNode, eventUnpacked utils.ChallengeStartedEvent) {
	if GetRoleByAddress(node, *node.BlockchainAddress) == utils.PEER_CHALLENGER {

		if *node.LogMode {
			log.Println(utils.Green + "Now Challenging Singular Node" + utils.Reset)
			if scatterlogs.IsNodeMalicious(node, eventUnpacked.NodeAddress.String()) {
				ChallengeNode(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName, eventUnpacked.NodeAddress.String(), true)
			} else {
				ChallengeNode(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName, eventUnpacked.NodeAddress.String(), false)
			}
		}

		ChallengeComplete(node, eventUnpacked.RequestorAddress.String(), eventUnpacked.TopicName)
	}

}

func RecordLogsHandler(node *utils.PeerNode, eventUnpacked utils.RecordLogsEvent) {
	balance, _ := new(big.Float).SetInt(GetScatterTokenBalance(node)).Float64()
	blockNum := GetBlockNumber()
	scatterlogs.CreateLogEvent(utils.LOG_EVENT_TOKEN_BALANCE, blockNum, balance, node)

	blockNum = GetBlockNumber()
	lottery, _ := new(big.Float).SetInt(eventUnpacked.LotteryAmount).Float64()
	scatterlogs.CreateLogEvent(utils.LOTTERY_BALANCE, blockNum, lottery, node)

	if node.PeerType == utils.PEER_VALIDATOR {
		stake, _ := new(big.Float).SetInt(GetScatterTokenStake(node, *node.BlockchainAddress)).Float64()
		blockNum, _ = new(big.Float).SetInt(eventUnpacked.BlockNumber).Float64()
		scatterlogs.CreateLogEvent(utils.STAKE_BALANCE, blockNum, stake, node)
	}
}

func JobCompleteHandler(node *utils.PeerNode, eventUnpacked utils.JobCompleteEvent) {
	if *node.LogMode {

		log.Println(utils.Green + "Training is Complete: Log is being written to database" + utils.Reset)
		balance, _ := new(big.Float).SetInt(GetScatterTokenBalance(node)).Float64()
		blockNum := GetBlockNumber()
		scatterlogs.CreateLogEvent(utils.LOG_EVENT_TOKEN_BALANCE, blockNum, balance, node)

		log.Println(utils.Green + "Training is Complete: Token supply is being updated" + utils.Reset)
		UpdateTokenSupply(node)

		lotteryBalance, _ := new(big.Float).SetInt(GetLotteryBalance(node)).Float64()
		blockNum = GetBlockNumber()
		scatterlogs.CreateLogEvent(utils.LOTTERY_BALANCE, blockNum, lotteryBalance, node)

		if node.PeerType == utils.PEER_VALIDATOR {
			blockNum = GetBlockNumber()
			stake, _ := new(big.Float).SetInt(GetScatterTokenStake(node, *node.BlockchainAddress)).Float64()
			scatterlogs.CreateLogEvent(utils.STAKE_BALANCE, blockNum, stake, node)
		}

	}

}
