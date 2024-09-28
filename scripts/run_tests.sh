#!/bin/bash

# Setup
## Start a headscale container
echo "Starting headscale container"
docker compose -f resources/docker-compose.yaml up -d
export HEADSCALE_ENDPOINT="http://127.0.0.1:8080"
export HEADSCALE_API_KEY=$(docker exec headscale headscale apikeys create)

## Create a user
echo "Creating a user"
TEST_USER=terraform
docker exec headscale headscale user create $TEST_USER
PREAUTHKEY=$(docker exec headscale headscale preauthkeys create --user $TEST_USER)

## Run tailscale container and connect a node
echo "Creating and connecting test device"
docker run -d \
  --name tailscale-container \
  --hostname tailscale-container \
  --privileged \
  --net=host \
  tailscale/tailscale \
  tailscaled

docker exec tailscale-container tailscale up --authkey $PREAUTHKEY --login-server $HEADSCALE_ENDPOINT

# Run tests
export TF_ACC=1
TF_ACC=1 go test ./headscale/test

# Clean up
echo "Cleaning up"

echo "Stopping headscale container and deleting data"
docker compose -f resources/docker-compose.yaml down
rm -rf resources/data

docker stop tailscale-container
docker rm tailscale-container
