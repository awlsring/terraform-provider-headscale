# Returns all devices
data "headscale_devices" "all" {}

# Returns all devices that belong to a specified user by ID
data "headscale_devices" "user_devices" {
    user_id = "1"
}

# Returns all devices that belong to a specified user by name
data "headscale_devices" "user_devices_by_name" {
    user_name = "terraform"
}

# Returns all devices with a specified name prefix
data "headscale_devices" "prefixed_devices" {
    name_prefix = "domain"
}

# Returns devices filtered by both user and name prefix
data "headscale_devices" "filtered_combined" {
    user_name = "terraform"
    name_prefix = "subnet"
}