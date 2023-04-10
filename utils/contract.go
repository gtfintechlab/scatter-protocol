package utils

import (
	"encoding/hex"
	"encoding/json"
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

// Function to generate a digital signature from a private key using Keccak256Hash
func GenerateDigitalSignature(privateKeyStr string, data map[string]interface{}) ([]byte, error) {

	// Generate a private key from the input bytes
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	// Marshal the data into a JSON string
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// Hash the data using Keccak256Hash
	hashedData := crypto.Keccak256Hash(dataBytes)

	// Generate a signature for the hashed data
	signature, err := crypto.Sign(hashedData.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func VerifyDigitalSignature(expectedPublicKey string, signature []byte, data map[string]interface{}) (bool, error) {
	publicKeyBytes, _ := hex.DecodeString(expectedPublicKey)

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return false, err
	}
	// Hash the data using Keccak256Hash
	hashedData := crypto.Keccak256Hash(dataBytes)

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery ID
	verified := crypto.VerifySignature(publicKeyBytes, hashedData.Bytes(), signatureNoRecoverID)
	return verified, nil
}
