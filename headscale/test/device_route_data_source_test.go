package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_DeviceRouteDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device_subnet_routes" "test" {
					device_id = "2"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("data.headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttrSet("data.headscale_device_subnet_routes.test", "routes.#"),
				),
			},
		},
	})
}

func Test_DeviceRouteDataSource_FilterByStatus(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device_subnet_routes" "disabled" {
					device_id = "2"
					status = "disabled"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_device_subnet_routes.disabled", "id"),
					resource.TestCheckResourceAttr("data.headscale_device_subnet_routes.disabled", "device_id", "2"),
				),
			},
			{
				Config: ProviderConfig + `
				data "headscale_device_subnet_routes" "enabled" {
					device_id = "2"
					status = "enabled"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_device_subnet_routes.enabled", "id"),
					resource.TestCheckResourceAttr("data.headscale_device_subnet_routes.enabled", "device_id", "2"),
				),
			},
		},
	})
}

func Test_DeviceRouteDataSource_ValidateRouteAttributes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device_subnet_routes" "test" {
					device_id = "2"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("data.headscale_device_subnet_routes.test", "device_id", "2"),
				),
			},
		},
	})
}

func Test_DeviceRouteDataSource_SpecificDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device_subnet_routes" "subnet_device" {
					device_id = "2"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_device_subnet_routes.subnet_device", "id"),
					resource.TestCheckResourceAttr("data.headscale_device_subnet_routes.subnet_device", "device_id", "2"),
				),
			},
		},
	})
}
