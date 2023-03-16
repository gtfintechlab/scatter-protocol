package networking

import (
	"encoding/json"
	"log"

	utils "github.com/gtfintechlab/scatter-protocol/utils"
	"github.com/libp2p/go-libp2p/core/network"
)

func SendMessage(stream *network.Stream, message utils.Message) {
	defer (*stream).Close()
	payload, _ := json.Marshal(&message)
	(*stream).Write(payload)
}

func DecodeMessage(stream *network.Stream) utils.Message {
	buffer := StreamReader(1024, stream)
	// Read the message as a JSON
	var message map[string]interface{}
	err := json.Unmarshal(buffer, &message)
	if err != nil {
		log.Fatal(err)
	}

	return utils.Message{
		MessageType: message["MessageType"].(string),
		Payload:     message["Payload"],
	}
}
