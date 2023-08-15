package peers

import (
	"context"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"

	"net/http"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	network "github.com/libp2p/go-libp2p/core/network"
)

func InitPeerNode(peerType string, serverAddress string, databaseUsername string, databasePassword string, databasePort int) *utils.PeerNode {
	// Create a new libp2p host for the new node
	node, _ := libp2p.New()
	table, _ := networking.NewDHT(context.Background(), node)
	// table, _ := dht.New(context.Background(), node)
	// The multiaddress of the bootstrap node

	table.Bootstrap(context.Background())
	log.Println("Peer Node:", node.ID())

	ps, _ := pubsub.NewGossipSub(context.Background(), node)
	blockchainAddress := os.Getenv("BLOCKCHAIN_ADDRESS")
	database := connectToPostgres(peerType, databaseUsername, databasePassword, databasePort)
	peerNode := utils.PeerNode{
		PeerType:          peerType,
		Start:             StartPeer,
		NodeId:            node.ID(),
		BlockchainAddress: &blockchainAddress,
		ExternalServer: &http.Server{
			Addr: serverAddress,
		},
		PeerToPeerServer:     &node,
		DistributedHashTable: table,
		PubSubService:        ps,
		DataStore:            database,
		PubSubTopics:         &map[string]*pubsub.Topic{},
		DatastoreLock:        &sync.Mutex{},
	}
	utils.SetNodeId(node.ID().String())
	getInitialTopics(utils.DATA_DIRECTORY, &peerNode)

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
		case utils.PEER_START_TRAINING:
			go peerStartTrainingHandler(node, &message, &stream)
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
	return serverMux
}
