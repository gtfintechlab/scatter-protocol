package peers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/gtfintechlab/scatter-protocol/cosmos"
	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	network "github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func InitPeerNode(peerType string, serverAddress string, useBoostrap bool) *utils.PeerNode {
	// Create a new libp2p host for the new node
	node, _ := libp2p.New()
	table, _ := dht.New(context.Background(), node)
	// The multiaddress of the bootstrap node

	if useBoostrap {
		bootstrapAddr, _ := multiaddr.NewMultiaddr(utils.BOOTSTRAP_NODE_MULTIADDR)
		peerInfo, _ := peer.AddrInfoFromP2pAddr(bootstrapAddr)
		err := node.Connect(context.Background(), *peerInfo)

		if err != nil {
			log.Fatal(err)
		}

		stream, _ := node.NewStream(context.Background(),
			peerInfo.ID, utils.PROTOCOL_IDENTIFIER)
		networking.SendMessage(&stream, networking.MESSAGE_JOIN_NETWORK)

	}
	table.Bootstrap(context.Background())
	fmt.Println("Peer Node:", node.ID())

	ps, _ := pubsub.NewGossipSub(context.Background(), node)

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
		TopicToDataPath:      &map[string]string{},
		TopicTrainerMap:      map[string][]string{},
		PubSubTopics:         &map[string]*pubsub.Topic{},
		InformationBox: &utils.InformationBox{
			CosmosTopics: &map[string]map[string]bool{},
		},
	}

	node.SetStreamHandler(utils.PROTOCOL_IDENTIFIER, peerStreamHandler(&peerNode))

	return &peerNode
}

func StartPeer(node *utils.PeerNode) {
	externalServerHandlers(node)
	utils.InitializePeerDiscovery(node.PeerToPeerServer)

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
			var topicList map[string]map[string]bool
			jsonData, _ := json.Marshal(message.Payload)
			json.Unmarshal(jsonData, &topicList)
			node.InformationBox.CosmosTopics = &topicList
		}
	}
}

// "Private" method handlers for you to communicate with your own node
func externalServerHandlers(node *utils.PeerNode) {
	http.HandleFunc("/node/health", health())
	http.HandleFunc("/node/role/switch", switchPeerNodeRole(node))
	http.HandleFunc("/node/topic/add", addTopic(node))
	http.HandleFunc("/node/topics/get", getOwnTopics(node))
	http.HandleFunc("/cosmos/topics/get", getCosmosTopics(node))
}
