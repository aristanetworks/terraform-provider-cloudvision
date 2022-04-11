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

func container_cv() *schema.Resource {
	return &schema.Resource{
		Create: container_cv_Create,
		Read:   container_cv_Read,
		Update: container_cv_Update,
		Delete: container_cv_Delete,

		Schema: map[string]*schema.Schema{
			"containername": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				Optional: false,
			},
			"parentcontainername": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				Optional: false,
			},
			"containerkey": &schema.Schema{
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
			"container_configlets": {
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
			Create: schema.DefaultTimeout(2 * time.Minute),
		},
	}
}

func container_cv_Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)
	//Get the parents Key fist
	TopContainer, err := c.API.GetContainerByName(d.Get("parentcontainername").(string))
	if err != nil {
		log.Fatal(err)
	}
	d.Set("parentcontainerkey", TopContainer.Key)

	//Set the ID
	d.SetId("cvtf-container-" + strconv.FormatInt(time.Now().Unix(), 10))

	//Set the container up
	err = c.API.AddContainer(d.Get("containername").(string), d.Get("parentcontainername").(string), d.Get("parentcontainerkey").(string))

	//Check for the container.

	resource.Retry(d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
		CheckContainer, err := c.API.GetContainerByName(d.Get("containername").(string))
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error Adding container%v\n", err))
		}

		if CheckContainer.Name != d.Get("containername") {
			return resource.RetryableError(fmt.Errorf("Container did not return the same name"))
		}
		err = container_cv_Read(d, m)
		if err != nil {
			return resource.NonRetryableError(err)
		} else {
			return nil
		}
	})

	ContainerItems := d.Get("container_configlets").([]interface{})

	if len(ContainerItems) > 0 {
		GetContainer, err := c.API.GetContainerByName(d.Get("containername").(string))
		if err != nil {
			log.Fatal(err)
		}
		AddConfigletsToContainer := []cvpapi.Configlet{}
		items := make([]string, len(ContainerItems))

		for i, raw := range ContainerItems {
			items[i] = raw.(string)
		}

		for _, cc := range items {
			configletscc, err := c.API.GetConfigletByName(cc)
			if err != nil {
				return err
			}
			AddConfigletsToContainer = append(AddConfigletsToContainer, *configletscc)
		}
		for _, configletname := range AddConfigletsToContainer {

			_, err := c.API.ApplyConfigletToContainer("cvptf", GetContainer, &configletname)
			if err != nil {
				return err
			}
		}
	}
	return nil

}

func container_cv_Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	//Need to read the current Key which is place that is needed in the event of deletion.
	ContainerKey, err := c.API.GetContainerByName(d.Get("containername").(string))
	if err != nil {
		log.Fatal(err)
	}
	d.Set("containerkey", ContainerKey.Key)

	return nil
}

func container_cv_Update(d *schema.ResourceData, m interface{}) error {
	// This would only kick in if there is a change in the order of configlets hopefully.

	c := m.(*client.CvpClient)
	ContainerItems := d.Get("container_configlets").([]interface{})

	if len(ContainerItems) > 0 {
		GetContainer, err := c.API.GetContainerByName(d.Get("containername").(string))
		if err != nil {
			log.Fatal(err)
		}
		SetConfigletsToContainer := []cvpapi.Configlet{}
		items := make([]string, len(ContainerItems))

		for i, raw := range ContainerItems {
			items[i] = raw.(string)
		}

		for _, cc := range items {
			configletscc, err := c.API.GetConfigletByName(cc)
			if err != nil {
				return err
			}
			SetConfigletsToContainer = append(SetConfigletsToContainer, *configletscc)
		}
		SetConfiglets, err := c.API.SetConfigletsToContainer("cvptf", GetContainer, true, SetConfigletsToContainer...)
		if err != nil {
			return err
		}

		_ = SetConfiglets

	}
	return nil
}

func container_cv_Delete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	ConnRemove := c.API.DeleteContainer(d.Get("containername").(string), d.Get("containerkey").(string), d.Get("parentcontainername").(string), d.Get("parentcontainerkey").(string))
	log.Print(ConnRemove)

	return nil
}
