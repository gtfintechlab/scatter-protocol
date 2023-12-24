package utils

import (
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
