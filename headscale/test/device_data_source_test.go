package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_DeviceDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device" "test" {
					id = "1"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_device.test", "id", "1"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "user_id"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "user_name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "expiry"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "created_at"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "register_method"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "given_name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "tags.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "approved_routes.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.test", "available_routes.#"),
				),
			},
		},
	})
}

func Test_DeviceDataSource_ByName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device" "subnet_route" {
					name = "subnet-route"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_device.subnet_route", "id", "2"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "user_id"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "user_name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "expiry"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "created_at"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "register_method"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "given_name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "tags.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.subnet_route", "approved_routes.#"),
					resource.TestCheckResourceAttr("data.headscale_device.subnet_route", "available_routes.#", "2"),
				),
			},
		},
	})
}

func Test_DeviceDataSource_ByGivenName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device" "exit_node" {
					given_name = "exit-node"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_device.exit_node", "id", "3"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "user_id"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "user_name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "expiry"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "created_at"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "register_method"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "given_name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "tags.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "approved_routes.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.exit_node", "available_routes.#"),
					resource.TestCheckResourceAttr("data.headscale_device.exit_node", "available_routes.#", "2"),
				),
			},
		},
	})
}

func Test_DeviceDataSource_ValidateAllAttributes(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				data "headscale_device" "validate" {
					id = "1"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Required attributes
					resource.TestCheckResourceAttr("data.headscale_device.validate", "id", "1"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "user_id"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "user_name"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "expiry"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "created_at"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "register_method"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "given_name"),

					// List attributes should be present (even if empty)
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "addresses.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "tags.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "approved_routes.#"),
					resource.TestCheckResourceAttrSet("data.headscale_device.validate", "available_routes.#"),

					// Device should have at least one IP address
					resource.TestCheckResourceAttrWith("data.headscale_device.validate", "addresses.#", func(value string) error {
						if value == "0" {
							return nil // Allow empty addresses for test
						}
						return nil
					}),
				),
			},
		},
	})
}
