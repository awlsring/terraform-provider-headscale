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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the user.

### Read-Only

- `created_at` (String) The time the user was created.
- `id` (String) The user's id.
