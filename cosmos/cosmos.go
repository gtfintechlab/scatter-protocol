package cosmos

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/gtfintechlab/scatter-protocol/utils"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func CreateCosmos(node *utils.PeerNode, ctx context.Context, topicName string) *utils.Cosmos {
	topicName = formatTopicName(node, topicName)
	topic, err := node.PubSubService.Join(topicName)

	if err != nil {
		log.Fatal(err)
	}

	subscription, err := topic.Subscribe()
	if err != nil {
		log.Fatal(err)
	}

	cosmosId := uuid.New()

	cosmos := &utils.Cosmos{
		Context:            ctx,
		CosmosName:         topicName,
		Topic:              topic,
		CosmosId:           cosmosId,
		Creator:            node,
		Subscription:       subscription,
		TrainingInProgress: false,
	}

	return cosmos
}

func JoinCosmos(ctx context.Context, node *utils.PeerNode, topicName string) *pubsub.Topic {
	topic, _ := node.PubSubService.Join(topicName)
	msg := utils.CosmosMessage{
		Type:     utils.PEER_TRAINER_JOIN,
		Message:  "Joined Network",
		SenderID: (*node.PeerToPeerServer).ID(),
	}

	msgBytes, _ := json.Marshal(msg)
	topic.Publish(ctx, msgBytes)
	return topic
}

func formatTopicName(node *utils.PeerNode, topicName string) string {
	return (*node.PeerToPeerServer).ID().String() + ": " + topicName
}

func CreateUniversalCosmos(node *utils.CelestialNode, ctx context.Context) *utils.UniversalCosmos {
	topicName := utils.UNIVERSAL_COSMOS
	topic, err := node.PubSubService.Join(topicName)

	if err != nil {
		log.Fatal(err)
	}

	subscription, err := topic.Subscribe()
	if err != nil {
		log.Fatal(err)
	}

	cosmosId := uuid.New()

	cosmos := &utils.UniversalCosmos{
		Context:            ctx,
		CosmosName:         topicName,
		Topic:              topic,
		CosmosId:           cosmosId,
		Creator:            node,
		Subscription:       subscription,
		TrainerPeerCount:   -1,
		TrainingInProgress: false,
	}

	return cosmos
}

func AddTopicToUniversalCosmos(node *utils.PeerNode, topicName string) {
	msg := utils.CosmosMessage{
		Type:     utils.PEER_REQUESTOR_ADD_TOPIC,
		Message:  topicName,
		SenderID: (*node.PeerToPeerServer).ID(),
	}

	msgBytes, _ := json.Marshal(msg)
	(*node.PubSubTopics)[utils.UNIVERSAL_COSMOS].Publish(context.Background(), msgBytes)
}

func GetTopicsFromUniversalCosmos(node *utils.PeerNode) {
	msg := utils.CosmosMessage{
		Type:     utils.PEER_GET_TOPICS,
		Message:  "",
		SenderID: (*node.PeerToPeerServer).ID(),
	}

	msgBytes, _ := json.Marshal(msg)
	(*node.PubSubTopics)[utils.UNIVERSAL_COSMOS].Publish(context.Background(), msgBytes)

}