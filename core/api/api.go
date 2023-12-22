package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gtfintechlab/scatter-protocol/core/bootstrap"
	"github.com/rs/cors"
)

func simulationServerHandlers() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/deploy-protocol", deployProtocol())
	r.HandleFunc("/transfer-initial-supply", transferInitialSupply())
	r.HandleFunc("/start-node", startNode())
	r.HandleFunc("/requestor/add-topic", addTopic())
	r.HandleFunc("/trainer/add-topic", addTopic())

	r.HandleFunc("/requestor/start-training", addTopic())
	r.HandleFunc("/stop-node", stopNode())

	return r
}

const PORT = 2000

func StartSimulationApi() {
	bootstrapNode := bootstrap.InitBootstrapNode("127.0.0.1", 7001)
	go bootstrapNode.Start(bootstrapNode)

	serverHandler := simulationServerHandlers()

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace with your Next.js app's URL
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Attach the CORS handler to your existing server handler
	handler := c.Handler(serverHandler)

	// Start your server
	go http.ListenAndServe(fmt.Sprintf(":%d", PORT), handler)
	select {}
}
