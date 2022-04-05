---
subcategory: ""
page_title: "Move new devices from undefined and attach base configlets"
description: |-
    An example of how to add brand new devices to Cloud vision portal and move them to the correct container with the correct configlets.
---

# Add Devices and move them to the correct container.

```terraform
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
```


## Argument Reference

The following arguments are supported in the `provider` block:

* `token` - This is the secure token which is needed for TF to talk to CVP.

* `port` - Port in which cvp is running.  Default should be 443.

* `device_fqdn` - Device fqdn which is the devices fqdn before moving.  For example, device123.example.com.

* `device_containername` - Container in which you are planning on moving the device to from another container..

* `device_configlet_base` - Base configlet.  This has to be a location of a configuration file and needs to be valid for the device. This also has to be tracked within terraform.
So for example, in the example section there is a intended folder with runningconfigs.  This is where the base file will be.


* `overwrite_compliant` - This by default is true.  The way this works is if a device is out of compliance it will not overwrite it.
For example, if a user makes a manual change to a device on the CLI this device is now out of state with cloudvision.  Cloudvision will
mark the device as out of compliance.  You may not want to push config to this device.  If that is the case mark this bool as false.

* `additional_configlets` - This is a list of configlets you want to attach at the device level.  For example, if you wanted all devices to also have a specific configlet attached to it this is where you would reference it.
It is assumed that the configlets already exist within cloudvision.  Within the example "syslogs" and "vlans" already exist.

The reason for depends_on is so we can move the device to the container and then apply config immediately there afterwards.