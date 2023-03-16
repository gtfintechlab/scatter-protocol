package bootstrap

import (
	"context"
	"fmt"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	network "github.com/libp2p/go-libp2p/core/network"
	"github.com/multiformats/go-multiaddr"
)

func InitBootstrapNode(ipv4Address string, tcpPort string) *utils.BootstrapNode {
	privateKey, _ := utils.LoadKeys()

	node, _ := libp2p.New(libp2p.Identity(privateKey),
		libp2p.ListenAddrStrings("/ip4/"+ipv4Address+"/tcp/"+tcpPort))
	// Create a new DHT
	distributedHashTable, _ := dht.New(context.Background(), node)
	fmt.Println("Bootstrap Node:", node.ID())

	// Set stream handler
	node.SetStreamHandler(utils.PROTOCOL_IDENTIFIER, bootstrapStreamHandler)

	bootstrapNode := utils.BootstrapNode{
		Start:                StartBootstrap,
		NodeId:               node.ID(),
		PeerToPeerServer:     &node,
		DistributedHashTable: distributedHashTable,
	}
	return &bootstrapNode
}

func StartBootstrap(node *utils.BootstrapNode) {
	addr, _ := multiaddr.NewMultiaddr(utils.BOOTSTRAP_NODE_MULTIADDR)
	go (*node.PeerToPeerServer).Network().Listen(addr)
	utils.InitializePeerDiscovery(node.PeerToPeerServer)
	select {}

}
func bootstrapStreamHandler(stream network.Stream) {
	message := networking.DecodeMessage(&stream)
	switch messageType := message.MessageType; messageType {
	case utils.MESSAGE_CODE_JOIN_NETWORK:
		fmt.Println("Peer connected:", stream.Conn().RemotePeer())
	}
}
