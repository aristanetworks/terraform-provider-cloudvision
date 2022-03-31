// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"fmt"
	"log"
	"strconv"
	"time"

	cvpapi "github.com/aristanetworks/go-cvprac/v3/api"
	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func device_cv_configlet() *schema.Resource {
	return &schema.Resource{
		Create: device_cv_configletCreate,
		Read:   device_cv_configletRead,
		Update: device_cv_configletUpdate,
		Delete: device_cv_configletDelete,

		Schema: map[string]*schema.Schema{
			"device_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: false,
				Optional: true,
			},
			"device_configlet_base": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: false,
				Optional: true,
			},
			"device_configlet_base_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"device_configlet_base_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"device_configlettaskid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"device_add_configlettaskid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"device_serial": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"overwrite_compliant": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
				Default:  true,
			},
			"additional_configlets": {
				Type:     schema.TypeList,
				Required: false,
				Optional: true,
				Computed: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create:  schema.DefaultTimeout(5 * time.Minute),
			Update:  schema.DefaultTimeout(5 * time.Minute),
			Default: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}
func device_cv_configletCreate(d *schema.ResourceData, m interface{}) error {

	c := m.(*client.CvpClient)

	device := d.Get("device_fqdn").(string)
	baseconfig_configlet_name := device + "-base"
	d.Set("device_configlet_base_name", baseconfig_configlet_name)

	//Set the ID
	d.SetId(d.Get("device_fqdn").(string) + "device-configlet-" + strconv.FormatInt(time.Now().Unix(), 10))

	AddConfiglet, err := c.API.AddConfiglet(baseconfig_configlet_name, d.Get("device_configlet_base").(string))
	if err != nil {
		return err
	}
	_ = AddConfiglet

	Configlets := []cvpapi.Configlet{}

	time.Sleep(20 * time.Second)

	resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		CheckConfiglet, err := c.API.GetConfigletByName(baseconfig_configlet_name)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error cannot find Configlet%v\n", err))
		}
		if CheckConfiglet.Name == baseconfig_configlet_name {
			d.Set("device_configlet_base_key", CheckConfiglet.Key)
			Configlets = append(Configlets, *CheckConfiglet)
		}
		if CheckConfiglet.Name != baseconfig_configlet_name {
			return resource.RetryableError(fmt.Errorf("Configlet not created just yet."))
		}
		if err != nil {
			return resource.NonRetryableError(err)
		} else {
			return nil
		}
	})

	//Set the base configlet to the device.

	Netelem, err := c.API.GetDeviceByName(d.Get("device_fqdn").(string))
	if err != nil {
		return err
	}

	DevApply, err := c.API.ApplyConfigletsToDevice("cvptf", Netelem, true, Configlets...)
	if err != nil {
		return err
	}

	_ = DevApply

	//Set the serial
	d.Set("device_serial", Netelem.SerialNumber)

	time.Sleep(5 * time.Second)
	configlettaskid, err := get_task(m, d)
	if err != nil {
		return err
	}

	d.Set("device_configlettaskid", configlettaskid)
	resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		_ = c.API.ExecuteTask(d.Get("device_configlettaskid").(int))
		gettask, err := c.API.GetTaskByID(d.Get("device_configlettaskid").(int))
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error Getting the taskID%v\n", err))
		}

		if gettask.TaskStatus != "COMPLETED" {
			return resource.RetryableError(fmt.Errorf("Was unable to execute the task."))
		}
		if err != nil {
			return resource.NonRetryableError(err)
		} else {
			return nil
		}

	})

	ConfigletItems := d.Get("additional_configlets").([]interface{})

	if len(ConfigletItems) > 0 {
		AddConfiglets := []cvpapi.Configlet{}
		items := make([]string, len(ConfigletItems))

		for i, raw := range ConfigletItems {
			items[i] = raw.(string)
		}

		for _, cc := range items {
			configletscc, err := c.API.GetConfigletByName(cc)
			if err != nil {
				return err
			}
			AddConfiglets = append(AddConfiglets, *configletscc)
		}

		AdditionalApply, err := c.API.ApplyConfigletsToDevice("cvptf", Netelem, true, AddConfiglets...)
		if err != nil {
			return err
		}
		_ = AdditionalApply

		time.Sleep(5 * time.Second)
		addconfiglettaskid, err := get_task(m, d)
		if err != nil {
			return err
		}

		d.Set("device_add_configlettaskid", addconfiglettaskid)
		return resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
			err = c.API.ExecuteTask(d.Get("device_add_configlettaskid").(int))

			gettask, err := c.API.GetTaskByID(d.Get("device_add_configlettaskid").(int))
			if err != nil {
				return resource.NonRetryableError(fmt.Errorf("Error Getting the taskID%v\n", err))
			}
			log.Print(gettask)

			if gettask.TaskStatus != "COMPLETED" {
				return resource.RetryableError(fmt.Errorf("Was unable to execute the task."))
			}

			if err != nil {
				return resource.NonRetryableError(err)
			} else {
				return nil
			}
		})

	} else {
		d.Set("device_add_configlettaskid", 0)
		return nil
	}
}

func device_cv_configletRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func device_cv_configletUpdate(d *schema.ResourceData, m interface{}) error {

	c := m.(*client.CvpClient)

	compliance := d.Get("overwrite_compliant").(bool)
	ComplianceCheck, err := c.API.GetDeviceByName(d.Get("device_fqdn").(string))
	if err != nil {
		return err
	}

	if compliance == false && ComplianceCheck.ComplianceCode != "0000" {
		return fmt.Errorf("Not able to push to device because it has the overwrite_compliant and it is not compliant it is out of sync.  Either turn this to false or fix the device")
	}

	UpdateConfiglet := c.API.UpdateConfiglet(d.Get("device_configlet_base").(string), d.Get("device_configlet_base_name").(string), d.Get("device_configlet_base_key").(string))
	log.Print(UpdateConfiglet)

	time.Sleep(5 * time.Second)
	configlettaskid, err := get_task(m, d)
	if err != nil {
		return err
	}
	d.Set("device_configlettaskid", configlettaskid)
	return resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		err = c.API.ExecuteTask(d.Get("device_configlettaskid").(int))

		gettask, err := c.API.GetTaskByID(d.Get("device_configlettaskid").(int))
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error Getting the taskID%v\n", err))
		}
		log.Print(gettask)

		if gettask.TaskStatus != "COMPLETED" {
			return resource.RetryableError(fmt.Errorf("Was unable to execute the task."))
		}
		if err != nil {
			return resource.NonRetryableError(err)
		} else {
			return nil
		}
	})

}
func device_cv_configletDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	DecomDev, err := c.API.DecomDevice(d.Get("device_serial").(string))
	if err != nil {
		return err
	}
	_ = DecomDev

	//Remove the base configlet.
	return resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		RemoveBaseConfiglet := c.API.DeleteConfiglet(d.Get("device_configlet_base_name").(string), d.Get("device_configlet_base_key").(string))
		if RemoveBaseConfiglet != nil {
			return resource.RetryableError(fmt.Errorf("Error while trying to remove base configlet"))
		} else {
			d.SetId("")
			return nil
		}
	})

}
