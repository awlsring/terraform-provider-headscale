package test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func Test_ApiKeyResource(t *testing.T) {
	var firstID string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_api_key" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrCaptured("headscale_api_key.test", "id", &firstID),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "key"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "prefix"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "expiration"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "created_at"),
					resource.TestCheckResourceAttr("headscale_api_key.test", "expired", "false"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_api_key" "test" {
					time_to_expire = "1h"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrChanged("headscale_api_key.test", "id", &firstID),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "key"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "prefix"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "expiration"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "created_at"),
					resource.TestCheckResourceAttr("headscale_api_key.test", "time_to_expire", "1h"),
					resource.TestCheckResourceAttr("headscale_api_key.test", "expired", "false"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_api_key" "test" {
					time_to_expire = "2h"
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrChanged("headscale_api_key.test", "id", &firstID),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "key"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "prefix"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "expiration"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "created_at"),
					resource.TestCheckResourceAttr("headscale_api_key.test", "time_to_expire", "2h"),
					resource.TestCheckResourceAttr("headscale_api_key.test", "expired", "false"),
				),
			},
		},
	})
}

func Test_ApiKeyResource_InvalidDuration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_api_key" "invalid" {
					time_to_expire = "banana"
				}`,
				ExpectError: regexp.MustCompile("must be a valid duration string"),
			},
		},
	})
}

func Test_ApiKeyResource_TimeToExpireCreateBeforeDestroy(t *testing.T) {
	// Specific test for https://github.com/awlsring/terraform-provider-headscale/issues/25
	var firstID string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
				resource "terraform_data" "rotate" {
				  input = "v1"
				}

				resource "headscale_api_key" "test" {
				  time_to_expire = "30d"

				  lifecycle {
					replace_triggered_by  = [terraform_data.rotate]
					create_before_destroy = true
				  }
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrCaptured("headscale_api_key.test", "id", &firstID),
					resource.TestCheckResourceAttr("headscale_api_key.test", "time_to_expire", "30d"),
					resource.TestCheckResourceAttr("headscale_api_key.test", "expired", "false"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "key"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "prefix"),
				),
			},
			{
				Config: ProviderConfig + `
				resource "terraform_data" "rotate" {
				  input = "v2"
				}

				resource "headscale_api_key" "test" {
				  time_to_expire = "30d"

				  lifecycle {
					replace_triggered_by  = [terraform_data.rotate]
					create_before_destroy = true
				  }
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrChanged("headscale_api_key.test", "id", &firstID),
					resource.TestCheckResourceAttr("headscale_api_key.test", "time_to_expire", "30d"),
					resource.TestCheckResourceAttr("headscale_api_key.test", "expired", "false"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "key"),
					resource.TestCheckResourceAttrSet("headscale_api_key.test", "prefix"),
				),
			},
		},
	})
}

func testAccCheckResourceAttrCaptured(resourceName string, attrName string, value *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}

		current, ok := rs.Primary.Attributes[attrName]
		if !ok {
			return fmt.Errorf("attribute %s missing on resource %s", attrName, resourceName)
		}
		if current == "" {
			return fmt.Errorf("attribute %s on resource %s is empty", attrName, resourceName)
		}

		*value = current
		return nil
	}
}

func testAccCheckResourceAttrChanged(resourceName string, attrName string, previous *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}

		current, ok := rs.Primary.Attributes[attrName]
		if !ok {
			return fmt.Errorf("attribute %s missing on resource %s", attrName, resourceName)
		}
		if current == "" {
			return fmt.Errorf("attribute %s on resource %s is empty", attrName, resourceName)
		}
		if *previous == "" {
			return fmt.Errorf("previous value for %s was not captured", attrName)
		}
		if current == *previous {
			return fmt.Errorf("expected %s to change for resource %s, value stayed %q", attrName, resourceName, current)
		}

		*previous = current
		return nil
	}
}
