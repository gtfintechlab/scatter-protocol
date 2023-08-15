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

func buildImage(requestorId string, topicName string) {
	err := exec.Command(
		"docker", "build",
		"--build-arg", fmt.Sprintf("NODE_ID=%s", requestorId),
		"--build-arg", fmt.Sprintf("TOPIC_NAME=%s", strings.ReplaceAll(topicName, " ", "-")),
		"-t", fmt.Sprintf("outer-image:%s", requestorId), "training/trainer",
	).Run()

	if err != nil {
		log.Fatalf("Failed to build image for %s:%s with error: %s", requestorId, topicName, err)
	}
}

func runContainer(requestorId string, topicName string) {
	basePath, _ := os.Getwd()
	cmd := exec.Command(
		"docker", "run",
		"--privileged",
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"-v", fmt.Sprintf("\"%s/training/trainer/jobs/%s/%s/output:/tmp/%s/%s/output/\"",
			basePath,
			requestorId, strings.ReplaceAll(topicName, " ", "-"),
			requestorId, strings.ReplaceAll(topicName, " ", "-")),
		"-it", fmt.Sprintf("outer-image:%s", requestorId),
	)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run container for %s:%s with error: %s", requestorId, topicName, err)
	}
}

func downloadTrainingJob(ipfsCid string, requestorId string) {
	filePath := fmt.Sprintf("training/trainer/jobs/%s/%s/", requestorId, ipfsCid)
	fileBytes, _ := utils.GetFileBytesFromIPFS(ipfsCid, filePath)
	networking.UnzipFolder(fileBytes, filePath)
}
