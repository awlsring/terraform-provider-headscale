package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_PreAuthKeyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `data "headscale_user" "test_user" {
			  name = "terraform"
			}

			data "headscale_pre_auth_keys" "test" {
			  user = data.headscale_user.test_user.id
			}
			`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(
						"data.headscale_pre_auth_keys.test", "user",
						"data.headscale_user.test_user", "id",
					),
				),
			},
		},
	})
}
