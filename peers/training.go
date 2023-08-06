package peers

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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
	fmt.Println(requestorId, topicName)
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

	fmt.Println(cmd.String())
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run container for %s:%s with error: %s", requestorId, topicName, err)
	}
}

func RunTrainingProcedure(requestorId string, topicName string) {
	dockerSetup()
	buildImage(requestorId, topicName)
	runContainer(requestorId, topicName)

}
