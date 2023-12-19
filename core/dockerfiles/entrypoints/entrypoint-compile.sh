#!/bin/sh

# Set the path to your contract-deploy.json file
JSON_FILE="contract-deploy.json"

# Compile contracts using solc
jq -r '.contracts[] | "solc contracts/\(.name).sol --overwrite --bin --abi -o build --base-path '\''/'\'' --include-path '\''node_modules/'\'' --optimize --optimize-runs 200"' $JSON_FILE | while read -r solc_command; do
  echo "Running solc: $solc_command"
  eval "$solc_command"
done

# Generate Go bindings using abigen
jq -r '.contracts[] | "abigen --bin=./build/\(.name).bin --abi=./build/\(.name).abi --pkg=\(.packageName) --out=protocol/\(.projectDirectory)/\(.name).go;"' $JSON_FILE | while read -r abigen_command; do
  echo "Running abigen: $abigen_command"
  eval "$abigen_command"
done

jq -r '.contracts[] | "chmod -R 777 protocol/\(.projectDirectory)/\(.name).go;"' $JSON_FILE | while read -r chmod_command; do
  echo "Running chmod: $chmod_command"
  eval "$chmod_command"
done
