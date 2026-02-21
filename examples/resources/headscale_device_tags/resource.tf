# Tags must be permitted in policy before assigning them to devices
resource "headscale_policy" "tag_owners" {
  policy = jsonencode({
    "tagOwners" : {
      "tag:terraform" : ["autogroup:admin"],
      "tag:server" : ["autogroup:admin"],
      "tag:production" : ["autogroup:admin"]
    }
  })
}

# Sets the tag `tag:terraform` to the device with ID 1
resource "headscale_device_tags" "device_1_tags" {
    depends_on = [headscale_policy.tag_owners]
    device_id = "1"
    tags = [ "tag:terraform" ]
}

# Applies multiple tags to a device
resource "headscale_device_tags" "device_tags" {
    depends_on = [headscale_policy.tag_owners]
    device_id = "1"
    tags = ["tag:server", "tag:production"]
}
