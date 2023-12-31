package utils

import (
	"crypto/rsa"
	"database/sql"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var BOOTSTRAP_NODE_MULTIADDR string = "/ip4/127.0.0.1/tcp/7001/p2p/QmSgdAwbFv5W1eLwCzmwFT8NqCnQTvWWrj3avtAfWPFjTm"

const (
	NODE_BOOTSTRAP = "bootstrap"
	NODE_PEER      = "peer"
)

const (
	REQUEST_DECRYPTION_KEY             = "request_decryption_key"
	AWKNOWLEDGE_DECRYPTION_KEY_REQUEST = "awknowledge_decryption_key_request"
)

var UTIL_GENERATE_KEYS = "keygen"
var UTIL_DEBUG_MODE = "debug"
var UTIL_SIMULATION_API = "api"
var UTIL_PEER_DATABASE_MIGRATION = "migrate:peer"
var UTIL_RUN_SIMULATION = "simulation"
var UTIL_RUN_IPFS_NODE = "ipfs"
var UTIL_DEPLOY_CONTRACTS = "deploy-contracts"

type Contracts struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty"`
	ScatterProtocolContract   string             `bson:"SCATTER_PROTOCOL_CONTRACT"`
	ScatterTokenContract      string             `bson:"SCATTER_TOKEN_CONTRACT"`
	TrainingTokenContract     string             `bson:"TRAINING_TOKEN_CONTRACT"`
	EvaluationTokenContract   string             `bson:"EVALUATION_TOKEN_CONTRACT"`
	ModelTokenContract        string             `bson:"MODEL_TOKEN_CONTRACT"`
	ReputationManagerContract string             `bson:"REPUTATION_MANAGER_CONTRACT"`
	VoteManagerContract       string             `bson:"VOTE_MANAGER_CONTRACT"`
}

var PROTOCOL_IDENTIFIER protocol.ID = "/scatter-protocol/1.0.0"
var DATA_DIRECTORY = "training/data"

// Boostrap Message Codes
const (
	MESSAGE_CODE_JOIN_NETWORK = "join network"
)

const (
	PEER_REQUESTOR  = "requestor"
	PEER_TRAINER    = "trainer"
	PEER_VALIDATOR  = "validator"
	PEER_CHALLENGER = "challenger"
	PEER_NO_ROLE    = "no role"
)

const (
	LOG_EVENT_TOKEN_BALANCE = "Token Balance"
	LOTTERY_BALANCE         = "Lottery Balance"
)

const (
	GOERLI  = 5
	SEPOLIA = 11155111
)

const VALIDATOR_STAKE = 25000

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

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
	ExternalServer       *http.Server                // Http Server to communicate with node
	PeerToPeerServer     *host.Host                  // Peer2Peer server to communicate with network
	DistributedHashTable *dht.IpfsDHT                // Distributed hash table for peer discovery
	TrainingLock         *map[string]map[string]bool // A training lock to ensure subsequent fired emits don't cause extra training
	JobQueue             *JobProcessor               // Asynchrnous job queue to process training requests
	CronJobRunner        *CronJobRunner              // Cron job runner to keep subscriptions alive
	DummyLoad            *bool                       // Use dummy data to make simulations go faster
	LogMode              *bool                       // Log mode to keep track of metrics during simulations
	WorkspaceId          *string                     // For debugging and simulation purposes
	Subscriptions        *SubscriptionManager        // Subscription manager for all event listeners
	Subscribe            func(*PeerNode)             // Function to subscribe to all event listeners
	AESKey               *[]byte                     // Private key for RSA
	RSAPublicKey         *rsa.PublicKey              // Public Key for RSA
	AESChannels          *map[string]map[string]map[string](chan bool)
}

type CronJobRunner struct {
	interval time.Duration
	jobs     []func()
	stop     chan struct{}
	wg       sync.WaitGroup
}

type SubscriptionManager struct {
	TrainingInitialized     *ethereum.Subscription
	RequestForEvaluationSet *ethereum.Subscription
	ModelReadyToValidate    *ethereum.Subscription
	DebugEvent              *ethereum.Subscription
	JobComplete             *ethereum.Subscription
}

// JobProcessor represents an asynchronous job processor
type JobProcessor struct {
	JobQueue   chan Job
	Wg         sync.WaitGroup
	Shutdown   chan struct{}
	IsShutdown bool
	Mu         sync.Mutex
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

type DecryptionKeyRequest struct {
	RequestorAddress string `json:"requestorAddress"`
	TopicName        string `json:"topicName"`
}

type AwknowledgeDecryptionKeyRequest struct {
	TrainerAddress   string `json:"trainerAddress"`
	RequestorAddress string `json:"requestorAddress"`
	TopicName        string `json:"topicName"`
	PublicKey        []byte `json:"publicKey"`
	PrivateKey       []byte `json:"privateKey"`
}

type AddTopicRequestBody struct {
	TopicName             string  `json:"topicName"`
	RequestorAddress      *string `json:"requestorAddress,omitempty"`
	TrainingJobPath       *string `json:"trainingJobPath,omitempty"`
	Reward                *int64  `json:"reward,omitempty"`
	Stake                 *int64  `json:"stake,omitempty"`
	ValidationThreshold   *int64  `json:"validationThreshold,omitempty"`
	EvaluationJobPath     *string `json:"evaluationJobPath,omitempty"`
	EvaluationJobDataPath *string `json:"evaluationJobDataPath,omitempty"`
	TrainingDataPath      *string `json:"trainingDataPath,omitempty"`
}

type InitializeTrainingRequestBody struct {
	TopicName string `json:"TopicName"`
}

type DiscoveryNotifee struct {
	Host host.Host
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
	ApiPort                   *int    `json:"apiPort"`
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
	DummyLoad               *bool   `json:"dummyLoad"`
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

type Job struct {
	ID       string
	Function interface{}
	Args     []interface{}
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

type EvaluationRequestEvent struct {
	Requestor common.Address
	TopicName string
}

type DebugEvent struct {
	Message string
}

type JobCompleteEvent struct {
	Requestor common.Address
	TopicName string
}

type SimulationStartNodeRequest struct {
	PeerType          string  `json:"peerType"`
	ApiPort           uint    `json:"apiPort"`
	PostgresUsername  string  `json:"postgresUsername"`
	PostgresPassword  string  `json:"postgresPassword"`
	DatabasePort      uint    `json:"databasePort"`
	BlockchainAddress string  `json:"blockchainAddress"`
	PrivateKey        string  `json:"privateKey"`
	DummyLoad         bool    `json:"dummyLoad"`
	UseMdns           bool    `json:"useMdns"`
	IsProtocolOwner   bool    `json:"isProtocolOwner"`
	WorkspaceId       *string `json:"workspaceId,omitempty"`
}

type SimulationAppTopicRequest struct {
	BlockchainAddress     string  `json:"blockchainAddress"`
	TopicName             string  `json:"topicName"`
	RequestorAddress      *string `json:"requestorAddress,omitempty"`
	TrainingJobPath       *string `json:"trainingJobPath,omitempty"`
	Reward                *int64  `json:"reward,omitempty"`
	Stake                 *int64  `json:"stake,omitempty"`
	ValidationThreshold   *int64  `json:"validationThreshold,omitempty"`
	EvaluationJobPath     *string `json:"evaluationJobPath,omitempty"`
	EvaluationJobDataPath *string `json:"evaluationJobDataPath,omitempty"`
	TrainingDataPath      *string `json:"trainingDataPath,omitempty"`
}

type SimulationStopNodeRequest struct {
	BlockchainAddress string `json:"blockchainAddress"`
}

type SimulationDeployProtocolRequest struct {
	BlockchainAddress string `json:"blockchainAddress"`
	PrivateKey        string `json:"privateKey"`
}

type SimulationStartTrainingRequest struct {
	BlockchainAddress string `json:"blockchainAddress"`
	TopicName         string `json:"topicName"`
}

type SimulationTransferInitialSupplyRequest struct {
	BlockchainAddress string         `json:"blockchainAddress"`
	PrivateKey        string         `json:"privateKey"`
	TransferAmounts   map[string]int `json:"transferAmounts"`
}

// Generic Log event for metrics --> can be simplified into x and y
// Units of measurement will be determined by the log type, client-side
type LogEvent struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	LogType           string             `bson:"logType"`
	WorkspaceID       primitive.ObjectID `bson:"workspaceId"`
	BlockchainAddress string             `bson:"blockchainAddress"`
	XDataPoint        float64            `bson:"xDataPoint"`
	YDataPoint        float64            `bson:"yDataPoint"`
	CreatedAt         primitive.DateTime `bson:"createdAt"`
}
