package node

import (
	"context"
	"fmt"

	utils "github.com/gtfintechlab/scatter-protocol/utils"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	core "github.com/libp2p/go-libp2p/core"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func InitPeerNode() {
	// Create a new libp2p host for the new node
	node, _ := libp2p.New()
	table, _ := dht.New(context.Background(), node)
	// The multiaddress of the bootstrap node
	bootstrapAddr, _ := multiaddr.NewMultiaddr(utils.BOOTSTRAP_NODE_MULTIADDR)
	node.Connect(context.Background(), core.PeerAddrInfo{Addrs: []multiaddr.Multiaddr{bootstrapAddr}})
	table.Bootstrap(context.Background())

	fmt.Println(node.Addrs())

}
