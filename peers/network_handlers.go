package peers

import (
	"encoding/json"
	"fmt"
	"os"

	networking "github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
	network "github.com/libp2p/go-libp2p/core/network"
)

func peerGetTopicsHandler(node *utils.PeerNode, message *utils.Message) {
	var topicList map[string]interface{}
	jsonData, _ := json.Marshal(message.Payload)
	json.Unmarshal(jsonData, &topicList)
	node.InformationBox.CosmosTopics = &topicList
}

func peerStartTrainingHandler(node *utils.PeerNode, message *utils.Message, stream *network.Stream) {
	var trainingInfo utils.TrainingInfoFromRequestor
	jsonData, _ := json.Marshal(message.Payload)
	json.Unmarshal(jsonData, &trainingInfo)
	zipFileBytes := trainingInfo.Files
	dataPath := fmt.Sprintf("training/trainer/jobs/%s/%s",
		(*stream).Conn().RemotePeer().String(),
		trainingInfo.Topic)
	fmt.Println(trainingInfo.Topic)
	os.MkdirAll(dataPath, os.ModePerm)
	networking.WriteBytesToFile(fmt.Sprintf("%s/training_zip.zip", dataPath), zipFileBytes)

}
