# Returns routes for device 1
data "headscale_device_subnet_routes" "device_1" {
    device_id = "1"
}

# Returns routes for a specific device filtered by status
data "headscale_device_subnet_routes" "disabled_devices" {
    device_id = "2"
    status = "disabled"
}

# Returns enabled routes for a device
data "headscale_device_subnet_routes" "enabled_devices" {
    device_id = "2"
    status = "enabled"
}