// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"log"
	"strconv"
	"time"

	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func get_inventory() *schema.Resource {
	return &schema.Resource{
		Read: data_read_inventory,
		Schema: map[string]*schema.Schema{
			"inventory": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipaddress": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"modelname": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"systemmacaddress": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"deviceinfo": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"devicestatus": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"parentcontainerkey": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"serialnumber": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: false,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func data_read_inventory(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	d.SetId("cvtf-datagetinv-" + strconv.FormatInt(time.Now().Unix(), 10))

	inv, err := c.API.GetInventory()
	if err != nil {
		log.Fatal(err)
	}

	netelem := make([]map[string]interface{}, len(inv))
	for inv, elem := range inv {
		netelem[inv] = map[string]interface{}{
			"ipaddress":          elem.IPAddress,
			"modelname":          elem.ModelName,
			"systemmacaddress":   elem.SystemMacAddress,
			"hostname":           elem.Hostname,
			"fqdn":               elem.Fqdn,
			"version":            elem.Version,
			"deviceinfo":         elem.DeviceInfo,
			"devicestatus":       elem.DeviceStatus,
			"parentcontainerkey": elem.ParentContainerKey,
			"serialnumber":       elem.SerialNumber,
			"type":               elem.Type,
		}
	}
	d.Set("inventory", netelem)

	return nil
}
