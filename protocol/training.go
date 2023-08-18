package protocol

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
	"github.com/gtfintechlab/scatter-protocol/networking"
	"github.com/gtfintechlab/scatter-protocol/utils"
)

func dockerSetup() {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal("Docker is not installed or the Docker daemon is not running.")
		return
	}

	// Ping the Docker server to check if it's accessible
	_, err = cli.Ping(context.Background())
	if err != nil {
		log.Fatal("Docker is not installed or the Docker daemon is not running.")
		return
	}
}

func buildImage(requestorId string, ipfsCid string, dataPath string) {
	basePath, _ := os.Getwd()
	requestorIdLower := strings.ToLower(requestorId)
	ipfsCidLower := strings.ToLower(ipfsCid)

	os.MkdirAll(fmt.Sprintf("%s/training/trainer/jobs/%s/%s/output",
		basePath,
		requestorIdLower,
		ipfsCidLower,
	),
		0700,
	)
	os.MkdirAll(fmt.Sprintf("%s/training/trainer/jobs/%s/%s/data",
		basePath,
		requestorIdLower,
		ipfsCidLower,
	),
		0700,
	)

	exec.Command("cp", "--recursive", dataPath,
		fmt.Sprintf("%s/training/trainer/jobs/%s/%s/data/",
			basePath,
			requestorIdLower,
			ipfsCidLower,
		),
	).Run()

	cmd := exec.Command(
		"docker", "build",
		"-t", fmt.Sprintf("%s:%s", requestorIdLower, ipfsCidLower),
		"-f", fmt.Sprintf("%s/training/trainer/jobs/%s/%s/Dockerfile", basePath, requestorIdLower, ipfsCidLower),
		fmt.Sprintf("%s/training/trainer/jobs/%s/%s/", basePath, requestorIdLower, ipfsCidLower),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal("An error occurred when building the image ", err)
	}

}

func runContainer(requestorId string, ipfsCid string) {
	requestorIdLower := strings.ToLower(requestorId)
	ipfsCidLower := strings.ToLower(ipfsCid)
	basePath, _ := os.Getwd()
	cmd := exec.Command(
		"docker", "run",
		"-v", fmt.Sprintf("%s/training/trainer/jobs/%s/%s/output:/tmp/output/",
			basePath,
			requestorIdLower, ipfsCidLower),
		fmt.Sprintf("%s:%s", requestorIdLower, ipfsCidLower),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run container for %s:%s with error: %s", requestorIdLower, ipfsCidLower, err)
	}
}

func submitModel(node *utils.PeerNode, requestorAddress string, ipfsCid string, topicName string) {
	requestorIdLower := strings.ToLower(requestorAddress)
	ipfsCidLower := strings.ToLower(ipfsCid)
	basePath, _ := os.Getwd()

	modelPath, _ := networking.FindFilePathWithExtension(fmt.Sprintf("%s/training/trainer/jobs/%s/%s/output",
		basePath,
		requestorIdLower, ipfsCidLower), ".pth")

	PublishModel(node, modelPath, requestorAddress, topicName)
}

func downloadTrainingJob(ipfsCid string, requestorId string) {
	filePath := fmt.Sprintf("training/trainer/jobs/%s/%s/", strings.ToLower(requestorId), strings.ToLower(ipfsCid))
	fileBytes, _ := utils.GetFileBytesFromIPFS(ipfsCid, filePath)
	networking.UnzipFolder(fileBytes, filePath)
}
