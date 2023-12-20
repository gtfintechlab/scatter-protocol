package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	api "github.com/gtfintechlab/scatter-protocol/core/api"
	bootstrap "github.com/gtfintechlab/scatter-protocol/core/bootstrap"
	peer "github.com/gtfintechlab/scatter-protocol/core/peers"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"
	"github.com/gtfintechlab/scatter-protocol/core/simulation"
	utils "github.com/gtfintechlab/scatter-protocol/core/utils"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var _ = godotenv.Load(".env")

	var util string
	var nodeType string
	var tcpPort int
	var ipv4Address string
	var peerType string
	var apiPort int
	var useMdns string
	var dummyLoad string
	var migrationDirection string
	var simulationName string
	var contractName string

	// All nodes
	flag.StringVar(&nodeType, "type", "", "Type of node you want to run (peer, bootstrap, or celestial)")
	flag.StringVar(&util, "utils", "", "Run a utility script")

	// Bootstrap Node Options
	flag.IntVar(&tcpPort, "tcpPort", 7001,
		"The TCP port you want this node to listen on")
	flag.StringVar(&ipv4Address, "ipv4Address", "127.0.0.1",
		"The IPV4 address you want this node to listen on")

	// Peer Node Options
	flag.StringVar(&peerType, "peerType", utils.PEER_REQUESTOR,
		"Type of peer node you want to run (requestor, trainer)")
	flag.IntVar(&apiPort, "apiPort", 5002,
		"External server address (for communication with the node)")

	flag.StringVar(&useMdns, "useMdns", "true",
		"Whether or not to connect to the bootstrap node")
	flag.StringVar(&migrationDirection, "direction", "up",
		"Migration direction (up/down)")

	flag.StringVar(&simulationName, "simulation", "benevolent_trainer",
		"Name of the simulation to run")
	flag.StringVar(&contractName, "contract", "all",
		"The smart contract you want to deploy")

	flag.StringVar(&dummyLoad, "dummyLoad", "false",
		"Whether to run training / eval procedures or use dummy data")

	flag.Parse()

	mdnsProtocol, _ := strconv.ParseBool(useMdns)
	dummyLoadBool, _ := strconv.ParseBool(dummyLoad)
	if nodeType == utils.NODE_BOOTSTRAP {
		bootstrapNode := bootstrap.InitBootstrapNode(ipv4Address, tcpPort)
		bootstrapNode.Start(bootstrapNode)
	} else if nodeType == utils.NODE_PEER {
		var dbPort int
		if peerType == utils.PEER_REQUESTOR {
			dbPort = 8701
		} else {
			dbPort = 8702
		}
		peerNode := peer.InitPeerNode(peerType, apiPort, "postgres", "postgres", dbPort,
			os.Getenv("BLOCKCHAIN_ADDRESS"), os.Getenv("PRIVATE_KEY"), dummyLoadBool)
		peerNode.Start(peerNode, mdnsProtocol)
	}

	if util == utils.UTIL_GENERATE_KEYS {
		utils.GenerateKeys()
	} else if util == utils.UTIL_DEBUG_MODE {
		log.Println("Debug Mode")
	} else if util == utils.UTIL_PEER_DATABASE_MIGRATION {
		var dbPort int
		if peerType == utils.PEER_REQUESTOR {
			dbPort = 8701
		} else {
			dbPort = 8702
		}
		peerDatabase.MigratePeerDB(migrationDirection, peerType, "postgres", "postgres", dbPort)
	} else if util == utils.UTIL_RUN_SIMULATION {
		simulation.RunSimulation(simulationName)
	} else if util == utils.UTIL_SIMULATION_API {
		api.StartSimulationApi()
	}
}
