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

func TestAcccvprovidercv_user(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testResourceInitialuserConfig,
				Check:  testResourceInitialuserconfigCheck,
			},
		},
	})
}

var testResourceInitialuserConfig = fmt.Sprintf(`
provider "cvprovider" {
    cvp = "1.2.3.4"
    token = "aaaa"
    port = 443
}

resource cvprovider_cv_container "example"{
  username = "dan"
  password = "iluvarista"
  email = "dan@dan.com"
  firstname = "dan"
  lastname = "hertzberg"
  usertype = "SSO"
  userstatus = "Enabled"
}
`)

func testResourceInitialuserconfigCheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["cvprovider_cv_user.example"]
	if resourceState == nil {
		return fmt.Errorf("cvprovider_cv_user.example resource not found")
	}
	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}
	if got, want := instanceState.Attributes["cvprovider_cv_user.containername"], "tf-example-user"; got != want {
		return fmt.Errorf("cvprovider_cv_user.example contains %s; want %s", got, want)
	}
	return nil
}
