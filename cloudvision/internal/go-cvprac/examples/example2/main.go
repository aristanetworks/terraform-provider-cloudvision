// Copyright (c) 2016-2017, Arista Networks, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//   * Redistributions of source code must retain the above copyright notice,
//   this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright
//   notice, this list of conditions and the following disclaimer in the
//   documentation and/or other materials provided with the distribution.
//
//   * Neither the name of Arista Networks nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL ARISTA NETWORKS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
// BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
// OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN
// IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
	"log"

	"github.com/aristanetworks/go-cvprac/v3/client"
)

func main() {
	name := "VLAN 100"
	content := "vlan 100\n   name DEMO_traffic\nend"
	node := "Leaf1"

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
	configlet, err := cvpClient.API.GetConfigletByName(name)
	if err != nil {
		log.Fatalf("Failed to get Configlet: %s", err)
	}
	if configlet == nil {
		// configlet doesn't exist lets create one
		configlet, err = cvpClient.API.AddConfiglet(name, content)
		if err != nil {
			log.Fatalf("Failed to Add Configlet: %s", err)
		}
	} else {
		// Configlet does exist, lets update the content
		if err = cvpClient.API.UpdateConfiglet(content, name, configlet.Key); err != nil {
			log.Fatalf("Failed to get Update Configlet: %s", err)
		}
	}

	netElement, err := cvpClient.API.GetDeviceByName(node)
	if err != nil {
		log.Fatalf("Failed to get Device %s: %s", node, err)
	}
	if netElement == nil {
		log.Fatalf("No device %s found", node)
	}

	configletList, err := cvpClient.API.GetConfigletsByDeviceID(netElement.SystemMacAddress)
	if err != nil {
		log.Fatalf("Failed to get Configlet list for device %s. %s",
			netElement.SystemMacAddress, err)
	}

	var found bool
	for _, cfglt := range configletList {
		if cfglt.Key == configlet.Key {
			found = true
		}
	}

	if !found {
		_, err = cvpClient.API.ApplyConfigletToDevice("AppName", netElement, configlet, true)
		if err != nil {
			log.Fatalf("Failed to apply configlet to device %s", node)
		}
	}
}
