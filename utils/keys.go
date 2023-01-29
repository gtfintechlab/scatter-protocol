package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func GenerateKeys() {
	privateKey, publicKey, _ := crypto.GenerateKeyPair(crypto.RSA, 2048)
	fmt.Println(privateKey)
	privateMarshaled, _ := crypto.MarshalPrivateKey(privateKey)
	publicMarshaled, _ := crypto.MarshalPublicKey(publicKey)
	ioutil.WriteFile("node/privateKey.pem", privateMarshaled, 0644)
	ioutil.WriteFile("node/publicKey.pem", publicMarshaled, 0644)

}

func LoadKeys() (crypto.PrivKey, crypto.PubKey) {
	privateKey, _ := ioutil.ReadFile("node/privateKey.pem")
	publicKey, _ := ioutil.ReadFile("node/privateKey.pem")

	privateUnmarshaled, _ := crypto.UnmarshalPrivateKey(privateKey)
	publicUnmarshaled, _ := crypto.UnmarshalPublicKey(publicKey)

	return privateUnmarshaled, publicUnmarshaled
}
