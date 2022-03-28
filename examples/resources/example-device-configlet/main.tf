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
resource "cvprovider_device_cv_container" "leaf1" {
  device_fqdn = "dh-tf-veos-leaf1.sjc.aristanetworks.com"
  device_containername = "cvptf"
}

resource "cvprovider_device_cv_configlet" "leaf1" {
  device_fqdn = "dh-tf-veos-leaf1.sjc.aristanetworks.com"
  device_configlet_base = file("${path.module}/intended/leaf1.cfg")
  overwrite_compliant = false
  additional_configlets = ["syslogs","vlans"]
  depends_on = [
    cvprovider_device_cv_container.leaf1
  ]
}

resource "cvprovider_device_cv_container" "leaf2" {
  device_fqdn = "dh-tf-veos-leaf2.sjc.aristanetworks.com"
  device_containername = "cvptf"
}

resource "cvprovider_device_cv_configlet" "leaf2" {
  device_fqdn = "dh-tf-veos-leaf2.sjc.aristanetworks.com"
  device_configlet_base = file("${path.module}/intended/leaf2.cfg")
  additional_configlets = ["syslogs"]
  depends_on = [
    cvprovider_device_cv_container.leaf2
  ]
}

resource "cvprovider_device_cv_container" "spine1" {
  device_fqdn = "dh-tf-veos-spine1.sjc.aristanetworks.com"
  device_containername = "cvptf"
}

resource "cvprovider_device_cv_configlet" "spine1" {
  device_fqdn = "dh-tf-veos-spine1.sjc.aristanetworks.com"
  device_configlet_base = file("${path.module}/intended/spine1.cfg")
  additional_configlets = ["syslogs"]
  depends_on = [
    cvprovider_device_cv_container.spine1
  ]
}

resource "cvprovider_device_cv_container" "spine2" {
  device_fqdn = "dh-tf-veos-spine2.sjc.aristanetworks.com"
  device_containername = "cvptf"
}

resource "cvprovider_device_cv_configlet" "spine2" {
  device_fqdn = "dh-tf-veos-spine2.sjc.aristanetworks.com"
  device_configlet_base = file("${path.module}/intended/spine2.cfg")
  depends_on = [
    cvprovider_device_cv_container.spine2
  ]
}