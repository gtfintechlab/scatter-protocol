## Scatter Protocol

Scatter Protocol is an incentivized, privacy-centered federated learning protocol for collaborative machine learning.

## Project Milestones

| **Milestone**                                         | **Progress**    |
| ----------------------------------------------------- | --------------- |
| General: Initial Project Setup                        | ✅ COMPLETE ✅    |
| Networking: Simple Peer2Peer Communication            | ✅ COMPLETE ✅    |
| Nodes: Persistent Peer2Peer Hashing                   | ✅ COMPLETE ✅    |
| Smart Contracts: ERC-20 Scatter Token                 | ✅ COMPLETE ✅    |
| Bootstrap Nodes: Distributed Hash Table Functionality | ✅ COMPLETE ✅    |
| Nodes: Peer Discovery Mechanism                       | ✅ COMPLETE ✅    |
| Celestial Nodes: Universal Cosmos                     | ✅ COMPLETE ✅    |
| Peer Nodes: Intialize Training Request                | ✅ COMPLETE ✅    |
| Peer Nodes: Smart Contract Interactions               | 🚧 IN PROGRESS 🚧 |
| Peer Nodes: Federated Learning Execution              | ✅ COMPLETE ✅    |
| Peer Nodes: Model Weights Exclusion                   |                 |
| Validator Nodes: Token Staking                        |                 |
| Validator Nodes: Training Consensus Mechanism         |                 |
| Validator Nodes: Node Eviction Mechanism              |                 |
| Smart Contracts: Token Reward System                  |                 |
| Node: Reputation System                               |                 |

## Generate Node Keys (For Bootstrap Nodes)

```
go run main.go --util keygen
```

## Run a bootstrap node

```
go run main.go --type bootstrap
```

## Run a peer node

```
go run main.go --type peer
```
