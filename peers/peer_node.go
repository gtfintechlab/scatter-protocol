package peers

import (
	"context"
	"fmt"
	"log"

	"net/http"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	network "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func InitPeerNode(peerType string, serverAddress string) *utils.PeerNode {
	// Create a new libp2p host for the new node
	node, _ := libp2p.New()
	table, _ := dht.New(context.Background(), node)
	// The multiaddress of the bootstrap node
	bootstrapAddr, _ := multiaddr.NewMultiaddr(utils.BOOTSTRAP_NODE_MULTIADDR)
	peerInfo, _ := peer.AddrInfoFromP2pAddr(bootstrapAddr)
	err := node.Connect(context.Background(), *peerInfo)

	if err != nil {
		log.Fatal(err)
	}
	table.Bootstrap(context.Background())
	fmt.Println("Peer Node:", node.ID())

	stream, _ := node.NewStream(context.Background(),
		peerInfo.ID, utils.PROTOCOL_IDENTIFIER)
	networking.SendMessage(&stream, networking.MESSAGE_JOIN_NETWORK)

	peerNode := utils.PeerNode{
		PeerType: peerType,
		Start:    StartPeer,
		ExternalServer: &http.Server{
			Addr: serverAddress,
		},
		PeerToPeerServer: &node,
		Topics:           map[string]string{},
	}

	return &peerNode
}

func StartPeer(node *utils.PeerNode) {
	externalServerHandlers(node)
	go node.ExternalServer.ListenAndServe()

	select {}
}

// "Public" method handlers for peers to communicate with your node
func networkStreamHandler(stream network.Stream) {
	message := networking.DecodeMessage(&stream)
	switch messageType := message.MessageType; messageType {
	case utils.PEER_SWITCH_ROLE:

	}
}

// "Private" method handlers for you to communicate with your own node
func externalServerHandlers(node *utils.PeerNode) {
	http.HandleFunc("/node/role/switch", switchPeerNodeRole(node))
	http.HandleFunc("/node/topic/add", addTopic(node))
}
