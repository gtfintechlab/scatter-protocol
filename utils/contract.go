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
	evaluationtoken "github.com/gtfintechlab/scatter-protocol/token/evaluation"
	scattertoken "github.com/gtfintechlab/scatter-protocol/token/scatter"
	trainingtoken "github.com/gtfintechlab/scatter-protocol/token/training"
	"github.com/joho/godotenv"
)

const (
	SCATTER_TOKEN_CONTRACT   = "0x7FB03B464be3b9749A7423Ee6D64D482aBD06aCe"
	TRAINING_TOKEN_CONTRACT  = "0x68f88DD73a75f9064777D5101207AE9F1f456D5E"
	EVALATION_TOKEN_CONTRACT = "0x274ad4d7749ECBd681C0BC3C8bbDCD806A4Fb48B"
)

var _ = godotenv.Load(".env")
var client, _ = ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))

var scatterContractAddress common.Address = common.HexToAddress(SCATTER_TOKEN_CONTRACT)
var trainingContractAddress common.Address = common.HexToAddress(TRAINING_TOKEN_CONTRACT)
var evaluationContractAddress common.Address = common.HexToAddress(EVALATION_TOKEN_CONTRACT)

var scatterTokenContract, _ = scattertoken.NewScattertoken(scatterContractAddress, client)
var trainingTokenContract, _ = trainingtoken.NewTrainingtoken(trainingContractAddress, client)
var contract, _ = evaluationtoken.NewEvaluationtoken(evaluationContractAddress, client)

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
