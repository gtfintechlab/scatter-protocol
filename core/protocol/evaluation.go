package protocol

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gtfintechlab/scatter-protocol/core/networking"
	peerDatabase "github.com/gtfintechlab/scatter-protocol/core/peers/db"
	"github.com/gtfintechlab/scatter-protocol/core/utils"
	"github.com/libp2p/go-libp2p/core/peer"
)

func downloadEvaluationJob(requestorAddress string, ipfsCid string) {
	filePath := fmt.Sprintf("training/validator/jobs/%s/%s/", strings.ToLower(requestorAddress), strings.ToLower(ipfsCid))
	fileBytes, _ := utils.GetFileBytesFromIPFS(ipfsCid)
	networking.UnzipFolder(fileBytes, filePath)
}

func downloadEvaluationJobData(requestorAddress string, evaluationJobCid string, evaluationDataCid string) {
	filePath := fmt.Sprintf("training/validator/jobs/%s/%s/data/", strings.ToLower(requestorAddress), strings.ToLower(evaluationJobCid))
	fileBytes, _ := utils.GetFileBytesFromIPFS(evaluationDataCid)
	networking.UnzipFolder(fileBytes, filePath)
}

func buildEvaluationImage(requestorId string, ipfsCid string) {
	basePath, _ := os.Getwd()
	requestorIdLower := strings.ToLower(requestorId)
	ipfsCidLower := strings.ToLower(ipfsCid)

	os.MkdirAll(fmt.Sprintf("%s/training/validator/jobs/%s/%s/output",
		basePath,
		requestorIdLower,
		ipfsCidLower,
	),
		0777,
	)

	cmd := exec.Command(
		"docker", "build",
		"-t", fmt.Sprintf("%s:%s", requestorIdLower, ipfsCidLower),
		"-f", fmt.Sprintf("%s/training/validator/jobs/%s/%s/Dockerfile", basePath, requestorIdLower, ipfsCidLower),
		fmt.Sprintf("%s/training/validator/jobs/%s/%s/", basePath, requestorIdLower, ipfsCidLower),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal("An error occurred when building the image ", err)
	}

}

func runEvaluationContainer(requestorId string, ipfsCid string) *big.Int {
	requestorIdLower := strings.ToLower(requestorId)
	ipfsCidLower := strings.ToLower(ipfsCid)
	cmd := exec.Command(
		"docker", "run",
		fmt.Sprintf("%s:%s", requestorIdLower, ipfsCidLower),
	)
	var dataBuffer bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &dataBuffer)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run container for %s:%s with error: %s", requestorIdLower, ipfsCidLower, err)
	}

	scoreString := strings.ReplaceAll(dataBuffer.String(), "\n", "")
	scoreString = strings.ReplaceAll(scoreString, " ", "")
	score, err := strconv.Atoi(scoreString)
	if err != nil {
		log.Fatalf("Output failed to output a valid score %s", err)
	}

	return big.NewInt(int64(score))
}

func downloadTrainerModel(node *utils.PeerNode, requestorAddress string, topicName string, ipfsCid string, trainerAddress string) {
	basePath, _ := os.Getwd()
	requestorIdLower := strings.ToLower(requestorAddress)
	ipfsCidLower := strings.ToLower(ipfsCid)
	modelPath := fmt.Sprintf("%s/training/validator/jobs/%s/%s/model.pth", basePath, requestorIdLower, ipfsCidLower)
	os.Remove(modelPath)

	modelCid := GetModelCidByTrainer(node, requestorAddress, topicName, trainerAddress)
	modelBytes, _ := utils.GetFileBytesFromIPFS(modelCid)

	nodeId := GetNodeIdFromAddress(node, trainerAddress)
	// Get Decryption key
	peerId, err := peer.Decode(nodeId)

	if err != nil {
		log.Fatal(err)
	}
	stream, err := (*node.PeerToPeerServer).NewStream(
		context.Background(),
		peerId,
		utils.PROTOCOL_IDENTIFIER,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create the boolean channel
	aesChannel := make(chan bool)
	// Initialize the structure
	if _, ok := (*node.AESChannels)[strings.ToLower(requestorAddress)]; !ok {
		(*node.AESChannels)[strings.ToLower(requestorAddress)] = make(map[string]map[string]chan bool)
	}
	if _, ok := (*node.AESChannels)[strings.ToLower(requestorAddress)][topicName]; !ok {
		(*node.AESChannels)[strings.ToLower(requestorAddress)][topicName] = make(map[string]chan bool)
	}
	if _, ok := (*node.AESChannels)[strings.ToLower(requestorAddress)][topicName][strings.ToLower(trainerAddress)]; !ok {
		(*node.AESChannels)[strings.ToLower(requestorAddress)][topicName][strings.ToLower(trainerAddress)] = aesChannel
	}
	networking.SendMessage(&stream, utils.Message{
		MessageType: utils.REQUEST_DECRYPTION_KEY,
		Payload: map[string]interface{}{
			"requestorAddress": requestorAddress,
			"topicName":        topicName,
		},
	})

	select {
	case <-(*node.AESChannels)[strings.ToLower(requestorAddress)][topicName][strings.ToLower(trainerAddress)]:
		privateKeyBytes := peerDatabase.RetrieveDecryptionKey(node, trainerAddress, topicName, requestorAddress)
		modelBytes, err = utils.DecryptData(modelBytes, privateKeyBytes)
		if err != nil {
			log.Fatal(err)
		}

		err = networking.WriteBytesToFile(modelPath, modelBytes)
		if err != nil {
			log.Fatalf("Failed to download model from ipfs with error: %s", err)
		}
		close((*node.AESChannels)[strings.ToLower(requestorAddress)][topicName][strings.ToLower(trainerAddress)])
		return
	}
}
