terraform {
  required_version = ">= 1.3"

  required_providers {
    cvprovider = {
      source  = "aristanetworks/cloudvision"
      version = ">= 0.1"
    }
  }
}

provider "cvprovider" {
  cvp   = var.cvphost
  token = var.cvptoken
  port  = 443
}

data "cvprovider_data_user" "admin" {
  username = "admin"
}


output "username" {
  description = "Username"
  value       = data.cvprovider_data_user.admin.username
}
