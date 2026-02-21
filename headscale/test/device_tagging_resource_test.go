package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_DeviceTaggingResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_policy" "tags" {
					policy = jsonencode({
						"tagOwners": {
							"tag:terraform": ["terraform@"],
							"tag:terraform:tests": ["terraform@"],
						},
					})
				  }

				  resource "headscale_device_tags" "test" {
					depends_on = [headscale_policy.tags]
					device_id = 1
					tags = ["tag:terraform"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_device_tags.test", "device_id", "1"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_policy" "tags" {
					policy = jsonencode({
						"tagOwners": {
							"tag:terraform": ["terraform@"],
							"tag:terraform:tests": ["terraform@"],
						},
					})
				  }

				  resource "headscale_device_tags" "test" {
					depends_on = [headscale_policy.tags]
					device_id = 1
					tags = ["tag:terraform:tests"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_device_tags.test", "device_id", "1"),
				),
			},
		},
	})
}
