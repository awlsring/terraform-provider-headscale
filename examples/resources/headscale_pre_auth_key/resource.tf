# A pre auth key that expires in the default 1 hour
resource "headscale_pre_auth_key" "default" {
    user = "awlsring"
}

# A pre auth key that expires in a week
resource "headscale_pre_auth_key" "week" {
    user = "awlsring"
    time_to_expire = "1w"
}

# A pre auth key that is reusable
resource "headscale_pre_auth_key" "reusable" {
    user = "awlsring"
    reusable = true
}

# A pre auth key that is ephemeral with tags
resource "headscale_pre_auth_key" "tags" {
    user = "awlsring"
    ephemeral = true
    acl_tags = ["tag:test"]
}
