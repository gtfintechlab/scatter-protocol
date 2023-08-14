package peers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gtfintechlab/scatter-protocol/cosmos"
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

		switch nodeType := node.PeerType; nodeType {
		case utils.PEER_REQUESTOR:
			node.PeerType = utils.PEER_TRAINER

		case utils.PEER_TRAINER:
			node.PeerType = utils.PEER_REQUESTOR
		}

		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"newRole": node.PeerType,
		})
	}
}

func addTopic(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)

		switch nodeType := node.PeerType; nodeType {
		case utils.PEER_REQUESTOR:
			var requestBody utils.AddTopicRequestBody
			json.NewDecoder(request.Body).Decode(&requestBody)

			if checkIfTopicExistsForNode(node, requestBody.Topic) {
				networking.SendJson(response, map[string]interface{}{
					"success": false,
					"Error":   "Topic already exists for node",
				})
				return
			}
			var cs *utils.Cosmos
			addTopicFromInfo(node, node.NodeId.String(), requestBody.Topic, node.PeerType, nil)
			cs = cosmos.CreateCosmos(node, context.Background(),
				fmt.Sprintf("%s:%s", node.NodeId.String(), requestBody.Topic))

			go func() {
				for {
					message, _ := cs.Subscription.Next(context.Background())
					HandleCosmosMessage(message, node)
				}
			}()
			zippedFileBytes, _ := networking.ZipFolder("training/requestor")
			zippedJobPath := fmt.Sprintf("training/requestor/%s_%s.zip", node.NodeId.String(), requestBody.Topic)
			networking.WriteBytesToFile(zippedJobPath, zippedFileBytes.Bytes())

			topicCid, _ := utils.PublishTrainingJob(zippedJobPath, requestBody.Topic)
			cosmos.AddTopicToUniversalCosmos(node, requestBody.Topic, &utils.PublishTrainingJobPayload{
				TopicCid: topicCid,
			})

			networking.SendJson(response, map[string]interface{}{
				"success":    true,
				"cosmosId":   cs.CosmosId,
				"cosmosName": cs.CosmosName,
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

			addTopicFromInfo(node, node.NodeId.String(), requestBody.Topic, node.PeerType, requestBody.Path)
			cosmos.JoinCosmos(context.Background(), node,
				fmt.Sprintf("%s:%s", *requestBody.RequestorId, requestBody.Topic))
		}
	}
}

func getOwnTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"topics":  getTopicsByNodeId(node, node.NodeId.String()),
		})
	}
}

func getCosmosTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		cosmos.GetTopicsFromUniversalCosmos(node)
		node.InformationBox.InformationBoxMutexLock.Lock()
		defer node.InformationBox.InformationBoxMutexLock.Unlock()
		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"topics":  *(node.InformationBox.CosmosTopics),
		})
	}
}

func getTopicTrainers(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		networking.SendJson(response, map[string]interface{}{
			"success":  true,
			"trainers": getTrainersForAllTopics(node),
		})
	}
}

func initializeTraining(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)

		var requestBody utils.InitializeTrainingRequestBody
		json.NewDecoder(request.Body).Decode(&requestBody)
		switch nodeType := node.PeerType; nodeType {
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
