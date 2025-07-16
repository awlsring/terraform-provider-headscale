package test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_DeviceRouteResource_BasicWithKnownDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = ["10.0.10.0/24"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.#", "1"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.0", "10.0.10.0/24"),
				),
			},
		},
	})
}

func Test_DeviceRouteResource_MultipleRoutes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = ["10.0.10.0/24", "192.168.1.0/24"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.#", "2"),
					resource.TestCheckTypeSetElemAttr("headscale_device_subnet_routes.test", "routes.*", "10.0.10.0/24"),
					resource.TestCheckTypeSetElemAttr("headscale_device_subnet_routes.test", "routes.*", "192.168.1.0/24"),
				),
			},
		},
	})
}

func Test_DeviceRouteResource_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = ["10.0.10.0/24"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.#", "1"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.0", "10.0.10.0/24"),
				),
			},
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = ["10.0.10.0/24", "192.168.1.0/24"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.#", "2"),
					resource.TestCheckTypeSetElemAttr("headscale_device_subnet_routes.test", "routes.*", "10.0.10.0/24"),
					resource.TestCheckTypeSetElemAttr("headscale_device_subnet_routes.test", "routes.*", "192.168.1.0/24"),
				),
			},
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = ["192.168.1.0/24"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.#", "1"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.0", "192.168.1.0/24"),
				),
			},
		},
	})
}

func Test_DeviceRouteResource_EmptyRoutes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = []
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.#", "0"),
				),
			},
		},
	})
}

func Test_DeviceRouteResource_ImportState(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = ["10.0.10.0/24"]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_device_subnet_routes.test", "id"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "device_id", "2"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.#", "1"),
					resource.TestCheckResourceAttr("headscale_device_subnet_routes.test", "routes.0", "10.0.10.0/24"),
				),
			},
			{
				ResourceName:      "headscale_device_subnet_routes.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func Test_DeviceRouteResource_ValidateRouteFormat(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				# Use known device ID 2 (subnet-route container)
				resource "headscale_device_subnet_routes" "test" {
					device_id = "2"
					routes = ["invalid-route-format"]
				}`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Value Match`),
			},
		},
	})
}
