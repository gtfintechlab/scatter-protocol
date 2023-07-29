package peers

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/docker/docker/client"
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
	exec.Command(
		"docker", "build",
		"--build-arg", fmt.Sprintf("NODE_ID=%s", requestorId),
		"--build-arg", fmt.Sprintf("TOPIC_NAME=%s", topicName),
		"-t", fmt.Sprintf("outer-image:%s", requestorId), "training/trainer",
	).Output()
}

func RunDockerContainer(requestorId string, topicName string) {
	dockerSetup()
	buildImage(requestorId, topicName)
}
