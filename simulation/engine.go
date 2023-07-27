package simulation

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gtfintechlab/scatter-protocol/bootstrap"
	"github.com/gtfintechlab/scatter-protocol/cosmos"
	"github.com/gtfintechlab/scatter-protocol/peers"
	"github.com/gtfintechlab/scatter-protocol/utils"
)

func RunSimulation(simulationName string) {
	var simulationConfiguration utils.SimulationConfiguration
	jsonData, err := os.ReadFile(fmt.Sprintf("simulations/%s", simulationName))
	if err != nil {
		fmt.Println(`An error occurred when reading the simulation. 
			Make sure you define the simulation configuration in /simulation/simulations`, err)
		return
	}

	err = json.Unmarshal(jsonData, &simulationConfiguration)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	nodeConfig := initializeAllNodes(simulationConfiguration.Nodes)
	executeSteps(nodeConfig, simulationConfiguration.Steps)
}

func initializeAllNodes(nodeList []utils.NodeConfig) map[string]interface{} {
	nodeConfig := map[string]interface{}{}
	for _, node := range nodeList {
		var configNode any

		switch nodeType := node.Type; nodeType {
		case utils.NODE_BOOTSTRAP:
			createdNode := bootstrap.InitBootstrapNode(*node.Ipv4Address, *node.TcpPort)
			bootstrap.StartBootstrap(createdNode)
			configNode = createdNode
		case utils.PEER_REQUESTOR, utils.PEER_TRAINER:
			createdNode := peers.InitPeerNode(node.Type, *node.ExtAddress)
			createdNode.Start(createdNode, *node.UseMdns) // true = use mdns
			configNode = createdNode
		case utils.NODE_CELESTIAL:
			createdNode := cosmos.InitCelestialNode()
			createdNode.Start(createdNode, *node.UseMdns) // true = use mdns
			configNode = createdNode
		}

		nodeConfig[node.Id] = map[string]interface{}{
			"Node":  configNode,
			"State": map[string]interface{}{},
		}

	}
	return nodeConfig
}

func executeSteps(nodeConfig map[string]interface{}, stepList []map[string]interface{}) {
}
