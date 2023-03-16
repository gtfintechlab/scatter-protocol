package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

func (node *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	fmt.Printf("Discovered new peer %s\n", pi.ID.Pretty())
	err := node.Host.Connect(context.Background(), pi)
	if err != nil {
		log.Fatal("Error Connecting to", pi.ID.Pretty(), err)
	}
}

func InitializePeerDiscovery(node *host.Host) {
	mdnsService := mdns.NewMdnsService(*node, string(PROTOCOL_IDENTIFIER),
		&DiscoveryNotifee{Host: *node})

	go mdnsService.Start()

}
