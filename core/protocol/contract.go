package protocol

import (
	"context"
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
	modeltoken "github.com/gtfintechlab/scatter-protocol/core/protocol/model"
	scatterprotocol "github.com/gtfintechlab/scatter-protocol/core/protocol/scatter-protocol"
	scattertoken "github.com/gtfintechlab/scatter-protocol/core/protocol/scatter-token"
	"github.com/gtfintechlab/scatter-protocol/core/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

var _ = godotenv.Load(".env")
var ethereumClient, _ = ethclient.Dial(os.Getenv("ETHEREUM_NODE"))

func initScatterProtocolContract() *scatterprotocol.Scatterprotocol {
	var scatterProtocolContractAddress common.Address = common.HexToAddress(GetContractInfo().ScatterProtocolContract)
	var scatterProtocolContract, _ = scatterprotocol.NewScatterprotocol(scatterProtocolContractAddress, ethereumClient)
	return scatterProtocolContract

}

func initScatterTokenContract() *scattertoken.Scattertoken {
	var scatterTokenContractAddress common.Address = common.HexToAddress(GetContractInfo().ScatterTokenContract)
	var scatterTokenContract, _ = scattertoken.NewScattertoken(scatterTokenContractAddress, ethereumClient)
	return scatterTokenContract
}

func initModelTokenContract() *modeltoken.Modeltoken {
	var modelTokenContractAddress common.Address = common.HexToAddress(GetContractInfo().ModelTokenContract)
	var modelTokenContract, _ = modeltoken.NewModeltoken(modelTokenContractAddress, ethereumClient)
	return modelTokenContract
}
func InitRequestorNode(node *utils.PeerNode) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	scatterProtocolContract.InitRequestorNode(auth)
}

func InitTrainerNode(node *utils.PeerNode) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	scatterProtocolContract.InitTrainerNode(auth)
}

func InitValidatorNode(node *utils.PeerNode) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	_, err := scatterProtocolContract.InitValidatorNode(auth)

	if err != nil {
		log.Fatal(err)
	}
}

func AddScatterTokenStake(node *utils.PeerNode, amount int64) {

	var scatterTokenContract = initScatterTokenContract()

	auth := getTransactor(node)
	_, err := scatterTokenContract.StakeToken(auth, big.NewInt(amount))

	if err != nil {
		log.Fatal(err)
	}
}

func GetScatterTokenStake(node *utils.PeerNode, account string) big.Int {

	var scatterTokenContract = initScatterTokenContract()

	auth := getTransactor(node)
	stake, err := scatterTokenContract.GetAccountStake(&bind.CallOpts{From: auth.From}, common.HexToAddress(account))

	if err != nil {
		log.Fatal(err)
	}

	return *stake
}

func CanBecomeValidator(node *utils.PeerNode, account string) bool {

	var scatterTokenContract = initScatterTokenContract()

	auth := getTransactor(node)
	isValid, err := scatterTokenContract.CanBecomeValidator(&bind.CallOpts{From: auth.From}, common.HexToAddress(account))

	if err != nil {
		log.Fatal(err)
	}

	return isValid
}

func GetProtocolRequestors(node *utils.PeerNode, skipAmount uint64) []string {

	var scatterProtocolContract = initScatterProtocolContract()

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

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	skip := big.NewInt(int64(skipAmount))
	topics, _ := scatterProtocolContract.GetTopicsByRequestorAddress(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address), skip)

	return utils.DeleteEmptyElements(strings.Split(strings.Trim(topics, " "), "\n"))
}

func GetTrainersByAddressAndTopic(node *utils.PeerNode, address string, topicName string, skipAmount uint64) []string {

	var scatterProtocolContract = initScatterProtocolContract()

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

	var scatterProtocolContract = initScatterProtocolContract()

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

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	_, err := scatterProtocolContract.StartTraining(auth, topicName)
	if err != nil {
		log.Fatal(err)
	}
}

func GetRoleByAddress(node *utils.PeerNode, address string) string {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	role, _ := scatterProtocolContract.GetRoleByAddress(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address))

	return role
}

func GetTrainingJobFromAddressAndTopic(node *utils.PeerNode, address string, topicName string) string {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	trainingInfo, err := scatterProtocolContract.AddressToFederatedJob(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(address), topicName)

	if err != nil {
		log.Fatal(err)
	}

	return trainingInfo.TrainingJobCid
}

func AddTopicForRequestor(node *utils.PeerNode, trainingJobPath string, evaluationJobPath string, topicName string, reward int64, validationThreshold int64) (string, string, string) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	trainingIpfsCid := utils.UploadFileToIpfs(trainingJobPath)
	evaluationIpfsCid := utils.UploadFileToIpfs(evaluationJobPath)
	transaction, err := scatterProtocolContract.RequestorAddTopic(
		auth, trainingIpfsCid, evaluationIpfsCid, topicName, big.NewInt(reward), big.NewInt(validationThreshold))

	if err != nil {
		return AddTopicForRequestor(node, trainingJobPath, evaluationJobPath, topicName, reward, validationThreshold)
	}
	return trainingIpfsCid, evaluationIpfsCid, transaction.Hash().Hex()
}

func AddTopicForTrainer(node *utils.PeerNode, requestorAddress string, topicName string, stakeAmount int64) {

	var scatterProtocolContract = initScatterProtocolContract()
	auth := getTransactor(node)

	_, err := scatterProtocolContract.TrainerAddTopic(
		auth, common.HexToAddress(requestorAddress), topicName, big.NewInt(stakeAmount))

	if err != nil {
		log.Fatal(err)
	}

}

func ChangeRoleToTrainer(node *utils.PeerNode) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	scatterProtocolContract.InitTrainerNode(auth)
}

func ChangeRoleToRequestor(node *utils.PeerNode) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	scatterProtocolContract.InitRequestorNode(auth)
}

func SetNodeId(node *utils.PeerNode, nodeId string) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	scatterProtocolContract.SetNodeId(auth, nodeId)
}

func GetNodeIdFromAddress(node *utils.PeerNode, blockchainAddress string) string {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	nodeId, _ := scatterProtocolContract.AddressToNodeId(&bind.CallOpts{From: auth.From}, common.HexToAddress(blockchainAddress))

	return nodeId
}

func IsValidatorForRequestorAndTopic(node *utils.PeerNode, requestorAddress string, topicName string) bool {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	isValidator, _ := scatterProtocolContract.ValidatorTrainingMap(
		&bind.CallOpts{From: auth.From},
		common.HexToAddress(requestorAddress),
		topicName,
		auth.From,
	)

	return isValidator
}
func PublishEvaluationData(node *utils.PeerNode, evaluationDataPath string, topicName string) string {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	ipfsCid := utils.UploadFileToIpfs(evaluationDataPath)
	_, err := scatterProtocolContract.SubmitEvaluationSet(auth, topicName, ipfsCid)

	if err != nil {
		// For some reason the nonce caculation always messes up here so I just do a recursive retry (lol.)
		return PublishEvaluationData(node, evaluationDataPath, topicName)
	}
	return ipfsCid
}

func GetEvaluationJobFromAddressAndTopic(node *utils.PeerNode, requestorAddress string, topicName string) string {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	federatedJobInfo, err := scatterProtocolContract.AddressToFederatedJob(&bind.CallOpts{From: auth.From},
		common.HexToAddress(requestorAddress), topicName)

	ipfsCid := federatedJobInfo.EvaluationJobCid

	if err != nil {
		log.Fatal(err)
	}
	return ipfsCid

}

func GetEvaluationDataFromAddressAndTopic(node *utils.PeerNode, requestorAddress string, topicName string) string {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	federatedJobInfo, err := scatterProtocolContract.AddressToFederatedJob(&bind.CallOpts{From: auth.From},
		common.HexToAddress(requestorAddress), topicName)

	ipfsCid := federatedJobInfo.EvaluationJobDataCid
	if err != nil {
		log.Fatal(err)
	}
	return ipfsCid
}

func SubmitEvaluationScore(node *utils.PeerNode, requestorAddress string, topicName string, trainerAddress string, score *big.Int) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)

	_, err := scatterProtocolContract.SubmitEvaluationScore(auth, common.HexToAddress(requestorAddress), topicName, common.HexToAddress(trainerAddress), score)

	if err != nil {
		log.Fatal(err)
	}
}

func PublishModel(node *utils.PeerNode, modelPath string, requestorAddress string, topicName string) {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	modelIpfsCid := utils.UploadFileToIpfs(modelPath)
	_, err := scatterProtocolContract.PublishModelToProtocol(auth,
		modelIpfsCid,
		common.HexToAddress(requestorAddress),
		topicName,
	)

	if err != nil {
		// For some reason the nonce caculation always messes up here so I just do a recursive retry (lol.)
		PublishModel(node, modelPath, requestorAddress, topicName)
		// log.Fatal(err)
	}
}

func GetModelCidByTrainer(node *utils.PeerNode, requestorAddress string, topicName string, trainerAddress string) string {

	var modelTokenContract = initModelTokenContract()
	auth := getTransactor(node)

	modelCid, err := modelTokenContract.ModelLogger(&bind.CallOpts{
		From: auth.From,
	}, common.HexToAddress(requestorAddress), topicName, common.HexToAddress(trainerAddress))

	if err != nil {
		log.Fatal(err)
	}
	return modelCid
}

func GetScatterTokenBalance(node *utils.PeerNode) *big.Int {

	var scatterTokenContract = initScatterTokenContract()

	auth := getTransactor(node)
	balance, err := scatterTokenContract.BalanceOf(&bind.CallOpts{From: auth.From}, auth.From)

	if err != nil {
		log.Fatal(err)
	}

	return balance
}

func GetPooledRewardByAddressAndTopic(node *utils.PeerNode, address string, topicName string) big.Int {

	var scatterProtocolContract = initScatterProtocolContract()

	auth := getTransactor(node)
	trainingJobInfo, err := scatterProtocolContract.AddressToFederatedJob(&bind.CallOpts{From: auth.From}, common.HexToAddress(address), topicName)
	if err != nil {
		log.Fatal(err)
	}

	return *trainingJobInfo.PooledReward
}

func GetLotteryBalance(node *utils.PeerNode) *big.Int {
	var scatterTokenContract = initScatterTokenContract()
	auth := getTransactor(node)

	lottery, err := scatterTokenContract.GetLotteryPool(&bind.CallOpts{From: auth.From})
	if err != nil {
		log.Fatal(err)
	}

	return lottery

}

func getTransactor(node *utils.PeerNode) *bind.TransactOpts {

	var CHAIN_ID, _ = strconv.Atoi(os.Getenv("CHAIN_ID"))
	var CHAIN = big.NewInt(int64(CHAIN_ID))

	privateKey, err := crypto.HexToECDSA(*node.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, CHAIN)
	auth.Value = big.NewInt(0)
	gas, _ := ethereumClient.SuggestGasPrice(context.Background())
	gasTip, _ := ethereumClient.SuggestGasTipCap(context.Background())

	auth.GasPrice = new(big.Int).Mul(gas, big.NewInt(500))
	auth.GasLimit = uint64(2000000)
	auth.GasTipCap = gasTip

	nonce, _ := ethereumClient.PendingNonceAt(context.Background(), auth.From)
	auth.Nonce = big.NewInt(int64(nonce))
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

func GetContractInfo() utils.Contracts {
	client, _ := utils.DbConnect()
	contractsCollection := client.Collection("contracts")
	singleResult := contractsCollection.FindOne(context.Background(), bson.D{})
	var contracts utils.Contracts
	if err := singleResult.Decode(&contracts); err != nil {
		return utils.Contracts{}
	}
	return contracts
}
