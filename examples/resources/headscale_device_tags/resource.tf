# Sets the tag `tag:terraform` to the device with ID 1
resource "headscale_device_tags" "device_1_tags" {
    device_id = "1"
    tags = [ "tag:terraform" ]
}

# Applies multiple tags to a device
resource "headscale_device_tags" "device_tags" {
    device_id = "1"
    tags = ["tag:server", "tag:production"]
}
