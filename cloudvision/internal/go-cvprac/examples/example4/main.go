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
	"strconv"

	cvpapi "github.com/aristanetworks/go-cvprac/v3/api"
	"github.com/aristanetworks/go-cvprac/v3/client"
)

func main() {
	hosts := []string{"10.16.129.98"}
	cvpClient, _ := client.NewCvpClient(
		client.Protocol("https"),
		client.Port(443),
		client.Hosts(hosts...),
		client.Debug(false))

	if err := cvpClient.Connect("cvpadmin", "cvp123"); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	mac := "04:47:cf:b3:2e:2b"
	destContainer := "Leafs-mnt1"

	log.Printf("Getting device: %s", mac)
	dev, err := cvpClient.API.GetDeviceByID(mac)
	if err != nil {
		log.Fatalf("Failed to Get Device: %s", err)
	}

	log.Printf("Getting Container: %s", destContainer)
	container, err := cvpClient.API.GetContainerByName(destContainer)
	if err != nil {
		log.Fatalf("Failed to Get Container: %s", err)
	}

	log.Printf("Moving device [%s] to container [%s]", dev.Fqdn, container.Name)
	taskInfo, err := cvpClient.API.MoveDeviceToContainer("TEST", dev, container, true)
	if err != nil {
		log.Fatalf("Failed to Move device: %s", err)
	}
	log.Printf("TaskInfo: %v", taskInfo)
	log.Printf("Getting ConfigletBuilder TestBuilder")
	builder, err := cvpClient.API.GetConfigletBuilderByName("TestBuilder")
	if err != nil {
		log.Fatalf("Failed to Get Builder: %s", err)
	}

	log.Printf("Generating Configlets for builder [%s]", builder.Name)
	c, err := cvpClient.API.GenerateConfigletForContainer(container, builder,
		[]cvpapi.NetElement{*dev})
	if err != nil {
		log.Fatalf("Failed to Generate Builder: %s", err)
	}

	for _, co := range c {
		taskInfo, err := cvpClient.API.ApplyConfigletToDevice("TEST", dev, &co, true)
		if err != nil {
			log.Fatalf("Failed to Apply Configlet: %s", err)
		}
		log.Printf("TaskInfo: %v", taskInfo)
		if taskInfo != nil {
			taskID, _ := strconv.Atoi(taskInfo.TaskIDs[0])
			if err := cvpClient.API.ExecuteTask(taskID); err != nil {
				log.Fatalf("Failed to Exec: %s", err)
			}
		}
	}

}
