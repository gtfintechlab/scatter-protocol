package protocol

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	evaluationtoken "github.com/gtfintechlab/scatter-protocol/protocol/evaluation"
	scatterprotocol "github.com/gtfintechlab/scatter-protocol/protocol/scatter-protocol"
	scattertoken "github.com/gtfintechlab/scatter-protocol/protocol/scatter-token"
	trainingtoken "github.com/gtfintechlab/scatter-protocol/protocol/training"
	"github.com/gtfintechlab/scatter-protocol/utils"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var ethereumClient, _ = ethclient.Dial(os.Getenv("ETHEREUM_NODE"))

var scatterProtocolContractAddress common.Address = common.HexToAddress(utils.SCATTER_PROTOCOL_CONTRACT)
var scatterTokenContractAddress common.Address = common.HexToAddress(utils.SCATTER_TOKEN_CONTRACT)
var trainingContractAddress common.Address = common.HexToAddress(utils.TRAINING_TOKEN_CONTRACT)
var evaluationContractAddress common.Address = common.HexToAddress(utils.EVALUATION_TOKEN_CONTRACT)

var scatterProtocolContract, _ = scatterprotocol.NewScatterprotocol(scatterProtocolContractAddress, ethereumClient)
var scatterTokenContract, _ = scattertoken.NewScattertoken(scatterTokenContractAddress, ethereumClient)
var trainingTokenContract, _ = trainingtoken.NewTrainingtoken(trainingContractAddress, ethereumClient)
var evaluationTokenContract, _ = evaluationtoken.NewEvaluationtoken(evaluationContractAddress, ethereumClient)

var CHAIN_ID, _ = strconv.Atoi(os.Getenv("CHAIN_ID"))
var CHAIN = big.NewInt(int64(CHAIN_ID))

func InitRequestorNode(node *utils.PeerNode) {
	auth := getTransactor(node)
	scatterProtocolContract.InitRequestorNode(auth)
}

func InitTrainerNode(node *utils.PeerNode) {
	auth := getTransactor(node)
	scatterProtocolContract.InitTrainerNode(auth)
}

func InitValidatorNode(node *utils.PeerNode) {
	auth := getTransactor(node)
	_, err := scatterProtocolContract.InitValidatorNode(auth)

	if err != nil {
		log.Fatal(err)
	}
}

func AddScatterTokenStake(node *utils.PeerNode, amount int64) {
	auth := getTransactor(node)
	_, err := scatterTokenContract.StakeToken(auth, big.NewInt(amount))

	if err != nil {
		log.Fatal(err)
	}
}

func GetScatterTokenStake(node *utils.PeerNode, account string) big.Int {
	auth := getTransactor(node)
	stake, err := scatterTokenContract.GetAccountStake(&bind.CallOpts{From: auth.From}, common.HexToAddress(account))

	if err != nil {
		log.Fatal(err)
	}

	return *stake
}

func CanBecomeValidator(node *utils.PeerNode, account string) bool {
	auth := getTransactor(node)
	isValid, err := scatterTokenContract.CanBecomeValidator(&bind.CallOpts{From: auth.From}, common.HexToAddress(account))

	if err != nil {
		log.Fatal(err)
	}

	return isValid
}

func GetProtocolRequestors(node *utils.PeerNode, skipAmount uint64) []string {
	auth := getTransactor(node)
	skip := big.NewInt(int64(skipAmount))
	addresses, _ := scatterProtocolContract.GetRequestorAddresses(&bind.CallOpts{
		From: auth.From,
	}, skip)

	return utils.DeleteEmptyElements(strings.Split(strings.Trim(addresses, " "), "\n"))
}

func GetAllProtocolRequestors(node *utils.PeerNode) []string {
	requestorList := []string{}
	skipAmount := 0
	currentRequestorList := GetProtocolRequestors(node, uint64(skipAmount))

	for len(currentRequestorList) != 0 {
		requestorList = append(requestorList, currentRequestorList...)
		skipAmount += 10
		currentRequestorList = GetProtocolRequestors(node, uint64(skipAmount))

	}

	return requestorList
}

func GetTopicsByAddress(node *utils.PeerNode, address string, skipAmount uint64) []string {
	auth := getTransactor(node)
	skip := big.NewInt(int64(skipAmount))
	topics, _ := scatterProtocolContract.GetTopicsByRequestorAddress(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address), skip)

	return utils.DeleteEmptyElements(strings.Split(strings.Trim(topics, " "), "\n"))
}

func GetTrainersByAddressAndTopic(node *utils.PeerNode, address string, topicName string, skipAmount uint64) []string {
	auth := getTransactor(node)
	skip := big.NewInt(int64(skipAmount))
	trainers, _ := scatterProtocolContract.GetTrainersByAddressAndTopic(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address), topicName, skip)
	return utils.DeleteEmptyElements(strings.Split(strings.Trim(trainers, " "), "\n"))
}

func GetAllTrainersByAddressAndTopic(node *utils.PeerNode, address string, topicName string) []string {
	trainerList := []string{}
	skipAmount := 0
	currentTrainerList := GetTrainersByAddressAndTopic(node, address, topicName, uint64(skipAmount))
	for len(currentTrainerList) != 0 {
		trainerList = append(trainerList, currentTrainerList...)
		skipAmount += 10
		currentTrainerList = GetTrainersByAddressAndTopic(node, address, topicName, uint64(skipAmount))
	}

	return trainerList
}

func CheckIfTopicExistsForRequestor(node *utils.PeerNode, address string, topicName string) bool {
	auth := getTransactor(node)
	exists, _ := scatterProtocolContract.CheckIfTopicExistsForRequestor(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address), topicName)

	return exists
}

func GetAllTopicsByAddress(node *utils.PeerNode, address string) []string {
	topicList := []string{}
	skipAmount := 0
	currentTopicList := GetTopicsByAddress(node, address, uint64(skipAmount))
	for len(currentTopicList) != 0 {
		topicList = append(topicList, currentTopicList...)
		skipAmount += 10
		currentTopicList = GetTopicsByAddress(node, address, uint64(skipAmount))
	}

	return topicList
}

func StartTraining(node *utils.PeerNode, topicName string) {
	auth := getTransactor(node)
	_, err := scatterProtocolContract.StartTraining(auth, topicName)
	if err != nil {
		log.Fatal(err)
	}
}

func GetRoleByAddress(node *utils.PeerNode, address string) string {
	auth := getTransactor(node)
	role, _ := scatterProtocolContract.GetRoleByAddress(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address))

	return role
}

func GetCidFromAddressAndTopic(node *utils.PeerNode, address string, topicName string) string {
	auth := getTransactor(node)
	trainingInfo, err := scatterProtocolContract.AddressToTrainingJobInfo(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address), topicName)

	if err != nil {
		log.Fatal(err)
	}

	return trainingInfo.TrainingJobCid
}

func AddTopicForRequestor(node *utils.PeerNode, trainingJobPath string, topicName string, reward int64) (string, string) {
	auth := getTransactor(node)
	ipfsCid := utils.UploadFileToIpfs(trainingJobPath)
	transaction, err := scatterProtocolContract.RequestorAddTopic(
		auth, ipfsCid, topicName, big.NewInt(reward))

	if err != nil {
		log.Fatal(err)
	}
	return ipfsCid, transaction.Hash().Hex()
}

func AddTopicForTrainer(node *utils.PeerNode, address string, topicName string, stakeAmount int64) {
	auth := getTransactor(node)
	_, err := scatterProtocolContract.TrainerAddTopic(
		auth, common.HexToAddress(address), topicName, big.NewInt(stakeAmount))

	if err != nil {
		log.Fatal(err)
	}

}

func ChangeRoleToTrainer(node *utils.PeerNode) {
	auth := getTransactor(node)
	scatterProtocolContract.InitTrainerNode(auth)
}

func ChangeRoleToRequestor(node *utils.PeerNode) {
	auth := getTransactor(node)
	scatterProtocolContract.InitRequestorNode(auth)
}

func SetNodeId(node *utils.PeerNode, nodeId string) {
	auth := getTransactor(node)
	scatterProtocolContract.SetNodeId(auth, nodeId)
}

func GetNodeIdFromAddress(node *utils.PeerNode, blockchainAddress string) string {
	auth := getTransactor(node)
	nodeId, _ := scatterProtocolContract.AddressToNodeId(&bind.CallOpts{From: auth.From}, common.HexToAddress(blockchainAddress))

	return nodeId
}

func IsValidatorForRequestorAndTopic(node *utils.PeerNode, requestorAddress string, topicName string) bool {
	auth := getTransactor(node)
	isValidator, _ := scatterProtocolContract.ValidatorTrainingMap(
		&bind.CallOpts{From: auth.From},
		common.HexToAddress(requestorAddress),
		topicName,
		auth.From,
	)

	return isValidator
}
func PublishEvaluationJob(node *utils.PeerNode, evaluationJobPath string, topicName string) string {
	auth := getTransactor(node)
	ipfsCid := utils.UploadFileToIpfs(evaluationJobPath)
	_, err := scatterProtocolContract.SubmitEvaluationSet(auth, topicName, ipfsCid)

	if err != nil {
		log.Fatal(err)
	}
	return ipfsCid
}

func GetEvaluationJobFromAddressAndTopic(node *utils.PeerNode, requestorAddress string, topicName string) string {
	auth := getTransactor(node)
	trainingInfo, err := scatterProtocolContract.AddressToTrainingJobInfo(&bind.CallOpts{From: auth.From},
		common.HexToAddress(requestorAddress), topicName)

	ipfsCid := trainingInfo.EvaluationJobCid

	if err != nil {
		log.Fatal(err)
	}
	return ipfsCid

}

func SubmitEvaluationScore(node *utils.PeerNode, requestorAddress string, topicName string, trainerAddress string, score *big.Int) {
	auth := getTransactor(node)
	scatterProtocolContract.SubmitEvaluationScore(auth, common.HexToAddress(requestorAddress), topicName, common.HexToAddress(trainerAddress), score)
}

func PublishModel(node *utils.PeerNode, modelPath string, requestorAddress string, topicName string) {
	auth := getTransactor(node)
	modelIpfsCid := utils.UploadFileToIpfs(modelPath)
	scatterProtocolContract.PublishModelToProtocol(auth,
		modelIpfsCid,
		common.HexToAddress(requestorAddress),
		topicName,
	)
}

func GetModelCidByTrainer(node *utils.PeerNode, requestorAddress string, topicName string, trainerAddress string) string {
	auth := getTransactor(node)

	modelCid, _ := scatterProtocolContract.ModelLogger(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(requestorAddress), topicName, common.HexToAddress(trainerAddress))

	return modelCid
}

func GetScatterTokenBalance(node *utils.PeerNode) *big.Int {
	auth := getTransactor(node)
	balance, _ := scatterTokenContract.BalanceOf(&bind.CallOpts{From: auth.From}, auth.From)
	return balance
}

func GetPooledRewardByAddressAndTopic(node *utils.PeerNode, address string, topicName string) big.Int {
	auth := getTransactor(node)
	trainingJobInfo, err := scatterProtocolContract.AddressToTrainingJobInfo(&bind.CallOpts{From: auth.From}, common.HexToAddress(address), topicName)
	if err != nil {
		log.Fatal(err)
	}

	return *trainingJobInfo.PooledReward
}

func getTransactor(node *utils.PeerNode) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(*node.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, CHAIN)
	auth.Value = big.NewInt(0)

	return auth
}

func GenerateDigitalSignature(privateKeyStr string, inputData map[string]interface{}) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(inputData)
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func VerifyDigitalSignature(address string, data map[string]interface{}, signature []byte) (bool, error) {
	ethereumAddress := common.HexToAddress(address)
	messageBytes, _ := json.Marshal(data)
	recoveredPublicKey, err := crypto.SigToPub(crypto.Keccak256(messageBytes), signature)
	if err != nil {
		return false, fmt.Errorf("error recovering public key: %v", err)
	}
	recoveredAddress := crypto.PubkeyToAddress(*recoveredPublicKey)
	isValid := recoveredAddress == ethereumAddress
	return isValid, nil
}
