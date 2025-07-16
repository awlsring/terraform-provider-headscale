package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func Test_PolicyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_policy" "test" {
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
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_policy.test", "id"),
					resource.TestCheckResourceAttrSet("headscale_policy.test", "policy"),
					resource.TestCheckResourceAttrSet("headscale_policy.test", "updated"),
				),
			},
		},
	})
}

func Test_PolicyResource_BasicPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_policy" "test_basic" {
					policy = jsonencode({"acls":[{"action":"accept"}]})
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("headscale_policy.test_basic", "policy", `{"acls":[{"action":"accept"}]}`),
					resource.TestCheckResourceAttrSet("headscale_policy.test_basic", "id"),
					resource.TestCheckResourceAttrSet("headscale_policy.test_basic", "updated"),
				),
			},
		},
	})
}

func Test_PolicyResource_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_policy" "test_update" {
					policy = jsonencode({
						"tagOwners": {
							"tag:test": ["user1@"],
						},
						"acls": [
							{
								"action": "accept",
								"src": ["tag:test"],
								"dst": ["*:22"],
							},
						],
					})
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_policy.test_update", "id"),
					resource.TestCheckResourceAttrSet("headscale_policy.test_update", "policy"),
					resource.TestCheckResourceAttrSet("headscale_policy.test_update", "updated"),
				),
			},
			{
				Config: ProviderConfig + `resource "headscale_policy" "test_update" {
					policy = jsonencode({
						"tagOwners": {
							"tag:test": ["user1@", "user2@"],
						},
						"acls": [
							{
								"action": "accept",
								"src": ["tag:test"],
								"dst": ["*:22,80,443"],
							},
						],
					})
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_policy.test_update", "id"),
					resource.TestCheckResourceAttrSet("headscale_policy.test_update", "policy"),
					resource.TestCheckResourceAttrSet("headscale_policy.test_update", "updated"),
				),
			},
		},
	})
}

func Test_PolicyResource_ComplexPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + `resource "headscale_policy" "test_complex" {
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
							{
								"action": "accept",
								"src": ["tag:prod"],
								"dst": ["*:443"],
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
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("headscale_policy.test_complex", "id"),
					resource.TestCheckResourceAttrSet("headscale_policy.test_complex", "policy"),
					resource.TestCheckResourceAttrSet("headscale_policy.test_complex", "updated"),
				),
			},
		},
	})
}
