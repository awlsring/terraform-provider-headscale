# Return only active pre auth keys for user awlsring
data "headscale_pre_auth_keys" "active_keys" {
    user = "awlsring"
}

# Return all pre auth keys for user awlsring
data "headscale_pre_auth_keys" "all_keys" {
    all = true
    user = awlsring
}