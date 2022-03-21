# Examples

This directory contains examples for the cloudvision Terraform Module.

#### For any of the examples to work properly you will need to have a CVP token issued.

To receive a service token please follow [this guide](https://www.arista.com/en/cg-cv/cv-service-accounts)

Once the token is issued it will need to be either referenced within a var as the following.

```
provider "cvprovider" {
    cvp = "10.90.226.175"
    cvptoken = "locationoftoken.txt"
    port = 443
}

```

The alternative and suggest method is to use a environmental variable.  For example,

```
export TF_VAR_cvptoken=123456789abcdefghi
```

To test

```
echo TF_VAR_cvptoken

```

* **provider/provider.tf** example file for the provider index page
* **data-sources/example-data-inventory** Will give a full listing of cloudvision inventory items and information about each for example.
```
	IPAddress
	ModelName
	InternalVersion
	SystemMacAddress
	MemTotal
	BootupTimeStamp
	MemFree
	Architecture
	InternalBuildID
	HardwareRevision
	Hostname
	Fqdn
	TaskIDList
	ZtpMode
	Version
	SerialNumber
	Key
	Type
	TempActionList
	IsDANZEnabled
	IsMLAGEnabled
	ComplianceIndication
	ComplianceCode
	LastSyncUp
	UnAuthorized
	DeviceInfo
	DeviceStatus
	ParentContainerKey
```

* **resources/example-configlet/main.tf** Will add a configlet to CVP.  The example has a configlet.cfg file in which anything will be added here to CVP.
* **resources/example-container/main.tf** Will add a container to CVP.  The example has a ```tf-example-container``` Which it will add to CVP in the example.
* **resources/example-device-configlet/main.tf** Will add any devices you have within here to CVP move them to the correct container and add any configuration that is attached inside of the intended folder to the device.

