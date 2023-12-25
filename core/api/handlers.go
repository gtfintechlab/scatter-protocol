package api

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os/exec"
	"time"

	scatterlogs "github.com/gtfintechlab/scatter-protocol/core/logs"
	"github.com/gtfintechlab/scatter-protocol/core/networking"
	"github.com/gtfintechlab/scatter-protocol/core/peers"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"
	"github.com/gtfintechlab/scatter-protocol/core/protocol"
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
			true,
			requestBody.WorkspaceId,
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

		networking.MakePostRequestToServer(node.ExternalServer, "/node/topic/add", bodyData)
		networking.SendJson(response, map[string]interface{}{"success": true})
	}
}

func initializeRoles() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		for _, node := range Nodes {
			if node.PeerType == utils.PEER_VALIDATOR {
				protocol.AddScatterTokenStake(node, utils.VALIDATOR_STAKE)
				protocol.InitValidatorNode(node)
			} else if node.PeerType == utils.PEER_REQUESTOR {
				protocol.InitRequestorNode(node)
			} else if node.PeerType == utils.PEER_TRAINER {
				protocol.InitTrainerNode(node)
			}
		}

		for _, node := range Nodes {
			if *node.LogMode {
				balance, _ := new(big.Float).SetInt(protocol.GetScatterTokenBalance(node)).Float64()
				timestamp := float64(time.Now().UnixMilli())
				scatterlogs.CreateLogEvent(utils.LOG_EVENT_TOKEN_BALANCE, timestamp, balance, node)
			}
		}
		networking.SendJson(response, map[string]interface{}{"success": true})
	}
}

func stopNode() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.SimulationStopNodeRequest
		json.NewDecoder(request.Body).Decode(&requestBody)

		if Nodes == nil {
			networking.SendJson(response, map[string]interface{}{"success": true})
			return
		}
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
		contracts := protocol.GetContractInfo()

		for _, node := range Nodes {
			node.Subscribe(node)
		}
		networking.SendJson(response, map[string]interface{}{"success": true, "payload": contracts})

	}
}

func startTraining() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.SimulationStartTrainingRequest
		json.NewDecoder(request.Body).Decode(&requestBody)
		node := Nodes[requestBody.BlockchainAddress]
		bodyData, err := json.Marshal(requestBody)
		if err != nil {
			log.Fatal(err)
		}

		networking.MakePostRequestToServer(node.ExternalServer, "/node/training/start", bodyData)
		networking.SendJson(response, map[string]interface{}{"success": true})
	}
}

func transferInitialSupply() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		networking.PostValidator(request, response)
		var requestBody utils.SimulationTransferInitialSupplyRequest
		json.NewDecoder(request.Body).Decode(&requestBody)

		for address, amount := range requestBody.TransferAmounts {
			transferToken(
				utils.RemoveHexPrefix(requestBody.PrivateKey),
				int64(amount),
				address,
				"ws://127.0.0.1:8545",
				protocol.GetContractInfo().ScatterTokenContract,
			)
		}

		for _, node := range Nodes {
			if *node.LogMode {
				balance, _ := new(big.Float).SetInt(protocol.GetScatterTokenBalance(node)).Float64()
				timestamp := float64(time.Now().UnixMilli())
				scatterlogs.CreateLogEvent(utils.LOG_EVENT_TOKEN_BALANCE, timestamp, balance, node)
			}
		}
		networking.SendJson(response, map[string]interface{}{"success": true})

	}
}
