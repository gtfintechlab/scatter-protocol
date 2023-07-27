package peers

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"

	"net/http"

	"github.com/gtfintechlab/scatter-protocol/cosmos"
	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	network "github.com/libp2p/go-libp2p/core/network"
)

func InitPeerNode(peerType string, serverAddress string) *utils.PeerNode {
	// Create a new libp2p host for the new node
	node, _ := libp2p.New()
	table, _ := networking.NewDHT(context.Background(), node)
	// table, _ := dht.New(context.Background(), node)
	// The multiaddress of the bootstrap node

	table.Bootstrap(context.Background())
	fmt.Println("Peer Node:", node.ID())

	ps, _ := pubsub.NewGossipSub(context.Background(), node)

	database := connectToPostgres(peerType)
	peerNode := utils.PeerNode{
		PeerType: peerType,
		Start:    StartPeer,
		NodeId:   node.ID(),
		ExternalServer: &http.Server{
			Addr: serverAddress,
		},
		PeerToPeerServer:     &node,
		DistributedHashTable: table,
		PubSubService:        ps,
		DataStore:            database,
		PubSubTopics:         &map[string]*pubsub.Topic{},
		InformationBox: &utils.InformationBox{
			CosmosTopics: &map[string]interface{}{},
		},
	}
	getInitialTopics(utils.DATA_DIRECTORY, &peerNode)

	node.SetStreamHandler(utils.PROTOCOL_IDENTIFIER, peerStreamHandler(&peerNode))

	return &peerNode
}

func StartPeer(node *utils.PeerNode, useMdns bool) {
	externalServerHandlers(node)
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
	universalCosmosTopic := cosmos.JoinCosmos(context.Background(), node, utils.UNIVERSAL_COSMOS)
	(*node.PubSubTopics)[utils.UNIVERSAL_COSMOS] = universalCosmosTopic

	go node.ExternalServer.ListenAndServe()
	select {}
}

func peerStreamHandler(node *utils.PeerNode) network.StreamHandler {
	// "Public" method handlers for peers to communicate with your node
	return func(stream network.Stream) {
		message := networking.DecodeMessage(&stream)
		switch messageType := message.MessageType; messageType {
		case utils.PEER_GET_TOPICS:
			go peerGetTopicsHandler(node, &message)
		case utils.PEER_START_TRAINING:
			go peerStartTrainingHandler(node, &message, &stream)
		}
	}
}

// "Private" method handlers for you to communicate with your own node
func externalServerHandlers(node *utils.PeerNode) {
	http.HandleFunc("/node/health", health(node))
	http.HandleFunc("/node/role/switch", switchPeerNodeRole(node))
	http.HandleFunc("/node/topic/add", addTopic(node))
	http.HandleFunc("/node/topics/get", getOwnTopics(node))
	http.HandleFunc("/cosmos/topics/get", getCosmosTopics(node))
	http.HandleFunc("/node/topics/trainers", getTopicTrainers(node))
	http.HandleFunc("/node/training/start", initializeTraining(node))
}
