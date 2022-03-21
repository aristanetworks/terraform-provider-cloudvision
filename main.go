package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	//"tf-cvp/cvprovider"
	cvprovider "github.com/aristanetworks/terraform-provider-cloudvision/cloudvision"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return cvprovider.Provider()
		},
	})
}
