package peers

import (
	"encoding/json"

	utils "github.com/gtfintechlab/scatter-protocol/utils"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func HandleCosmosMessage(message *pubsub.Message, node *utils.PeerNode) {

	var cosmosMessage utils.CosmosMessage

	json.Unmarshal(message.Data, &cosmosMessage)

	switch messageType := cosmosMessage.Type; messageType {
	case utils.PEER_TRAINER_JOIN:
		if _, ok := node.TopicTrainerMap[message.GetTopic()]; ok {
			node.TopicTrainerMap[message.GetTopic()] = append(node.TopicTrainerMap[message.GetTopic()], message.ReceivedFrom.String())
		} else {
			node.TopicTrainerMap[message.GetTopic()] = []string{}
			node.TopicTrainerMap[message.GetTopic()] = append(node.TopicTrainerMap[message.GetTopic()], message.ReceivedFrom.String())
		}
	}
}
