package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type EnvironmentConfig struct {
	DeployProtocol          *bool   `json:"deployProtocol"`
	ProtocolOwnerAddress    *string `json:"protocolOwnerAddress"`
	ProtocolOwnerPrivateKey *string `json:"protocolOwnerPrivateKey"`
	EthereumNode            *string `json:"ethereumNode"`
}

type SimulationConfiguration struct {
	Environment EnvironmentConfig `json:"environment"`
}

func replaceEnvValue(filename, key, newValue string) error {
	// Open the .env file for reading
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a temporary file for writing
	tempFile, err := os.CreateTemp("", "tempenv")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// Create a scanner to read the .env file
	scanner := bufio.NewScanner(file)

	// Iterate through each line in the .env file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			// Check if the current line contains the key we want to replace
			if parts[0] == key {
				// Replace the old value with the new value
				line = key + "=" + newValue
			}
		}

		// Write the line to the temporary file
		_, _ = tempFile.WriteString(line + "\n")
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return err
	}

	// Rename the temporary file to the original filename
	err = os.Rename(tempFile.Name(), filename)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var simulationName string
	flag.StringVar(&simulationName, "simulation", "lazy_load",
		"Name of the simulation to run")
	flag.Parse()

	var simulationConfiguration SimulationConfiguration
	jsonData, err := os.ReadFile(fmt.Sprintf("simulation/simulations/%s.json", simulationName))
	if err != nil {
		log.Fatal(`An error occurred when reading the simulation. 
			Make sure you define the simulation configuration in /simulation/simulations`, err)
		return
	}

	err = json.Unmarshal(jsonData, &simulationConfiguration)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
		return
	}

	if *simulationConfiguration.Environment.DeployProtocol {
		replaceEnvValue(".env", "PRIVATE_KEY", *simulationConfiguration.Environment.ProtocolOwnerPrivateKey)
		exec.Command("docker-compose", "-f", "docker-compose-contracts.yml", "up").Run()

		// exec.Command(
		// 	"go", "run",
		// 	"scripts/deploy.go",
		// 	"--pkey", *simulationConfiguration.Environment.ProtocolOwnerPrivateKey,
		// ).Run()
	}

}
