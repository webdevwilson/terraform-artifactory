package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/kieferrj/terraform-provider-artifactory/builtin/providers/artifactory"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: artifactory.Provider,
	}
	plugin.Serve(&opts)
}
