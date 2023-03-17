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
	"golang.org/x/exp/maps"
)

func health() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		healthy := map[string]interface{}{
			"Hello": "World",
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
			(*node.TopicToDataPath)[requestBody.Topic] = ""
			cs := cosmos.CreateCosmos(node, context.Background(), requestBody.Topic)
			go func() {
				for {
					message, _ := cs.Subscription.Next(context.Background())
					HandleCosmosMessage(message, node)
				}
			}()

			cosmos.AddTopicToUniversalCosmos(node, requestBody.Topic)

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

			(*node.TopicToDataPath)[requestBody.Topic] = *requestBody.Path
			if _, ok := (*node.RequestorTopicMap)[requestBody.RequestorId]; ok {
				(*node.RequestorTopicMap)[requestBody.RequestorId] =
					append((*node.RequestorTopicMap)[requestBody.RequestorId],
						requestBody.Topic)
			} else {
				(*node.RequestorTopicMap)[requestBody.RequestorId] = []string{}
				(*node.RequestorTopicMap)[requestBody.RequestorId] =
					append((*node.RequestorTopicMap)[requestBody.RequestorId],
						requestBody.Topic)
			}
			cosmos.JoinCosmos(context.Background(), node,
				fmt.Sprintf("%s: %s", requestBody.RequestorId, requestBody.Topic))
		}
	}
}

func getOwnTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		networking.SendJson(response, map[string]interface{}{
			"success": true,
			"topics":  maps.Keys((*node.TopicToDataPath)),
		})
	}
}

func getCosmosTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)

		cosmos.GetTopicsFromUniversalCosmos(node)
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
			"trainers": node.TopicTrainerMap,
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
			trainers := (*node.TopicTrainerMap)[fmt.Sprintf("%s: %s",
				node.NodeId.String(),
				requestBody.Topic)]

			for _, trainer := range trainers {
				peerId, _ := peer.Decode(trainer)
				stream, _ := (*node.PeerToPeerServer).NewStream(
					context.Background(),
					peerId,
					utils.PROTOCOL_IDENTIFIER,
				)
				networking.SendMessage(&stream, utils.Message{
					MessageType: utils.PEER_START_TRAINING,
					Payload:     map[string]interface{}{},
				})
			}
		}

	}
}
