package cosmos

import (
	"context"
	"fmt"

	"github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func InitCelestialNode() *utils.CelestialNode {
	node, _ := libp2p.New()
	table, _ := dht.New(context.Background(), node)

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

func StartCelestial(node *utils.CelestialNode, useMdns bool) {
	go (*node.PeerToPeerServer).Network().Listen((*node.PeerToPeerServer).Addrs()[0])

	if useMdns {
		utils.InitializePeerDiscovery(node.PeerToPeerServer)

	} else {
		go networking.InitializePeerDiscoveryDHT(
			context.Background(),
			*node.PeerToPeerServer,
			node.DistributedHashTable,
			string(utils.PROTOCOL_IDENTIFIER),
		)
	}

	go func() {
		universalCosmos := CreateUniversalCosmos(node, context.Background())
		for {
			message, _ := universalCosmos.Subscription.Next(context.Background())
			HandleUniversalCosmosMessage(message, node)
		}
	}()
	select {}

}
