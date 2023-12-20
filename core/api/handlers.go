package api

import (
	"encoding/json"
	"net/http"

	"github.com/gtfintechlab/scatter-protocol/core/networking"
	"github.com/gtfintechlab/scatter-protocol/core/peers"
	"github.com/gtfintechlab/scatter-protocol/core/utils"
)

func startNode() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.StartNodeRequest
		json.NewDecoder(request.Body).Decode(&requestBody)

		node := peers.InitPeerNode(
			requestBody.PeerType,
			int(requestBody.ApiPort),
			requestBody.PostgresUsername,
			requestBody.PostgresPassword,
			int(requestBody.DatabasePort),
			requestBody.BlockchainAddress,
			requestBody.PrivateKey,
			requestBody.DummyLoad,
		)

		go node.Start(node, requestBody.UseMdns)
	}

}
