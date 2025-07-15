package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_UserDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_data_source" {
					name = "test_data_source"
				  }

				  data "headscale_user" "test_data_source" {
					name = headscale_user.test_data_source.name
				  }
				  `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source", "name", "test_data_source"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source", "created_at"),
				),
			},
		},
	})
}

func Test_UserDataSource_ById(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_data_source_by_id" {
					name = "test_data_source_by_id"
				  }

				  data "headscale_user" "test_data_source_by_id" {
					id = headscale_user.test_data_source_by_id.id
				  }
				  `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source_by_id", "name", "test_data_source_by_id"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source_by_id", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source_by_id", "created_at"),
				),
			},
		},
	})
}

func Test_UserDataSource_WithOptionalFields(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_data_source_full" {
					name = "test_data_source_full"
					display_name = "Test Data Source User"
					email = "test-datasource@example.com"
					profile_picture_url = "https://example.com/datasource-avatar.jpg"
				  }

				  data "headscale_user" "test_data_source_full" {
					name = headscale_user.test_data_source_full.name
				  }
				  `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source_full", "name", "test_data_source_full"),
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source_full", "display_name", "Test Data Source User"),
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source_full", "email", "test-datasource@example.com"),
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source_full", "profile_picture_url", "https://example.com/datasource-avatar.jpg"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source_full", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source_full", "created_at"),
				),
			},
		},
	})
}

func Test_UserDataSource_ByEmail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_data_source_by_email" {
					name = "test_data_source_by_email"
					email = "test-by-email@example.com"
				  }

				  data "headscale_user" "test_data_source_by_email" {
					email = headscale_user.test_data_source_by_email.email
				  }
				  `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source_by_email", "name", "test_data_source_by_email"),
					resource.TestCheckResourceAttr("data.headscale_user.test_data_source_by_email", "email", "test-by-email@example.com"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source_by_email", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_user.test_data_source_by_email", "created_at"),
				),
			},
		},
	})
}

func Test_UsersDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_user" "test_users_1" {
					name = "test_users_1"
					display_name = "Test User 1"
					email = "test1@example.com"
				  }

				  resource "headscale_user" "test_users_2" {
					name = "test_users_2"
					display_name = "Test User 2"
					email = "test2@example.com"
				  }

				  data "headscale_users" "test_users_data_source" {
					depends_on = [headscale_user.test_users_1, headscale_user.test_users_2]
				  }
				  `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_users.test_users_data_source", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_users.test_users_data_source", "users.#"),
					resource.TestCheckTypeSetElemNestedAttrs("data.headscale_users.test_users_data_source", "users.*", map[string]string{
						"name":         "test_users_1",
						"display_name": "Test User 1",
						"email":        "test1@example.com",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("data.headscale_users.test_users_data_source", "users.*", map[string]string{
						"name":         "test_users_2",
						"display_name": "Test User 2",
						"email":        "test2@example.com",
					}),
				),
			},
		},
	})
}
