# Scatter Protocol Simulation Engine Notes

- If you have an environment specification, you need to run `go run simulation/setup/environment.go` before running the engine
  - `yarn simulation` will run this script automatically
- We specify all simulation configurations in the `simulation/simulations/*`
- `nodes` specifies the different nodes in our simulation
  - Some nodes have a `blockchainAddress` and `privateKey` associated with them -- these are simulated keys from Ganache 
- `steps` specify the different steps in our simulation (generally API calls to the different endpoints)