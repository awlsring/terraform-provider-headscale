# Returns all devices
data "headscale_devices" "all" {}

# Returns all devices that belong to a specified user
data "headscale_devices" "user_devices" {
    user_id = "1"
}

# Returns all devices with a specified name prefix
data "headscale_devices" "user_devices" {
    name_prefix = "domain"
}