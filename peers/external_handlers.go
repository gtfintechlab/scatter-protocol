package peers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gtfintechlab/scatter-protocol/cosmos"
	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
)

func health() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		healthy := map[string]string{
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

		networking.SendJson(response, map[string]any{
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
			node.Topics[requestBody.Topic] = requestBody.Topic
			cs := cosmos.CreateCosmos(node, context.Background(), requestBody.Topic)
			go func() {
				for {
					message, _ := cs.Subscription.Next(context.Background())
					HandleCosmosMessage(message, node)
				}
			}()

			cosmos.AddTopicToUniversalCosmos(node, requestBody.Topic)

			networking.SendJson(response, map[string]any{
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

			node.Topics[requestBody.Topic] = *requestBody.Path

			cosmos.JoinCosmos(context.Background(), node, requestBody.Topic)
		}
	}
}

func getOwnTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)
		networking.SendJson(response, map[string]any{
			"success": true,
			"topics":  node.Topics,
		})
	}
}

func getCosmosTopics(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)

		cosmos.GetTopicsFromUniversalCosmos(node)
		networking.SendJson(response, map[string]any{
			"success": true,
			"topics":  node.CosmosTopics,
		})
	}
}
