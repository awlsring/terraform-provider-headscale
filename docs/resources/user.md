---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "headscale_user Resource - terraform-provider-headscale"
subcategory: ""
description: |-
  The user resource allows you to register a user on the Headscale instance.
---

# headscale_user (Resource)

The user resource allows you to register a user on the Headscale instance.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the user.

### Optional

- `display_name` (String) The display name of the user.
- `email` (String) The email address of the user.
- `force_delete` (Boolean) If the user should be deleted even if it has nodes attached to it. Defaults to `false`.
- `profile_picture_url` (String) The URL of the user's profile picture.

### Read-Only

- `created_at` (String) The time the user was created.
- `id` (String) The user's id.
