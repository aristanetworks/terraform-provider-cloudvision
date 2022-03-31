// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func configlet_cv() *schema.Resource {
	return &schema.Resource{
		Create: configlet_cv_Create,
		Read:   configlet_cv_Read,
		Update: configlet_cv_Update,
		Delete: configlet_cv_Delete,

		Schema: map[string]*schema.Schema{
			"configletname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				Optional: false,
			},
			"configletdata": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				Optional: false,
			},
			"configletkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
		},
	}
}

func configlet_cv_Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)
	//Create the configlet.
	AddConfiglet, err := c.API.AddConfiglet(d.Get("configletname").(string), d.Get("configletdata").(string))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Adding Configlet ", AddConfiglet)

	//Set the ID
	d.SetId("cvtf-configlet-" + strconv.FormatInt(time.Now().Unix(), 10))

	//Set the configlet up
	//Check for the configlet.

	return resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		Checkconfiglet, err := c.API.GetConfigletByName(d.Get("configletname").(string))
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error Adding configlet%v\n", err))
		}

		if Checkconfiglet.Name != d.Get("configletname") {
			return resource.RetryableError(fmt.Errorf("configlet did not return the same name"))
		}
		err = configlet_cv_Read(d, m)
		if err != nil {
			return resource.NonRetryableError(err)
		} else {
			return nil
		}
	})
}

func configlet_cv_Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	//Need to read the current Key which is place that is needed in the event of deletion.
	configletKey, err := c.API.GetConfigletByName(d.Get("configletname").(string))
	if err != nil {
		log.Fatal(err)
	}
	d.Set("configletkey", configletKey.Key)

	return nil
}

func configlet_cv_Update(d *schema.ResourceData, m interface{}) error {

	c := m.(*client.CvpClient)

	UpdateConfiglet := c.API.UpdateConfiglet(d.Get("configletdata").(string), d.Get("configletname").(string), d.Get("configletkey").(string))
	log.Print(UpdateConfiglet)

	return configlet_cv_Read(d, m)

}

func configlet_cv_Delete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	ConTainerRemove := c.API.DeleteConfiglet(d.Get("configletname").(string), d.Get("configletkey").(string))
	log.Print(ConTainerRemove)

	return nil
}
