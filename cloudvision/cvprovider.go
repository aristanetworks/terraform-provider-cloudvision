// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"log"
	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider info
func Provider() *schema.Provider {
	c := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cvp": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"token": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				Computed: false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cvprovider_onboard_inventory":   onboard_cvinvntory(),
			"cvprovider_cv_container":        container_cv(),
			"cvprovider_cv_configlet":        configlet_cv(),
			"cvprovider_device_cv_container": device_cv_container(),
			"cvprovider_device_cv_configlet": device_cv_configlet(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cvprovider_data_inventory": get_inventory(),
			"cvprovider_data_user":      get_user(),
		},
	}
	c.ConfigureFunc = providerConfigure
	return c

}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	hosts := d.Get("cvp").(string)
	port := d.Get("port").(int)
	token := d.Get("token").(string)
	cvps := []string{hosts}
	cvpClient, _ := client.NewCvpClient(
		client.Protocol("https"),
		client.Port(port),
		client.Hosts(cvps...),
		client.Debug(false))

	if err := cvpClient.ConnectWithToken(token); err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return cvpClient, nil
}
