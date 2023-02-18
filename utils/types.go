package utils

import (
	"net/http"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/protocol"
)

var BOOTSTRAP_NODE_MULTIADDR string = "/ip4/127.0.0.1/tcp/7001/p2p/QmSgdAwbFv5W1eLwCzmwFT8NqCnQTvWWrj3avtAfWPFjTm"
var NODE_BOOTSTRAP string = "bootstrap"
var NODE_PEER string = "peer"

var UTIL_GENERATE_KEYS = "keygen"

var PROTOCOL_IDENTIFIER protocol.ID = "/scatter-protocol/1.0.0"

// Boostrap Message Codes
const (
	MESSAGE_CODE_JOIN_NETWORK = "join network"
)

// Peer Node Message Codes
const (
	PEER_SWITCH_ROLE = "peer switch role"
)

const (
	PEER_REQUESTOR = "requestor"
	PEER_TRAINER   = "trainer"
)

type Message struct {
	MessageType string
	Payload     map[string]interface{}
}

type PeerNode struct {
	PeerType         string
	Start            func(*PeerNode)
	ExternalServer   *http.Server
	PeerToPeerServer *host.Host
	Topics           map[string]string
}

type BootstrapNode struct {
	Start            func(*BootstrapNode)
	PeerToPeerServer *host.Host
}
type ValidatorNode struct {
	ValidatorType string
	Start         func(*ValidatorNode)
}

type CosmosNode struct {
	NodeType string
	Start    func(*CosmosNode)
}

type AddTopicRequestBody struct {
	Topic string  `json:"topic"`
	Path  *string `json:"path,omitempty"`
}
