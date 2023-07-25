package cosmos

import (
	"context"
	"encoding/json"

	"github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func HandleUniversalCosmosMessage(message *pubsub.Message, node *utils.CelestialNode) {
	var cosmosMessage utils.CosmosMessage
	json.Unmarshal(message.Data, &cosmosMessage)
	switch messageType := cosmosMessage.Type; messageType {
	case utils.PEER_REQUESTOR_ADD_TOPIC:
		if checkIfTopicMappingExists(node, message.ReceivedFrom.String(), cosmosMessage.Message) {
			break
		}

		addTopicMappingFromInfo(node, message.ReceivedFrom.String(), utils.PEER_REQUESTOR, cosmosMessage.Message)

	case utils.PEER_REQUESTOR_REMOVE_TOPIC:
		deleteTopicByInfo(node, message.ReceivedFrom.String(), utils.PEER_REQUESTOR)

	case utils.PEER_GET_TOPICS:
		networkStream, _ := (*node.PeerToPeerServer).NewStream(context.Background(),
			message.ReceivedFrom, utils.PROTOCOL_IDENTIFIER)

		networking.SendMessage(&networkStream, utils.Message{
			MessageType: utils.PEER_GET_TOPICS,
			Payload: map[string]interface{}{
				"allTopics": getAllTopics(node),
			},
		})
	}
}
