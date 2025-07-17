# Terraform Provider Headscale

This is a Terraform provider for [Headscale](https://github.com/juanfont/headscale). This provider allows interaction with the Headscale API to manage and gather data on resources.

You can find this provider on the [Terraform Registry](https://registry.terraform.io/providers/awlsring/headscale/latest).

## Versions

Various versions of this provider are created as backwards incompatible changes occur in the Headscale API. This provider will primarily only support the latest version of Headscale. To allow for patching bug fixes for previous version, each minor version will have a new branch created. For example, the 0.1.x releases will be on the `v0.1.x` branch, the 0.2.x releases will be on the `v0.2.x` branch, etc.

Here is a table illustrating provider versions and the Headscale versions they support:

| Provider Version                                                             | Headscale Version |
| ---------------------------------------------------------------------------- | ----------------- |
| [0.1.x](https://github.com/awlsring/terraform-provider-headscale/tree/0.1.x) | 0.20.x-0.22.x     |
| [0.2.x](https://github.com/awlsring/terraform-provider-headscale/tree/0.2.x) | 0.23.x-0.24.x     |
| [0.3.x](https://github.com/awlsring/terraform-provider-headscale/tree/0.3.x) | 0.25.x            |
| [0.4.x](https://github.com/awlsring/terraform-provider-headscale/tree/0.4.x) | 0.26.x            |

## Differences between the Tailscale and Headscale Providers

As Headscale has a different API than Tailscale, the functionality of this provider differs from what the [Tailscale provider](https://registry.terraform.io/providers/tailscale/tailscale) offers.

Some data sources and resources may offer similar functionality between these two providers, but for many of these the configuration options and functionality will be different. This provider contains data sources and resources the Tailscale provider does not offer and lacks some that it does.

## Getting Started

To install this provider in your project, you can copy the code snippet below into your project then run `terraform init`.

```hcl
terraform {
  required_providers {
    headscale = {
      source = "awlsring/headscale"
      version = "0.3.0"
    }
  }
}

provider "headscale" {
  api_key = "api_key"
  endpoint = "https://headscale.example.com"
}
```

In the `provider` block you will need to replace `api_key` and `endpoint` with the values for your Headscale instance. These can also be set via the environment variables `HEADSCALE_API_KEY` and `HEADSCALE_ENDPOINT`.

For further details on how to use this provider, please see the documentation in the `/docs` section of this repo or on the [Terraform Registry page](https://registry.terraform.io/providers/awlsring/headscale/latest/docs).
