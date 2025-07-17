# Return only active pre auth keys for user with ID of 1
data "headscale_pre_auth_keys" "active_keys" {
    user = "1"
}

# Return all pre auth keys for user with ID of 1
data "headscale_pre_auth_keys" "all_keys" {
    all = true
    user = "1"
}