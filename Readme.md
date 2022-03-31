Terraform Provider Cloudvision Portal(CVP)
=========================


![Alt text](images/overall.jpg?raw=true "Overall")

## Description

This is a terraform provider that will use Terraform to drive cloudvision resources.

The provider gives the ability to
- Add new Arista EOS devices to Cloud vision portal.
- Manage Cloud vision portal configlets.
- Manage Cloud vision portal containers.
- Move Arista EOS devices to their proper containers.
- Attach additional configlets to each device.

The provider uses [go-cvprac](https://github.com/aristanetworks/go-cvprac) which provides connectivy to either the Cloudvision on prem or Cloudvision as a service([cvaas](https://www.arista.com/en/cg-cv/cv-cloudvision-as-a-service))

## Creating a token for Cloud vision.
Before proceeding please leverage the cloudvision service token.

To receive a service token please follow [this guide](https://www.arista.com/en/cg-cv/cv-service-accounts)

Once the token is issued it will need to be either referenced within a var as the following.

```
provider "cvprovider" {
    cvp = "10.90.226.175"
    cvptoken = "locationoftoken.txt"
    port = 443
}

```

The alternative and suggested method is to use a environmental variable.  For example,

```
export TF_VAR_cvptoken=123456789abcdefghi
```

## Building the provider

Linux amd-64
```
make linux
```

Darwin amd-64
```
make darwin
```

## Examples

#### Demo on a linux device.

```
make linux
```
All demos can be found within the examples/ directory.

```
 cd examples/data-sources/example-data-inventory
 ```

```
terraform init
```

```
terraform plan
```

Truncated output

```
Changes to Outputs:
  + test = [
      + "SPINE-2",
      + "SPINE-1",
      + "LEAF3",
```

If this works you should be able to successfully run resources.  This test is a simple test making sure that you can talk to Cloudvision portal.

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 1.1.5+
-	[Go](https://golang.org/doc/install) 1.18.x (to build the provider plugin)
-   [Cloudvision portal](https://www.arista.com/en/cg-cv/cv-cloudvision-portal-cvp-overview) 2021.3.0

## Usage

Detailed documentation for the GitHub provider can be found [here](https://www.terraform.io/docs/providers/github/index.html).

## Contributing

Detailed documentation for contributing to the GitHub provider can be found [here](contributing.md).
