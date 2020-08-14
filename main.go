package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/icaho/terraform-provider-healthchecksio/healthchecksio"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return healthchecksio.Provider()
		},
	})
}
