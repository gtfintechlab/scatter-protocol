package peers

import (
	"encoding/json"
	"fmt"
	"net/http"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
)

func switchPeerNodeRole(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.GetValidator(request, response)

		switch nodeType := node.PeerType; nodeType {
		case utils.PEER_REQUESTOR:
			node.PeerType = utils.PEER_TRAINER

		case utils.PEER_TRAINER:
			node.PeerType = utils.PEER_REQUESTOR
		}
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
		}

		fmt.Print(node.Topics)
	}
}

func publishTopic(node *utils.PeerNode) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if node.PeerType == utils.PEER_REQUESTOR {

		}
	}
}
