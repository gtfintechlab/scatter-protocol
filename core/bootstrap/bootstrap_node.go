package bootstrap

import (
	"context"
	"fmt"
	"log"

	networking "github.com/gtfintechlab/scatter-protocol/core/networking"
	utils "github.com/gtfintechlab/scatter-protocol/core/utils"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	network "github.com/libp2p/go-libp2p/core/network"
	"github.com/multiformats/go-multiaddr"
)

func InitBootstrapNode(ipv4Address string, tcpPort int) *utils.BootstrapNode {
	privateKey, _ := utils.LoadKeys()

	node, _ := libp2p.New(libp2p.Identity(privateKey),
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/%s/tcp/%d", ipv4Address, tcpPort)))
	// Create a new DHT
	distributedHashTable, _ := dht.New(context.Background(), node)
	distributedHashTable.Bootstrap(context.Background())
	log.Println("Bootstrap Node:", node.ID())

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
	// go networking.InitializePeerDiscoveryDHT(
	// 	context.Background(),
	// 	*node.PeerToPeerServer,
	// 	node.DistributedHashTable,
	// 	string(utils.PROTOCOL_IDENTIFIER),
	// )
	select {}

}
func bootstrapStreamHandler(stream network.Stream) {
	message := networking.DecodeMessage(&stream)
	switch messageType := message.MessageType; messageType {
	case utils.MESSAGE_CODE_JOIN_NETWORK:
		log.Println("Peer connected:", stream.Conn().RemotePeer())
	}
}
