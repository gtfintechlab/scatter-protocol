package simulation

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gtfintechlab/scatter-protocol/bootstrap"
	"github.com/gtfintechlab/scatter-protocol/peers"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/peers/db"
	"github.com/gtfintechlab/scatter-protocol/protocol"
	scattertoken "github.com/gtfintechlab/scatter-protocol/protocol/scatter-token"
	"github.com/gtfintechlab/scatter-protocol/utils"
)

func clearDatabase(datastore *sql.DB) {
	_, err := datastore.Query(`DROP SCHEMA public CASCADE;`)

	if err != nil {
		log.Fatal("Failed to clear database")
	}

	_, err = datastore.Query(`CREATE SCHEMA public;`)
	if err != nil {
		log.Fatal("Failed to restore public schema")
	}

}

func RunSimulation(simulationName string) {
	var simulationConfiguration utils.SimulationConfiguration
	jsonData, err := os.ReadFile(fmt.Sprintf("simulation/simulations/%s.json", simulationName))
	if err != nil {
		log.Fatal(`An error occurred when reading the simulation. 
			Make sure you define the simulation configuration in /simulation/simulations`, err)
		return
	}

	err = json.Unmarshal(jsonData, &simulationConfiguration)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
		return
	}
	nodeConfig := initializeAllNodes(simulationConfiguration.Nodes, simulationConfiguration.Environment)
	executeSteps(nodeConfig, simulationConfiguration.Steps)
}

func initializeAllNodes(nodeList []utils.NodeConfig, environment utils.EnvironmentConfig) map[string]utils.SimulationNodeConfig {
	nodeConfig := make(map[string]utils.SimulationNodeConfig)
	for _, node := range nodeList {
		var simulationNode utils.SimulationNode

		switch nodeType := node.Type; nodeType {
		case utils.NODE_BOOTSTRAP:
			createdNode := bootstrap.InitBootstrapNode(*node.Ipv4Address, *node.TcpPort)
			go createdNode.Start(createdNode)
			simulationNode = utils.SimulationNode{BootstrapNode: createdNode}
		case utils.PEER_REQUESTOR, utils.PEER_TRAINER:
			createdNode := peers.InitPeerNode(
				node.Type,
				*node.ExtAddress,
				*node.DatastoreUsername,
				*node.DatastorePassword,
				*node.DatastorePort,
				*node.BlockchainAddress,
				*node.PrivateKey,
			)
			transferToken(
				*environment.ProtocolOwnerPrivateKey,
				int64(*node.InitialScatterTokenSupply),
				*node.BlockchainAddress,
				*environment.EthereumNode,
				utils.SCATTER_TOKEN_CONTRACT,
			)
			clearDatabase(createdNode.DataStore)
			peerDatabase.MigratePeerDB("up",
				node.Type,
				*node.DatastoreUsername,
				*node.DatastorePassword,
				*node.DatastorePort,
			)

			if nodeType == utils.PEER_REQUESTOR {
				protocol.InitRequestorNode(createdNode)
			} else {
				protocol.InitTrainerNode(createdNode)
			}
			go createdNode.Start(createdNode, *node.UseMdns)
			simulationNode = utils.SimulationNode{PeerNode: createdNode}
		}

		nodeConfig[node.Id] = utils.SimulationNodeConfig{
			Node:  simulationNode,
			Type:  node.Type,
			State: map[string]interface{}{},
		}
	}
	time.Sleep(time.Second * 10)
	return nodeConfig
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func executeSteps(nodeConfig map[string]utils.SimulationNodeConfig, stepList []utils.StepConfig) {
	for _, step := range stepList {
		log.Println(Green + step.Description + Reset)
		method := strings.Split((*step.Action), " ")[0]
		endpoint := strings.Split((*step.Action), " ")[1]
		var url string

		// Useless switch statement for (potentially) later extensibility
		switch nodeConfig[step.NodeId].Type {
		case utils.PEER_TRAINER, utils.PEER_REQUESTOR:
			node := nodeConfig[step.NodeId].Node.PeerNode
			url = fmt.Sprintf("http://127.0.0.1%s", node.ExternalServer.Addr)
		}

		var response map[string]interface{}
		var err error
		if step.Body != nil {
			response, err = simpleRequest(endpoint, method, url, parseBody(step.Body, nodeConfig[step.NodeId].State))
		} else {
			response, err = simpleRequest(endpoint, method, url, nil)
		}

		if err != nil {
			log.Fatal(err)
		}

		if step.StateKey != nil {
			nodeConfig[step.NodeId].State[*step.StateKey] = response
		}

		// Let other steps finish first
		time.Sleep(time.Second * 10)
	}
	select {}
}

func parseBody(body *map[string]interface{}, nodeState map[string]interface{}) *map[string]interface{} {
	if body == nil {
		return nil
	}
	parsedBody := map[string]interface{}{}
	useStateRegex := regexp.MustCompile(`^{.*}$`)

	for key, value := range *body {
		assertedValue := value.(string)
		if useStateRegex.MatchString(assertedValue) {
			// Remove all the curly braces so we can start processing this
			assertedValue = strings.Replace(assertedValue, "{", "", -1)
			assertedValue = strings.Replace(assertedValue, "}", "", -1)

			// Split the expressions into parts
			parts := strings.Split(assertedValue, ".")

			var currentExpression any
			indexingRegex := regexp.MustCompile(`\[.*?]`)

			// Iterate through the parts and process it seperately
			for index, part := range parts {
				processedPart := part
				var indexedPart *int = nil
				if indexingRegex.MatchString(part) {
					indexedPartString := strings.Split(processedPart, "[")[1]
					processedPart = strings.Split(processedPart, "[")[0]

					indexedPartInt, _ := strconv.Atoi(strings.Replace(indexedPartString, "]", "", -1))
					indexedPart = &indexedPartInt
				}

				if currentExpression == nil && index != 0 {
					break
				}

				if index == 0 {
					currentExpression = nodeState[processedPart].(map[string]interface{})
					continue
				}

				if _, ok := currentExpression.(map[string]any)[processedPart]; !ok {
					break
				}

				if indexedPart != nil {
					currentExpression = currentExpression.(map[string]any)[processedPart].([]any)[*indexedPart]
				} else {
					currentExpression = currentExpression.(map[string]any)[processedPart]
				}
			}

			parsedBody[key] = currentExpression
		} else {
			parsedBody[key] = value
		}
	}

	return &parsedBody
}
func simpleRequest(endpoint string, method string, url string, body *map[string]interface{}) (map[string]interface{}, error) {
	client := &http.Client{}

	// Encode the request body as JSON if provided
	var requestBody []byte
	if body != nil {
		requestBodyJSON, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		requestBody = requestBodyJSON
	}

	req, err := http.NewRequest(method, url+endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(responseBody) == 0 {
		return map[string]interface{}{}, nil
	}

	// Create an empty map to hold the JSON data
	jsonData := make(map[string]interface{})

	// Unmarshal the JSON data into the map
	err = json.Unmarshal(responseBody, &jsonData)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}
	return jsonData, nil
}

func transferToken(privateKey string, amount int64, recipient string, ethereumNode string, contractAddress string) {
	auth := getTransactor(privateKey)
	var ethereumClient, _ = ethclient.Dial(ethereumNode)
	contract, _ := scattertoken.NewScattertoken(common.HexToAddress(contractAddress), ethereumClient)
	contract.Transfer(auth, common.HexToAddress(recipient), big.NewInt(amount))
}

func getTransactor(ownerPrivateKey string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(ownerPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, protocol.CHAIN)
	auth.Value = big.NewInt(0)

	return auth
}
