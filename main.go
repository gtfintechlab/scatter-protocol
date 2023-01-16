package main

import (
	"flag"

	node "github.com/gtfintechlab/scatter-protocol/node"
	"github.com/gtfintechlab/scatter-protocol/utils"
)

func main() {
	var nodeType string
	flag.StringVar(&nodeType, "type", utils.NODE_PEER, "Type of node you want to run (peer or bootstrap)")
	flag.Parse()

	if nodeType == utils.NODE_BOOTSTRAP {
		node.InitBootstrapNode()
	} else if nodeType == utils.NODE_PEER {
		node.InitPeerNode()
	}
}
