# An API Key that expires in the default 90 days
resource "headscale_api_key" "default" {}

# An API Key that expires in 1 week
resource "headscale_api_key" "week" {
    time_to_expire = "1w"
}

# Creates an API key with specific expiration date
resource "headscale_api_key" "terraform_key" {
    expiration = "2024-12-31T23:59:59Z"
}

# Creates an API key with no expiration
resource "headscale_api_key" "permanent_key" {}
