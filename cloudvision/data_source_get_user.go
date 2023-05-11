package cvprovider

import (
	"log"

	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func get_user() *schema.Resource {
	return &schema.Resource{
		Read: data_read_user,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firstname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lastname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"usertype": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"userstatus": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func data_read_user(d *schema.ResourceData, m interface{}) error {
	username := d.Get("username").(string)

	c := m.(*client.CvpClient)

	user, err := c.API.GetUser(username)
	if err != nil {
		log.Fatal(err)
	}

	d.SetId(user.UserData.UserID)
	d.Set("username", user.UserData.UserID)
	d.Set("email", user.UserData.Email)
	d.Set("firstname", user.UserData.FirstName)
	d.Set("lastname", user.UserData.LastName)
	d.Set("usertype", user.UserData.UserType)
	d.Set("userstatus", user.UserData.UserStatus)
	d.Set("roles", user.Roles)

	return nil
}
