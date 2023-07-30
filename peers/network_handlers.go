package peers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
	network "github.com/libp2p/go-libp2p/core/network"
)

func peerGetTopicsHandler(node *utils.PeerNode, message *utils.Message) {
	var topicList map[string]interface{}
	jsonData, _ := json.Marshal(message.Payload)
	json.Unmarshal(jsonData, &topicList)
	node.InformationBox.CosmosTopics = &topicList
	// Accompanying lock is in GetTopicsFromUniversalCosmos()
	node.InformationBox.InformationBoxMutexLock.Unlock()
}

func peerStartTrainingHandler(node *utils.PeerNode, message *utils.Message, stream *network.Stream) {
	var trainingInfo utils.TrainingInfoFromRequestor
	jsonData, _ := json.Marshal(message.Payload)
	json.Unmarshal(jsonData, &trainingInfo)
	zipFileBytes := trainingInfo.Files

	requestorId := (*stream).Conn().RemotePeer().String()
	dataPath := fmt.Sprintf("training/trainer/jobs/%s/%s",
		requestorId,
		strings.ReplaceAll(trainingInfo.Topic, " ", "-"),
	)

	os.MkdirAll(dataPath, os.ModePerm)
	networking.UnzipFolder(zipFileBytes, dataPath)
	os.Mkdir(
		fmt.Sprintf("%s/output", dataPath),
		os.ModePerm,
	)
	RunTrainingProcedure(requestorId, trainingInfo.Topic)
}
