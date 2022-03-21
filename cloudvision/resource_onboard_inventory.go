// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func onboard_cvinvntory() *schema.Resource {
	return &schema.Resource{
		Create: onboard_cvinvntoryCreate,
		Read:   onboard_cvinvntoryRead,
		Update: onboard_cvinvntoryUpdate,
		Delete: onboard_cvinvntoryDelete,

		Schema: map[string]*schema.Schema{
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"typeofdev": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"requestid": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"decomid": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}
func onboard_cvinvntoryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	OnboardDev, err := c.API.OnboardDevice(d.Get("ipaddress").(string), d.Get("typeofdev").(string))
	if err != nil {
		return err
	}
	d.Set("requestid", OnboardDev.Value.Key.RequestID)
	d.SetId("cvtf-" + strconv.FormatInt(time.Now().Unix(), 10))
	return resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		statusdev, err := c.API.OnboardStatus(d.Get("requestid").(string))
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error Adding device", err))
		}
		if statusdev != "ONBOARDING_STATUS_SUCCESS" {
			return resource.RetryableError(fmt.Errorf("Device not ready yet ", statusdev))
		}

		err = onboard_cvinvntoryRead(d, m)
		if err != nil {
			return resource.NonRetryableError(err)
		} else {
			return nil
		}
	})

}

func onboard_cvinvntoryRead(d *schema.ResourceData, m interface{}) error {

	c := m.(*client.CvpClient)
	inv, err := c.API.GetDeviceByName(d.Get("fqdn").(string))
	if err != nil {
		return err
	}
	d.Set("serial", inv.SerialNumber)
	return nil
}

func onboard_cvinvntoryUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func onboard_cvinvntoryDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	DecomDev, err := c.API.DecomDevice(d.Get("serial").(string))
	if err != nil {
		return err
	}
	_ = DecomDev

	d.SetId("")

	return nil

}
