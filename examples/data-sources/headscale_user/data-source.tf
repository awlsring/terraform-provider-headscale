# Returns data on the user "awlsring"
data "headscale_user" "awlsring" {
    name = "awlsring"
}

# Returns user by ID
data "headscale_user" "by_id" {
    id = "1"
}

# Returns user by email
data "headscale_user" "by_email" {
    email = "user@example.com"
}

# Returns user by display name
data "headscale_user" "by_display_name" {
    display_name = "John Doe"
}