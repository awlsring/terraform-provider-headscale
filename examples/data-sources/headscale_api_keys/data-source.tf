# Return only active API keys
data "headscale_api_keys" "active_keys" {}

# Return all API keys
data "headscale_api_keys" "all_keys" {
    all = true
}