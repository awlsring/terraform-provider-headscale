package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_PolicyDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
					resource "headscale_policy" "test" {
						policy = jsonencode({
							"tagOwners": {
								"tag:example": ["user1@"],
							},
							"acls": [
								{
									"action": "accept",
									"src": ["tag:example"],
									"dst": ["*:*"],
								},
							],
						})
					}

					data "headscale_policy" "test" {
						depends_on = [headscale_policy.test]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_policy.test", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_policy.test", "policy"),
					resource.TestCheckResourceAttrSet("data.headscale_policy.test", "updated"),
				),
			},
		},
	})
}

func Test_PolicyDataSource_EmptyPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
					resource "headscale_policy" "test_empty" {
						policy = ""
					}

					data "headscale_policy" "test_empty" {
						depends_on = [headscale_policy.test_empty]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_policy.test_empty", "id"),
					resource.TestCheckResourceAttr("data.headscale_policy.test_empty", "policy", ""),
					resource.TestCheckResourceAttrSet("data.headscale_policy.test_empty", "updated"),
				),
			},
		},
	})
}

func Test_PolicyDataSource_ComplexPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
					resource "headscale_policy" "test_complex" {
						policy = jsonencode({
							"groups": {
								"group:admin": ["user1@", "user2@"],
								"group:dev": ["user3@", "user4@"],
							},
							"tagOwners": {
								"tag:prod": ["group:admin"],
								"tag:dev": ["group:dev"],
							},
							"acls": [
								{
									"action": "accept",
									"src": ["group:admin"],
									"dst": ["*:*"],
								},
								{
									"action": "accept",
									"src": ["group:dev"],
									"dst": ["tag:dev:*"],
								},
							],
							"ssh": [
								{
									"action": "accept",
									"src": ["group:admin"],
									"dst": ["*"],
									"users": ["root", "ubuntu"],
								},
							],
						})
					}

					data "headscale_policy" "test_complex" {
						depends_on = [headscale_policy.test_complex]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_policy.test_complex", "id"),
					resource.TestCheckResourceAttrSet("data.headscale_policy.test_complex", "policy"),
					resource.TestCheckResourceAttrSet("data.headscale_policy.test_complex", "updated"),
				),
			},
		},
	})
}

func Test_PolicyDataSource_Standalone(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `
					data "headscale_policy" "test_standalone" {}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.headscale_policy.test_standalone", "id"),
					resource.TestCheckResourceAttr("data.headscale_policy.test_standalone", "policy", ""),
					resource.TestCheckResourceAttrSet("data.headscale_policy.test_standalone", "updated"),
				),
			},
		},
	})
}
