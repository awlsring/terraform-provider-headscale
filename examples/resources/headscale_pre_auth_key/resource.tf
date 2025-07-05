# A pre auth key that expires in the default 1 hour
resource "headscale_user" "terraform" {
  name = "terraform"
}

resource "headscale_pre_auth_key" "default" {
  user = headscale_user.terraform.id
}

# A pre auth key that expires in a week
resource "headscale_pre_auth_key" "week" {
  user           = headscale_user.terraform.id
  time_to_expire = "1w"
}

# A pre auth key that is reusable
resource "headscale_pre_auth_key" "reusable" {
  user     = headscale_user.terraform.id
  reusable = true
}

# A pre auth key that is ephemeral with tags
resource "headscale_pre_auth_key" "tags" {
  user      = headscale_user.terraform.id
  ephemeral = true
  acl_tags  = ["tag:test"]
}
