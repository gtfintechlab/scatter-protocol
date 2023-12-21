package api

import (
	"encoding/json"
	"net/http"

	"github.com/gtfintechlab/scatter-protocol/core/networking"
	"github.com/gtfintechlab/scatter-protocol/core/peers"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"
	"github.com/gtfintechlab/scatter-protocol/core/simulation"
	"github.com/gtfintechlab/scatter-protocol/core/utils"
)

var Nodes map[string]*utils.PeerNode

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
		simulation.ClearDatabase(node.DataStore)

		peerDatabase.MigratePeerDB(
			"up",
			requestBody.PeerType,
			requestBody.PostgresUsername,
			requestBody.PostgresPassword,
			int(requestBody.DatabasePort),
		)

		Nodes[requestBody.BlockchainAddress] = node
		go node.Start(node, requestBody.UseMdns)

		networking.SendJson(response, map[string]interface{}{"success": true})
	}

}

func stopNode() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.StopNodeRequest
		json.NewDecoder(request.Body).Decode(&requestBody)

		peers.StopPeer(Nodes[requestBody.BlockchainAddress])
		delete(Nodes, requestBody.BlockchainAddress)
		networking.SendJson(response, map[string]interface{}{"success": true})
	}
}
