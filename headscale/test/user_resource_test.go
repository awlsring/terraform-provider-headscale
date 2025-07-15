package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_UserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test" {
					name = "test"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_user.test", "name", "test"),
					resource.TestCheckResourceAttr("headscale_user.test", "force_delete", "false"),
					resource.TestCheckResourceAttrSet("headscale_user.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_user.test", "created_at"),
				),
			},
		},
	})
}

func Test_UserResource_WithOptionalFields(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_full" {
					name = "test-full"
					display_name = "Test User"
					email = "test@example.com"
					profile_picture_url = "https://example.com/avatar.jpg"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_user.test_full", "name", "test-full"),
					resource.TestCheckResourceAttr("headscale_user.test_full", "display_name", "Test User"),
					resource.TestCheckResourceAttr("headscale_user.test_full", "email", "test@example.com"),
					resource.TestCheckResourceAttr("headscale_user.test_full", "profile_picture_url", "https://example.com/avatar.jpg"),
					resource.TestCheckResourceAttr("headscale_user.test_full", "force_delete", "false"),
					resource.TestCheckResourceAttrSet("headscale_user.test_full", "id"),
					resource.TestCheckResourceAttrSet("headscale_user.test_full", "created_at"),
				),
			},
		},
	})
}

func Test_UserResource_WithForceDelete(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_with_force_delete" {
					name = "test-with-force-delete"
					force_delete = true
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_user.test_with_force_delete", "name", "test-with-force-delete"),
					resource.TestCheckResourceAttr("headscale_user.test_with_force_delete", "force_delete", "true"),
					resource.TestCheckResourceAttrSet("headscale_user.test_with_force_delete", "id"),
					resource.TestCheckResourceAttrSet("headscale_user.test_with_force_delete", "created_at"),
				),
			},
		},
	})
}

func Test_UserResource_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_update" {
					name = "test-update"
					force_delete = false
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_user.test_update", "name", "test-update"),
					resource.TestCheckResourceAttr("headscale_user.test_update", "force_delete", "false"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_user" "test_update" {
					name = "test-update-renamed"
					force_delete = true
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_user.test_update", "name", "test-update-renamed"),
					resource.TestCheckResourceAttr("headscale_user.test_update", "force_delete", "true"),
				),
			},
		},
	})
}

func Test_UserResource_Import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_import" {
					name = "test-import"
				  }`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_user.test_import", "name", "test-import"),
					resource.TestCheckResourceAttrSet("headscale_user.test_import", "id"),
				),
			},
			{
				ResourceName:            "headscale_user.test_import",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_delete"},
			},
		},
	})
}
