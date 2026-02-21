package test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func Test_PreAuthKeyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					reusable = true
					ephemeral = false
					time_to_expire = "24h"
					acl_tags = ["tag:preauth-ds", "tag:terraform"]
				  }

				  data "headscale_pre_auth_keys" "test" {
					depends_on = [headscale_pre_auth_key.test]
					user = "1"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_pre_auth_keys.test", "user", "1"),
					testAccCheckPreAuthKeyDataSourceContainsTaggedKey("data.headscale_pre_auth_keys.test", "tag:preauth-ds"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_pre_auth_key" "test" {
					user = "1"
					reusable = true
					ephemeral = false
					time_to_expire = "24h"
					acl_tags = ["tag:preauth-ds", "tag:terraform"]
				  }

				  data "headscale_pre_auth_keys" "test" {
					depends_on = [headscale_pre_auth_key.test]
					user = "1"
					all = true
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_pre_auth_keys.test", "user", "1"),
					testAccCheckPreAuthKeyDataSourceContainsTaggedKey("data.headscale_pre_auth_keys.test", "tag:preauth-ds"),
				),
			},
		},
	})
}

func Test_PreAuthKeyDataSource_InvalidUserID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `data "headscale_pre_auth_keys" "invalid" {
					user = "not-a-number"
				  }`,
				ExpectError: regexp.MustCompile("must be a valid user ID"),
			},
		},
	})
}

func testAccCheckPreAuthKeyDataSourceContainsTaggedKey(resourceName string, tag string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found in state", resourceName)
		}

		countRaw, ok := rs.Primary.Attributes["pre_auth_keys.#"]
		if !ok {
			return fmt.Errorf("attribute pre_auth_keys.# missing on resource %s", resourceName)
		}

		count, err := strconv.Atoi(countRaw)
		if err != nil {
			return fmt.Errorf("unable to parse pre_auth_keys.# value %q: %w", countRaw, err)
		}
		if count == 0 {
			return fmt.Errorf("expected at least one pre-auth key in %s", resourceName)
		}

		for i := 0; i < count; i++ {
			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.user", i)] != "1" {
				continue
			}

			tagCountRaw := rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.acl_tags.#", i)]
			tagCount, err := strconv.Atoi(tagCountRaw)
			if err != nil {
				continue
			}

			foundTag := false
			for j := 0; j < tagCount; j++ {
				if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.acl_tags.%d", i, j)] == tag {
					foundTag = true
					break
				}
			}
			if !foundTag {
				continue
			}

			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.id", i)] == "" {
				return fmt.Errorf("matched pre-auth key in %s has empty id", resourceName)
			}
			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.key", i)] == "" {
				return fmt.Errorf("matched pre-auth key in %s has empty key", resourceName)
			}
			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.expiration", i)] == "" {
				return fmt.Errorf("matched pre-auth key in %s has empty expiration", resourceName)
			}
			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.created_at", i)] == "" {
				return fmt.Errorf("matched pre-auth key in %s has empty created_at", resourceName)
			}
			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.reusable", i)] != "true" {
				return fmt.Errorf("matched pre-auth key in %s was expected reusable=true", resourceName)
			}
			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.ephemeral", i)] != "false" {
				return fmt.Errorf("matched pre-auth key in %s was expected ephemeral=false", resourceName)
			}
			if rs.Primary.Attributes[fmt.Sprintf("pre_auth_keys.%d.expired", i)] != "false" {
				return fmt.Errorf("matched pre-auth key in %s was expected expired=false", resourceName)
			}

			return nil
		}

		return fmt.Errorf("did not find pre-auth key with tag %q in %s", tag, resourceName)
	}
}
