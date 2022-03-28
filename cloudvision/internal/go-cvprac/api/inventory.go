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
	"encoding/json"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// NetElement represents a CVP network element returned as part of
// inventory query
type NetElement struct {
	IPAddress            string       `json:"ipAddress"`
	ModelName            string       `json:"modelName"`
	InternalVersion      string       `json:"internalVersion"`
	SystemMacAddress     string       `json:"systemMacAddress"`
	MemTotal             int64        `json:"memTotal"`
	BootupTimeStamp      float64      `json:"bootupTimeStamp"`
	MemFree              int64        `json:"memFree"`
	Architecture         string       `json:"architecture"`
	InternalBuildID      string       `json:"internalBuildId"`
	HardwareRevision     string       `json:"hardwareRevision"`
	Hostname             string       `json:"hostname"`
	Fqdn                 string       `json:"fqdn"`
	TaskIDList           []CvpTask    `json:"taskIdList"`
	ZtpMode              interface{}  `json:"ztpMode"`
	Version              string       `json:"version"`
	SerialNumber         string       `json:"serialNumber"`
	Key                  string       `json:"key"`
	Type                 string       `json:"type"`
	TempActionList       []TempAction `json:"tempAction"`
	IsDANZEnabled        string       `json:"isDANZEnabled"`
	IsMLAGEnabled        string       `json:"isMLAGEnabled"`
	ComplianceIndication string       `json:"complianceIndication"`
	ComplianceCode       string       `json:"complianceCode"`
	LastSyncUp           int64        `json:"lastSyncUp"`
	UnAuthorized         bool         `json:"unAuthorized"`
	DeviceInfo           string       `json:"deviceInfo"`
	DeviceStatus         string       `json:"deviceStatus"`
	ParentContainerKey   string       `json:"parentContainerKey"`
}

// TempAction is
type TempAction struct {
	CcID                            string   `json:"ccId"`
	SessionID                       string   `json:"sessionId"`
	ContainerKey                    string   `json:"containerKey"`
	TaskID                          int      `json:"taskId"`
	Info                            string   `json:"info"`
	InfoPreview                     string   `json:"infoPreview"`
	Note                            string   `json:"note"`
	Action                          string   `json:"action"`
	NodeType                        string   `json:"nodeType"`
	NodeID                          string   `json:"nodeId"`
	ToID                            string   `json:"toId"`
	FromID                          string   `json:"fromId"`
	NodeName                        string   `json:"nodeName"`
	ToName                          string   `json:"toName"`
	FromName                        string   `json:"fromName"`
	ChildTasks                      []string `json:"childTasks"`
	ParentTask                      string   `json:"parentTask"`
	OldNodeName                     string   `json:"oldNodeName"`
	ToIDType                        string   `json:"toIdType"`
	ConfigletList                   []string `json:"configletList"`
	IgnoreConfigletList             []string `json:"ignoreConfigletList"`
	ConfigletNamesList              []string `json:"configletNamesList"`
	IgnoreConfigletNamesList        []string `json:"ignoreConfigletNamesList"`
	NodeList                        []string `json:"nodeList"`
	IgnoreNodeList                  []string `json:"ignoreNodeList"`
	NodeNamesList                   []string `json:"nodeNamesList"`
	IgnoreNodeNamesList             []string `json:"ignoreNodeNamesList"`
	NodeIPAddress                   string   `json:"nodeIpAddress"`
	NodeTargetIPAddress             string   `json:"nodeTargetIpAddress"`
	Key                             string   `json:"key"`
	IgnoreNodeID                    string   `json:"ignoreNodeId"`
	IgnoreNodeName                  string   `json:"ignoreNodeName"`
	ImageBundleID                   string   `json:"imageBundleId"`
	Mode                            string   `json:"mode"`
	Timestamp                       int64    `json:"timestamp"`
	ConfigletBuilderList            []string `json:"configletBuilderList"`
	ConfigletBuilderNamesList       []string `json:"configletBuilderNamesList"`
	IgnoreConfigletBuilderList      []string `json:"ignoreConfigletBuilderList"`
	IgnoreConfigletBuilderNamesList []string `json:"ignoreConfigletBuilderNamesList"`
	PageType                        string   `json:"pageType"`
	ViaContainer                    bool     `json:"viaContainer"`
	BestImageContainerID            string   `json:"bestImageContainerId"`
	User                            string   `json:"user"`
	FactoryID                       int      `json:"factoryId"`
	ID                              int      `json:"id"`
}

// CvpInventoryList is a list of NetElements and Containers
type CvpInventoryList struct {
	Total          int               `json:"total"`
	ContainerList  map[string]string `json:"containerList"`
	NetElementList []NetElement      `json:"netElementList"`

	ErrorResponse
}

// CvpInventoryConfiguration is the config and warnings for a device
type CvpInventoryConfiguration struct {
	Output   string   `json:"output"`
	Warnings []string `json:"warnings"`

	ErrorResponse
}

// SaveInventoryResp is the response returned for saveInventory API call
type SaveInventoryResp struct {
	Data SaveInventoryData `json:"data"`
	ErrorResponse
}

// SaveInventoryData relates to saveInventory status
type SaveInventoryData struct {
	Total                            string `json:"total"`
	UpgradeRequired                  string `json:"Upgrade required"`
	InvalidContainer                 string `json:"Invalid-Container"`
	Connected                        string `json:"Connected"`
	RegistrationInProcessByOtherUser string `json:"Registration in process by other user"`
	Duplicate                        string `json:"Duplicate"`
	Retry                            string `json:"Retry"`
	UnauthorizedAccess               string `json:"Unauthorized access"`
	Message                          string `json:"message"`
	Connecting                       string `json:"Connecting"`
}

// GetInventory returns a CvpInventoryList based on a provided query and range.
//
// Failed search returns empty
// {
//   "total": 0,
//   "containerList": {},
//   "netElementList": []
// }
func (c CvpRestAPI) GetInventory() ([]NetElement, error) {
	var info []NetElement
	query := &url.Values{
		"provisioned": {"true"},
	}

	resp, err := c.client.Get("/cvpservice/inventory/devices", query)
	if err != nil {
		return nil, errors.Errorf("GetInventory: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetInventory: %s Payload:\n%s", err, resp)

	}
	return info, nil
}

// GetInventoryConfiguration returns a CvpInventoryConfiguration based on a provided MAC Address.
//
// Failed search returns empty
// {
//   "output": "",
//   "warnings": [],
// }
func (c CvpRestAPI) GetInventoryConfiguration(
	macAddress string) (*CvpInventoryConfiguration, error) {
	var info CvpInventoryConfiguration
	query := &url.Values{
		"netElementId": {macAddress},
	}

	resp, err := c.client.Get("/cvpservice/inventory/getInventoryConfiguration.do", query)
	if err != nil {
		return nil, errors.Errorf("GetInventoryConfiguration: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetInventoryConfiguration: %s Payload:\n%s", err, resp)

	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetInventoryConfiguration: %s", err)
	}

	return &info, nil
}

// GetAllDevices returns CvpInventoryList of all current inventory
func (c CvpRestAPI) GetAllDevices() ([]NetElement, error) {
	ret, err := c.GetInventory()
	return ret, errors.Wrap(err, "GetAllDevices")
}

// GetDeviceByName returns a CvpInventoryList based on device name provided
func (c CvpRestAPI) GetDeviceByName(fqdn string) (*NetElement, error) {
	data, err := c.GetInventory()
	if err != nil {
		return nil, errors.Errorf("GetDeviceByName: %s", err)
	}

	for idx, device := range data {
		if device.Fqdn == fqdn {
			return &data[idx], nil
		}
	}
	return nil, nil
}

// GetDeviceByID returns NetElement info related to a device mac.
func (c CvpRestAPI) GetDeviceByID(mac string) (*NetElement, error) {
	data, err := c.GetInventory()
	if err != nil {
		return nil, errors.Errorf("GetDeviceByName: %s", err)
	}

	for idx, device := range data {
		if device.SystemMacAddress == mac {
			return &data[idx], nil
		}
	}
	return nil, nil
}

// GetDevicesInContainer returns a CvpInventoryList based on container name provided
func (c CvpRestAPI) GetDevicesInContainer(name string) ([]NetElement, error) {
	containerInfo, err := c.GetContainerByName(name)
	if err != nil {
		return nil, errors.Errorf("GetDevicesInContainer: %s", err)
	} else if containerInfo == nil {
		return nil, nil
	}

	data, err := c.GetAllDevices()
	if err != nil {
		return nil, errors.Errorf("GetDevicesInContainer: %s", err)
	} else if data == nil {
		return nil, nil
	}

	var netElements []NetElement
	for idx, ele := range data {
		if ele.ParentContainerKey == containerInfo.Key {
			netElements = append(netElements, data[idx])
		}
	}
	return netElements, nil
}

// GetUndefinedDevices returns a NetElement list of devices within the Undefined container
func (c CvpRestAPI) GetUndefinedDevices() ([]NetElement, error) {
	var res []NetElement

	data, err := c.GetInventory()
	if err != nil {
		return nil, errors.Errorf("GetUndefinedDevices: %s", err)
	}

	for _, netElement := range data {
		if netElement.ParentContainerKey == "undefined_container" {
			res = append(res, netElement)
		}
	}
	return res, nil
}

// GetDeviceContainer returns a Container this device is allocated to
func (c CvpRestAPI) GetDeviceContainer(mac string) (*Container, error) {
	data, err := c.SearchTopology(mac)
	if err != nil {
		return nil, errors.Errorf("GetDeviceContainer: %s", err)
	}

	var containerName string
	for _, device := range data.NetElementContainerList {
		if device.NetElementKey == mac {
			containerName = device.ContainerName
			break
		}
	}

	if containerName == "" {
		return nil, errors.Errorf("Device [%s] not of any Container", mac)
	}
	return c.GetContainerByName(containerName)
}

// Container is
type Container struct {
	Key       string `json:"Key"`
	Name      string `json:"Name"`
	CreatedBy string `json:"CreatedBy"`
	CreatedOn int64  `json:"CreatedOn"`
	Mode      string `json:"Mode"`
}

// GetContainer returns
// The endpoint searchContainers.do will not return the Undefined_Container in the list
func (c CvpRestAPI) GetContainer(name string) ([]Container, error) {
	var info []Container
	var query *url.Values

	if name != "" {
		query = &url.Values{
			"name": {name},
		}
	}

	resp, err := c.client.Get("/cvpservice/inventory/containers", query)
	if err != nil {
		return nil, errors.Errorf("GetContainer: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetContainer: %s Payload:\n%s", err, resp)
	}

	return info, nil
}

// GetAllContainers returns all current inventory Containers
func (c CvpRestAPI) GetAllContainers() ([]Container, error) {
	return c.GetContainer("")
}

// GetContainerByName returns a Container
func (c CvpRestAPI) GetContainerByName(name string) (*Container, error) {
	containers, err := c.GetContainer(name)
	if err != nil {
		return nil, errors.Errorf("GetContainerByName: %s", err)
	}
	for _, container := range containers {
		// Container names are not case sensitive
		if strings.ToLower(container.Name) == strings.ToLower(name) {
			return &container, nil
		}
	}
	return nil, nil
}

// ContainerInfo ...
type ContainerInfo struct {
	Date                 int64  `json:"date"`
	Name                 string `json:"name"`
	ParentName           string `json:"parentName"`
	UserID               string `json:"userId"`
	BundleName           string `json:"bundleName"`
	AssociatedSwitches   int    `json:"associatedSwitches"`
	AssociatedConfiglets int    `json:"associatedConfiglets"`
	DanzEnabledSwitches  int    `json:"danzEnabledSwitches"`
}

// GetContainerInfoByID returns ContainerInfo
func (c CvpRestAPI) GetContainerInfoByID(id string) (*ContainerInfo, error) {
	var query *url.Values

	query = &url.Values{
		"containerId": {id},
	}

	infoResp := struct {
		ContainerInfo
		ErrorResponse
	}{}

	resp, err := c.client.Get("/cvpservice/provisioning/getContainerInfoById.do", query)
	if err != nil {
		return nil, errors.Errorf("GetContainerInfoByID: %s", err)
	}

	if err = json.Unmarshal(resp, &infoResp); err != nil {
		return nil, errors.Errorf("GetContainerInfoByID: %s Payload:\n%s", err, resp)
	}

	if err := infoResp.Error(); err != nil {
		return nil, errors.Errorf("GetNonConnectedDeviceCount: %s", err)
	}
	return &infoResp.ContainerInfo, nil
}

// GetNonConnectedDeviceCount returns number of devices not connected
func (c CvpRestAPI) GetNonConnectedDeviceCount() (int, error) {
	resp, err := c.client.Get("/cvpservice/inventory/add/getNonConnectedDeviceCount.do", nil)
	if err != nil {
		return -1, errors.Errorf("GetNonConnectedDeviceCount: %s", err)
	}

	info := struct {
		Data int `json:"data"`
		ErrorResponse
	}{}

	if err = json.Unmarshal(resp, &info); err != nil {
		return -1, errors.Errorf("GetNonConnectedDeviceCount: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return -1, errors.Errorf("GetNonConnectedDeviceCount: %s", err)
	}

	return info.Data, nil
}

// SaveInventory saves the current CVP inventory
func (c CvpRestAPI) SaveInventory() (*SaveInventoryData, error) {
	var info SaveInventoryResp

	resp, err := c.client.Post("/cvpservice/inventory/v2/saveInventory.do", nil, []string{})
	if err != nil {
		return nil, errors.Errorf("SaveInventory: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetNonConnectedDeviceCount: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetNonConnectedDeviceCount: %s", err)
	}

	return &info.Data, nil
}

// AddToInventory Add device to the Cvp inventory. Warning -- Method doesn't check the
// existance of the parent container
func (c CvpRestAPI) AddToInventory(deviceIPAddress, parentContainerName,
	parentContainerID string) error {
	urlParams := &url.Values{
		"startIndex": {"0"},
		"endIndex":   {"0"},
	}

	containerList := []interface{}{}

	data := struct {
		Data []map[string]interface{} `json:"data"`
	}{
		Data: []map[string]interface{}{
			{
				"containerName": parentContainerName,
				"containerId":   parentContainerID,
				"containerType": "Existing",
				"ipAddress":     deviceIPAddress,
				"containerList": containerList,
			},
		},
	}

	_, err := c.client.Post("/cvpservice/inventory/add/addToInventory.do", urlParams, data)
	return errors.Wrapf(err, "AddToInventor:")
}

// DeleteDeviceByMac Remove device from the Cvp inventory using Mac address
func (c CvpRestAPI) DeleteDeviceByMac(deviceMac string) error {
	err := c.DeleteDevicesByMac([]string{deviceMac})
	return errors.Wrapf(err, "DeleteDeviceByMac:")
}

// DeleteDevicesByMac Remove list of devices from the Cvp inventory using Mac
// Addresses
func (c CvpRestAPI) DeleteDevicesByMac(deviceMacs []string) error {
	serialNums := make([]string, len(deviceMacs))
	for idx, mac := range deviceMacs {
		device, err := c.GetDeviceByID(mac)
		if err != nil {
			return errors.Wrap(err, "DeleteDevices")
		}
		serialNums[idx] = device.SerialNumber
	}
	err := c.DeleteDevicesBySerial(serialNums)
	return errors.Wrapf(err, "DeleteDevices:")
}

// DeleteDeviceBySerial Remove device based on serial
func (c CvpRestAPI) DeleteDeviceBySerial(serialNumber string) error {
	err := c.DeleteDevicesBySerial([]string{serialNumber})
	return errors.Wrapf(err, "DeleteDeviceBySerial:")
}

// DeleteDevicesBySerial Remove devices based on serial
func (c CvpRestAPI) DeleteDevicesBySerial(serialNumbers []string) error {
	data := struct {
		Data []string `json:"data"`
	}{
		Data: serialNumbers,
	}

	_, err := c.client.Delete("/cvpservice/inventory/devices", nil, data)
	return errors.Wrapf(err, "DeleteDevicesBySerial:")
}
