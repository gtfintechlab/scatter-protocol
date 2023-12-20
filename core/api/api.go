package api

import (
	"fmt"
	"net/http"
)

func simulationServerHandlers() *http.ServeMux {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/start-node", startNode())

	return serverMux
}

const PORT = 2000

func StartSimulationApi() {
	serverMux := simulationServerHandlers()
	go http.ListenAndServe(fmt.Sprintf(":%d", PORT), serverMux)
	select {}
}
