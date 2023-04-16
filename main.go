package main

import (
	"context"
	"log"

	"github.com/awlsring/terraform-provider-headscale/headscale"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate
func main() {
	err := providerserver.Serve(context.Background(), headscale.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/awlsring/headscale",
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
