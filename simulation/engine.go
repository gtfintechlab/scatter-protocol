package simulation

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gtfintechlab/scatter-protocol/bootstrap"
	"github.com/gtfintechlab/scatter-protocol/cosmos"
	celestialDatabase "github.com/gtfintechlab/scatter-protocol/cosmos/db"
	"github.com/gtfintechlab/scatter-protocol/peers"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/peers/db"
	"github.com/gtfintechlab/scatter-protocol/utils"
)

func cleanUp() {
	cmd := exec.Command("docker", "rm", "-vf", "$(docker ps -a -q)")
	cmd.Output()
	cmd = exec.Command("docker", "rmi", "-f", "$(docker images -a -q)")
	cmd.Output()
}

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
	cleanUp()
	var simulationConfiguration utils.SimulationConfiguration
	jsonData, err := os.ReadFile(fmt.Sprintf("simulation/simulations/%s.json", simulationName))
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
			go createdNode.Start(createdNode)
			configNode = createdNode
		case utils.PEER_REQUESTOR, utils.PEER_TRAINER:
			createdNode := peers.InitPeerNode(
				node.Type,
				*node.ExtAddress,
				*node.DatastoreUsername,
				*node.DatastorePassword,
				*node.DatastorePort,
			)
			clearDatabase(createdNode.DataStore)
			peerDatabase.MigratePeerDB("up",
				node.Type,
				*node.DatastoreUsername,
				*node.DatastorePassword,
				*node.DatastorePort,
			)

			go createdNode.Start(createdNode, *node.UseMdns)
			configNode = createdNode
		case utils.NODE_CELESTIAL:
			createdNode := cosmos.InitCelestialNode(
				*node.DatastoreUsername,
				*node.DatastorePassword,
				*node.DatastorePort,
			)
			clearDatabase(createdNode.DataStore)
			celestialDatabase.MigrateCelestialDB("up",
				*node.DatastoreUsername,
				*node.DatastorePassword,
				*node.DatastorePort,
			)
			go createdNode.Start(createdNode, *node.UseMdns)
			configNode = createdNode
		}

		nodeConfig[node.Id] = map[string]interface{}{
			"Node":  configNode,
			"State": map[string]interface{}{},
		}
	}
	select {}
	return nodeConfig
}

func executeSteps(nodeConfig map[string]interface{}, stepList []map[string]interface{}) {
}
