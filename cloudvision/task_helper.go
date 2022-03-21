// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"fmt"
	"strconv"

	"github.com/aristanetworks/go-cvprac/v3/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func get_task(m interface{}, d *schema.ResourceData) (int, error) {
	c := m.(*client.CvpClient)

	getTasks, err := c.API.GetTaskByStatus("Pending")
	if err != nil {
		fmt.Print(err)
	}

	for _, i := range getTasks {
		if i.WorkOrderDetails.NetElementHostName == d.Get("device_fqdn").(string) {
			changid, err := strconv.Atoi(i.WorkOrderID)
			if err != nil {
				return 0, err
			}
			return changid, nil

		}
	}
	return 0, nil
}
