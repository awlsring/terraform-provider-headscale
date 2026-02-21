package test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func Test_PreAuthKeyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "1"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "false"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					reusable = true
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "1"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "true"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "false"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					ephemeral = true
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "1"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "true"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					acl_tags = ["tag:terraform", "tag:terra-form", "tag:terra_form"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "1"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "reusable", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "ephemeral", "false"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "acl_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terraform"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terra-form"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terra_form"),
				),
			},
		},
	})
}

func Test_PreAuthKeyResource_Issue24_ACLTagsReorderedNoReplacement(t *testing.T) {
	var firstID string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					acl_tags = ["tag:terraform", "tag:terra-form", "tag:terra_form"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrCapturedPreAuth("headscale_pre_auth_key.test", "id", &firstID),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "acl_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terraform"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terra-form"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terra_form"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					acl_tags = ["tag:terra_form", "tag:terraform", "tag:terra-form"]
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrUnchangedPreAuth("headscale_pre_auth_key.test", "id", &firstID),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "acl_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terraform"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terra-form"),
					resource.TestCheckTypeSetElemAttr("headscale_pre_auth_key.test", "acl_tags.*", "tag:terra_form"),
				),
			},
		},
	})
}

func Test_PreAuthKeyResource_TimeToExpireReplacement(t *testing.T) {
	var firstID string

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					time_to_expire = "1h"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrCapturedPreAuth("headscale_pre_auth_key.test", "id", &firstID),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "1"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "time_to_expire", "1h"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "key"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "expiration"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "created_at"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					time_to_expire = "2h"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckResourceAttrChangedPreAuth("headscale_pre_auth_key.test", "id", &firstID),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "user", "1"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "time_to_expire", "2h"),
					resource.TestCheckResourceAttr("headscale_pre_auth_key.test", "expired", "false"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "key"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "expiration"),
					resource.TestCheckResourceAttrSet("headscale_pre_auth_key.test", "created_at"),
				),
			},
		},
	})
}

func Test_PreAuthKeyResource_InvalidDuration(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "invalid" {
					user = "1"
					time_to_expire = "banana"
				}`,
				ExpectError: regexp.MustCompile("must be a valid duration string"),
			},
		},
	})
}

func Test_PreAuthKeyResource_InvalidUserID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "invalid" {
					user = "not-a-number"
				}`,
				ExpectError: regexp.MustCompile("must be a valid user ID"),
			},
		},
	})
}

func testAccCheckResourceAttrCapturedPreAuth(resourceName string, attrName string, value *string) resource.TestCheckFunc {
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

func testAccCheckResourceAttrUnchangedPreAuth(resourceName string, attrName string, previous *string) resource.TestCheckFunc {
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
		if current != *previous {
			return fmt.Errorf("expected %s to stay the same for resource %s, got %q (previous %q)", attrName, resourceName, current, *previous)
		}

		return nil
	}
}

func testAccCheckResourceAttrChangedPreAuth(resourceName string, attrName string, previous *string) resource.TestCheckFunc {
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
