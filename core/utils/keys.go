package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func GenerateKeys() {
	privateKey, publicKey, _ := crypto.GenerateKeyPair(crypto.RSA, 2048)
	privateMarshaled, _ := crypto.MarshalPrivateKey(privateKey)
	publicMarshaled, _ := crypto.MarshalPublicKey(publicKey)
	os.WriteFile("keys/privateKey.pem", privateMarshaled, 0644)
	os.WriteFile("keys/publicKey.pem", publicMarshaled, 0644)
}

func LoadKeys() (crypto.PrivKey, crypto.PubKey) {
	privateKey, _ := os.ReadFile("keys/privateKey.pem")
	publicKey, _ := os.ReadFile("keys/privateKey.pem")

	privateUnmarshaled, _ := crypto.UnmarshalPrivateKey(privateKey)
	publicUnmarshaled, _ := crypto.UnmarshalPublicKey(publicKey)

	return privateUnmarshaled, publicUnmarshaled
}

// GenerateAESKey generates a random AES key of the specified key size in bits
func GenerateAESKey(bits int) ([]byte, error) {
	keySize := bits / 8
	key := make([]byte, keySize)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// Save AES key to file
func SaveAESKeyToFile(key []byte, fileName string) error {
	return os.WriteFile(fileName, key, 0600)
}

// Load AES key from file
func LoadAESKeyFromFile(fileName string) ([]byte, error) {
	key, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// Pad data to make it a multiple of the block size
func pad(data []byte) []byte {
	padding := aes.BlockSize - (len(data) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Remove padding from data
func unpad(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

// Encrypt data using AES with a given key
func encryptData(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedData := pad(data)

	ciphertext := make([]byte, aes.BlockSize+len(paddedData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedData)

	return ciphertext, nil
}

// Decrypt data using AES with a given key
func DecryptData(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	unpaddedData := unpad(ciphertext)

	return unpaddedData, nil
}
