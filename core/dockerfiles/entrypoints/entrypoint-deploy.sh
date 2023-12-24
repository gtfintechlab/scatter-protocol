#!/bin/bash

# Check if PRIVATE_KEY is defined
if [ -z "$PRIVATE_KEY" ]; then
  go run scripts/deploy.go --contract all
else
  go run scripts/deploy.go --contract all --pkey "$PRIVATE_KEY"
fi
``
