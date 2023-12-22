package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

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
		var requestBody utils.SimulationStartNodeRequest
		json.NewDecoder(request.Body).Decode(&requestBody)

		node := peers.InitPeerNode(
			requestBody.PeerType,
			int(requestBody.ApiPort),
			requestBody.PostgresUsername,
			requestBody.PostgresPassword,
			int(requestBody.DatabasePort),
			requestBody.BlockchainAddress,
			utils.RemoveHexPrefix(requestBody.PrivateKey),
			requestBody.DummyLoad,
		)
		if int(requestBody.DatabasePort) != 0 {
			simulation.ClearDatabase(node.DataStore)
			peerDatabase.MigratePeerDB(
				"up",
				requestBody.PeerType,
				requestBody.PostgresUsername,
				requestBody.PostgresPassword,
				int(requestBody.DatabasePort),
			)
		}
		if Nodes == nil {
			Nodes = make(map[string]*utils.PeerNode)
		}

		Nodes[requestBody.BlockchainAddress] = node
		go node.Start(node, requestBody.UseMdns)

		networking.SendJson(response, map[string]interface{}{"success": true})
	}

}

func addTopic() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.SimulationAppTopicRequest
		json.NewDecoder(request.Body).Decode(&requestBody)
		node := Nodes[requestBody.BlockchainAddress]
		bodyData, err := json.Marshal(requestBody)

		if err != nil {
			log.Fatal(err)
		}

		networking.MakePostRequestToServer(node.ExternalServer, "/add-topic", bodyData)
		networking.SendJson(response, map[string]interface{}{"success": true})
	}
}

func stopNode() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.SimulationStopNodeRequest
		json.NewDecoder(request.Body).Decode(&requestBody)

		peers.StopPeer(Nodes[requestBody.BlockchainAddress])
		delete(Nodes, requestBody.BlockchainAddress)

		networking.SendJson(response, map[string]interface{}{"success": true})
	}
}

func deployProtocol() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.SimulationDeployProtocolRequest
		json.NewDecoder(request.Body).Decode(&requestBody)
		utils.ReplaceEnvValue(".env", "BLOCKCHAIN_ADDRESS", utils.RemoveHexPrefix(requestBody.BlockchainAddress))
		utils.ReplaceEnvValue(".env", "PRIVATE_KEY", utils.RemoveHexPrefix(requestBody.PrivateKey))

		exec.Command("docker-compose", "-f", "docker-compose-contracts.yml", "up").Run()
		contracts := utils.ReadContractInfo()
		networking.SendJson(response, map[string]interface{}{"success": true, "payload": contracts})

	}
}

func transferInitialSupply() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.SimulationTransferInitialSupplyRequest
		json.NewDecoder(request.Body).Decode(&requestBody)

		for address, amount := range requestBody.TransferAmounts {
			transferToken(
				requestBody.PrivateKey,
				int64(amount),
				address,
				"ws://127.0.0.1:8545",
				utils.ReadContractInfo().ScatterTokenContract,
			)
		}

	}
}
