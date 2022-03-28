terraform {
  required_providers {
    cvprovider = {
      source = "arista.com/dev/cloudvision"
    }
  }
}

provider "cvprovider" {
    cvp = "www.cv-staging.corp.arista.io"
    token = "${var.cvptoken}"
    port = 443
}

data "cvprovider_data_inventory" "all" {
}


output "test" {
   value = [for k in data.cvprovider_data_inventory.all.inventory : k.hostname]
}
