package cvprovider_test

import (
	"os"
	"testing"

	cvprovider "github.com/aristanetworks/terraform-provider-cloudvision/cloudvision"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = cvprovider.Provider()
	testAccProviders = map[string]*schema.Provider{
		"cvaas": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := cvprovider.Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = cvprovider.Provider()
}

func testAccPreCheck(t *testing.T) {
	t.Helper()

	if err := os.Getenv("TF_VAR_cvptoken"); err == "" {
		t.Fatal("Need the cvptoken")
	}
}
