# Terraform Provider Headscale

This is a Terraform provider for [Headscale](https://github.com/juanfont/headscale). This provider allows interaction with the Headscale API to manage and gather data on resources.

You can find this provider on the [Terraform Registry](https://registry.terraform.io/providers/awlsring/headscale/latest).

## Versions

As of release 0.2.0, this provider supports Headscale v0.23.0. This release changed the API with backwards incompatible changes. For compatibility with the previous API, use releases like 0.1.x.

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
      version = "0.1.1"
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
