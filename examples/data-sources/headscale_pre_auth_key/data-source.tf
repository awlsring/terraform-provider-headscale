# Return only active pre auth keys for the user
resource "headscale_user" "test" {
  name = "test"
}

data "headscale_pre_auth_keys" "active_keys" {
  user = headscale_user.test.id
}

# Return all pre auth keys for the user, including expired
data "headscale_pre_auth_keys" "all_keys" {
  user = headscale_user.test.id
  all  = true
}