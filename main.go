package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	improvmxTerraform "github.com/issyl0/terraform-provider-improvmx/improvmx"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return improvmxTerraform.Provider()
		},
	})
}
