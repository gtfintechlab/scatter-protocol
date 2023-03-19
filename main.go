package main

import (
	"flag"
	"strconv"

	bootstrap "github.com/gtfintechlab/scatter-protocol/bootstrap"
	"github.com/gtfintechlab/scatter-protocol/cosmos"
	peer "github.com/gtfintechlab/scatter-protocol/peers"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
)

func main() {
	var util string
	var nodeType string
	var tcpPort string
	var ipv4Address string
	var peerType string
	var extAddress string
	var useMdns string
	// All nodes
	flag.StringVar(&nodeType, "type", "", "Type of node you want to run (peer, bootstrap, or celestial)")
	flag.StringVar(&util, "utils", "", "Run a utility script")

	// Bootstrap Node Options
	flag.StringVar(&tcpPort, "tcpPort", "7001",
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

	flag.Parse()

	mdnsProtocol, _ := strconv.ParseBool(useMdns)
	if nodeType == utils.NODE_BOOTSTRAP {
		bootstrapNode := bootstrap.InitBootstrapNode(ipv4Address, tcpPort)
		bootstrapNode.Start(bootstrapNode)
	} else if nodeType == utils.NODE_PEER {
		peerNode := peer.InitPeerNode(peerType, extAddress)
		peerNode.Start(peerNode, mdnsProtocol)
	} else if nodeType == utils.NODE_CELESTIAL {
		celestialNode := cosmos.InitCelestialNode()
		celestialNode.Start(celestialNode, mdnsProtocol)
	}

	if util == utils.UTIL_GENERATE_KEYS {
		utils.GenerateKeys()
	}
}
