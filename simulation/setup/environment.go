package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		exec.Command("npm", "run", "deploy-evaluation:no-push").Run()
		exec.Command("npm", "run", "deploy-training:no-push").Run()
		exec.Command("npm", "run", "deploy-model:no-push").Run()
		exec.Command("npm", "run", "deploy-token:no-push").Run()
		exec.Command("npm", "run", "deploy-protocol:no-push").Run()

		exec.Command(
			"go", "run",
			"scripts/deploy.go",
			"--pkey", *simulationConfiguration.Environment.ProtocolOwnerPrivateKey,
		).Run()
	}

}
