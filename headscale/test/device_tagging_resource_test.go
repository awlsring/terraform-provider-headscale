package test

import (
	"regexp"
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
					resource.TestCheckResourceAttr("headscale_device_tags.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("headscale_device_tags.test", "tags.0", "tag:terraform"),
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
					resource.TestCheckResourceAttr("headscale_device_tags.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("headscale_device_tags.test", "tags.0", "tag:terraform:tests"),
				),
			},
		},
	})
}

func Test_DeviceTaggingResource_Import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_policy" "tags" {
					policy = jsonencode({
						"tagOwners": {
							"tag:terraform": ["terraform@"],
						},
					})
				  }

				  resource "headscale_device_tags" "test_import" {
					depends_on = [headscale_policy.tags]
					device_id = 1
					tags = ["tag:terraform"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_device_tags.test_import", "device_id", "1"),
					resource.TestCheckResourceAttr("headscale_device_tags.test_import", "tags.#", "1"),
					resource.TestCheckResourceAttr("headscale_device_tags.test_import", "tags.0", "tag:terraform"),
				),
			},
			{
				ResourceName:      "headscale_device_tags.test_import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func Test_DeviceTaggingResource_InvalidTag(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_device_tags" "invalid" {
					device_id = 1
					tags = ["terraform"]
				}`,
				ExpectError: regexp.MustCompile("tag must follow scheme"),
			},
		},
	})
}

func Test_DeviceTaggingResource_DuplicateTags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_device_tags" "invalid" {
					device_id = 1
					tags = ["tag:terraform", "tag:terraform"]
				}`,
				ExpectError: regexp.MustCompile("duplicate|unique values"),
			},
		},
	})
}
