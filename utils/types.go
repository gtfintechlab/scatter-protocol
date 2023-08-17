package utils

import (
	"database/sql"
	"encoding/json"
	"math/big"
	"net/http"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/common"
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
var UTIL_DEBUG_MODE = "debug"
var UTIL_CELESTIAL_DATABASE_MIGRATION = "migrate:celestial"
var UTIL_PEER_DATABASE_MIGRATION = "migrate:peer"
var UTIL_RUN_SIMULATION = "simulation"
var UTIL_RUN_IPFS_NODE = "ipfs"
var UTIL_DEPLOY_CONTRACTS = "deploy-contracts"

type Contracts struct {
	ScatterProtocolContract string `json:"SCATTER_PROTOCOL_CONTRACT"`
	ScatterTokenContract    string `json:"SCATTER_TOKEN_CONTRACT"`
	TrainingTokenContract   string `json:"TRAINING_TOKEN_CONTRACT"`
	EvaluationTokenContract string `json:"EVALUATION_TOKEN_CONTRACT"`
	ModelTokenContract      string `json:"MODEL_TOKEN_CONTRACT"`
}

func ReadContractInfo() Contracts {
	jsonData, _ := os.ReadFile("utils/contracts.json")
	var contractData Contracts
	json.Unmarshal(jsonData, &contractData)
	return contractData
}

var CONTRACTS = ReadContractInfo()

var (
	SCATTER_PROTOCOL_CONTRACT = CONTRACTS.ScatterProtocolContract
	TRAINING_TOKEN_CONTRACT   = CONTRACTS.TrainingTokenContract
	EVALUATION_TOKEN_CONTRACT = CONTRACTS.EvaluationTokenContract
	SCATTER_TOKEN_CONTRACT    = CONTRACTS.ScatterTokenContract
	MODEL_TOKEN_CONTRACT      = CONTRACTS.ModelTokenContract
)

var PROTOCOL_IDENTIFIER protocol.ID = "/scatter-protocol/1.0.0"
var DATA_DIRECTORY = "training/data"

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
	PEER_REQUESTOR = "requestor"
	PEER_TRAINER   = "trainer"
)

const (
	GOERLI  = 5
	SEPOLIA = 11155111
)

var IDENTITY_VERIFICATION_DATA map[string]interface{} = map[string]interface{}{
	"Verify": "Identity",
}

type Message struct {
	MessageType string
	Payload     map[string]interface{}
}

type PeerNode struct {
	PeerType             string                      // Type of Peer (requestor or trainer)
	BlockchainAddress    *string                     // Address of the wallet of this node
	PrivateKey           *string                     // Private key of the wallet associated with this account
	NodeId               peer.ID                     // ID of Node
	Start                func(*PeerNode, bool)       // Start Function for node
	DataStore            *sql.DB                     // DataStore to store information
	DatastoreLock        *sync.Mutex                 // Mutex Lock for datastore
	ExternalServer       *http.Server                // Http Server to communicate with node
	PeerToPeerServer     *host.Host                  // Peer2Peer server to communicate with network
	DistributedHashTable *dht.IpfsDHT                // Distributed hash table for peer discovery
	PubSubService        *pubsub.PubSub              // PubSub Service for the node
	PubSubTopics         *map[string]*pubsub.Topic   // PubSub Topics for topics we have subscribed to
	TrainingLock         *map[string]map[string]bool // A training lock to ensure subsequent fired emits don't cause extra training
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

type TrainingInfoFromRequestor struct {
	Files []byte `json:"files"`
	Topic string `json:"topic"`
}

type AddTopicRequestBody struct {
	Topic            string  `json:"topic"`
	RequestorAddress *string `json:"requestorAddress,omitempty"`
	Path             *string `json:"path,omitempty"`
	Reward           *int64  `json:"reward,omitempty"`
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

type PublishTrainingJobPayload struct {
	TopicCid string `json:"topicCid"`
}

type SimulationConfiguration struct {
	Environment EnvironmentConfig `json:"environment"`
	Nodes       []NodeConfig      `json:"nodes"`
	Steps       []StepConfig      `json:"steps"`
}

type StepConfig struct {
	Description string                  `json:"description"`
	NodeId      string                  `json:"nodeId"`
	Action      *string                 `json:"action"`
	Body        *map[string]interface{} `json:"body"`
	StateKey    *string                 `json:"stateKey"`
}

type NodeConfig struct {
	Id                        string  `json:"id"`
	Type                      string  `json:"type"`
	Ipv4Address               *string `json:"ipv4Address"`
	TcpPort                   *int    `json:"tcpPort"`
	ExtAddress                *string `json:"extAddress"`
	DatastorePort             *int    `json:"datastorePort"`
	DatastoreUsername         *string `json:"datastoreUsername"`
	DatastorePassword         *string `json:"datastorePassword"`
	UseMdns                   *bool   `json:"useMdns"`
	BlockchainAddress         *string `json:"blockchainAddress"`
	PrivateKey                *string `json:"privateKey"`
	InitialScatterTokenSupply *uint64 `json:"initialScatterTokenSupply"`
}

type EnvironmentConfig struct {
	DeployProtocol          *bool   `json:"deployProtocol"`
	ProtocolOwnerAddress    *string `json:"protocolOwnerAddress"`
	ProtocolOwnerPrivateKey *string `json:"protocolOwnerPrivateKey"`
	EthereumNode            *string `json:"ethereumNode"`
}

type SimulationNode struct {
	PeerNode      *PeerNode
	BootstrapNode *BootstrapNode
}

type SimulationNodeConfig struct {
	Node  SimulationNode
	Type  string
	State map[string]interface{}
}

type TopicInformation struct {
	NodeAddress      string
	NodeType         string
	TopicName        string
	TrainingTokenCID string
	PooledReward     big.Int
}

type TrainingInitializedEvent struct {
	Requestor common.Address
	TopicName string
}

type ModelReadyToValidateEvent struct {
	Trainer   common.Address
	Requestor common.Address
	TopicName string
}
