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
	os.MkdirAll(fmt.Sprintf("%s/training/trainer/jobs/%s/%s/output", basePath, requestorId, ipfsCid), 0700)
	os.MkdirAll(fmt.Sprintf("%s/training/trainer/jobs/%s/%s/data", basePath, requestorId, ipfsCid), 0700)

	exec.Command("cp", "--recursive", dataPath, fmt.Sprintf("%s/training/trainer/jobs/%s/%s/data/", basePath, requestorId, ipfsCid)).Run()

	cmd := exec.Command(
		"docker", "build",
		// "--build-arg", fmt.Sprintf("DATA_PATH=%s", dataPath),
		"-t", fmt.Sprintf("%s:%s", strings.ToLower(requestorId), strings.ToLower(ipfsCid)),
		"-f", fmt.Sprintf("%s/training/trainer/jobs/%s/%s/Dockerfile", basePath, requestorId, ipfsCid),
		fmt.Sprintf("%s/training/trainer/jobs/%s/%s/", basePath, requestorId, ipfsCid),
	)

	fmt.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		log.Fatal("An error occurred when building the image ", err)
	}

}

func runContainer(requestorId string, ipfsCid string) {
	basePath, _ := os.Getwd()
	cmd := exec.Command(
		"docker", "run",
		"-v", fmt.Sprintf("\"%s/training/trainer/jobs/%s/%s/output:/tmp/output/\"",
			basePath,
			requestorId, ipfsCid),
		"-it", fmt.Sprintf("%s:%s", strings.ToLower(requestorId), strings.ToLower(ipfsCid)),
	)
	fmt.Println(cmd)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run container for %s:%s with error: %s", requestorId, ipfsCid, err)
	}
}

func downloadTrainingJob(ipfsCid string, requestorId string) {
	filePath := fmt.Sprintf("training/trainer/jobs/%s/%s/", requestorId, ipfsCid)
	fileBytes, _ := utils.GetFileBytesFromIPFS(ipfsCid, filePath)
	networking.UnzipFolder(fileBytes, filePath)
}
