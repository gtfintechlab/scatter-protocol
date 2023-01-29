package main

import (
	"flag"

	node "github.com/gtfintechlab/scatter-protocol/node"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
)

func main() {
	var util string
	var nodeType string
	flag.StringVar(&nodeType, "type", "", "Type of node you want to run (peer or bootstrap)")
	flag.StringVar(&util, "utils", "", "Run a utility script")
	flag.Parse()

	if nodeType == utils.NODE_BOOTSTRAP {
		node.InitBootstrapNode()
	} else if nodeType == utils.NODE_PEER {
		node.InitPeerNode()
	}

	if util == utils.UTIL_GENERATE_KEYS {
		utils.GenerateKeys()
	}
}
