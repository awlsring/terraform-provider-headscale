# An API Key that expires in the default 90 days
resource "headscale_api_key" "default" {}

# An API Key that expires in 1 week
resource "headscale_api_key" "week" {
    time_to_expire = "1w"
}
