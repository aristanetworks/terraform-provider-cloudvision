terraform {
  required_providers {
    cvprovider = {
      source = "arista.com/dev/cvprovider"
    }
  }
}

provider "cvprovider" {
    cvp = "10.90.226.175"
    token = "${var.cvptoken}"
    port = 443
}

data "cvprovider_data_inventory" "all" {
}


output "test" {
   value = [for k in data.cvprovider_data_inventory.all.inventory : k.hostname]
}
