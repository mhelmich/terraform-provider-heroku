package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mhelmich/terraform-provider-heroku/heroku"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: heroku.Provider})
}
