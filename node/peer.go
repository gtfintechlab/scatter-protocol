package node

import (
	"context"
	"fmt"
	"log"

	"github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	peer "github.com/libp2p/go-libp2p/core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func InitPeerNode() {
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

	select {}
}
