terraform {
  required_providers {
    cvprovider = {
      source = "arista.com/dev/cloudvision"
    }
  }
}

provider "cvprovider" {
    cvp = "10.90.226.175"
    token = "${var.cvptoken}"
    port = 443
}
