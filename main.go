package main

import (
	"github.com/elc-online/terraform-provider-confluence/confluence"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: confluence.Provider})
}
