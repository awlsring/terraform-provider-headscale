#!/bin/bash

# set -e;

# Array to keep track of created containers for cleanup
TAILSCALE_CONTAINERS=()

# Cleanup function
cleanup() {
  echo "[*] Cleaning up tailscale containers"
  for container in "${TAILSCALE_CONTAINERS[@]}"; do
    echo "[*] Stopping container: $container"
    docker stop "$container" 2>/dev/null || echo "Container $container already stopped"
  done

  echo "[*] Stopping headscale container and deleting data"
  docker compose -f resources/docker-compose.yaml down
  rm -rf resources/data
}

# Set up trap to ensure cleanup happens on script exit
trap cleanup EXIT

# Function to create and connect a tailscale container
create_tailscale_container() {
  local container_name="$1"
  local tailscale_flags="$2"
  
  echo "[*] Creating and connecting test device: $container_name"
  
  # Create the tailscale container
  docker run -d --rm \
    --name "$container_name" \
    --hostname "$container_name" \
    --network headscale-test-network \
    --cap-add=NET_ADMIN \
    --cap-add=SYS_MODULE \
    --device=/dev/net/tun \
    -v /dev/net/tun:/dev/net/tun \
    tailscale/tailscale \
    tailscaled
  
  if [ $? -ne 0 ]; then
    echo "Error: Failed to create container $container_name"
    return 1
  fi
  
  # Add container to tracking array
  TAILSCALE_CONTAINERS+=("$container_name")
  
  echo "[*] Connecting $container_name to headscale container"

  local preauthkey=$(docker exec headscale headscale preauthkeys create --user "$USER_ID")
  if [ -z "$preauthkey" ]; then
    echo "Error: Failed to create preauth key for user $TEST_USER"
    return 1
  fi
  
  # Connect to headscale with the specified flags
  docker exec "$container_name" tailscale up --authkey "$preauthkey" --login-server "$HEADSCALE_INTERNAL_ENDPOINT" $tailscale_flags
  
  if [ $? -ne 0 ]; then
    echo "Error: Failed to connect $container_name to headscale"
    return 1
  fi
  
  echo "[*] Checking tailscale status for $container_name"
  docker exec "$container_name" tailscale status
  
  echo "[*] Successfully created and connected $container_name"
}

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
USER_ID=$(docker exec headscale headscale user create $TEST_USER -o json | jq .id)

# Create devices for tailnet
create_tailscale_container "basic" ""
create_tailscale_container "subnet-route" "--advertise-routes=10.0.10.0/24,192.168.1.0/24"
create_tailscale_container "exit-node" "--advertise-exit-node"

# # Run tests
export TF_ACC=1
echo "[*] Running tests"
TF_ACC=1 go test ./headscale/test
TEST_EXIT_CODE=$?

# Check if tests failed
if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo "[*] Tests failed with exit code $TEST_EXIT_CODE"
  exit $TEST_EXIT_CODE
fi

echo "[*] All tests passed successfully"
