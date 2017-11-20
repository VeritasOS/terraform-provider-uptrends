package main

import (
	"github.com/VeritasOS/terraform-provider-uptrends.git/uptrends"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: uptrends.Provider})
}
