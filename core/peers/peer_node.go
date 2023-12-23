package peers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"net/http"

	networking "github.com/gtfintechlab/scatter-protocol/core/networking"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"
	protocol "github.com/gtfintechlab/scatter-protocol/core/protocol"
	utils "github.com/gtfintechlab/scatter-protocol/core/utils"
	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	network "github.com/libp2p/go-libp2p/core/network"
)

func InitPeerNode(peerType string, apiPort int, databaseUsername string,
	databasePassword string, databasePort int, blockchainAddress string,
	privateKey string, dummyLoad bool, logMode bool, workspaceId *string) *utils.PeerNode {
	// Create a new libp2p host for the new node
	node, _ := libp2p.New()
	table, _ := networking.NewDHT(context.Background(), node)
	// table, _ := dht.New(context.Background(), node)
	// The multiaddress of the bootstrap node

	table.Bootstrap(context.Background())
	log.Println("Peer Node:", node.ID())

	ps, _ := pubsub.NewGossipSub(context.Background(), node)
	var database *sql.DB
	if databasePort != 0 {
		database = peerDatabase.ConnectToPostgres(peerType, databaseUsername, databasePassword, databasePort)
	} else {
		database = nil
	}
	jq := utils.NewJobProcessor(1)
	jq.StartWorkers(1)

	peerNode := utils.PeerNode{
		PeerType:          peerType,
		Start:             StartPeer,
		NodeId:            node.ID(),
		BlockchainAddress: &blockchainAddress,
		PrivateKey:        &privateKey,
		ExternalServer: &http.Server{
			Addr: fmt.Sprintf(":%d", apiPort),
		},
		PeerToPeerServer:     &node,
		DistributedHashTable: table,
		PubSubService:        ps,
		DataStore:            database,
		PubSubTopics:         &map[string]*pubsub.Topic{},
		DatastoreLock:        &sync.Mutex{},
		TrainingLock:         &map[string]map[string]bool{},
		JobQueue:             jq,
		DummyLoad:            &dummyLoad,
		LogMode:              &logMode,
		WorkspaceId:          workspaceId,
		Subscriptions:        &utils.SubscriptionManager{},
		Subscribe:            Subscribe,
	}

	protocol.SetNodeId(&peerNode, node.ID().String())

	go protocol.TrainingEventListener(&peerNode)
	go protocol.EvaluationRequestListener(&peerNode)
	go protocol.ModelValidationListener(&peerNode)
	go protocol.DebugEventListener(&peerNode)
	go protocol.JobCompleteEventListener(&peerNode)

	if protocol.GetRoleByAddress(&peerNode, *peerNode.BlockchainAddress) == utils.PEER_NO_ROLE {
		switch peerType {
		case utils.PEER_REQUESTOR:
			protocol.InitRequestorNode(&peerNode)
		case utils.PEER_TRAINER:
			protocol.InitTrainerNode(&peerNode)
		}
	}

	node.SetStreamHandler(utils.PROTOCOL_IDENTIFIER, peerStreamHandler(&peerNode))

	return &peerNode
}

func Subscribe(node *utils.PeerNode) {
	go protocol.TrainingEventListener(node)
	go protocol.EvaluationRequestListener(node)
	go protocol.ModelValidationListener(node)
	go protocol.DebugEventListener(node)
	go protocol.JobCompleteEventListener(node)
}

func StartPeer(node *utils.PeerNode, useMdns bool) {
	serverMux := externalServerHandlers(node)

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace with your Next.js app's URL
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	// Attach the CORS handler to your existing server handler
	handler := c.Handler(serverMux)

	if useMdns {
		utils.InitializePeerDiscovery(node.PeerToPeerServer)
	} else {
		go networking.InitializePeerDiscoveryDHT(
			context.Background(),
			node.PeerToPeerServer,
			node.DistributedHashTable,
			string(utils.PROTOCOL_IDENTIFIER),
		)
	}

	go http.ListenAndServe(node.ExternalServer.Addr, handler)
	select {}
}

func StopPeer(node *utils.PeerNode) {
	node.ExternalServer.Close()
}

func peerStreamHandler(node *utils.PeerNode) network.StreamHandler {
	// "Public" method handlers for peers to communicate with your node
	return func(stream network.Stream) {
		message := networking.DecodeMessage(&stream)
		switch messageType := message.MessageType; messageType {
		}
	}
}

// "Private" method handlers for you to communicate with your own node
func externalServerHandlers(node *utils.PeerNode) http.Handler {
	serverMux := mux.NewRouter()
	serverMux.HandleFunc("/node/health", health(node))
	serverMux.HandleFunc("/node/role/switch", switchPeerNodeRole(node))
	serverMux.HandleFunc("/node/topic/add", addTopic(node))
	serverMux.HandleFunc("/node/topics/get", getOwnTopics(node))
	serverMux.HandleFunc("/node/topics/published", getPublishedTopics(node))
	serverMux.HandleFunc("/node/topics/trainers", getTopicTrainers(node))
	serverMux.HandleFunc("/node/training/start", initializeTraining(node))
	serverMux.HandleFunc("/node/token/balance", getScatterTokenBalance(node))

	return serverMux
}
