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
				Config: ProviderConfig + `resource "headscale_user" "test" {
					name = "test"
				  }
				
				  data "headscale_user" "test" {
					name = headscale_user.test.name
					force_delete = false
				  }
				  `,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.headscale_user.test", "name", "test"),
				),
			},
		},
	})
}
