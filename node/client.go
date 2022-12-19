package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {
	node, err := libp2p.New()
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println("Listen addresses:", node.Addrs())

	if err := node.Close(); err != nil {
		fmt.Println(err)
	}
}
