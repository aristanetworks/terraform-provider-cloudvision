// +build systest
//
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

package cvpapi

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"
)

func GetNextTaskID(c *CvpRestAPI) int {
	tasks, err := c.GetTasks("", 0, 1)
	if err != nil {
		panic(err)
	}
	id, err := strconv.Atoi(tasks[0].WorkOrderID)
	if err != nil {
		panic(err)
	}
	return id + 1
}

func monitorTask(c *CvpRestAPI, taskID int, status string) error {
	sleepCount := 0

	for ; sleepCount < 120; sleepCount++ {
		taskResult, err := c.GetTaskByID(taskID)
		if err != nil {
			return err
		}
		if taskResult.WorkOrderUserDefinedStatus == status {
			break
		}
		time.Sleep(time.Second * 3)
	}
	if sleepCount >= 30 {
		return fmt.Errorf("Create task timeout after 120 seconds")
	}
	return nil
}

var lldpRegex = regexp.MustCompile(`(?m)lldp timer (\d+)`)

func CreateTask(c *CvpRestAPI) (int, string, error) {
	var configlet *Configlet
	nextTaskID := GetNextTaskID(c)

	// for each configlet for our selected device
	for idx, cflet := range devConfiglets {

		// XXX: Hack to get around the fact that getConfigletsByNetElementId.do returns
		// a list of configlets where each configlets NetElementCount/ContainerCount is
		// always zero.
		tmpConfiglet, err := c.GetConfigletByName(cflet.Name)
		if err != nil {
			panic(err)
		}
		// Choose a configlet that's only applied to one device and is not a
		// Reconciled configlet (Can't change Reconciled configlets)
		if tmpConfiglet.NetElementCount == 1 && !tmpConfiglet.Reconciled {
			// Make sure we are referencing the actual memory here
			// so when we update the configlet it's refected throughout
			configlet = &devConfiglets[idx]
		}
	}
	if configlet == nil {
		msg := fmt.Sprintf("No device level configlet found for device %s:%s",
			dev.Fqdn, dev.IPAddress)
		panic(msg)
	}

	config := configlet.Config
	origConfig := configlet.Config

	matched := lldpRegex.FindStringSubmatch(config)
	if matched != nil {
		value, _ := strconv.Atoi(matched[1])
		value = value + 1
		repl := fmt.Sprintf("lldp timer %d", value)
		config = lldpRegex.ReplaceAllString(config, repl)
	} else {
		config = "lldp timer 13\n" + config
	}
	configlet.Config = config

	if err := c.UpdateConfiglet(config, configlet.Name, configlet.Key); err != nil {
		return 0, "", err
	}

	sleepCount := 0
	for ; sleepCount < 30; sleepCount++ {
		time.Sleep(time.Second * 2)
		result, err := c.GetTaskByID(nextTaskID)
		if err != nil {
			err = fmt.Errorf("TaskID:%d Error:%s", nextTaskID, err)
			return nextTaskID, origConfig, err
		}
		if result.WorkOrderID == strconv.Itoa(nextTaskID) {
			break
		}
	}
	if sleepCount >= 30 {
		err := fmt.Errorf("Create task timeout")
		return 0, "", err
	}
	return nextTaskID, origConfig, nil
}

var api *CvpRestAPI
var dev *NetElement
var devContainer *Container
var devConfiglets []Configlet

var debugFlag = flag.Bool("debug", false, "Enable debug")
var unitTest = flag.Bool("unittest", false, "Run Unit Tests")
var sysTest = flag.Bool("systest", false, "Run System Tests")

func TestMain(m *testing.M) {
	var err error
	flag.Parse()

	// Get config data for setup
	_, err = LoadConfigFile("cvp_node.gcfg")
	if err != nil {
		log.Fatal(err)
	}
	config := GetConfig()
	// Setup our client
	nodeID := config.GetNodeIds()[0]
	node := config.Nodes[nodeID]
	testClient := NewRealClient(nodeID, "https", 443)
	testClient.Client.Debug = *debugFlag

	fmt.Printf("Connecting to %s\n", nodeID)
	fmt.Printf("Using creds %s/%s for testing\n", node.getUsername(), node.getPassword())

	api = NewCvpRestAPI(testClient)

	if _, err := api.Login(node.getUsername(), node.getPassword()); err != nil {
		log.Printf("Login Failure: %s", err)
		os.Exit(m.Run())
	}

	// verify we have at least one device in inventory
	inventory, err := api.GetAllDevices()
	if err != nil || len(inventory) == 0 {
		log.Fatalf("No devices found. Error: %v", err)
	}

	// get the first device that is not in undefined_container
	for idx, device := range inventory {
		if device.ParentContainerKey != "undefined_container" {
			// Get our device
			dev = &inventory[idx]
			break
		}
	}
	if dev == nil {
		log.Fatalf("All devices are in Undefined_Container")
	}

	devContainer, err = api.GetParentContainerForDevice(dev.SystemMacAddress)
	if err != nil {
		log.Fatal(err)
	}
	if devContainer == nil {
		log.Fatalf("Device [%s:%s] not assigned to container", dev.Fqdn, dev.SystemMacAddress)
	}
	devConfiglets, err = api.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		log.Fatal(err)
	}
	if len(devConfiglets) == 0 {
		log.Fatalf("Device [%s:%s] no assigned configlets", dev.Fqdn, dev.SystemMacAddress)
	}

	fmt.Printf("Device:            %s\n", dev.Fqdn)
	fmt.Printf("Device Container:  %s\n", devContainer.Name)
	fmt.Printf("Device Configlets: %v\n", devConfiglets)
	os.Exit(m.Run())
}
