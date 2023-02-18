package bootstrap

import (
	"context"
	"fmt"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	network "github.com/libp2p/go-libp2p/core/network"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func InitBootstrapNode() {
	privateKey, _ := utils.LoadKeys()

	node, _ := libp2p.New(libp2p.Identity(privateKey),
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/7001"))
	// Create a new DHT
	dht.New(context.Background(), node)
	fmt.Println("Bootstrap Node:", node.ID())

	// Listen for incoming connections
	addr, _ := multiaddr.NewMultiaddr(utils.BOOTSTRAP_NODE_MULTIADDR)
	node.SetStreamHandler(utils.PROTOCOL_IDENTIFIER, bootstrapStreamHandler)
	node.Network().Listen(addr)
	select {}
}

func bootstrapStreamHandler(stream network.Stream) {
	message := networking.DecodeMessage(&stream)
	switch messageType := message.MessageType; messageType {
	case utils.MESSAGE_CODE_JOIN_NETWORK:
		fmt.Println("Peer connected:", stream.Conn().RemotePeer())
	}
}
