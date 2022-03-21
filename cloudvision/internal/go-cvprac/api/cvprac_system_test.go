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
	"strconv"
	"strings"
	"testing"
)

func TestCvpRac_GetCvpInfo_SystemTest(t *testing.T) {
	info, err := api.GetCvpInfo()
	ok(t, err)
	assert(t, *info != (CvpInfo{}), "No cvpinfo found")
	t.Logf("%s", info)
}

func TestCvpRac_GetAllDevices_SystemTest(t *testing.T) {
	inv, err := api.GetAllDevices()
	ok(t, err)
	assert(t, inv != nil, "Nil inventory returned")
	assert(t, len(inv) != 0, "Inventory total equals 0")
}

func TestCvpRac_GetLogsByID_SystemTest(t *testing.T) {
	tasks, err := api.GetTaskByStatus("Completed")
	ok(t, err)
	assert(t, len(tasks) != 0, "No tasks returned.")
	currTask, err := strconv.Atoi(tasks[0].WorkOrderID)
	ok(t, err)

	cvpLogs, err := api.GetLogsByID(currTask)
	ok(t, err)
	assert(t, cvpLogs != nil, "Nil cvpLog returned")
	assert(t, len(cvpLogs) != 0, "No task logs returned")
}

func TestCvpRac_GetLogsByID_Bad_SystemTest(t *testing.T) {
	tasks, err := api.GetAllTasks()
	ok(t, err)
	assert(t, len(tasks) != 0, "No tasks in CvpTask list returned.")
	currTask, err := strconv.Atoi(tasks[0].WorkOrderID)
	ok(t, err)

	cvpLogs, err := api.GetLogsByID(currTask + 100)
	assert(t, err != nil, "No error returned for invalid task Id")
	assert(t, cvpLogs == nil, "Expected nil cvpLogs, Got: %v", cvpLogs)
}

func TestCvpRac_GetAllTasks_SystemTest(t *testing.T) {
	cvpTasks, err := api.GetAllTasks()
	ok(t, err)
	assert(t, cvpTasks != nil, "Nil cvpTasks returned.")
	assert(t, len(cvpTasks) != 0, "No cvpTasks found in cvp task list")
}

func TestCvpRac_GetTaskByID_Bad_SystemTest(t *testing.T) {
	task, err := api.GetTaskByID(100000)
	assert(t, err != nil, "No error returned for invalid taskId")
	assert(t, task == nil, "Expected nil task, Got: %v", task)

	task, err = api.GetTaskByID(-1)
	assert(t, err != nil, "No error returned for invalid taskId")
	assert(t, task == nil, "Expected nil task, Got: %v", task)
}

func TestCvpRac_GetTaskByStatus_Bad_SystemTest(t *testing.T) {
	tasks, err := api.GetTaskByStatus("BOGUS")
	ok(t, err)
	assert(t, tasks != nil, "Nil task returned")
}

func TestCvpRac_GetConfigletByName_SystemTest(t *testing.T) {
	configlet := devConfiglets[0]
	cf, err := api.GetConfigletByName(configlet.Name)
	ok(t, err)
	assert(t, cf != nil, "No Configlet for name \"%s\" found", configlet.Name)
}

func TestCvpRac_GetConfigletsByDeviceID_SystemTest(t *testing.T) {
	configletList, err := api.GetConfigletsByDeviceID(dev.SystemMacAddress)
	ok(t, err)
	assert(t, configletList != nil, "No Configlet list found for \"%s\"",
		dev.SystemMacAddress)
	assert(t, len(configletList) != 0, "No Configlets in list returned for \"%s\"",
		dev.SystemMacAddress)
}

func TestCvpRac_GetConfigletHistory_SystemTest(t *testing.T) {
	configlet := devConfiglets[0]
	hlist, err := api.GetAllConfigletHistory(configlet.Key)
	ok(t, err)
	assert(t, hlist != nil, "Nil history list returned for \"%s\"", configlet.Key)
}

func TestCvpRac_GetConfigletHistory_Bad_SystemTest(t *testing.T) {
	hlist, err := api.GetAllConfigletHistory("configlet_999_999999999999")
	assert(t, err != nil, "Expected error, Got: %s", err)
	assert(t, hlist == nil, "Expected Nil history list, Got: %v", hlist)
}

func TestCvpRac_GetDeviceByName_SystemTest(t *testing.T) {
	inv, err := api.GetDeviceByName(dev.Fqdn)
	ok(t, err)
	assert(t, inv != nil, "Nil inventory list returned for device: %s", dev.Fqdn)
}

func TestCvpRac_GetDeviceByName_Bad_SystemTest(t *testing.T) {
	inv, err := api.GetDeviceByName("bogus_host_name")
	ok(t, err)
	assert(t, inv == nil, "non-Nil netelement returned")
}

func TestCvpRac_Add_Update_Delete_Configlet_SystemTest(t *testing.T) {
	name := "CvpRacTestConfigletOps"
	config := "lldp timer 9"
	var configletAdded *Configlet

	configlet, err := api.GetConfigletByName(name)
	ok(t, err)
	if configlet != nil {
		t.Logf("Configlet \"%s\" Exists. Deleting.\n", configlet.Name)
		err = api.DeleteConfiglet(name, configlet.Key)
		ok(t, err)
	}

	// Add the configlet
	configletAdded, err = api.AddConfiglet(name, config)
	ok(t, err)
	assert(t, configletAdded != nil, "Nil Configlet returned after add")

	// Verify configlet was added
	configlet, err = api.GetConfigletByName(name)
	ok(t, err)
	assert(t, configlet != nil, "Nil Configlet returned.")
	assert(t, configlet.Name == name, "Invalid name. Expected: \"%s\", Got: \"%s\"",
		name, configlet.Name)
	assert(t, configlet.Config == config,
		"Config didn't update: Expected: \"%s\" Got: \"%s\"", config, configlet.Config)
	assert(t, configlet.Key == configletAdded.Key, "Invalid key. Expected: \"%s\", Got: \"%s\"",
		configletAdded.Key, configlet.Key)

	// Update the configlet
	newConfig := "!! this is a test configlet generated by cvprac system test\n" + config
	err = api.UpdateConfiglet(newConfig, configlet.Name, configlet.Key)
	ok(t, err)

	// Verify Update
	configlet, err = api.GetConfigletByName(name)
	ok(t, err)
	assert(t, configlet != nil, "Nil Configlet returned.")
	assert(t, configlet.Name == name, "Invalid name. Expected: \"%s\", Got: \"%s\"",
		name, configlet.Name)
	assert(t, configlet.Config == newConfig,
		"Config didn't update: Expected: \"%s\" Got: \"%s\"", newConfig, configlet.Config)
	assert(t, configlet.Key == configletAdded.Key, "Invalid key. Expected: \"%s\", Got: \"%s\"",
		configletAdded.Key, configlet.Key)

	// Delete the configlet
	err = api.DeleteConfiglet(name, configlet.Key)
	ok(t, err)

	// Verify configlet was deleted
	configlet, err = api.GetConfigletByName(name)
	ok(t, err)
	assert(t, configlet == nil, "Expected: nil, Got: \"%v\"", configlet)
}

func TestCvpRac_ExecuteTask_SystemTest(t *testing.T) {
	note := "CvpRacTest Note Update"
	// Create task
	taskID, _, err := CreateTask(api)
	ok(t, err)

	err = api.AddNoteToTask(taskID, note)
	ok(t, err)

	task, err := api.GetTaskByID(taskID)
	ok(t, err)
	assert(t, task.Note == note, "Task Note not added. Expected: \"%s\" Got: \"%s\"",
		note, task.Note)

	err = api.ExecuteTask(taskID)
	ok(t, err)

	err = monitorTask(api, taskID, "Completed")
	assert(t, err == nil, "Create task monitor error: %s", err)
}

func TestCvpRac_CancelTask_SystemTest(t *testing.T) {
	note := "Cancel Test"
	// Create task
	taskID, _, err := CreateTask(api)
	ok(t, err)

	err = api.AddNoteToTask(taskID, note)
	ok(t, err)

	task, err := api.GetTaskByID(taskID)
	ok(t, err)
	assert(t, task.WorkOrderUserDefinedStatus == "Pending",
		"Invalid task state. Expected: \"Pending\", Got: \"%s\"",
		task.WorkOrderUserDefinedStatus)

	err = api.CancelTask(taskID)
	ok(t, err)

	err = monitorTask(api, taskID, "Cancelled")
	assert(t, err == nil, "Cancel task monitor error: %s", err)
}

func TestCvpRac_Containers_SystemTest(t *testing.T) {
	name := "CVPRACTEST"
	parent := devContainer
	// create test container
	err := api.AddContainer(name, parent.Name, parent.Key)
	ok(t, err)

	// Verify container created
	searchRes, err := api.SearchTopology(name)
	ok(t, err)
	assert(t, searchRes.Total == 1, "Expected: 1, Got: %d", searchRes.Total)
	containerData := searchRes.ContainerList[0]
	assert(t, containerData.Name == name, "")

	container, err := api.GetContainerByName(name)
	ok(t, err)
	assert(t, container.Name == name, "Expected container name [%s] Got [%s]",
		container.Name, name)

	containerInfo, err := api.GetContainerInfoByID(container.Key)
	ok(t, err)
	assert(t, containerInfo.ParentName == parent.Name, "Expected parent [%s] Got [%s]",
		parent.Name, containerInfo.ParentName)

	// Delete container
	err = api.DeleteContainer(name, container.Key, parent.Name, parent.Key)
	ok(t, err)

	// Verify container deleted
	searchRes, err = api.SearchTopology(name)
	ok(t, err)
	assert(t, searchRes.Total == 0, "Expected: 0, Got: %d", searchRes.Total)
}

func TestCvpRac_ConfigletsToDevice_SystemTest(t *testing.T) {
	name := "CvpRacTestConfigletToDevice"
	config := `!! this is a test configlet generated by cvprac system test
alias srie show running-config interface ethernet 1`
	label := "cvprac test"
	var taskInfo *TaskInfo

	configlet, err := api.GetConfigletByName(name)
	ok(t, err)
	if configlet != nil {
		t.Logf("Configlet \"%s\" Exists. Deleting.\n", configlet.Name)
		err = api.DeleteConfiglet(name, configlet.Key)
		ok(t, err)
	}

	configlet, err = api.AddConfiglet(name, config)
	ok(t, err)
	assert(t, configlet.Key != "", "Null configlet key")

	nextTaskID := GetNextTaskID(api)

	newConfiglet := &Configlet{
		Name: configlet.Name,
		Key:  configlet.Key,
		Type: "Static",
	}

	validationResp, err := api.ValidateConfigletsForDevice(dev.SystemMacAddress,
		[]string{newConfiglet.Key})
	ok(t, err)
	assert(t, validationResp != nil, "Expected validation response to not be nil")
	assert(t, validationResp.New == 1,
		"Expected 1 new line in validation check of new configlets for device. Found %d",
		validationResp.New)

	taskInfo, err = api.ApplyConfigletToDevice(label, dev, newConfiglet, true)
	ok(t, err)
	assert(t, taskInfo != nil, "Expected valid taskInfo, Got: nil")
	assert(t, taskInfo.Status == "success", "Expected: \"success\", Got: \"%s\"", taskInfo.Status)
	assert(t, len(taskInfo.TaskIDs) == 1, "Expected: 1, Got: %d, taskInfo: %v",
		len(taskInfo.TaskIDs), taskInfo)
	taskID, err := strconv.Atoi(taskInfo.TaskIDs[0])
	ok(t, err)
	assert(t, taskID == nextTaskID, "TaskId Expected: %d Got: %d", nextTaskID, taskID)

	cvpTask, err := api.GetTaskByID(nextTaskID)
	ok(t, err)
	assert(t, cvpTask != nil, "")
	assert(t, cvpTask.WorkOrderID == strconv.Itoa(nextTaskID), "Expected: %d Got: %s",
		nextTaskID, cvpTask.WorkOrderID)
	assert(t, strings.Contains(cvpTask.Description, label),
		"Expected task description: \"%s\", Got: \"%s\"",
		label, cvpTask.Description)

	err = api.ExecuteTask(nextTaskID)
	ok(t, err)

	err = monitorTask(api, nextTaskID, "Completed")
	assert(t, err == nil, "Execute task error: %s", err)

	taskInfo, err = api.RemoveConfigletFromDevice(label, dev, newConfiglet, true)
	ok(t, err)
	assert(t, taskInfo != nil, "Expected valid taskInfo, Got: nil")
	assert(t, taskInfo.Status == "success", "Expected: \"success\", Got: \"%s\"", taskInfo.Status)
	assert(t, len(taskInfo.TaskIDs) == 1, "Expected: 1, Got: %d, taskInfo: %v",
		len(taskInfo.TaskIDs), taskInfo)
	taskID, err = strconv.Atoi(taskInfo.TaskIDs[0])
	ok(t, err)

	cvpTask, err = api.GetTaskByID(taskID)
	ok(t, err)
	assert(t, cvpTask != nil, "Expected valid cvpTask, Got: nil")
	assert(t, cvpTask.WorkOrderID == strconv.Itoa(taskID),
		"Expected taskID: %d, Got: %s",
		taskID, cvpTask.WorkOrderID)
	assert(t, strings.Contains(cvpTask.Description, label),
		"Expected task description: \"%s\", Got: \"%s\"",
		label, cvpTask.Description)

	err = api.ExecuteTask(taskID)
	ok(t, err)

	err = monitorTask(api, taskID, "Completed")
	assert(t, err == nil, "Execute task error: %s", err)

	err = api.DeleteConfiglet(name, configlet.Key)
	ok(t, err)

	configlet, err = api.GetConfigletByName(name)
	ok(t, err)
	assert(t, configlet == nil, "Expected: nil, Got: %v", configlet)
}

func TestCvpRac_CheckCompliance_SystemTest(t *testing.T) {
	res, err := api.CheckCompliance(dev.SystemMacAddress, "netelement")
	ok(t, err)
	assert(t, res.ComplianceCode == "0000", "Expected: \"0000\", Got: \"%s\"",
		res.ComplianceCode)
	assert(t, res.ComplianceIndication == "", "Expected: \"\", Got: \"%s\"",
		res.ComplianceIndication)
}

func TestCvpRac_GetParentContainerForDevice_SystemTest(t *testing.T) {
	res, err := api.GetParentContainerForDevice(dev.SystemMacAddress)
	ok(t, err)
	assert(t, res != nil, "Expected container, Got: nil")
}

func TestCvpRac_GetDevicesInContainer_SystemTest(t *testing.T) {
	var found bool
	res, err := api.GetParentContainerForDevice(dev.SystemMacAddress)
	ok(t, err)
	assert(t, res != nil, "Expected container, Got: nil")
	netEle, err := api.GetDevicesInContainer(res.Name)
	ok(t, err)
	assert(t, len(netEle) > 0, "No NetElements returned")
	for _, ele := range netEle {
		if ele.SystemMacAddress == dev.SystemMacAddress {
			found = true
		}
	}
	assert(t, found, "No match found for %s", res.Name)
}

func TestCvpRac_GetUndefinedDevices_SystemTest(t *testing.T) {
	_, err := api.GetUndefinedDevices()
	ok(t, err)
}

func TestCvpRac_GetAllContainers_SystemTest(t *testing.T) {
	_, err := api.GetAllContainers()
	ok(t, err)
}

func TestCvpRac_GetValidContainer_SystemTest(t *testing.T) {
	containers, err := api.GetAllContainers()
	ok(t, err)
	if len(containers) > 0 {
		containerName := containers[0].Name
		c, err := api.GetContainerByName(containerName)
		ok(t, err)
		assert(t, c.Name == containerName, "Expected: %s, Got: %s", containerName, c.Name)
	}
}

func TestCvpRac_GetInValidContainer_SystemTest(t *testing.T) {
	containerName := "bogus"
	c, err := api.GetContainerByName(containerName)
	assert(t, err == nil, "Bogus container name should return nil")
	assert(t, c == nil, "Expected: nil, Got: %v", c)

}

func TestCvpRac_ImageFuncs_SystemTest(t *testing.T) {
	imageBundleList, err := api.GetAllImageBundles()
	ok(t, err)
	assert(t, len(imageBundleList) != 0, "No image bundle list")

	imageInfo, err := api.GetImages("", 0, 0)
	ok(t, err)
	assert(t, len(imageInfo) != 0, "No image info")

	imageBundleInfo, err := api.GetImageBundleByName(imageBundleList[0].Name)
	ok(t, err)
	assert(t, imageBundleInfo != nil, "No image bundle list")
}
