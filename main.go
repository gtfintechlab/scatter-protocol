package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	bootstrap "github.com/gtfintechlab/scatter-protocol/bootstrap"
	"github.com/gtfintechlab/scatter-protocol/cosmos"
	celestialDatabase "github.com/gtfintechlab/scatter-protocol/cosmos/db"
	peer "github.com/gtfintechlab/scatter-protocol/peers"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/peers/db"
	"github.com/gtfintechlab/scatter-protocol/simulation"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
)

func main() {
	var util string
	var nodeType string
	var tcpPort int
	var ipv4Address string
	var peerType string
	var extAddress string
	var useMdns string
	var migrationDirection string
	var simulationName string
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
	flag.StringVar(&extAddress, "extAddress", ":5002",
		"External server address (for communication with the node)")

	flag.StringVar(&useMdns, "useMdns", "true",
		"Whether or not to connect to the bootstrap node")
	flag.StringVar(&migrationDirection, "direction", "up",
		"Migration direction (up/down)")

	flag.StringVar(&simulationName, "simulation", "train_simple_model",
		"Name of the simulation to run")

	flag.Parse()

	mdnsProtocol, _ := strconv.ParseBool(useMdns)
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
		peerNode := peer.InitPeerNode(peerType, extAddress, "postgres", "postgres", dbPort)
		peerNode.Start(peerNode, mdnsProtocol)
	} else if nodeType == utils.NODE_CELESTIAL {
		celestialNode := cosmos.InitCelestialNode("postgres", "postgres", 5432)
		celestialNode.Start(celestialNode, mdnsProtocol)
	}

	if util == utils.UTIL_GENERATE_KEYS {
		utils.GenerateKeys()
	} else if util == utils.UTIL_DEBUG_MODE {
		fmt.Println("Debug Mode")
		os.MkdirAll("training/trainer/jobs/test2/test2", os.ModePerm)
	} else if util == utils.UTIL_CELESTIAL_DATABASE_MIGRATION {
		celestialDatabase.MigrateCelestialDB(migrationDirection, "postgres", "postgres", 5432)
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
	}

}
