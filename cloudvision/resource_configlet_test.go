// Copyright (c) 2020 Arista Networks, Inc.
// Use of this source code is governed by the Mozilla Public License Version 2.0
// that can be found in the LICENSE file.

package cvprovider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAcccvprovidercv_configlet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testResourceInitialConfigletConfig,
				Check:  testResourceInitialConfigletonfigCheck,
			},
			{
				Config: testResourceUpdatedConfigletConfig,
				Check:  testResourceInitialConfigletonfigCheck,
			},
		},
	})
}

var configlet string = "logging host 1.2.3.5"

var testResourceInitialConfigletConfig = fmt.Sprintf(`
provider "cvprovider" {
    cvp = "1.2.3.4"
    token = %s
    port = 443
}

resource cvprovider_cv_configlet "example"{
  configletname = "tf-example-configlet"
  configletdata = %s
}
`, os.Getenv("TF_VAR_token"), configlet)

var testResourceUpdatedConfigletConfig = fmt.Sprintf(`
provider "cvprovider" {
    cvp = "1.2.3.4"
    token = %s
    port = 443
}

resource cvprovider_cv_configlet "example"{
  configletname = "tf-example-configlet"
  configletdata = %s
}
`, os.Getenv("TF_VAR_cvptoken"), configlet)

func testResourceInitialConfigletonfigCheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["cvprovider_cv_configlet.example"]
	if resourceState == nil {
		return fmt.Errorf("cvprovider_cv_configlet.example resource not found")
	}
	return nil

}

func testAccCheckExampleResourceDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cvprovider_cv_configlet" {
			continue
		}
		return nil
	}
	return nil
}
