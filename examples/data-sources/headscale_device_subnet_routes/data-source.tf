# Returns routes for device 1
data "headscale_device_subnet_routes" "device_1" {
    device_id = "1"
}

# Returns all devices that are disabled
data "headscale_device_subnet_routes" "disabled_devices" {
    user_id = "1"
    status = "disabled"
}

# Returns all devices that are enabled
data "headscale_device_subnet_routes" "enabled_devices" {
    user_id = "1"
    status = "enabled"
}