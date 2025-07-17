# Creates the user terraform
resource "headscale_user" "terraform_user" {
    name = "terraform"
}

# Creates a user with all optional fields
resource "headscale_user" "full_user" {
    name = "full-user"
    display_name = "Full User"
    email = "full@example.com"
    profile_picture_url = "https://example.com/avatar.jpg"
}

# Creates a user with force delete enabled
resource "headscale_user" "force_delete_user" {
    name = "deletable-user"
    force_delete = true
}