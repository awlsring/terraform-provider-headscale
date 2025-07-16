package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_RouteDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `data "headscale_subnet_routes" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.test", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.test", "routes.#"),
				),
			},
		},
	})
}

func Test_RouteDataSource_FilterByStatus(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `data "headscale_subnet_routes" "disabled" {
					status = "disabled"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.disabled", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.disabled", "routes.#"),
					resource.TestCheckResourceAttr("data.headscale_subnet_routes.disabled", "routes.0.enabled", "false"),
				),
			},
			{
				Config: ProviderConfig + `
				resource "headscale_device_subnet_routes" "enable_routes" {
					device_id = "2"
					routes = ["10.0.10.0/24"]
				}

				data "headscale_subnet_routes" "enabled" {
					status = "enabled"
					depends_on = [headscale_device_subnet_routes.enable_routes]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.enabled", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.enabled", "routes.#"),
				),
			},
		},
	})
}

func Test_RouteDataSource_FilterByDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_subnet_routes" "device_routes" {
					device_id = "2"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.device_routes", "id"),
					resource.TestCheckResourceAttr("data.headscale_subnet_routes.device_routes", "device_id", "2"),
				),
			},
		},
	})
}

func Test_RouteDataSource_FilterByDeviceAndStatus(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_subnet_routes" "device_disabled_routes" {
					device_id = "2"
					status = "disabled"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.device_disabled_routes", "id"),
					resource.TestCheckResourceAttr("data.headscale_subnet_routes.device_disabled_routes", "device_id", "2"),
				),
			},
		},
	})
}

func Test_RouteDataSource_ValidateRouteAttributes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `data "headscale_subnet_routes" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_subnet_routes.test", "id"),
				),
			},
		},
	})
}
