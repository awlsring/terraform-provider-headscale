# Returns data about the host with ID 1
data "headscale_device" "by_id" {
    id = "1"
}

# Returns data about the host with name "my-node"
data "headscale_device" "by_name" {
    name = "my-node"
}

# Returns data about the host with a given name of "my-node"
data "headscale_device" "by_given_name" {
    given_name = "my-node"
}