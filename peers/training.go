package peers

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/docker/docker/client"
	"github.com/google/uuid"
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
func buildImage(dockerfilePath string, imageName string) {
	// Create the Docker build command
	buildCmd := exec.Command("docker", "build", "-t", imageName, dockerfilePath)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr

	err := buildCmd.Run()
	if err != nil {
		panic(err)
	}
}

func RunDockerContainer() string {
	dockerSetup()
	imageName := fmt.Sprintf("training:%s", uuid.New().String())
	buildImage("training/", imageName)
	runCmd := exec.Command("docker", "run", "-d", imageName)
	// Execute the run command and capture the output
	output, err := runCmd.Output()
	if err != nil {
		panic(err)
	}

	// Get the container ID from the output
	containerID := string(output)
	containerID = containerID[:len(containerID)-1] // Remove newline character
	return containerID
}
