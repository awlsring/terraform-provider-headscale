# Sets the tag `tag:terrafrom` to the device with ID 1
resource "headscale_device_tags" "device_1_tags" {
    device_id = "1"
    tags = [ "tag:terrafrom" ]
}