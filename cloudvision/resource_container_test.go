// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAcccvprovidercv_container(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testResourceInitialcontainerConfig,
				Check:  testResourceInitialcontaineronfigCheck,
			},
		},
	})
}

var testResourceInitialcontainerConfig = fmt.Sprintf(`
provider "cvprovider" {
    cvp = "1.2.3.4"
    token = "aaaa"
    port = 443
}

resource cvprovider_cv_container "example"{
  containername = "tf-example-container"
  parentcontainername = "Tenant"
}
`)

func testResourceInitialcontaineronfigCheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["cvprovider_cv_container.example"]
	if resourceState == nil {
		return fmt.Errorf("cvprovider_cv_container.example resource not found")
	}
	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}
	if got, want := instanceState.Attributes["cvprovider_cv_container.containername"], "tf-example-container"; got != want {
		return fmt.Errorf("cvprovider_cv_container.example contains %s; want %s", got, want)
	}
	return nil
}
