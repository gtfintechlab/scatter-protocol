package scatterlogs

import (
	"context"
	"log"
	"time"

	"github.com/gtfintechlab/scatter-protocol/core/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateLogEvent(logType string, x float64, y float64, node *utils.PeerNode) {
	client, _ := utils.DbConnect()
	workspaceId, _ := primitive.ObjectIDFromHex(*node.WorkspaceId)

	logEvent := utils.LogEvent{
		LogType:           logType,
		WorkspaceID:       workspaceId,
		BlockchainAddress: *node.BlockchainAddress,
		XDataPoint:        x,
		YDataPoint:        y,
		CreatedAt:         primitive.NewDateTimeFromTime(time.Now()),
	}

	logEventsCollection := client.Collection("logevents")
	_, err := logEventsCollection.InsertOne(context.Background(), logEvent)

	if err != nil {
		log.Fatal(err)
	}

}
