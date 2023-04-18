# An API Key that doesn't expire
resource "headscale_api_key" "infinite" {}

# An API Key that expires in 1 week
resource "headscale_api_key" "week" {
    days_to_expire = 7
}
