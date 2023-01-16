package node

import (
	"context"
	"fmt"
	"net"
	"strings"

	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func InitBootstrapNode() {
	node, _ := libp2p.New()
	// Create a new DHT
	dht.New(context.Background(), node)
	// Get the host's multiaddress
	_, port := multiaddr.SplitFirst(node.Addrs()[0])
	portString := strings.Split(port.String(), "/")[2]

	localIP := getLocalIP()

	// Build the final multiaddress string
	finalAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s", localIP, portString))

	// Print the host's multiaddress
	fmt.Println("Bootstrap node:", finalAddr)

	// Listen for incoming connections
	node.Network().Listen(finalAddr)

}

func getLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
