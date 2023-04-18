# Enables the route with prefix "10.0.10.0/24" advertised by device 1
resource "headscale_device_subnet_routes" "device_1" {
    device_id = "1"
    routes = [ "10.0.10.0/24" ]
}