package peers

import (
	"context"
	"log"
	"sync"

	_ "github.com/lib/pq"

	"net/http"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/peers/db"
	protocol "github.com/gtfintechlab/scatter-protocol/protocol"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	network "github.com/libp2p/go-libp2p/core/network"
)

func InitPeerNode(peerType string, serverAddress string, databaseUsername string,
	databasePassword string, databasePort int, blockchainAddress string,
	privateKey string, dummyLoad bool) *utils.PeerNode {
	// Create a new libp2p host for the new node
	node, _ := libp2p.New()
	table, _ := networking.NewDHT(context.Background(), node)
	// table, _ := dht.New(context.Background(), node)
	// The multiaddress of the bootstrap node

	table.Bootstrap(context.Background())
	log.Println("Peer Node:", node.ID())

	ps, _ := pubsub.NewGossipSub(context.Background(), node)
	database := peerDatabase.ConnectToPostgres(peerType, databaseUsername, databasePassword, databasePort)
	jq := utils.NewJobProcessor(1)
	jq.StartWorkers(1)
	peerNode := utils.PeerNode{
		PeerType:          peerType,
		Start:             StartPeer,
		NodeId:            node.ID(),
		BlockchainAddress: &blockchainAddress,
		PrivateKey:        &privateKey,
		ExternalServer: &http.Server{
			Addr: serverAddress,
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
	}
	protocol.SetNodeId(&peerNode, node.ID().String())

	go protocol.TrainingEventListener(&peerNode)
	go protocol.EvaluationRequestListener(&peerNode)
	go protocol.ModelValidationListener(&peerNode)
	go protocol.DebugEventListener(&peerNode)

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

func StartPeer(node *utils.PeerNode, useMdns bool) {
	serverMux := externalServerHandlers(node)
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

	go http.ListenAndServe(node.ExternalServer.Addr, serverMux)
	select {}
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
func externalServerHandlers(node *utils.PeerNode) *http.ServeMux {
	serverMux := http.NewServeMux()
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
