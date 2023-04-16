package test

import (
	"github.com/awlsring/terraform-provider-headscale/headscale"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	// Requires HEADSCALE_API_KEY and HEADSCALE_ENDPOINT env vars set
	ProviderConfig = `
provider "headscale" {}
`
)

var (
	TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"headscale": providerserver.NewProtocol6WithError(headscale.New()),
	}
)
