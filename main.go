package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/mesomorphic/terraform-provider-confluence/confluence"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: confluence.Provider})
}
