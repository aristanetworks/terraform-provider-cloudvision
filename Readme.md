Terraform Provider Cloudvision
=========================


![Alt text](docs/images/overall.jpg?raw=true "Overall")

## Description

This is a terraform provider that will use Terraform to drive cloudvision resources.

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

The alternative and suggest method is to use a environmental variable.  For example,

```
export TF_VAR_cvptoken=123456789abcdefghi
```

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.18.x (to build the provider plugin)

## Usage

Detailed documentation for the GitHub provider can be found [here](https://www.terraform.io/docs/providers/github/index.html).

## Contributing

Detailed documentation for contributing to the GitHub provider can be found [here](contributing.md).
