package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_DeviceTagging(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_device_tags" "test" {
					device_id = 1
					tags = ["tag:terraform"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_device_tags.test", "device_id", "1"),
				),
			},
		},
	})
}
