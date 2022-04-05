terraform {
  required_providers {
    cvprovider = {
      source = "aristanetworks/cloudvision"
    }
  }
}

provider "cvprovider" {
    cvp = "www.cv-staging.corp.arista.io"
    token = "${var.cvptoken}"
    port = 443
}
