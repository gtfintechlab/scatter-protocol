package protocol

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strings"

	"github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
)

func downloadEvaluationJob(requestorAddress string, ipfsCid string) {
	filePath := fmt.Sprintf("training/validator/jobs/%s/%s/", strings.ToLower(requestorAddress), strings.ToLower(ipfsCid))
	fileBytes, _ := utils.GetFileBytesFromIPFS(ipfsCid)
	networking.UnzipFolder(fileBytes, filePath)
}

func buildEvaluationImage(requestorId string, ipfsCid string) {
	basePath, _ := os.Getwd()
	requestorIdLower := strings.ToLower(requestorId)
	ipfsCidLower := strings.ToLower(ipfsCid)

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

	err := cmd.Run()

	if err != nil {
		log.Fatalf("Failed to run container for %s:%s with error: %s", requestorIdLower, ipfsCidLower, err)
	}

	return big.NewInt(0)
}

func downloadTrainerModel(node *utils.PeerNode, requestorAddress string, topicName string, ipfsCid string, trainerAddress string) {
	basePath, _ := os.Getwd()
	requestorIdLower := strings.ToLower(requestorAddress)
	ipfsCidLower := strings.ToLower(ipfsCid)
	modelPath := fmt.Sprintf("%s/training/validator/%s/%s/model.pth", basePath, requestorIdLower, ipfsCidLower)
	os.Remove(modelPath)

	modelCid := GetModelCidByTrainer(node, requestorAddress, topicName, trainerAddress)
	modelBytes, _ := utils.GetFileBytesFromIPFS(modelCid)
	networking.WriteBytesToFile(modelPath, modelBytes)
}
