# go-cvprac - Golang CloudVision Portal RESTful API and Client

#### Table of Contents

1. [Overview](#overview)
2. [Requirements](#requirements)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Development](#development)
6. [Testing](#testing)
7. [Versioning](#versioning)
8. [Contributing](#contributing)
9. [Support](#support)
10. [License](#license)


## Overview

This module provides a RESTful API client for Cloudvision® Portal (CVP) which can be used for building applications that work with Arista CVP.

There are two pieces to go-cvprac:

* api - Provides all the formated request/response to CVP in a nice to use interface.
* client - Client implementation using the API.

If you would like to spin your own version of a client, then you only need to implement the API interface.

## Requirements

* Go 1.6.3+

CVP Versions supported for each go-cvprac release:

CVP Version | go-cvprac Version
------------ | -------------
2016.x-2017.x | v1
2018.1.x, 2018.2.x, 2019.1.x | v2
2020.1.x, 2020.2.x | v3

## Installation
First, it is assumed you have a standard Go workspace, as described in http://golang.org/doc/code.html, with proper GOPATH set.

Please refer section [Versioning](#versioning) for detailed info.

To download and install go-cvprac:

```bash
$ go get -u gopkg.in/aristanetworks/go-cvprac.v2
```

After setting up Go and installing go-cvprac, any required build tools can be installed by bootstrapping your environment via:

```bash
$ make bootstrap
```

## Usage

Basic usage:

The included client can be used to connect/interact with CVP:

If using a standard username and password with CVP as in examples/example1/main.go use the following.

```golang
package main

import (
	"fmt"
	"log"

	"gopkg.in/aristanetworks/go-cvprac.v2/client"
)

func main() {
	hosts := []string{"10.81.110.85"}
	cvpClient, _ := client.NewCvpClient(
		client.Protocol("https"),
		client.Port(443),
		client.Hosts(hosts...),
		client.Debug(false))

	if err := cvpClient.Connect("cvpadmin", "cvp123"); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// verify we have at least one device in inventory
	data, err := cvpClient.API.GetCvpInfo()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	fmt.Printf("Data: %v\n", data)
}
```

If using a [Service Token]("https://www.arista.com/en/cg-cv/cv-service-accounts") provided from CVP please using the following from examples/example6/main.go to pass in the bearer token within the Authorization field.

```golang
package main

import (
	"fmt"
	"log"

	"github.com/aristanetworks/go-cvprac/v3/client"
)

func main() {
	TokenCvp := "cvp token goes here"
	hosts := []string{"10.20.30.186"}
	cvpClient, _ := client.NewCvpClient(
		client.Protocol("https"),
		client.Port(443),
		client.Hosts(hosts...),
		client.Debug(false))

	if err := cvpClient.ConnectWithToken(TokenCvp); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// Find out the cvp info
	data, err := cvpClient.API.GetCvpInfo()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Printf("Data: %v\n", data)
	// Should return the following when using 2021.2.0 for example.
	// Data: version:2021.2.0, appVersion:
}
```


If you want to use your own client (to leverage some custom behavior), you merely need to implement the provided ClientInterface:

```golang
type ClientInterface interface {
	Get(string, *url.Values) ([]byte, error)
	Post(string, *url.Values, interface{}) ([]byte, error)
}
```

You then can access/interact with CVP using your clients underlying behavior. Example:

```golang
import (
	"fmt"
	"log"

	"gopkg.in/aristanetworks/go-cvprac.v2/api"
)

type YourCustomClient struct {
  ...
}
func NewYourCustomClient(host string) *YourCustomClient {
  ...
}
func (c *YourCustomClient) Get(url string, params *url.Values) ([]byte, error) {
  ...
}
func (c *YourCustomClient) Post(url string, params *url.Values, data interface{}) ([]byte, error) {
  ...
}

yourClient := NewYourCustomClient("10.10.1.2")
cvpClient := cvpapi.NewCvpRestAPI(yourClient)
cvpClient.Login(user, passwd)

```

## Development

Please refer to Contributing section on contribution guidelines.
To install the needed packages for lint/vet/etc. run the bootstrap provided:

```bash
$ make bootstrap
```

## Testing

The go-cvprac library provides various tests. To run System specific tests, you will need to update the cvp_node.gcfg file (found in api/) to include CVP specifics/credentials for your setup.

**System Test Requirements**:

* Need one CVP node for test with a test user account. Create the same account on the switch used for testing.
* Test has dedicated access to the CVP node.
* CVP node contains at least one device in a container.
* Container or device has at least one configlet applied.

For running System tests, issue the following from the root of the go-cvprac directory:

```bash
$ make systest
```

Similarly, Unit tests can be run via:

```bash
$ make unittest
```

Note: Test cases live in respective XXX_test.go files and have the following function signature:

Unit Tests: TestXXX_UnitTest(t *testing.T){...
System Tests: TestXXX_SystemTest(t *testing.T){...

Any tests written must conform to this standard.

## Versioning

Releases are done according to [Semantic Versioning](https://semver.org/)

* gopkg.in/aristanetworks/go-cvprac.v{X} points to appropriate tagged versions; {X} denotes version series number and it's a stable release for production use. For e.g. gopkg.in/arsitanetworks/go-cvprac.v1

## Contributing

Bug reports and pull requests are welcome on [GitHub](https://github.com/aristanetworks/go-cvprac).
Please note that all contributions that modify the library behavior require corresponding test cases. Otherwise the pull request will be rejected.

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [Contributor
Covenant](http://contributor-covenant.org) code of conduct.

## Support

For support, please open an
[issue](https://github.com/aristanetworks/go-cvprac) on GitHub or
contact eosplus@arista.com.  Commercial support agreements are available
through your Arista account team.

## License
BSD 3-Clause License

Copyright (c) 2017, Arista Networks EOS+
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name Arista nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
