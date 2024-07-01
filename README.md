# Scatter Protocol

Scatter Protocol is an incentivized and trustless protocol for decentralized, federated learning. The Scatter Protocol repository has two primary components to it:
- **Core:** This contains all the core code for the client node and protocol. The client node consists of an HTTP server, which allows you to interact with the node and protocol more broadly. Other components include a peer-to-peer server for inter-node communication, an asynchronous job queue, a data store, and event listeners. The protocol is defined through smart contracts, as provided by the `core/contracts` folder.
- **Web Interface:** We provide a simple web interface to run simulations. This interface allows you to control the number of nodes, roles each node plays, steps to execute for each node, etc.

If you find this project useful in your research, please use the following BibTeX entry for citation.
```
@INPROCEEDINGS{Saho2408:Scatter,
    AUTHOR="Samrat Sahoo and Sudheer Chava",
    TITLE="Scatter Protocol: An Incentivized and Trustless Protocol for Decentralized Federated Learning",
    BOOKTITLE="2024 IEEE International Conference on Blockchain (Blockchain) (Blockchain-2024)",
    ADDRESS="Copenhagen, Denmark",
    PAGES=8,
    DAYS=18,
    MONTH=aug,
    YEAR=2024,
    ABSTRACT="Federated Learning is a form of privacy-preserving machine learning where multiple entities train local models, which are then aggregated into a global model. Current forms of federated learning rely on a centralized server to orchestrate the process, leading to issues such as requiring trust in the orchestrator, the necessity of a middleman, and a single point of failure. Blockchains provide a way to record information on a transparent, distributed ledger that is accessible and verifiable by any entity. We leverage these properties of blockchains to produce a decentralized, federated learning marketplace-style protocol for training models collaboratively.  Our core contributions are as follows: first, we introduce novel staking, incentivization, and penalization mechanisms to deter malicious nodes and encourage benign behavior. Second, we introduce a dual-layered validation mechanism to ensure the authenticity of the models trained. Third, we test different components of our system to verify sufficient incentivization, penalization, and resistance to malicious attacks."
}
```

## Technology Specifications
- Blockchain: Any EVM Compatible Chain (Production) or Hardhat Network (Development)
- Node Datastore: PostgreSQL
- Peer-to-Peer Communication: go-libp2p
- Training Job Runtime Environment: Docker + Open Container Initiative (OCI) Runtime
- Smart Contracts: Solidity + Solc + Go-Ethereum Client 
- Web Interface: MongoDB + Next.js

## Setup Repository Locally

This repository requires Docker to run external resources (i.e., PostgreSQL, Hardhat Network, etc.). To instantiate these resources:

1. Install [Docker](https://docs.docker.com/engine/install/)
2. Start the external resources with Docker Compose: `docker compose up`

We recommend interacting with the protocol via the web interface and the wrapper API we provide to control nodes. You can set this up as follows.

**Step 1:** Install dependencies for the core module: 
```
cd core
go mod download
```
**Step 2:** Generate keys for the bootstrap node (make sure you are still in the core folder)
```
go run main.go --util keygen
```
**Step 3:** Modify the multiaddress for the bootstrap node (`BOOTSTRAP_NODE_MULTIADDR` in `core/utils/types.go`). The format of the multiaddress takes the following: `/ip4/127.0.0.1/tcp/{PORT NUMBER}/p2p/{NODE ID}`. The bootstrap node is currently configured to be on port 7001 automatically. You can get the node ID by spinning up the bootstrap node and seeing the ID in the console. You can do this with the following command:
```
go run main.go --type bootstrap
> 2023/12/30 20:02:35 bootstrap_node.go:24: Bootstrap Node: QmSgdAwbFv5W1eLwCzmwFT8NqCnQTvWWrj3avtAfWPFjTm
```
**Step 4:** Install OpenZepplin contracts (using yarn)
```
yarn install
```
**Step 5:** Run the Wrapper API
```
yarn api
```
**Step 6:** Open a new terminal and install dependencies for the web interface
```
cd web
yarn install
```
**Step 7:** Start the web interface
```
yarn dev
```
