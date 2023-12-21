package api

import (
	"fmt"
	"net/http"

	"github.com/gtfintechlab/scatter-protocol/core/bootstrap"
)

func simulationServerHandlers() *http.ServeMux {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/start-node", startNode())
	serverMux.HandleFunc("/stop-node", stopNode())

	return serverMux
}

const PORT = 2000

func StartSimulationApi() {
	bootstrapNode := bootstrap.InitBootstrapNode("127.0.0.1", 7001)
	bootstrapNode.Start(bootstrapNode)

	serverMux := simulationServerHandlers()
	go http.ListenAndServe(fmt.Sprintf(":%d", PORT), serverMux)
	select {}
}
