// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"fmt"
	"log"
	"strconv"
	"time"

	//"time"

	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func device_cv_container() *schema.Resource {
	return &schema.Resource{
		Create: device_cv_containerCreate,
		Read:   device_cv_containerRead,
		Update: device_cv_containerUpdate,
		Delete: device_cv_containerDelete,

		Schema: map[string]*schema.Schema{
			"device_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: false,
				Optional: true,
			},
			"device_containername": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"device_containertaskid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"device_taskstatus": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"device_taskstatus2": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"device_containerkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create:  schema.DefaultTimeout(5 * time.Minute),
			Update:  schema.DefaultTimeout(5 * time.Minute),
			Default: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}
func device_cv_containerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	//Set the ID
	d.SetId(d.Get("device_fqdn").(string) + "container-" + strconv.FormatInt(time.Now().Unix(), 10))

	device, err := c.API.GetDeviceByName(d.Get("device_fqdn").(string))
	if err != nil {
		return err
	}

	container, err := c.API.GetContainerByName(d.Get("device_containername").(string))
	if err != nil {
		return err
	}

	movedev, err := c.API.MoveDeviceToContainer("cvptf"+strconv.FormatInt(time.Now().Unix(), 10), device, container, true)
	if err != nil {
		return err
	}
	_ = movedev

	time.Sleep(5 * time.Second)
	contaskid, err := get_task(m, d)
	if err != nil {
		return err
	}
	d.Set("device_containertaskid", contaskid)
	return resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		err = c.API.ExecuteTask(d.Get("device_containertaskid").(int))

		gettask, err := c.API.GetTaskByID(d.Get("device_containertaskid").(int))
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error Getting the taskID", err))
		}
		log.Print(gettask)

		if gettask.TaskStatus != "COMPLETED" {
			return resource.RetryableError(fmt.Errorf("Was unable to execute the task."))
		}
		err = device_cv_containerRead(d, m)

		if err != nil {
			return resource.NonRetryableError(err)
		} else {
			return nil
		}
	})

}

func device_cv_containerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func device_cv_containerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func device_cv_containerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
