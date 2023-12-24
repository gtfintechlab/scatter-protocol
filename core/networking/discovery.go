package networking

import (
	"context"
	"log"
	"sync"
	"time"

	utils "github.com/gtfintechlab/scatter-protocol/core/utils"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"
)

func NewDHT(ctx context.Context, host host.Host) (*dht.IpfsDHT, error) {
	bootstrapMultiAddr, _ := multiaddr.NewMultiaddr(utils.BOOTSTRAP_NODE_MULTIADDR)
	bootstrapPeers := []multiaddr.Multiaddr{bootstrapMultiAddr}

	var options []dht.Option

	// if no bootstrap peers give this peer act as a bootstraping node
	// other peers can use this peers ipfs address for peer discovery via dht
	if len(bootstrapPeers) == 0 {
		options = append(options, dht.Mode(dht.ModeServer))
	}

	kdht, err := dht.New(ctx, host, options...)
	if err != nil {
		return nil, err
	}

	if err = kdht.Bootstrap(ctx); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	for _, peerAddr := range bootstrapPeers {
		peerInfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerInfo); err != nil {
				log.Printf("Error while connecting to node %q: %-v\n", peerInfo, err)
			} else {
				log.Printf("Connected to bootstrap node: %q\n", peerInfo.ID.String())
				stream, _ := host.NewStream(context.Background(),
					peerInfo.ID, utils.PROTOCOL_IDENTIFIER)
				SendMessage(&stream, MESSAGE_JOIN_NETWORK)

			}
		}()
	}
	wg.Wait()

	return kdht, nil
}
func InitializePeerDiscoveryDHT(ctx context.Context, host *host.Host, dht *dht.IpfsDHT, rendezvous string) {
	routingDiscovery := drouting.NewRoutingDiscovery(dht)
	dutil.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-(ctx).Done():
			return
		case <-ticker.C:
			peerChannels, _ := routingDiscovery.FindPeers(ctx, rendezvous)

			for peerNode := range peerChannels {

				if peerNode.ID == (*host).ID() {
					continue
				}

				if (*host).Network().Connectedness(peerNode.ID) != network.Connected {
					(*host).Connect(ctx, peerNode)
					log.Printf("Connected to peer node %s\n", peerNode.ID.Pretty())
				}
			}
		}

	}
}
