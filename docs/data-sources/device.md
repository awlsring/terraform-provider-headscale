---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "headscale_device Data Source - terraform-provider-headscale"
subcategory: ""
description: |-
  The device data source allows you to get information about a device registered in Headscale instance.
---

# headscale_device (Data Source)

The device data source allows you to get information about a device registered in Headscale instance.

## Example Usage

```terraform
# Returns data about the host with ID 1
data "headscale_device" "first_host" {
    id = "1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the device

### Optional

- `register_method` (String) The method used to register the device.

### Read-Only

- `addresses` (List of String) List of the device's ip addresses.
- `created_at` (String) The time the device entry was created.
- `expiry` (String) The expiry date of the device.
- `given_name` (String) The device's given name.
- `name` (String) The device's name.
- `tags` (List of String) The tags applied to the device.
- `user` (String) The ID of the user who owns the device.

