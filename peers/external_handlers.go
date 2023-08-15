package peers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	"github.com/libp2p/go-libp2p/core/peer"
)

func health(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		healthy := map[string]interface{}{
			"Hello":  "World",
			"nodeId": node.NodeId,
		}

		networking.SendJson(response, healthy)
	}
}

func switchPeerNodeRole(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)

		switch nodeType := utils.GetRoleByAddress(*node.BlockchainAddress); nodeType {
		case utils.PEER_REQUESTOR:
			utils.ChangeRoleToTrainer()

		case utils.PEER_TRAINER:
			utils.ChangeRoleToRequestor()
		}

		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"newRole": utils.GetRoleByAddress(*node.BlockchainAddress),
		})
	}
}

func addTopic(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)

		switch nodeType := utils.GetRoleByAddress(*node.BlockchainAddress); nodeType {
		case utils.PEER_REQUESTOR:
			var requestBody utils.AddTopicRequestBody
			json.NewDecoder(request.Body).Decode(&requestBody)

			if utils.CheckIfTopicExistsForRequestor(*node.BlockchainAddress, requestBody.Topic) {
				networking.SendJson(response, map[string]interface{}{
					"success": false,
					"Error":   "Topic already exists for node",
				})
				return
			}

			zippedFileBytes, _ := networking.ZipFolder(*requestBody.Path)
			zippedJobPath := fmt.Sprintf("%s/%s_%s.zip", *requestBody.Path, *node.BlockchainAddress, requestBody.Topic)
			networking.WriteBytesToFile(zippedJobPath, zippedFileBytes.Bytes())

			topicCid, _ := utils.AddTopicForRequestor(zippedJobPath, requestBody.Topic)

			networking.SendJson(response, map[string]interface{}{
				"success":  true,
				"topicCid": topicCid,
			})
		case utils.PEER_TRAINER:
			var requestBody utils.AddTopicRequestBody
			json.NewDecoder(request.Body).Decode(&requestBody)

			if requestBody.Path == nil {
				response.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(response,
					"Trainer nodes must specify a path to the data of the topic they want to subscribe to")
				return
			}
			utils.AddTopicForTrainer(*requestBody.RequestorAddress, requestBody.Topic)
			addTopicFromInfo(
				node,
				*requestBody.RequestorAddress,
				requestBody.Topic,
				utils.GetRoleByAddress(*node.BlockchainAddress),
				requestBody.Path,
			)

			networking.SendJson(response, map[string]interface{}{
				"success":            true,
				"requestor_address":  *requestBody.RequestorAddress,
				"topic":              requestBody.Topic,
				"training_data_path": requestBody.Path,
			})
		}
	}
}

// An endpoint to get one's own topics on the blockchain
func getOwnTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"topics":  utils.GetAllTopicsByAddress(*node.BlockchainAddress),
		})
	}
}

// An endpoint to get all published topics on the blockchain
func getPublishedTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)

		addressSkipCount := request.URL.Query().Get("skip")
		var publishedTopics []utils.TopicInformation
		var participants []string

		if addressSkipCount != "" {
			parsedSkipCount, _ := strconv.Atoi(addressSkipCount)
			participants = utils.GetProtocolRequestors(uint64(parsedSkipCount))
		} else {
			participants = utils.GetAllProtocolRequestors()
		}

		for _, participant := range participants {
			topics := utils.GetAllTopicsByAddress(participant)

			for _, topic := range topics {
				publishedTopics = append(publishedTopics, utils.TopicInformation{
					NodeAddress:      participant,
					NodeType:         utils.GetRoleByAddress(participant),
					TopicName:        topic,
					TrainingTokenCID: utils.GetCidFromAddressAndTopic(participant, topic),
				})
			}
		}

		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"topics":  publishedTopics,
		})
	}
}

// An endpoint to get trainers for your own topics on the blockchain
func getTopicTrainers(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		topics := utils.GetAllTopicsByAddress(*node.BlockchainAddress)
		topicTrainerMap := map[string][]string{}

		for _, topic := range topics {
			trainers := utils.GetAllTrainersByAddressAndTopic(*node.BlockchainAddress, topic)
			topicTrainerMap[topic] = trainers
		}
		networking.SendJson(response, map[string]interface{}{
			"success":  true,
			"trainers": topicTrainerMap,
		})
	}
}

func initializeTraining(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)

		var requestBody utils.InitializeTrainingRequestBody
		json.NewDecoder(request.Body).Decode(&requestBody)
		switch nodeType := utils.GetRoleByAddress(*node.BlockchainAddress); nodeType {
		case utils.PEER_REQUESTOR:
			trainers := getTrainersByTopic(node, requestBody.Topic)
			zippedFileBytes, _ := networking.ZipFolder("training/requestor")
			for _, trainer := range trainers {
				peerId, _ := peer.Decode(trainer)
				stream, _ := (*node.PeerToPeerServer).NewStream(
					context.Background(),
					peerId,
					utils.PROTOCOL_IDENTIFIER,
				)
				networking.SendMessage(&stream, utils.Message{
					MessageType: utils.PEER_START_TRAINING,
					Payload: map[string]interface{}{
						"files": zippedFileBytes.Bytes(),
						"topic": requestBody.Topic,
					},
				})
			}
		}

	}
}
