#!/bin/bash

# set -e;

# Setup
## Ensure headscale network exists
if ! docker network inspect headscale-test-network >/dev/null 2>&1; then
  echo "[*] Creating headscale network"
  docker network create --driver bridge headscale-test-network
fi

## Start a headscale container
echo "[*] Starting headscale container"
docker compose -f resources/docker-compose.yaml up -d
# Endpoint available inside docker network
export HEADSCALE_INTERNAL_ENDPOINT="http://headscale:8080"
# Endpoint available for host for go test
export HEADSCALE_ENDPOINT="http://127.0.0.1:8080"
export HEADSCALE_API_KEY=$(docker exec headscale headscale apikeys create)

## Create a user
echo "[*] Creating a user"
TEST_USER=terraform
docker exec headscale headscale user create $TEST_USER -o json

# Create a preauth key for the user
echo "[*] Creating a preauth key"
PREAUTHKEY=$(docker exec headscale headscale preauthkeys create --user $TEST_USER)
if [ -z "$PREAUTHKEY" ]; then
  echo "Failed to create preauth key for user $TEST_USER"
  exit 1
fi
echo "[*] Preauth key created: $PREAUTHKEY"

## Run tailscale container and connect a node
echo "[*] Creating and connecting test device"
docker run -d --rm \
  --name tailscale-container \
  --hostname tailscale-container \
  --network headscale-test-network \
  --cap-add=NET_ADMIN \
  --cap-add=SYS_MODULE \
  --device=/dev/net/tun \
  -v /dev/net/tun:/dev/net/tun \
  tailscale/tailscale \
  tailscaled

echo "[*] Connecting tailscale container to headscale container"

docker exec tailscale-container tailscale up --authkey $PREAUTHKEY --login-server $HEADSCALE_INTERNAL_ENDPOINT

echo "[*] Checking tailscale status"
docker exec tailscale-container tailscale status

# Run tests
export TF_ACC=1
echo "[*] Running tests"
TF_ACC=1 go test ./headscale/test

# Clean up
echo "[*] Cleaning up tailscale container"
docker stop tailscale-container

echo "[*] Stopping headscale container and deleting data"
docker compose -f resources/docker-compose.yaml down
rm -rf resources/data
