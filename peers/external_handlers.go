package peers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/peers/db"
	protocol "github.com/gtfintechlab/scatter-protocol/protocol"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
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

		switch nodeType := protocol.GetRoleByAddress(node, *node.BlockchainAddress); nodeType {
		case utils.PEER_REQUESTOR:
			protocol.ChangeRoleToTrainer(node)

		case utils.PEER_TRAINER:
			protocol.ChangeRoleToRequestor(node)
		}

		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"newRole": protocol.GetRoleByAddress(node, *node.BlockchainAddress),
		})
	}
}

func addTopic(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)

		switch nodeType := protocol.GetRoleByAddress(node, *node.BlockchainAddress); nodeType {
		case utils.PEER_REQUESTOR:
			var requestBody utils.AddTopicRequestBody
			json.NewDecoder(request.Body).Decode(&requestBody)

			if protocol.CheckIfTopicExistsForRequestor(node, *node.BlockchainAddress, requestBody.Topic) {
				networking.SendJson(response, map[string]interface{}{
					"success": false,
					"Error":   "Topic already exists for node",
				})
				return
			}

			if requestBody.EvaluationJob == nil {
				networking.SendJson(response, map[string]interface{}{
					"success": false,
					"Error":   "requestors must include an evaluation job",
				})
				return
			}

			zippedFileBytes, _ := networking.ZipFolder(*requestBody.Path)
			zippedJobPath := fmt.Sprintf("%s/%s_%s.zip", *requestBody.Path, *node.BlockchainAddress, requestBody.Topic)
			networking.WriteBytesToFile(zippedJobPath, zippedFileBytes.Bytes())
			topicCid, _ := protocol.AddTopicForRequestor(node, zippedJobPath, requestBody.Topic, *requestBody.Reward)
			peerDatabase.AddTopicFromInfo(
				node,
				*node.BlockchainAddress,
				topicCid,
				requestBody.Topic,
				nil,
				requestBody.EvaluationJob,
			)

			os.RemoveAll(zippedJobPath)
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
			protocol.AddTopicForTrainer(node, *requestBody.RequestorAddress, requestBody.Topic)
			peerDatabase.AddTopicFromInfo(
				node,
				*requestBody.RequestorAddress,
				protocol.GetCidFromAddressAndTopic(node, *requestBody.RequestorAddress, requestBody.Topic),
				requestBody.Topic,
				requestBody.Path,
				nil,
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
			"topics":  protocol.GetAllTopicsByAddress(node, *node.BlockchainAddress),
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
			participants = protocol.GetProtocolRequestors(node, uint64(parsedSkipCount))
		} else {
			participants = protocol.GetAllProtocolRequestors(node)
		}

		for _, participant := range participants {
			topics := protocol.GetAllTopicsByAddress(node, participant)

			for _, topic := range topics {
				publishedTopics = append(publishedTopics, utils.TopicInformation{
					NodeAddress:      participant,
					NodeType:         protocol.GetRoleByAddress(node, participant),
					TopicName:        topic,
					TrainingTokenCID: protocol.GetCidFromAddressAndTopic(node, participant, topic),
					PooledReward:     protocol.GetPooledRewardByAddressAndTopic(node, participant, topic),
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
		topics := protocol.GetAllTopicsByAddress(node, *node.BlockchainAddress)
		topicTrainerMap := map[string][]string{}

		for _, topic := range topics {
			trainers := protocol.GetAllTrainersByAddressAndTopic(node, *node.BlockchainAddress, topic)
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
		protocol.StartTraining(node, requestBody.Topic)

	}
}

func getScatterTokenBalance(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"balance": protocol.GetScatterTokenBalance(node),
		})

	}
}
