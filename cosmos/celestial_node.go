package cosmos

import (
	"context"
	"fmt"
	"log"

	"github.com/gtfintechlab/scatter-protocol/utils"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func InitCelestialNode() *utils.CelestialNode {
	node, _ := libp2p.New()
	table, _ := dht.New(context.Background(), node)

	bootstrapAddr, _ := multiaddr.NewMultiaddr(utils.BOOTSTRAP_NODE_MULTIADDR)
	peerInfo, _ := peer.AddrInfoFromP2pAddr(bootstrapAddr)
	err := node.Connect(context.Background(), *peerInfo)

	if err != nil {
		log.Fatal(err)
	}

	table.Bootstrap(context.Background())
	fmt.Println("Celestial Node:", node.ID())
	pubsub, _ := pubsub.NewGossipSub(context.Background(), node)

	celestialNode := utils.CelestialNode{
		NodeId:               node.ID(),
		PeerToPeerServer:     &node,
		Start:                StartCelestial,
		PubSubService:        pubsub,
		DistributedHashTable: table,
		NodeTopicMappings:    map[string]map[string]bool{},
	}

	return &celestialNode

}

func StartCelestial(node *utils.CelestialNode) {
	go (*node.PeerToPeerServer).Network().Listen((*node.PeerToPeerServer).Addrs()[0])
	utils.InitializePeerDiscovery(node.PeerToPeerServer)
	go func() {
		universalCosmos := CreateUniversalCosmos(node, context.Background())
		for {
			message, _ := universalCosmos.Subscription.Next(context.Background())
			HandleUniversalCosmosMessage(message, node)
		}
	}()
	select {}

}
