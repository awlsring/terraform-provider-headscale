# Returns all advertised routes
data "headscale_subnet_routes" "all" {}

# Returns all advertised routes that belong to device 1
data "headscale_subnet_routes" "device_1_routes" {
    device_id = "1"
}

# Returns all advertised routes that are disabled
data "headscale_subnet_routes" "disabled_routes" {
    status = "disabled"
}

# Returns all advertised routes that are enabled
data "headscale_subnet_routes" "enabled_routes" {
    status = "enabled"
}

# Returns all advertised routes that are enabled on device 1
data "headscale_subnet_routes" "device_1_enabled_routes" {
    status = "enabled"
    device_id = "1"
}
