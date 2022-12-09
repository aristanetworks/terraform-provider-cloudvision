// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (

	//"strconv"
	"time"

	cvpapi "github.com/aristanetworks/go-cvprac/v3/api"
	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func user_cv() *schema.Resource {
	return &schema.Resource{
		Create: user_cv_Create,
		Read:   user_cv_Read,
		Update: user_cv_Update,
		Delete: user_cv_Delete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				Optional: false,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
				Optional: false,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"firstname": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"lastname": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"usertype": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"userstatus": &schema.Schema{
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

func user_cv_Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)
	User := cvpapi.SingleUser{}
	// Gather the user data for the fields.
	User.UserData.UserID = d.Get("username").(string)
	User.UserData.Email = d.Get("email").(string)
	User.UserData.FirstName = d.Get("firstname").(string)
	User.UserData.LastName = d.Get("lastname").(string)
	User.UserData.UserType = d.Get("usertype").(string)
	User.UserData.UserStatus = d.Get("userstatus").(string)

	// Create the user
	AddUser := c.API.AddUser(&User)
	_ = AddUser

	return nil
}

func user_cv_Read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func user_cv_Update(d *schema.ResourceData, m interface{}) error {

	return nil

}

func user_cv_Delete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.CvpClient)

	user := d.Get("username").(string)

	DeletedUser := []string{user}

	RemoveUser := c.API.DeleteUsers(DeletedUser)

	_ = RemoveUser

	return nil
}
