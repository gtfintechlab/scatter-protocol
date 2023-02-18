package main

import (
	"flag"

	bootstrap "github.com/gtfintechlab/scatter-protocol/bootstrap"
	peer "github.com/gtfintechlab/scatter-protocol/peers"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
)

func main() {
	var util string
	var nodeType string

	flag.StringVar(&nodeType, "type", "", "Type of node you want to run (peer or bootstrap)")
	flag.StringVar(&util, "utils", "", "Run a utility script")
	flag.Parse()

	if nodeType == utils.NODE_BOOTSTRAP {
		var tcpPort string
		var ipv4Address string
		flag.StringVar(&tcpPort, "tcpPort", "7001",
			"The TCP port you want this node to listen on")

		flag.StringVar(&ipv4Address, "ipv4Address", "127.0.0.1",
			"The IPV4 address you want this node to listen on")

		flag.Parse()

		bootstrapNode := bootstrap.InitBootstrapNode(ipv4Address, tcpPort)
		bootstrapNode.Start(bootstrapNode)

	} else if nodeType == utils.NODE_PEER {
		var peerType string
		var extAddress string
		flag.StringVar(&peerType, "peerType", utils.PEER_REQUESTOR,
			"Type of peer node you want to run (requestor, trainer)")

		flag.StringVar(&extAddress, "extAddress", ":5002",
			"External server address (for communication with the node)")

		flag.Parse()

		peerNode := peer.InitPeerNode(peerType, extAddress)
		peerNode.Start(peerNode)
	}

	if util == utils.UTIL_GENERATE_KEYS {
		utils.GenerateKeys()
	}
}
