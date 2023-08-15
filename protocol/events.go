package protocol

import (
	"context"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	scatterprotocol "github.com/gtfintechlab/scatter-protocol/protocol/scatter"
	"github.com/gtfintechlab/scatter-protocol/utils"
	"golang.org/x/exp/slices"
)

func TrainingEventListener(node *utils.PeerNode) {
	contractABI, _ := abi.JSON(strings.NewReader(string(scatterprotocol.ScatterprotocolABI)))

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(utils.SCATTER_PROTCOL_CONTRACT)},
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
				RunTrainingProcedure(node, eventUnpacked.Requestor.String(), eventUnpacked.TopicName)
			}
		}
	}
}

func RunTrainingProcedure(node *utils.PeerNode, requestorId string, topicName string) {
	dockerSetup()
	ipfsCid := GetCidFromAddressAndTopic(node, requestorId, topicName)
	downloadTrainingJob(ipfsCid, requestorId)

	// buildImage(requestorId, topicName)
	// runContainer(requestorId, topicName)
}
