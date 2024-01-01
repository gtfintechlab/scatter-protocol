package scatterlogs

import (
	"context"

	"github.com/gtfintechlab/scatter-protocol/core/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsNodeMalicious(node *utils.PeerNode, blockchainAddress string) bool {
	client, err := utils.DbConnect()
	for client == nil || err != nil {
		client, err = utils.DbConnect()
	}
	defer client.Client().Disconnect(context.Background())
	workspaceId, _ := primitive.ObjectIDFromHex(*node.WorkspaceId)
	filter := bson.M{
		"workspaceId":       workspaceId,
		"blockchainAddress": bson.M{"$regex": primitive.Regex{Pattern: blockchainAddress, Options: "i"}},
	}

	var loggedNode utils.LoggedSimulationNode
	client.Collection("protocolnodes").FindOne(context.TODO(), filter).Decode(&loggedNode)

	return loggedNode.IsMalicious
}
