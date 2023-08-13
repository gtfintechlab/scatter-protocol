package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	evaluationtoken "github.com/gtfintechlab/scatter-protocol/protocol/evaluation"
	scattertoken "github.com/gtfintechlab/scatter-protocol/protocol/scatter"
	trainingtoken "github.com/gtfintechlab/scatter-protocol/protocol/training"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var client, _ = ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))

var scatterContractAddress common.Address = common.HexToAddress(SCATTER_PROTCOL_CONTRACT)
var trainingContractAddress common.Address = common.HexToAddress(TRAINING_TOKEN_CONTRACT)
var evaluationContractAddress common.Address = common.HexToAddress(EVALATION_TOKEN_CONTRACT)

var scatterTokenContract, _ = scattertoken.NewScatterprotocol(scatterContractAddress, client)
var trainingTokenContract, _ = trainingtoken.NewTrainingtoken(trainingContractAddress, client)
var evaluationTokenContract, _ = evaluationtoken.NewEvaluationtoken(evaluationContractAddress, client)

func getTransactor() *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(GOERLI))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth
}
func AddScatterTokenStake(stakeAmount *big.Int) *types.Transaction {
	auth := getTransactor()
	transaction, err := scatterTokenContract.AddStake(auth, stakeAmount)

	if err != nil {
		log.Fatal(err)
	}
	return transaction
}

func GetScatterTokenStake() *big.Int {
	auth := getTransactor()
	stake, err := scatterTokenContract.GetOwnStake(&bind.CallOpts{
		From: auth.From,
	})

	if err != nil {
		log.Fatal(err)
	}
	return stake
}

func PublishTrainingJob(trainingJobPath string, topicName string) string {
	auth := getTransactor()
	ipfsCid := UploadFileToIpfs(trainingJobPath)
	_, err := trainingTokenContract.PublishTrainingJob(
		auth, ipfsCid, topicName)

	if err != nil {
		log.Fatal(err)
	}
	return ipfsCid
}

func PublishEvaluationJob(evaluationJobPath string) string {
	auth := getTransactor()
	ipfsCid := UploadFileToIpfs(evaluationJobPath)
	transaction, err := evaluationTokenContract.PublishEvaluationJob(auth, auth.From, ipfsCid)

	if err != nil {
		log.Fatal(err)
	}
	return transaction.Hash().String()
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
