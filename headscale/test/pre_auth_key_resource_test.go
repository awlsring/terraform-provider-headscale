package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_PreAuthKeyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "terraform"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "terraform"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "false"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "terraform"
					reusable = true
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "terraform"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "true"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "false"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "terraform"
					ephemeral = true
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "terraform"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "true"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "terraform"
					acl_tags = ["tag:terraform"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "terraform"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "acl_tags.0", "tag:terraform"),
				),
			},
		},
	})
}
