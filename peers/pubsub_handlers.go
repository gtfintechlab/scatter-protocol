package peers

import (
	"encoding/json"
	"strings"

	utils "github.com/gtfintechlab/scatter-protocol/utils"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func HandleCosmosMessage(message *pubsub.Message, node *utils.PeerNode) {

	var cosmosMessage utils.CosmosMessage

	json.Unmarshal(message.Data, &cosmosMessage)

	switch messageType := cosmosMessage.Type; messageType {
	// Trainer telling requestors that it plans on joining the training group
	case utils.PEER_TRAINER_JOIN:
		topicName := strings.Split(message.GetTopic(), ":")[1]
		addTopicFromInfo(node, message.ReceivedFrom.String(), topicName, utils.PEER_TRAINER, nil)
	}
}
