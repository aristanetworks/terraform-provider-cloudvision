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
	"fmt"
	"log"

	"github.com/aristanetworks/go-cvprac/v3/client"
	//"github.com/aristanetworks/go-cvprac/v3/client"
	//"github.com/aristanetworks/go-cvprac/v3/cvpapi"
)

func main() {
	TokenCvp := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOjY5ODQ0OTU5MDQzMTcyNDM3MDMsImRzbiI6ImRhbi10Zi1jdnAiLCJkc3QiOiJhY2NvdW50IiwiZXhwIjoxNjY5OTA3MjY5LCJpYXQiOjE2NDYwNTcyODcsInNpZCI6IjEwZmEzMmMwMGRlYTVlYzFmMmIxN2JlYmJmZjdlZTVmMzA5M2EzYmI5MDNiZmI4MTkwMzZiY2JjZjQ2NGZjM2QtVGxrdDNUNHYtd2RZQXNoVTZrX3pXY3A5V1BWaHZrUjRjUExseGl2SiJ9.ieSXK1o6TwNQYFtYQ6PsYMl7E9ippTMgdogkE3XbNdFzEBTYSUb2L-pXOCukN5nqyEAXCfOKinLMuK-4rByAPzX2Go8fi8kEbh0PdEg7DlI6MJhfZk9BmnDou-YDAI_-OGjtZE1MQdjrQ713Vp-g2RNEXcgD84s_I0gcw6SlfzmcBpM6R21uvspDt9s3kbMsSMmwLYgxyHI7xaPPqizfX9VyWasehUct3dHrTGEpUwfIBEXwHHAlJPKK4PkzEUn97Xk-AlmACJZWM7CPeV4Hf1wKNOtTBSKB2lliVIbDogtpEz78cQvNa0gwkCqvramKBkVvH5iCJuWL-I1GxJI00N1QK5dQ8RJabwdFORogxp4qSxBtN5bjGRSJDYE8G9ud_bl6neGdqDjArRCeO_Vn-P9qy85xlkzEEy1n5q9h5IIK32M1M2SYZInKE7lujljMT-I7xutBRV4mxLg6om47ruGrlRsop7LjCwysf0wuJA7RFVB1uaMrxQZOiOFH-eu5zqYrvFe56Hr__rGGlypLwj1wXO2tnCelVJ1u-NJBxnf7EZ7HR_6bynrJSmWhG4k2wb7M8vvYfWE-dslOL3RXTeiqdP_D-oZyfRfKvPhfHCEfv2Hd-L-FUpe6b7Rm2_ttgoYFlRi7o5F8e-C2X_68tLS-FMqUgEBEUxpYgZCtTYg"
	hosts := []string{"10.90.226.175"}
	cvpClient, _ := client.NewCvpClient(
		client.Protocol("https"),
		client.Port(443),
		client.Hosts(hosts...),
		client.Debug(false),
		client.Cvaas(false, ""))
	if err := cvpClient.ConnectWithToken(TokenCvp); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// Find out the cvp info
	//data, err := cvpClient.API.GetCvpInfo()
	//if err != nil {
	//	log.Fatalf("ERROR: %s", err)
	//}
	//fmt.Printf("Data: %v\n", data)
	// Should return the following when using 2021.2.0 for example.
	// Data: version:2021.2.0, appVersion:
	//DecomDevice, err := cvpClient.API.DecomDevice("6EEC95020E0AF62E6BD556A79EA8385A")
	//OnboardDevice, err := cvpClient.API.OnboardDevice("10.90.226.202", "eos")
	//if err != nil {
	//	log.Fatalf("ERROR: %s", err)
	//}
	//fmt.Println("This is RequestID   ", OnboardDevice.Value.Key.RequestID)
	//log.Printf("Removing Device status : %q", DecomDevice.Value.Status)

	//device, err := cvpClient.API.GetDeviceByName("dh-tf-veos-leaf1.sjc.aristanetworks.com")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//name, err := cvpClient.API.GetDeviceByName("dh-tf-veos-leaf2.sjc.aristanetworks.com")
	//inv, err := cvpClient.API.GetInventory()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, i := range inv {
	//	fmt.Println(i.Hostname)
	//}
	//fmt.Print(reflect.TypeOf(device))
	//fmt.Print(device.Fqdn)
	//statusdev, err := cvpClient.API.OnboardStatus("eebe030c-7c1d-48e9-b1d6-3699e1a839dd")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Println("The status is", statusdev)
	//TopContainer, err := cvpClient.API.GetContainerByName("Tenant")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//testdan, err := cvpClient.API.GetContainerByName("testdan")
	//if err != nil {
	//	log.Print(err)
	//}

	//fmt.Println(reflect.TypeOf(testdan))

	//CurrentKey, err := cvpClient.API.GetContainerByName("TF-Container")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//ParentKey := (TopContainer.Key)
	//TFkey := (CurrentKey.Key)

	//err = cvpClient.API.AddContainer("tfcontainer", "Tenant", ParentKey)
	//ConnRemove := cvpClient.API.DeleteContainer("TF-Container", TFkey, "Tenant", ParentKey)
	//fmt.Println(ConnRemove)

	//CheckContainer, err := cvpClient.API.GetContainerByName("Tenant")
	//if err != nil {
	//	fmt.Print(err)
	//}
	//fmt.Println(CheckContainer.Key)
	//ConnRemove := cvpClient.API.DeleteContainer("tfcontainer", "root", d.Get("parentcontainername").(string), d.Get("parentcontainerkey").(string))
	//log.Print(ConnRemove)
	//GetConfiglet, err := cvpClient.API.GetConfigletByName("dantest")
	///if err != nil {
	//fmt.Println(err)
	//}
	//fmt.Println(GetConfiglet.Key)

	//AddConfiglet, err := cvpClient.API.AddConfiglet("gocvprac", "Hiiii")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//log.Print("Adding stuff", AddConfiglet)
	//configlet_ba9db7bb-9ca5-4cdf-8f86-6c83948a58c0
	//DeleteConfiglet := cvpClient.API.DeleteConfiglet("gocvprac", "configlet_6b576997-f27e-4571-a527-4e0663ecd196")
	//log.Println(DeleteConfiglet)

	//UpdateConfiglet := cvpClient.API.UpdateConfiglet("Something else", "dantest", "configlet_be0ec3e0-8191-4ae9-9aaa-89ca85cd2869")
	//fmt.Println(UpdateConfiglet)
	//movedev, err := cvpClient.API.MoveDeviceToContainer("cvpractest", device, testdan, true)
	//if err != nil {
	//	log.Print(err)
	//}
	//log.Print(movedev.TaskIDs[0])
	//Task := movedev.TaskIDs
	//fmt.Println(Task)
	//Itask, err := strconv.Atoi(Task[0])
	//if err != nil {
	//	log.Print(err)
	//}
	//fmt.Println(Itask)

	//log.Print(Task[0])
	//for _,t := range Task {
	//	fmt.Print(t)
	//}

	//ExecuteTask := cvpClient.API.ExecuteTask(Itask)
	//log.Print(ExecuteTask)
	/*
		movedev, err := cvpClient.API.GetDeviceByName("dh-tf-veos-leaf1.sjc.aristanetworks.com")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(reflect.TypeOf(movedev))
		fmt.Println(movedev.Hostname)

		base, err := cvpClient.API.GetConfigletByName("dh-tf-veos-leaf1-base")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(reflect.TypeOf(base))

		vlans, err := cvpClient.API.GetConfigletByName("vlans")
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(reflect.TypeOf(vlans))

		syslogs, err := cvpClient.API.GetConfigletByName("syslogs")
		if err != nil {
			fmt.Print(err)
		}

		Configlets := []cvpapi.Configlet{*base, *vlans, *syslogs}
		fmt.Println(Configlets)

		//DevApply, err := cvpClient.API.ApplyConfigletsToDevice("tf-testing", movedev, true, Configlets...)
		DevApply, err := cvpClient.API.ApplyConfigletsToDevice("tf-testing", movedev, true, Configlets...)

		if err != nil {
			fmt.Print(err)
		}
		log.Print(DevApply)
		Task := DevApply.TaskIDs
		fmt.Println(Task)
		DevTask, err := strconv.Atoi(Task[0])
		ExecuteTask := cvpClient.API.ExecuteTask(DevTask)
		log.Print(ExecuteTask)

		time.Sleep(5 * time.Second)

		CheckTask, err := cvpClient.API.GetTaskByID(DevTask)
		if err != nil {
			fmt.Println(err)
		}
		for _ = range time.Tick(time.Second * 10) {
			fmt.Println(CheckTask.TaskStatus)
		}
	*/
	device, err := cvpClient.API.GetDeviceByName("dh-tf-veos-leaf1.sjc.aristanetworks.com")
	if err != nil {
		fmt.Print(err)
	}

	container, err := cvpClient.API.GetContainerByName("cvp-tf")
	if err != nil {
		log.Fatal(err)
	}

	movedev, err := cvpClient.API.MoveDeviceToContainer("cv-tv", device, container, true)
	if err != nil {
		log.Print(err)
	}

	log.Print("Now moving the container ", movedev)

}
