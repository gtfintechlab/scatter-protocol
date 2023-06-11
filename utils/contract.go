package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	scattertoken "github.com/gtfintechlab/scatter-protocol/token"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")
var contractAddress common.Address = common.HexToAddress(os.Getenv("ETHEREUM_GOERLI_CONTRACT_ADDRESS"))
var client, _ = ethclient.Dial(os.Getenv("ETHEREUM_GOERLI_NODE"))
var contract, _ = scattertoken.NewScattertoken(contractAddress, client)

func getOwnStake() *big.Int {
	stake, err := contract.GetOwnStake(nil)

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
