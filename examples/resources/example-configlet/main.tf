terraform {
  required_providers {
    cvprovider = {
      source = "arista.com/dev/cvprovider"
    }
  }
}

provider "cvprovider" {
    cvp = "1.2.3.4"
    token = "${var.cvptoken}"
    port = 443
}

resource cvprovider_cv_configlet "example"{
  configletname = "tf-example-configlet"
  configletdata = file("${path.module}/configlet.cfg")
}
