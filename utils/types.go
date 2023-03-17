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
	PEER_START_TRAINING         = "start training"
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
	Payload     map[string]interface{}
}

type InformationBox struct {
	CosmosTopics *map[string]map[string]bool
}

type PeerNode struct {
	PeerType             string                    // Type of Peer (requestor or trainer)
	NodeId               peer.ID                   // ID of Node
	Start                func(*PeerNode, bool)     // Start Function for node
	ExternalServer       *http.Server              // Http Server to communicate with node
	PeerToPeerServer     *host.Host                // Peer2Peer server to communicate with network
	TopicToDataPath      *map[string]string        // Topics to Data Path Mapping
	DistributedHashTable *dht.IpfsDHT              // Distributed hash table for peer discovery
	PubSubService        *pubsub.PubSub            // PubSub Service for the node
	PubSubTopics         *map[string]*pubsub.Topic // PubSub Topics for topics we have subscribed to
	TopicTrainerMap      *map[string][]string      // Map of the topic --> list of trainer nodes
	RequestorTopicMap    *map[string][]string      // Map of requestor --> list of topics
	InformationBox       *InformationBox           // Network Information
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
	Start                func(*CelestialNode, bool)
	PubSubService        *pubsub.PubSub
	DistributedHashTable *dht.IpfsDHT
}

type AddTopicRequestBody struct {
	Topic       string  `json:"topic"`
	RequestorId string  `json:"requestorId,omitempty"`
	Path        *string `json:"path,omitempty"`
}

type PublishTopicRequestBody struct {
	Topic string `json:"topic"`
}

type InitializeTrainingRequestBody struct {
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
