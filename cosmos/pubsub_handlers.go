package cosmos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func HandleUniversalCosmosMessage(message *pubsub.Message, node *utils.CelestialNode) {
	var cosmosMessage utils.CosmosMessage

	json.Unmarshal(message.Data, &cosmosMessage)
	switch messageType := cosmosMessage.Type; messageType {
	case utils.PEER_REQUESTOR_ADD_TOPIC:
		if _, ok := node.NodeTopicMappings[message.ReceivedFrom.String()]; ok {
			node.NodeTopicMappings[message.ReceivedFrom.String()][cosmosMessage.Message] = true
		} else {
			node.NodeTopicMappings[message.ReceivedFrom.String()] = map[string]bool{}
			node.NodeTopicMappings[message.ReceivedFrom.String()][cosmosMessage.Message] = true
		}

	case utils.PEER_REQUESTOR_REMOVE_TOPIC:
		delete(node.NodeTopicMappings[message.ReceivedFrom.String()], cosmosMessage.Message)

	case utils.PEER_GET_TOPICS:
		fmt.Println("HERE")
		networkStream, _ := (*node.PeerToPeerServer).NewStream(context.Background(),
			message.ReceivedFrom, utils.PROTOCOL_IDENTIFIER)

		fmt.Println("HERE")
		networking.SendMessage(&networkStream, utils.Message{
			MessageType: utils.PEER_GET_TOPICS,
			Payload:     node.NodeTopicMappings,
		})
	}
}