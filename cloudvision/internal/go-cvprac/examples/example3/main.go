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
	hosts := []string{"10.81.110.85"}
	cvpClient, _ := client.NewCvpClient(
		client.Protocol("https"),
		client.Port(443),
		client.Hosts(hosts...),
		client.Debug(false))

	if err := cvpClient.Connect("cvpadmin", "cvp123"); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	mac := "44:4c:a8:a5:0d:a1"

	dev, err := cvpClient.API.GetDeviceByID(mac)
	if err != nil {
		log.Fatalf("Failed to Get Device: %s", err)
	}

	cont, err := cvpClient.API.GetParentContainerForDevice(mac)

	cvpClient.API.DeleteDeviceByMac(mac)

	tmp, err := cvpClient.API.GetDeviceByID(mac)
	if tmp == nil {
		log.Println("Not found...good")
	}

	log.Printf("%s %s %s\n", dev.IPAddress, cont.Name, cont.Key)
	if err := cvpClient.API.AddToInventory(dev.IPAddress, cont.Name, cont.Key); err != nil {
		log.Fatalf("Failed to Add to Inventory: %s", err)
	}

	if i, err := cvpClient.API.GetNonConnectedDeviceCount(); err != nil {
		log.Fatalf("Failed to Get NonConnected device count: %s", err)
	} else {
		log.Printf("%d\n", i)
	}

	if _, err := cvpClient.API.SaveInventory(); err != nil {
		log.Fatalf("Failed to Save: %s", err)
	}

	tmp, err = cvpClient.API.GetDeviceByID(mac)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	if tmp == nil {
		log.Fatal("Not found")
	}
}
