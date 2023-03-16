package utils

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

var BOOTSTRAP_NODE_MULTIADDR string = "/ip4/127.0.0.1/tcp/7001/p2p/QmSgdAwbFv5W1eLwCzmwFT8NqCnQTvWWrj3avtAfWPFjTm"

const (
	NODE_BOOTSTRAP = "bootstrap"
	NODE_PEER      = "peer"
	NODE_CELESTIAL = "celestial"
)

var UTIL_GENERATE_KEYS = "keygen"

var PROTOCOL_IDENTIFIER protocol.ID = "/scatter-protocol/1.0.0"

// Boostrap Message Codes
const (
	MESSAGE_CODE_JOIN_NETWORK = "join network"
)

// Peer Node Message Codes
const (
	PEER_SWITCH_ROLE            = "peer switch role"
	PEER_TRAINER_JOIN           = "trainer join"
	PEER_REQUESTOR_ADD_TOPIC    = "requestor add topic"
	PEER_REQUESTOR_REMOVE_TOPIC = "requestor remove topic"
	PEER_GET_TOPICS             = "trainer get topics"
)

const (
	UNIVERSAL_COSMOS = "Universal Cosmos"
)

const (
	PEER_REQUESTOR = "requestor"
	PEER_TRAINER   = "trainer"
)

type Message struct {
	MessageType string
	Payload     any
}

type PeerNode struct {
	PeerType             string
	NodeId               peer.ID
	Start                func(*PeerNode)
	ExternalServer       *http.Server
	PeerToPeerServer     *host.Host
	Topics               map[string]string
	DistributedHashTable *dht.IpfsDHT
	PubSubService        *pubsub.PubSub
	PubSubTopics         *map[string]*pubsub.Topic
	TopicTrainerMap      map[string][]string
	CosmosTopics         any // Will fix all any types later lol
}

type BootstrapNode struct {
	Start                func(*BootstrapNode)
	NodeId               peer.ID
	PeerToPeerServer     *host.Host
	DistributedHashTable *dht.IpfsDHT
}
type ValidatorNode struct {
	ValidatorType        string
	Start                func(*ValidatorNode)
	PeerToPeerServer     *host.Host
	DistributedHashTable *dht.IpfsDHT
}

type CelestialNode struct {
	NodeId               peer.ID
	NodeTopicMappings    map[string]map[string]bool
	PeerToPeerServer     *host.Host
	Start                func(*CelestialNode)
	PubSubService        *pubsub.PubSub
	DistributedHashTable *dht.IpfsDHT
}

type AddTopicRequestBody struct {
	Topic string  `json:"topic"`
	Path  *string `json:"path,omitempty"`
}

type PublishTopicRequestBody struct {
	Topic string `json:"topic"`
}

type DiscoveryNotifee struct {
	Host host.Host
}

type Cosmos struct {
	Context            context.Context
	CosmosName         string
	Topic              *pubsub.Topic
	CosmosId           uuid.UUID
	Creator            *PeerNode
	Subscription       *pubsub.Subscription
	TrainingInProgress bool
}

type UniversalCosmos struct {
	Context            context.Context
	CosmosName         string
	Topic              *pubsub.Topic
	CosmosId           uuid.UUID
	Creator            *CelestialNode
	Subscription       *pubsub.Subscription
	TrainerPeerCount   int
	TrainingInProgress bool
}

type CosmosMessage struct {
	Type     string  `json:"Type"`
	Message  string  `json:"Message"`
	SenderID peer.ID `json:"SenderID"`
}
