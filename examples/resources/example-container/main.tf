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

resource cvprovider_cv_container "example"{
  containername = "tf-example-container"
  parentcontainername = "Tenant"
}
