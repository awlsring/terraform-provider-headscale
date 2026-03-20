# A pre auth key that expires in the default 1 hour
resource "headscale_pre_auth_key" "default" {
    user = "1"
}

# A pre auth key that expires in a week
resource "headscale_pre_auth_key" "week" {
    user = "1"
    time_to_expire = "1w"
}

# A pre auth key that is reusable
resource "headscale_pre_auth_key" "reusable" {
    user = "1"
    reusable = true
}

# A pre auth key that is ephemeral with tags
resource "headscale_pre_auth_key" "tags" {
    user = "1"
    ephemeral = true
    acl_tags = ["tag:test"]
}

# Creates a tags-only pre-auth key for infrastructure nodes on Headscale v0.28.0+
resource "headscale_pre_auth_key" "with_tags" {
    reusable = true
    acl_tags = ["tag:server", "tag:production"]
}
