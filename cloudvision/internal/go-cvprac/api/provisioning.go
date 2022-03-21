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
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ConfigCompareCount request structure for saveTopology
type ConfigCompareCount struct {
	Mismatch  int `json:"mismatch"`
	New       int `json:"new"`
	Reconcile int `json:"reconcile"`
}

// Action request structure for saveTopology
type Action struct {
	Action string `json:"action"`

	ConfigCompareCount *ConfigCompareCount `json:"configCompareCount,omitempty"`

	ConfigletBuilderList            []string `json:"configletBuilderList,omitempty"`
	ConfigletBuilderNamesList       []string `json:"configletBuilderNamesList,omitempty"`
	ConfigletList                   []string `json:"configletList,omitempty"`
	ConfigletNamesList              []string `json:"configletNamesList,omitempty"`
	FromID                          string   `json:"fromId"`
	FromName                        string   `json:"fromName"`
	IgnoreConfigletBuilderList      []string `json:"ignoreConfigletBuilderList,omitempty"`
	IgnoreConfigletBuilderNamesList []string `json:"ignoreConfigletBuilderNamesList,omitempty"`
	IgnoreConfigletList             []string `json:"ignoreConfigletList,omitempty"`
	IgnoreConfigletNamesList        []string `json:"ignoreConfigletNamesList,omitempty"`
	IgnoreNodeID                    string   `json:"ignoreNodeId,omitempty"`
	IgnoreNodeName                  string   `json:"ignoreNodeName,omitempty"`
	Info                            string   `json:"info"`
	InfoPreview                     string   `json:"infoPreview"`
	NodeID                          string   `json:"nodeId"`
	NodeIPAddress                   string   `json:"nodeIpAddress,omitempty"`
	NodeName                        string   `json:"nodeName"`
	NodeTargetIPAddress             string   `json:"nodeTargetIpAddress,omitempty"`
	NodeType                        string   `json:"nodeType"`
	ToID                            string   `json:"toId"`
	ToIDType                        string   `json:"toIdType"`
	ToName                          string   `json:"toName"`

	CCID                 string   `json:"ccId,omitempty"`
	ID                   int      `json:"id,omitempty"`
	Note                 string   `json:"note,omitempty"`
	ChildTasks           []string `json:"childTasks,omitempty"`
	ParentTask           string   `json:"parentTask,omitempty"`
	FactoryID            int      `json:"factoryId,omitempty"`
	BestImageContainerID string   `json:"bestImageContainerId,omitempty"`
	SessionID            string   `json:"sessionId,omitempty"`
	ContainerKey         string   `json:"containerKey,omitempty"`
	TaskID               int      `json:"taskId,omitempty"`
	OldNodeName          string   `json:"oldNodeName,omitempty"`
	NodeList             []string `json:"nodeList,omitempty"`
	IgnoreNodeList       []string `json:"ignoreNodeList,omitempty"`
	NodeNamesList        []string `json:"nodeNamesList,omitempty"`
	IgnoreNodeNamesList  []string `json:"ignoreNodeNamesList,omitempty"`
	UserID               string   `json:"userId,omitempty"`
	Key                  string   `json:"key,omitempty"`
	ImageBundleID        string   `json:"imageBundleId,omitempty"`
	Mode                 string   `json:"mode,omitempty"`
	Timestamp            float64  `json:"timestamp,omitempty"`
	PageType             string   `json:"pageType,omitempty"`
	ViaContainer         bool     `json:"viaContainer,omitempty"`
}

// Topology ..
type Topology struct {
	Key                      string        `json:"key"`
	Name                     string        `json:"name"`
	Type                     string        `json:"type"`
	ChildContainerCount      int           `json:"childContainerCount"`
	ChildNetElementCount     int           `json:"childNetElementCount"`
	ParentContainerID        interface{}   `json:"parentContainerId"`
	Mode                     string        `json:"mode"`
	DevStatus                DeviceStatus  `json:"deviceStatus"`
	ChildTaskCount           int           `json:"childTaskCount"`
	ChildContainerList       []Topology    `json:"childContainerList"`
	ChildNetElementList      []NetElement  `json:"childNetElementList"`
	HierarchyNetElementCount int           `json:"hierarchyNetElementCount"`
	TempAction               []interface{} `json:"tempAction"`
	TempEvent                []interface{} `json:"tempEvent"`
}

// DeviceStatus represents status for a device
type DeviceStatus struct {
	Critical     int `json:"critical"`
	Warning      int `json:"warning"`
	Normal       int `json:"normal"`
	ImageUpgrade int `json:"imageUpgrade"`
	Task         int `json:"task"`
	UnAuthorized int `json:"unAuthorized"`
}

// ConfigletMapping represents basic info related to a Configlet
type ConfigletMapping struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// ConfigletInfo represents the configlets for a netelemet
type ConfigletInfo struct {
	Total           int                         `json:"total"`
	ConfigletMapper map[string]ConfigletMapping `json:"configletMapper"`
	ConfigletList   []Configlet                 `json:"configletList"`

	ErrorResponse
}

// ContainerData represents a container within CVP
type ContainerData struct {
	Undefined            bool   `json:"undefined"`
	UserID               string `json:"userId"`
	DateTimeInLongFormat int64  `json:"dateTimeInLongFormat"`
	FactoryID            int    `json:"factoryId"`
	Root                 bool   `json:"root"`
	Mode                 string `json:"mode"`
	Name                 string `json:"name"`
	Key                  string `json:"key"`
	ID                   int    `json:"id"`
}

// NetElementContainer maps a specific netelemet to a container
type NetElementContainer struct {
	ContainerKey  string `json:"containerKey"`
	ContainerName string `json:"containerName"`
	NetElementKey string `json:"netElementKey"`
}

// SearchTopologyResp represents the response from SearchTopology request
type SearchTopologyResp struct {
	NetElementContainerList []NetElementContainer `json:"netElementContainerList"`
	Total                   int                   `json:"total"`
	KeywordList             []string              `json:"keywordList"`
	ContainerList           []ContainerData       `json:"containerList"`
	NetElementList          []NetElement          `json:"netElementList"`

	ErrorResponse
}

// ComplianceResp represents a response from a Compliance check
type ComplianceResp struct {
	Architecture         string    `json:"architecture"`
	BootupTimeStamp      float64   `json:"bootupTimeStamp"`
	ComplianceCode       string    `json:"complianceCode"`
	ComplianceIndication string    `json:"complianceIndication"`
	DeviceStatus         string    `json:"deviceStatus"`
	DeviceStatusInfo     string    `json:"deviceStatusInfo"`
	Fqdn                 string    `json:"fqdn"`
	HardwareRevision     string    `json:"hardwareRevision"`
	InternalBuildID      string    `json:"internalBuildId"`
	InternalVersion      string    `json:"internalVersion"`
	IPAddress            string    `json:"ipAddress"`
	IsDANZEnabled        string    `json:"isDANZEnabled"`
	IsMLAGEnabled        string    `json:"isMLAGEnabled"`
	Key                  string    `json:"key"`
	LastSyncUp           int64     `json:"lastSyncUp"`
	MemFree              int       `json:"memFree"`
	MemTotal             int       `json:"memTotal"`
	ModelName            string    `json:"modelName"`
	SerialNumber         string    `json:"serialNumber"`
	SystemMacAddress     string    `json:"systemMacAddress"`
	TaskIDList           []CvpTask `json:"taskIdList"`
	Type                 string    `json:"type"`
	UnAuthorized         bool      `json:"unAuthorized"`
	Version              string    `json:"version"`
	ZtpMode              string    `json:"ztpMode"`
	//tempAction  null `json:"tempAction"`

	ErrorResponse
}

// ReconciledConfig ...
type ReconciledConfig struct {
	Key                  string  `json:"key"`
	Name                 string  `json:"name"`
	Reconciled           bool    `json:"reconciled"`
	Config               string  `json:"config"`
	User                 string  `json:"user"`
	Note                 string  `json:"note"`
	ContainerCount       int     `json:"containerCount"`
	NetElementCount      int     `json:"netElementCount"`
	DateTimeInLongFormat float64 `json:"dateTimeInLongFormat"`
	IsDefault            string  `json:"isDefault"`
	IsAutoBuilder        string  `json:"isAutoBuilder"`
	Type                 string  `json:"type"`
	Editable             bool    `json:"editable"`
	SSLConfig            bool    `json:"sslConfig"`
	Visible              bool    `json:"visible"`
	IsDraft              bool    `json:"isDraft"`
	FactoryID            int     `json:"factoryId"`
	ID                   int     `json:"id"`
}

// ConfigBlock ...
type ConfigBlock struct {
	Command         string `json:"command"`
	RowID           int    `json:"rowId"`
	ParentRowID     int    `json:"parentRowId"`
	BlockID         string `json:"blockId"`
	Code            string `json:"code"`
	ShouldReconcile string `json:"shouldReconcile"`
}

// ValidateAndCompareConfigletsResp ...
type ValidateAndCompareConfigletsResp struct {
	ReconciledConfig   ReconciledConfig `json:"reconciledConfig"`
	Reconcile          int              `json:"reconcile"`
	New                int              `json:"new"`
	DesignedConfig     []ConfigBlock    `json:"designedConfig"`
	Total              int              `json:"total"`
	RunningConfig      []ConfigBlock    `json:"runningConfig"`
	IsReconcileInvoked bool             `json:"isReconcileInvoked"`
	Mismatch           int              `json:"mismatch"`
	Warnings           []string         `json:"warnings"`
	Errors             []ValidateError  `json:"errors"`
}

// ValidateError is the entry for device config errors
type ValidateError struct {
	ConfigletLineNo int    `json:"configletLineNo"`
	ErrorMsg        string `json:"error"`
	ConfigletID     string `json:"configletId"`
}

func (v ValidateError) String() string {
	return "\"Line:" + strconv.Itoa(v.ConfigletLineNo) + " Msg:" + v.ErrorMsg + "\""
}

// UnmarshalJSON ...
func (vcc *ValidateAndCompareConfigletsResp) UnmarshalJSON(data []byte) error {
	// Check if response is an ErrorResponse
	// This check is necessary because certain invalid calls to CVP will
	// return an ErrorResponse. If a good response is returned this
	// Unmarshal will fail and then be ignored.
	var errorResp ErrorResponse
	json.Unmarshal(data, &errorResp)
	if errorResp.ErrorMessage != "" {
		return errors.Errorf("ValidateAndCompareConfigletsResp UnmarshalJSON Error %s - %s",
			errorResp.ErrorCode, errorResp.ErrorMessage)
	}
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(data, &objMap)
	if err != nil {
		return err
	}

	// Check data for Errors as list of strings. Return if found.
	var newErrors []ValidateError
	if err = json.Unmarshal(*objMap["errors"], &newErrors); err != nil {
		// Check for Errors as string
		// For some dumb reason if there is no error, then 'errors' is a string type.
		// However, if there is an error, then errors is structured data like []ValidateError.
		// So we will try to handle two different types of errors here.
		var newErrorsString string
		err = json.Unmarshal(*objMap["errors"], &newErrorsString)
		// If Errors found as non empty string save it as list of strings
		if err == nil && newErrorsString != "" {
			newErrors = []ValidateError{{ErrorMsg: newErrorsString}}
		}
	}
	if len(newErrors) > 0 {
		vcc.Errors = newErrors
		return nil
	}

	// Check data for ReconciledConfig
	// This check is necessary because the ReconciledConfig is returned as an object when
	// there is data but as an empty string when there is none
	var newRecConf ReconciledConfig
	json.Unmarshal(*objMap["reconciledConfig"], &newRecConf)
	vcc.ReconciledConfig = newRecConf
	// Check data for Reconcile
	var newReconcile int
	json.Unmarshal(*objMap["reconcile"], &newReconcile)
	vcc.Reconcile = newReconcile
	// Check data for New
	var newNew int
	json.Unmarshal(*objMap["new"], &newNew)
	vcc.New = newNew
	// Check data for Mismatch
	var newMismatch int
	json.Unmarshal(*objMap["mismatch"], &newMismatch)
	vcc.Mismatch = newMismatch
	// Check data for Total
	var newTotal int
	json.Unmarshal(*objMap["total"], &newTotal)
	vcc.Total = newTotal
	// Check data for IsReconcileInvoked
	var newInvoked bool
	json.Unmarshal(*objMap["isReconcileInvoked"], &newInvoked)
	vcc.IsReconcileInvoked = newInvoked
	// Check data for DesignedConfig
	var newDesignedConfig []ConfigBlock
	json.Unmarshal(*objMap["designedConfig"], &newDesignedConfig)
	vcc.DesignedConfig = newDesignedConfig
	// Check data for RunningConfig
	var newRunningConfig []ConfigBlock
	json.Unmarshal(*objMap["runningConfig"], &newRunningConfig)
	vcc.RunningConfig = newRunningConfig
	// Check data for Warnings
	var newWarnings []string
	json.Unmarshal(*objMap["warnings"], &newWarnings)
	vcc.Warnings = newWarnings
	return nil
}

// TaskInfo represents task info
type TaskInfo struct {
	TaskIDs []string `json:"taskIds"`
	Status  string   `json:"status"`
}

// GetDeviceConfigletInfo returns all configlet info related to a device.
func (c CvpRestAPI) GetDeviceConfigletInfo(mac string) (*ConfigletInfo, error) {
	var info ConfigletInfo
	query := &url.Values{
		"netElementId": {mac},
		"queryParam":   {""},
		"startIndex":   {"0"},
		"endIndex":     {"0"},
	}

	resp, err := c.client.Get("/cvpservice/provisioning/getConfigletsByNetElementId.do", query)
	if err != nil {
		return nil, errors.Errorf("GetDeviceConfigletInfo: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetDeviceConfigletInfo: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetDeviceConfigletInfo: %s", err)
	}
	return &info, nil
}

// GetConfigletsByDeviceID returns the list of configlets applied to a device.
func (c CvpRestAPI) GetConfigletsByDeviceID(mac string) ([]Configlet, error) {
	info, err := c.GetDeviceConfigletInfo(mac)
	if err != nil {
		return nil, errors.Errorf("GetConfigletsByDeviceID: %s", err)
	}
	return info.ConfigletList, nil
}

func (c CvpRestAPI) addTempAction(data interface{}) error {
	var resp ErrorResponse

	query := &url.Values{
		"format":     {"topology"},
		"queryParam": {""},
		"nodeId":     {"root"},
	}

	reqResp, err := c.client.Post("/cvpservice/ztp/addTempAction.do", query, data)
	if err != nil {
		return errors.Errorf("addTempAction: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return errors.Errorf("addTempAction: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return errors.Errorf("addTempAction: %s", err)
	}
	return nil
}

// SaveTopology Schedule tasks for many operations like configlet and image bundle
// mapping/removal to/from device or container, addition/deletion of containers,
// deletion of device. Return a list of taskIds created in response to saving
// the topology.
func (c CvpRestAPI) SaveTopology() (*TaskInfo, error) {
	resp := struct {
		Data TaskInfo `json:"data"`
	}{}

	reqResp, err := c.client.Post("/cvpservice/ztp/v2/saveTopology.do", nil, []string{})
	if err != nil {
		return nil, errors.Errorf("SaveTopology: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("SaveTopology: %s Payload:\n%s", err, reqResp)
	}

	return &resp.Data, nil
}

// configletAndBuilderKeyNames helper struct for the quadtuple of configlet keys, names and
// configletBuilder keys, names.
type configletAndBuilderKeyNames struct {
	keys,
	names,
	bKeys,
	bNames []string
}

func (c *configletAndBuilderKeyNames) Equals(other *configletAndBuilderKeyNames) bool {
	if len(c.keys) != len(other.keys) || len(c.bKeys) != len(other.bKeys) {
		return false
	}

	for i, v := range c.keys {
		if v != other.keys[i] {
			return false
		}
	}

	return true
}

func changesNeeded(applied, configlets []Configlet) (
	*configletAndBuilderKeyNames, *configletAndBuilderKeyNames, error,
) {
	// configlets to set
	configletAndBuilders, err := splitToConfigletAndBuilder(configlets)
	if err != nil {
		return nil, nil, err
	}

	rmConfiglets := configletDifference(applied, configlets)

	rmConfigletAndBuilders, err := splitToConfigletAndBuilder(rmConfiglets)
	if err != nil {
		return nil, nil, err
	}

	return &configletAndBuilders, &rmConfigletAndBuilders, nil
}

// SetConfigletsToDevice Sets the configlets to the device,
// and removes configlets from the device not referenced in `configlets`.
func (c CvpRestAPI) SetConfigletsToDevice(appName string, dev *NetElement, commit bool,
	configlets ...Configlet) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("SetConfigletsToDevice: nil NetElement")
	}

	// configlets to be removed; applied minus not in configlets
	currentConfiglets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}

	newCAndB, rmCAndB, err := changesNeeded(currentConfiglets, configlets)
	if err != nil {
		return nil, err
	}

	info := appName + ": Configlet Assign: to Device " + dev.Fqdn
	infoPreview := "<b>Configlet Assign:</b> to Device" + dev.Fqdn

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ID:                              1,
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            newCAndB.bKeys,
			ConfigletBuilderNamesList:       newCAndB.bNames,
			ConfigletList:                   newCAndB.keys,
			ConfigletNamesList:              newCAndB.names,
			IgnoreConfigletBuilderNamesList: rmCAndB.bNames,
			IgnoreConfigletBuilderList:      rmCAndB.bKeys,
			IgnoreConfigletNamesList:        rmCAndB.names,
			IgnoreConfigletList:             rmCAndB.keys,
			ToID:                            dev.SystemMacAddress,
			ToIDType:                        "netelement",
			FromID:                          "",
			NodeIPAddress:                   dev.IPAddress,
			NodeName:                        "",
			NodeTargetIPAddress:             dev.IPAddress,
			FromName:                        "",
			ToName:                          dev.Fqdn,
			ChildTasks:                      []string{},
			ParentTask:                      "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("SetConfigletsToDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}

	return nil, nil
}

// ApplyConfigletsToDevice apply the configlets to the device.
func (c CvpRestAPI) ApplyConfigletsToDevice(appName string, dev *NetElement, commit bool,
	newConfiglets ...Configlet) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: nil NetElement")
	}

	configlets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}

	action, configletAndBuilders, err := checkConfigMapping(configlets, newConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "ApplyConfigletsToDevice")
	}

	if !action {
		return nil, nil
	}

	info := appName + ": Configlet Assign: to Device " + dev.Fqdn
	infoPreview := "<b>Configlet Assign:</b> to Device" + dev.Fqdn

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            configletAndBuilders.bKeys,
			ConfigletBuilderNamesList:       configletAndBuilders.bNames,
			ConfigletList:                   configletAndBuilders.keys,
			ConfigletNamesList:              configletAndBuilders.names,
			IgnoreConfigletBuilderNamesList: []string{},
			IgnoreConfigletBuilderList:      []string{},
			IgnoreConfigletNamesList:        []string{},
			IgnoreConfigletList:             []string{},
			ToID:                            dev.SystemMacAddress,
			ToIDType:                        "netelement",
			FromID:                          "",
			NodeIPAddress:                   dev.IPAddress,
			NodeName:                        "",
			NodeTargetIPAddress:             dev.IPAddress,
			FromName:                        "",
			ToName:                          dev.Fqdn,
			ChildTasks:                      []string{},
			ParentTask:                      "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}
	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ApplyConfigletToDevice apply the configlets to the device.
func (c CvpRestAPI) ApplyConfigletToDevice(appName string, dev *NetElement,
	newConfiglet *Configlet, commit bool) (*TaskInfo, error) {
	var newConfigletList []Configlet
	newConfigletList = append(newConfigletList, *newConfiglet)
	return c.ApplyConfigletsToDevice(appName, dev, commit, newConfigletList...)
}

// ValidateConfigletsForDevice validate provided configlets for device.
func (c CvpRestAPI) ValidateConfigletsForDevice(deviceMac string, configletKeys []string) (
	*ValidateAndCompareConfigletsResp, error) {
	var resp ValidateAndCompareConfigletsResp
	data := struct {
		NetElementID string   `json:"netElementId"`
		ConfigIDList []string `json:"configIdList"`
		PageType     string   `json:"pageType"`
	}{
		NetElementID: deviceMac,
		ConfigIDList: configletKeys,
		PageType:     "validateConfig",
	}

	reqResp, err := c.client.Post("/cvpservice/provisioning/v2/validateAndCompareConfiglets.do", nil, data)
	if err != nil {
		return nil, errors.Errorf("ValidateConfigletsForDevice: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("ValidateConfigletsForDevice: %s Payload:\n%s", err, reqResp)
	}

	if len(resp.Errors) > 0 {
		var errorList []string
		for _, vErr := range resp.Errors {
			errorList = append(errorList, vErr.String())
		}
		return nil, errors.Errorf("ValidateConfigletsForDevice: %s", strings.Join(errorList, ", "))
	}
	return &resp, nil
}

// ValidateAndApplyConfigletsToDevice validate and apply the configlets to the device.
func (c CvpRestAPI) ValidateAndApplyConfigletsToDevice(appName string, dev *NetElement, commit bool,
	newConfiglets ...Configlet) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: nil NetElement")
	}

	configlets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}
	action, configletAndBuilders, err := checkConfigMapping(configlets, newConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "ApplyConfigletsToDevice")
	}

	if !action {
		return nil, nil
	}

	// Run Validation of new configlets to be applied
	validateResp, err := c.ValidateConfigletsForDevice(dev.SystemMacAddress, configletAndBuilders.keys)
	if err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}
	// If validation returned a proper validation response pull the config compare count values
	// to be applied to the Action data
	var confCompCount *ConfigCompareCount
	if validateResp != nil {
		confCompCount = &ConfigCompareCount{
			Mismatch:  validateResp.Mismatch,
			New:       validateResp.New,
			Reconcile: validateResp.Reconcile,
		}
	}

	info := appName + ": Configlet Assign: to Device " + dev.Fqdn
	infoPreview := "<b>Configlet Assign:</b> to Device" + dev.Fqdn

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ConfigCompareCount:              confCompCount,
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            configletAndBuilders.bKeys,
			ConfigletBuilderNamesList:       configletAndBuilders.bNames,
			ConfigletList:                   configletAndBuilders.keys,
			ConfigletNamesList:              configletAndBuilders.names,
			IgnoreConfigletBuilderNamesList: []string{},
			IgnoreConfigletBuilderList:      []string{},
			IgnoreConfigletNamesList:        []string{},
			IgnoreConfigletList:             []string{},
			ToID:                            dev.SystemMacAddress,
			ToIDType:                        "netelement",
			FromID:                          "",
			NodeIPAddress:                   dev.IPAddress,
			NodeName:                        "",
			NodeTargetIPAddress:             dev.IPAddress,
			FromName:                        "",
			ToName:                          dev.Fqdn,
			ChildTasks:                      []string{},
			ParentTask:                      "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyConfigletsToDevice: %s", err)
	}
	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// RemoveConfigletsFromDevice Remove the configlets from the device.
func (c CvpRestAPI) RemoveConfigletsFromDevice(appName string, dev *NetElement, commit bool,
	remConfiglets ...Configlet) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("RemoveConfigletsFromDevice: nil NetElement")
	}

	configlets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("RemoveConfigletsFromDevice: %s", err)
	}

	action, configletAndBuilders, rmConfigletAndBuilders, err :=
		checkRemoveConfigMapping(configlets, remConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "RemoveConfigletsFromDevice")
	}

	if !action {
		return nil, nil
	}

	info := appName + ": Configlet Remove: from Device " + dev.Fqdn
	infoPreview := "<b>Configlet Remove:</b> from Device" + dev.Fqdn

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ID:                              1,
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletList:                   configletAndBuilders.keys,
			ConfigletNamesList:              configletAndBuilders.names,
			ConfigletBuilderList:            configletAndBuilders.bKeys,
			ConfigletBuilderNamesList:       configletAndBuilders.bNames,
			IgnoreConfigletList:             rmConfigletAndBuilders.keys,
			IgnoreConfigletNamesList:        rmConfigletAndBuilders.names,
			IgnoreConfigletBuilderList:      rmConfigletAndBuilders.bKeys,
			IgnoreConfigletBuilderNamesList: rmConfigletAndBuilders.bNames,
			ToID:                            dev.SystemMacAddress,
			ToIDType:                        "netelement",
			FromID:                          "",
			NodeName:                        "",
			NodeIPAddress:                   dev.IPAddress,
			NodeTargetIPAddress:             dev.IPAddress,
			FromName:                        "",
			ToName:                          dev.Fqdn,
			ChildTasks:                      []string{},
			ParentTask:                      "",
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("RemoveConfigletsFromDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// RemoveConfigletFromDevice Remove the configlets from the device.
func (c CvpRestAPI) RemoveConfigletFromDevice(appName string, dev *NetElement,
	remConfiglet *Configlet, commit bool) (*TaskInfo, error) {
	var remConfigletList []Configlet
	remConfigletList = append(remConfigletList, *remConfiglet)
	return c.RemoveConfigletsFromDevice(appName, dev, commit, remConfigletList...)
}

// SetConfigletsToContainer Sets the configlets to the container,
// and removes configlets from the container not referenced in `configlets`.
func (c CvpRestAPI) SetConfigletsToContainer(appName string, cont *Container, commit bool,
	configlets ...Configlet) (*TaskInfo, error) {
	if cont == nil {
		return nil, errors.Errorf("SetConfigletsToContainer: nil Container")
	}

	// configlets to be removed; applied minus not in configlets
	currentConfiglets, err := c.GetContainerConfiglets(cont.Key)
	if err != nil {
		return nil, errors.Errorf("SetConfigletsToContainer: %s", err)
	}

	newCAndB, rmCAndB, err := changesNeeded(currentConfiglets, configlets)
	if err != nil {
		return nil, err
	}

	info := appName + ": Configlet Assign: to Container " + cont.Name
	infoPreview := "<b>Configlet Assign:</b> to Container" + cont.Name

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ID:                              1,
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            newCAndB.bKeys,
			ConfigletBuilderNamesList:       newCAndB.bNames,
			ConfigletList:                   newCAndB.keys,
			ConfigletNamesList:              newCAndB.names,
			IgnoreConfigletBuilderNamesList: rmCAndB.bNames,
			IgnoreConfigletBuilderList:      rmCAndB.bKeys,
			IgnoreConfigletNamesList:        rmCAndB.names,
			IgnoreConfigletList:             rmCAndB.keys,
			ToID:                            cont.Key,
			ToIDType:                        "container",
			FromID:                          "",
			FromName:                        "",
			ToName:                          cont.Name,
			ChildTasks:                      []string{},
			ParentTask:                      "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("SetConfigletsToDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}

	return nil, nil
}

// ApplyConfigletsToContainer apply the configlets to the container.
func (c CvpRestAPI) ApplyConfigletsToContainer(appName string, cont *Container,
	newConfiglets ...Configlet) (*TaskInfo, error) {
	if cont == nil {
		return nil, errors.Errorf("ApplyConfigletsToContainer: nil Container")
	}

	configlets, err := c.GetContainerConfiglets(cont.Key)
	if err != nil {
		return nil, errors.Errorf("ApplyConfigletsToContainer: %s", err)
	}

	action, configletAndBuilders, err := checkConfigMapping(configlets, newConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "ApplyConfigletsToContainer")
	}

	if !action {
		return nil, nil
	}

	info := appName + ": Configlet Assign: to Container " + cont.Name
	infoPreview := "<b>Configlet Assign:</b> to Container " + cont.Name

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            configletAndBuilders.bKeys,
			ConfigletBuilderNamesList:       configletAndBuilders.bNames,
			ConfigletList:                   configletAndBuilders.keys,
			ConfigletNamesList:              configletAndBuilders.names,
			IgnoreConfigletBuilderNamesList: []string{},
			IgnoreConfigletBuilderList:      []string{},
			IgnoreConfigletNamesList:        []string{},
			IgnoreConfigletList:             []string{},
			ToID:                            cont.Key,
			ToIDType:                        "container",
			FromID:                          "",
			FromName:                        "",
			//NodeIPAddress:                   dev.IPAddress,
			//NodeTargetIPAddress:             dev.IPAddress,
			NodeName:   "",
			ToName:     cont.Name,
			ChildTasks: []string{},
			ParentTask: "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyConfigletsToContainer: %s", err)
	}
	return c.SaveTopology()
}

// ApplyConfigletToContainer apply the configlets to the container.
func (c CvpRestAPI) ApplyConfigletToContainer(appName string, cont *Container,
	newConfiglet *Configlet) (*TaskInfo, error) {
	var newConfigletList []Configlet
	newConfigletList = append(newConfigletList, *newConfiglet)
	return c.ApplyConfigletsToContainer(appName, cont, newConfigletList...)
}

// RemoveConfigletsFromContainer Remove the configlets from the container.
func (c CvpRestAPI) RemoveConfigletsFromContainer(appName string, cont *Container,
	remConfiglets ...Configlet) (*TaskInfo, error) {
	if cont == nil {
		return nil, errors.Errorf("RemoveConfigletsFromContainer: nil Container")
	}

	configlets, err := c.GetContainerConfiglets(cont.Key)
	if err != nil {
		return nil, errors.Errorf("RemoveConfigletsFromContainer: %s", err)
	}

	action, configletAndBuilders, rmConfigletAndBuilders, err :=
		checkRemoveConfigMapping(configlets, remConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "RemoveConfigletsFromContainer")
	}

	if !action {
		return nil, nil
	}

	info := appName + ": Configlet Remove: from Container " + cont.Name
	infoPreview := "<b>Configlet Remove:</b> from Container" + cont.Name

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ID:                              1,
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletList:                   configletAndBuilders.keys,
			ConfigletNamesList:              configletAndBuilders.names,
			ConfigletBuilderList:            configletAndBuilders.bKeys,
			ConfigletBuilderNamesList:       configletAndBuilders.bNames,
			IgnoreConfigletList:             rmConfigletAndBuilders.keys,
			IgnoreConfigletNamesList:        rmConfigletAndBuilders.names,
			IgnoreConfigletBuilderList:      rmConfigletAndBuilders.bKeys,
			IgnoreConfigletBuilderNamesList: rmConfigletAndBuilders.bNames,
			ToID:                            cont.Key,
			ToIDType:                        "container",
			FromID:                          "",
			NodeName:                        "",
			//NodeIPAddress:                   dev.IPAddress,
			//NodeTargetIPAddress:             dev.IPAddress,
			FromName:   "",
			ToName:     cont.Name,
			ChildTasks: []string{},
			ParentTask: "",
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("RemoveConfigletsFromContainer: %s", err)
	}
	return c.SaveTopology()
}

// RemoveConfigletFromContainer Remove the configlets from the device.
func (c CvpRestAPI) RemoveConfigletFromContainer(appName string, cont *Container,
	remConfiglet *Configlet) (*TaskInfo, error) {
	var remConfigletList []Configlet
	remConfigletList = append(remConfigletList, *remConfiglet)
	return c.RemoveConfigletsFromContainer(appName, cont, remConfigletList...)
}

// GetContainerConfiglets returns a list of configlets for a given container key
func (c CvpRestAPI) GetContainerConfiglets(cid string) ([]Configlet, error) {
	result, err := c.GetContainerConfigletsWithRange(cid, 0, 0)
	return result, errors.Wrap(err, "GetContainerConfiglets")
}

// GetContainerConfigletsWithRange returns a list of configlets for a given
// container key and start/end range
func (c CvpRestAPI) GetContainerConfigletsWithRange(cid string, start int,
	end int) ([]Configlet, error) {
	var resp ConfigletInfo
	query := &url.Values{
		"containerId": {cid},
		"startIndex":  {strconv.Itoa(start)},
		"endIndex":    {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/cvpservice/provisioning/getConfigletsByContainerId.do", query)
	if err != nil {
		return nil, errors.Errorf("GetContainerConfigletsWithRange: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetContainerConfigletsWithRange: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetContainerConfigletsWithRange: %s", err)
	}
	return resp.ConfigletList, nil
}

func (c CvpRestAPI) containerOp(containerName, containerKey, parentName,
	parentKey, operation string) (*TaskInfo, error) {

	msg := operation + " container " + containerName + " under container " + parentName

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			Info:        msg,
			InfoPreview: msg,
			Action:      operation,
			NodeType:    "container",
			NodeID:      containerKey,
			ToID:        parentKey,
			ToIDType:    "container",
			ToName:      parentName,
			FromID:      "",
			FromName:    "",
			NodeName:    containerName,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("containerOp: %s", err)
	}
	return c.SaveTopology()
}

// AddContainer adds the container to the specified parent.
func (c CvpRestAPI) AddContainer(containerName, parentName,
	parentKey string) error {
	_, err := c.containerOp(containerName, "New_container1", parentName, parentKey, "add")
	return errors.Wrap(err, "AddContainer")
}

// DeleteContainer deletes the container from the specified parent.
func (c CvpRestAPI) DeleteContainer(containerName, containerKey,
	parentName, parentKey string) error {
	_, err := c.containerOp(containerName, containerKey, parentName, parentKey, "delete")
	return errors.Wrap(err, "DeleteContainer")
}

// ResetDevice Resets/Reboots the device to factory setting.
func (c CvpRestAPI) ResetDevice(appName string, dev *NetElement,
	container *Container, commit bool) (*TaskInfo, error) {
	if dev == nil {
		return nil, errors.Errorf("ResetDevice: nil NetElement ref provided")
	}
	if container == nil {
		return nil, errors.Errorf("ResetDevice: nil Container ref provided")
	}

	info := appName + ": Reset: Device Reset: %s - To be Reset" + dev.Fqdn
	infoPreview := "<b>Device Reset:</b> %s - To be Reset" + dev.Fqdn

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ID:          1,
			Action:      "reset",
			FromID:      dev.ParentContainerKey,
			FromName:    container.Name,
			Info:        info,
			InfoPreview: infoPreview,
			NodeID:      dev.SystemMacAddress,
			NodeName:    dev.Fqdn,
			NodeType:    "netelement",
			ToID:        "undefined_container",
			ToIDType:    "container",
			ParentTask:  "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ResetDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// SearchTopologyWithRange searches the topology for items matching the query parameter
// and returning those within the specified range.
//
// If query yields no hits, then result is (SearchTopologyResp{})
func (c CvpRestAPI) SearchTopologyWithRange(querystr string, start int,
	end int) (*SearchTopologyResp, error) {
	var resp SearchTopologyResp
	query := &url.Values{
		"queryParam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/cvpservice/provisioning/searchTopology.do", query)
	if err != nil {
		return nil, errors.Errorf("SearchTopologyWithRange: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("SearchTopologyWithRange: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("SearchTopologyWithRange: %s", err)
	}
	return &resp, nil
}

// SearchTopology searches the topology for items matching the query parameter.
func (c CvpRestAPI) SearchTopology(query string) (*SearchTopologyResp, error) {
	return c.SearchTopologyWithRange(query, 0, 0)
}

// CheckCompliance Check that a device is in compliance, that is the configlets
// applied to the device match the devices running configuration.
//
// Supported only for NetElements
//
func (c CvpRestAPI) CheckCompliance(nodeKey string, nodeType string) (*ComplianceResp, error) {
	var info ComplianceResp
	data := map[string]string{
		"nodeId":   nodeKey,
		"nodeType": nodeType,
	}

	resp, err := c.client.Post("/cvpservice/provisioning/checkCompliance.do", nil, data)
	if err != nil {
		return nil, errors.Errorf("CheckCompliance: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("CheckCompliance: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("CheckCompliance: %s", err)
	}

	return &info, nil
}

// GetParentContainerForDevice returns the Container for specified deviceMAC
func (c CvpRestAPI) GetParentContainerForDevice(deviceMAC string) (*Container, error) {
	results, err := c.SearchTopologyWithRange(deviceMAC, 0, 0)
	if err != nil {
		return nil, errors.Errorf("GetParentContainerForDevice: %s", err)
	}
	for _, netContainerInfo := range results.NetElementContainerList {
		if netContainerInfo.NetElementKey == deviceMAC {
			return c.GetContainerByName(netContainerInfo.ContainerName)
		}
	}
	return nil, nil
}

// MoveDeviceToContainer moves a specified netelement to a container.
func (c CvpRestAPI) MoveDeviceToContainer(appName string, device *NetElement,
	container *Container, commit bool) (*TaskInfo, error) {
	if device == nil {
		return nil, errors.Errorf("MoveDeviceToContainer: nil NetElement")
	}
	if container == nil {
		return nil, errors.Errorf("MoveDeviceToContainer: nil Container")
	}

	var fromID string
	var fromName string
	if device.ParentContainerKey != "" {
		container, err := c.GetContainerInfoByID(device.ParentContainerKey)
		if err != nil {
			return nil, errors.Errorf("MoveDeviceToContainer: %s", err)
		}
		if container == nil {
			return nil, errors.Errorf("MoveDeviceToContainer: No container found for "+
				"device [%s] containerID [%s]", device.Fqdn, fromID)
		}
		fromID = device.ParentContainerKey
		fromName = container.Name
	} else {
		parentCont, err := c.GetParentContainerForDevice(device.SystemMacAddress)
		if err != nil {
			return nil, errors.Errorf("MoveDeviceToContainer: %s", err)
		}
		if parentCont == nil {
			return nil, errors.Errorf("MoveDeviceToContainer: No parent container found for "+
				"device [%s]", device.SystemMacAddress)
		}
		fromID = parentCont.Key
		fromName = parentCont.Name
	}

	info := appName + ": Moving device " + device.Fqdn + " from container " + fromID +
		" to container " + container.Name
	infoPreview := "<b>Moving device:</b> " + device.Fqdn + " from container " + fromID +
		" to container " + container.Name

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			Info:        info,
			InfoPreview: infoPreview,
			Action:      "update",
			NodeType:    "netelement",
			NodeID:      device.SystemMacAddress,
			ToID:        container.Key,
			ToName:      container.Name,
			ToIDType:    "container",
			FromID:      fromID,
			FromName:    fromName,
			NodeName:    device.Fqdn,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("MoveDeviceToContainer: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ImageInfo represents information related to an Image within CVP
type ImageInfo struct {
	AppliedContainersCount   int    `json:"appliedContainersCount"`
	AppliedDevicesCount      int    `json:"appliedDevicesCount"`
	FactoryID                int    `json:"factoryId"`
	ID                       int    `json:"id"`
	ImageID                  string `json:"imageId"`
	ImageFile                string `json:"imageFile"`
	ImageFileName            string `json:"imageFileName"`
	ImageSize                string `json:"imageSize"`
	IsHotFix                 string `json:"isHotFix"`
	IsRebootRequired         string `json:"isRebootRequired"`
	Key                      string `json:"key"`
	MD5                      string `json:"md5"`
	SHA512                   string `json:"sha512"`
	Name                     string `json:"name"`
	SwiMaxHwepoch            string `json:"swiMaxHwepoch"`
	SwiVarient               string `json:"swiVarient"`
	UploadedDateinLongFormat int64  `json:"uploadedDateinLongFormat"`
	User                     string `json:"user"`
	Version                  string `json:"version"`
}

// ImageResp response from Image request
type ImageResp struct {
	Total int         `json:"total"`
	Data  []ImageInfo `json:"data"`

	ErrorResponse
}

// ImageBundleInfo represents ImageBundle object within CVP
type ImageBundleInfo struct {
	AppliedContainersCount   int         `json:"appliedContainersCount"`
	AppliedDevicesCount      int         `json:"appliedDevicesCount"`
	FactoryID                int         `json:"factoryId"`
	ID                       int         `json:"id"`
	IsCertifiedImageBundle   string      `json:"isCertifiedImageBundle"`
	ImageIds                 []string    `json:"imageIds"`
	Images                   []ImageInfo `json:"images,omitempty"`
	Key                      string      `json:"key"`
	Name                     string      `json:"name"`
	Note                     string      `json:"note"`
	UploadedBy               string      `json:"uploadedBy,omitempty"`
	UploadedDateinLongFormat int64       `json:"uploadedDateinLongFormat"`
	User                     string      `json:"user"`

	ErrorResponse
}

// ImageBundleResp response data/payload for ImageBundle query
type ImageBundleResp struct {
	Total                 int                   `json:"total"`
	Data                  []ImageBundleInfo     `json:"data"`
	ImageBundleMapper     map[string]*ImageInfo `json:"imageBundleMapper"`
	AssignedImageBundleID string                `json:"assignedImageBundleId"`

	ErrorResponse
}

// GetImages returns a list of Images based on a specific query string and range
func (c CvpRestAPI) GetImages(querystr string, start int, end int) ([]ImageInfo, error) {
	var resp ImageResp
	query := &url.Values{
		"queryParam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/cvpservice/image/getImages.do", query)
	if err != nil {
		return nil, errors.Errorf("GetImages: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetImages: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetImages: %s", err)
	}
	return resp.Data, nil
}

// GetImageByName returns an ImageInfo object based on name provided
func (c CvpRestAPI) GetImageByName(name string) (*ImageInfo, error) {
	resp, err := c.GetImages(name, 0, 0)
	if err != nil {
		return nil, errors.Errorf("GetImageByName: %s", err)
	}

	for _, image := range resp {
		if image.Name == name {
			return &image, nil
		}
	}
	return nil, nil
}

// GetImageBundles returns a list of ImageBundles based on a specific query string and range
func (c CvpRestAPI) GetImageBundles(querystr string, start, end int) ([]ImageBundleInfo, error) {
	var resp ImageBundleResp
	query := &url.Values{
		"queryParam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/cvpservice/image/getImageBundles.do", query)
	if err != nil {
		return nil, errors.Errorf("GetImageBundles: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetImageBundles: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetImageBundles: %s", err)
	}
	return resp.Data, nil

}

// GetAllImageBundles gets all ImageBundles
func (c CvpRestAPI) GetAllImageBundles() ([]ImageBundleInfo, error) {
	return c.GetImageBundles("", 0, 0)
}

// GetImageBundleByName gets ImageBundle by specified name
func (c CvpRestAPI) GetImageBundleByName(name string) (*ImageBundleInfo, error) {
	// Hack around string returned for ID
	type tmp struct {
		AppliedContainersCount   int         `json:"appliedContainersCount"`
		AppliedDevicesCount      int         `json:"appliedDevicesCount"`
		FactoryID                int         `json:"factoryId"`
		ID                       string      `json:"id"`
		IsCertifiedImageBundle   string      `json:"isCertifiedImageBundle"`
		ImageIds                 []string    `json:"imageIds"`
		Images                   []ImageInfo `json:"images,omitempty"`
		Key                      string      `json:"key"`
		Name                     string      `json:"name"`
		Note                     string      `json:"note"`
		UploadedBy               string      `json:"uploadedBy,omitempty"`
		UploadedDateinLongFormat int64       `json:"uploadedDateinLongFormat"`
		User                     string      `json:"user"`
		ErrorResponse
	}
	var resp tmp

	query := &url.Values{
		"name": {name},
	}

	reqResp, err := c.client.Get("/cvpservice/image/getImageBundleByName.do", query)
	if err != nil {
		return nil, errors.Errorf("GetImageBundleByName: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetImageBundleByName: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetImageBundleByName: %s", err)
	}
	ret := &ImageBundleInfo{
		AppliedContainersCount:   resp.AppliedContainersCount,
		AppliedDevicesCount:      resp.AppliedDevicesCount,
		FactoryID:                resp.FactoryID,
		ID:                       1,
		IsCertifiedImageBundle:   resp.IsCertifiedImageBundle,
		ImageIds:                 resp.ImageIds,
		Images:                   resp.Images,
		Key:                      resp.Key,
		Name:                     resp.Name,
		Note:                     resp.Note,
		UploadedBy:               resp.UploadedBy,
		UploadedDateinLongFormat: resp.UploadedDateinLongFormat,
		User:                     resp.User,
	}
	return ret, nil

}

// ApplyImageToDevice Applies image bundle to device
func (c CvpRestAPI) ApplyImageToDevice(appName string, imageInfo *ImageBundleInfo,
	netElement *NetElement, commit bool) (*TaskInfo, error) {
	if imageInfo == nil {
		return nil, errors.Errorf("ApplyImageToDevice: nil ImageBundleInfo")
	}
	if netElement == nil {
		return nil, errors.Errorf("ApplyImageToDevice: nil NetElement")
	}

	msg := appName + ": Apply image " + imageInfo.Name + " to NetElement " + netElement.Fqdn

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			Info:        msg,
			InfoPreview: msg,
			Note:        "",
			Action:      "associate",
			NodeType:    "imagebundle",
			NodeID:      strconv.Itoa(imageInfo.ID),
			ToID:        netElement.SystemMacAddress,
			ToIDType:    "netelement",
			FromID:      "",
			NodeName:    imageInfo.Name,
			FromName:    "",
			ToName:      netElement.Fqdn,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyImageToDevice: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// ApplyImageToContainer Applies image bundle to container
func (c CvpRestAPI) ApplyImageToContainer(appName string, imageInfo *ImageBundleInfo,
	container *Container, commit bool) (*TaskInfo, error) {
	if imageInfo == nil {
		return nil, errors.Errorf("ApplyImageToContainer: nil ImageBundleInfo")
	}
	if container == nil {
		return nil, errors.Errorf("ApplyImageToContainer: nil Container")
	}

	msg := appName + ": Apply image " + imageInfo.Name + " to Container " + container.Name
	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ID:          1,
			Info:        msg,
			InfoPreview: msg,
			Note:        "",
			Action:      "associate",
			NodeType:    "imagebundle",
			NodeID:      strconv.Itoa(imageInfo.ID),
			ToID:        container.Key,
			ToIDType:    "container",
			FromID:      "",
			NodeName:    imageInfo.Name,
			FromName:    "",
			ToName:      container.Name,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("ApplyImageToContainer: %s", err)
	}

	if commit {
		return c.SaveTopology()
	}
	return nil, nil
}

// RemoveImageFromContainer removes image bundle from container
func (c CvpRestAPI) RemoveImageFromContainer(appName string, imageInfo *ImageBundleInfo,
	container *Container) (*TaskInfo, error) {
	if imageInfo == nil {
		return nil, errors.Errorf("RemoveImageFromContainer: nil ImageBundleInfo")
	}
	if container == nil {
		return nil, errors.Errorf("RemoveImageFromContainer: nil Container")
	}

	msg := appName + ": Remove image " + imageInfo.Name + " from Container " + container.Name

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			ID:             1,
			Info:           msg,
			InfoPreview:    msg,
			Note:           "",
			Action:         "associate",
			NodeType:       "imagebundle",
			NodeID:         "",
			ToID:           container.Key,
			ToIDType:       "container",
			FromID:         "",
			NodeName:       "",
			FromName:       "",
			ToName:         container.Name,
			IgnoreNodeID:   strconv.Itoa(imageInfo.ID),
			IgnoreNodeName: imageInfo.Name,
		},
	}}
	if err := c.addTempAction(data); err != nil {
		return nil, errors.Errorf("RemoveImageFromContainer: %s", err)
	}
	return c.SaveTopology()
}

// DeployDevice Move a device from the undefined container to a target container.
// Optionally, apply device-specific configlets to the device.
func (c CvpRestAPI) DeployDevice(appName string, dev *NetElement, devTargetIP string,
	container *Container, configlets ...Configlet) (*TaskInfo, error) {
	return c.DeployDeviceWithImage(appName, dev, devTargetIP, container, "", configlets...)
}

// DeployDeviceWithImage Move a device from the undefined container to a target container
// and apply image. Optionally, apply device-specific configlets to the device.
func (c CvpRestAPI) DeployDeviceWithImage(appName string, dev *NetElement, devTargetIP string,
	container *Container, image string, configlets ...Configlet) (*TaskInfo, error) {

	if _, err := c.MoveDeviceToContainer(appName, dev, container, false); err != nil {
		c.ClearAllTempActions()
		return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
	}

	applyConfiglets, err := c.GenerateHierarchicalConfiglets(dev, container)
	if err != nil {
		return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
	}

	conf, err := c.GetTempConfigByNetElementID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
	}

	applyConfiglets = append(applyConfiglets, conf.ProposedConfiglets...)
	if configlets != nil {
		applyConfiglets = append(applyConfiglets, configlets...)
	}

	curConfiglets, err := c.GetConfigletsByDeviceID(dev.SystemMacAddress)
	if err != nil {
		return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
	}

	_, configletAndBuilders, err := checkConfigMapping(curConfiglets,
		applyConfiglets)
	if err != nil {
		return nil, errors.Wrap(err, "DeployDeviceWithImage")
	}

	info := appName + ": DeployDevice: Configlet Assign - Device " + dev.Fqdn
	infoPreview := "<b>DeployDevice: Configlet Assign</b> - Device" + dev.Fqdn

	data := struct {
		Data []Action `json:"data,omitempty"`
	}{Data: []Action{
		{
			Info:                            info,
			InfoPreview:                     infoPreview,
			Note:                            "",
			Action:                          "associate",
			NodeType:                        "configlet",
			NodeID:                          "",
			ConfigletBuilderList:            configletAndBuilders.bKeys,
			ConfigletBuilderNamesList:       configletAndBuilders.bNames,
			ConfigletList:                   configletAndBuilders.keys,
			ConfigletNamesList:              configletAndBuilders.names,
			IgnoreConfigletBuilderNamesList: []string{},
			IgnoreConfigletBuilderList:      []string{},
			IgnoreConfigletNamesList:        []string{},
			IgnoreConfigletList:             []string{},
			ToID:                            dev.SystemMacAddress,
			ToIDType:                        "netelement",
			FromID:                          "",
			NodeIPAddress:                   dev.IPAddress,
			NodeName:                        "",
			// The target IP needs to be set. AssignConfigletToDevice() doesn't perform
			// this thus we setup the Action here.
			NodeTargetIPAddress: devTargetIP,
			FromName:            "",
			ToName:              dev.Fqdn,
			ChildTasks:          []string{},
			ParentTask:          "",
		},
	}}

	if err := c.addTempAction(data); err != nil {
		c.ClearAllTempActions()
		return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
	}

	if image != "" {
		imageBundle, err := c.GetImageBundleByName(image)
		if err != nil {
			return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
		}
		if _, err = c.ApplyImageToDevice(appName, imageBundle, dev, false); err != nil {
			c.ClearAllTempActions()
			return nil, errors.Errorf("DeployDeviceWithImage: %s", err)
		}
	}

	// Clear the temp Actions if we have issues
	var taskInfo *TaskInfo
	if taskInfo, err = c.SaveTopology(); err != nil {
		c.ClearAllTempActions()
		return nil, errors.Wrap(err, "DeployDeviceWithImage")
	}
	return taskInfo, nil
}

// GenerateHierarchicalConfiglets ...
func (c CvpRestAPI) GenerateHierarchicalConfiglets(dev *NetElement,
	container *Container) ([]Configlet, error) {
	var applyConfiglets []Configlet

	if dev == nil {
		return nil, errors.Errorf("GenerateHierarchicalConfiglets: nil NetElement")
	}
	if container == nil {
		return nil, errors.Errorf("GenerateHierarchicalConfiglets: nil Container")
	}

	// The the hierarchical Configlet BuilderInfo
	cblInfoList, err := c.GetHierarchicalConfigletBuilders(container)
	if err != nil {
		return nil, errors.Wrap(err, "GenerateHierarchicalConfiglets")
	}

	// Generate configlets using the builders
	for _, builder := range cblInfoList.BuildMapperList {
		cb, err := c.GetConfigletBuilderByKey(builder.BuilderID)
		if err != nil {
			return nil, errors.Wrap(err, "GenerateHierarchicalConfiglets")
		}
		// Builders with FormLists or SSLConfig skip
		if len(cb.FormList) != 0 || cb.SSLConfig {
			continue
		}

		cbInfo, err := c.GenerateAutoConfiglet([]string{dev.SystemMacAddress}, builder.BuilderID,
			container.Key, "netelement")
		if err != nil {
			return nil, errors.Wrap(err, "GenerateHierarchicalConfiglets")
		}
		if len(cbInfo) != 1 {
			return nil, errors.Errorf("GenerateHierarchicalConfiglets: No generated Configlet "+
				"for builder [%s]", builder.BuilderName)
		}
		applyConfiglets = append(applyConfiglets, cbInfo[0].Configlet)
	}
	return applyConfiglets, nil
}

// TempConfig ...
type TempConfig struct {
	ExistingConfiglets        []string    `json:"existingConfiglets"`
	IgnoredConfiglets         []string    `json:"ignoredConfiglets"`
	AssignedConfiglets        []string    `json:"assignedConfiglets"`
	ProposedConfiglets        []Configlet `json:"proposedConfiglets"`
	DeviceConfigletBuilders   []string    `json:"deviceConfigletBuilders"`
	AssignedConfigletBuilders []string    `json:"assignedConfigletBuilders"`
	IgnoredConfigletBuilders  []string    `json:"ignoredConfigletBuilders"`
	DeviceConfiglets          []string    `json:"deviceConfiglets"`
	ExistingConfigletBuilders []string    `json:"existingConfigletBuilders"`

	ErrorResponse
}

// GetTempConfigByNetElementID gets the current temporary config for the supplied netElement
func (c CvpRestAPI) GetTempConfigByNetElementID(netElementID string) (*TempConfig, error) {
	var resp TempConfig
	query := &url.Values{
		"netElementId": {netElementID},
	}

	reqResp, err := c.client.Get("/cvpservice/provisioning/getTempConfigsByNetElementId.do", query)
	if err != nil {
		return nil, errors.Errorf("GetTempConfigByNetElementID: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetTempConfigByNetElementID: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetTempConfigByNetElementID: %s", err)
	}
	return &resp, nil

}

// ClearAllTempActions clears outstanding actions
func (c CvpRestAPI) ClearAllTempActions() (string, error) {
	var resp struct {
		Data string `json:"data"`
		ErrorResponse
	}

	reqResp, err := c.client.Delete("/cvpservice/ztp/deleteAllTempAction.do", nil, nil)
	if err != nil {
		return "", errors.Errorf("GetTempActions: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return "", errors.Errorf("GetTempActions: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return "", errors.Errorf("GetTempActions: %s", err)
	}
	return resp.Data, nil

}

// GetAllTempActions gets the list of current actions outstanding
func (c CvpRestAPI) GetAllTempActions(start, end int) ([]Action, error) {
	var resp struct {
		Total int
		Data  []Action
		ErrorResponse
	}

	query := &url.Values{
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	reqResp, err := c.client.Get("/cvpservice/provisioning/getAllTempActions.do", query)
	if err != nil {
		return nil, errors.Errorf("GetTempActions: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("GetTempActions: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("GetTempActions: %s", err)
	}
	return resp.Data, nil

}

// GetTempAction returns the first outstanding action
func (c CvpRestAPI) GetTempAction() (*Action, error) {
	results, err := c.GetAllTempActions(0, 1)
	if err != nil {
		return nil, errors.Errorf("GetTempAction: %s", err)
	}
	if len(results) > 0 {
		return &results[0], nil
	}
	return nil, nil
}

// FilterTopologyWithRange filters the topology for items matching the query parameter
// and returning those within the specified range.
func (c CvpRestAPI) FilterTopologyWithRange(nodeID, querystr, format string, start int,
	end int) (*Topology, error) {

	query := &url.Values{
		"nodeId":     {nodeID},
		"queryParam": {querystr},
		"format":     {format},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp := struct {
		Topology Topology `json:"topology"`
		Type     string   `json:"type"`

		ErrorResponse
	}{}

	reqResp, err := c.client.Get("/cvpservice/ztp/filterTopology.do", query)
	if err != nil {
		return nil, errors.Errorf("FilterTopologyWithRange: %s", err)
	}

	if err = json.Unmarshal(reqResp, &resp); err != nil {
		return nil, errors.Errorf("FilterTopologyWithRange: %s Payload:\n%s", err, reqResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("FilterTopologyWithRange: %s", err)
	}
	return &resp.Topology, nil
}

// FilterTopology filters the topology for items matching the query parameter.
func (c CvpRestAPI) FilterTopology(nodeID, query string) (*Topology, error) {
	return c.FilterTopologyWithRange(nodeID, query, "topology", 0, 0)
}

// configletUnion Do the set configletUnion of c1 and c2
// elements not in c1, but are in c2 are appended to c1 and returned
func configletUnion(c1, c2 []Configlet) []Configlet {
	configletMap := make(map[string]string)
	for _, configlet := range c1 {
		configletMap[configlet.Key] = configlet.Name
	}

	for _, c := range c2 {
		if _, found := configletMap[c.Key]; !found {
			c1 = append(c1, c)
		}
	}

	return c1
}

// configletDifference Do the set difference of c1 and c2.
// Elements of c1 are returned only if they are not in c2.
func configletDifference(c1, c2 []Configlet) []Configlet {
	rmKeys := map[string]Configlet{}
	for _, c := range c2 {
		rmKeys[c.Key] = c
	}

	var configletsToRemain []Configlet

	for _, c := range c1 {
		if _, found := rmKeys[c.Key]; !found {
			configletsToRemain = append(configletsToRemain, c)
		}
	}

	return configletsToRemain
}

// checkConfigMapping Checks whether the new configlets to be applied are
// already applied or not. Returns actionReqd ( bool ), configletMap,
// and builderMap.
func checkConfigMapping(applied []Configlet, newconfiglets []Configlet) (bool,
	configletAndBuilderKeyNames, error,
) {
	configlets := configletUnion(applied, newconfiglets)

	configletAndBuilders, err := splitToConfigletAndBuilder(configlets)
	if err != nil {
		return false, configletAndBuilders, err
	}

	actionReqd := len(configlets) != len(applied)

	return actionReqd, configletAndBuilders, nil
}

func splitToConfigletAndBuilder(configlets []Configlet) (configletAndBuilderKeyNames, error) {
	var (
		configletNames     []string
		configletKeys      []string
		builderNames       []string
		builderKeys        []string
		reconcileConfiglet *Configlet
	)

	for i, configlet := range configlets {
		// Reconcile configlet must be last in the list
		// Store it separately to be appended at very end of current and new configlets
		if configlet.Reconciled {
			reconcileConfiglet = &configlets[i]

			continue
		}

		switch configlet.Type {
		case "Static", "Generated", "Reconciled":
			configletNames = append(configletNames, configlet.Name)
			configletKeys = append(configletKeys, configlet.Key)

		case "Builder":
			builderNames = append(builderNames, configlet.Name)
			builderKeys = append(builderKeys, configlet.Key)

		default:
			return configletAndBuilderKeyNames{},
				errors.Errorf("Configlet [%s] Invalid Type [%s]", configlet.Name, configlet.Type)
		}
	}

	// Add reconcileConfiglet as last in the list
	if reconcileConfiglet != nil {
		configletNames = append(configletNames, reconcileConfiglet.Name)
		configletKeys = append(configletKeys, reconcileConfiglet.Key)
	}

	return configletAndBuilderKeyNames{
		configletKeys,
		configletNames,
		builderKeys,
		builderNames,
	}, nil
}

// checkRemoveConfigMapping Creates map of configlets that needs to be there after removal of
// specific configlets
func checkRemoveConfigMapping(applied []Configlet, rmConfiglets []Configlet) (bool,
	configletAndBuilderKeyNames, configletAndBuilderKeyNames, error,
) {
	configletsToRemain := configletDifference(applied, rmConfiglets)

	rmConfigletAndBuilders, err := splitToConfigletAndBuilder(rmConfiglets)
	if err != nil {
		return false, configletAndBuilderKeyNames{}, configletAndBuilderKeyNames{}, err
	}

	configletAndBuilders, err2 := splitToConfigletAndBuilder(configletsToRemain)
	if err2 != nil {
		return false, configletAndBuilderKeyNames{}, configletAndBuilderKeyNames{}, err2
	}

	actionReqd := len(configletsToRemain) != len(applied)

	return actionReqd, configletAndBuilders, rmConfigletAndBuilders, nil
}

func (c CvpRestAPI) getConfigletKeys(configletNames []string) ([]string, error) {
	configletInfo, err := c.GetConfiglets()
	if err != nil {
		return nil, errors.Wrap(err, "getConfigletKeys")
	}

	configletMap := make(map[string]string, len(configletInfo))
	configletKeys := make([]string, len(configletNames))

	// create our mapping for name to key
	for _, configlet := range configletInfo {
		configletMap[configlet.Name] = configlet.Key
	}

	// create our list of keys
	for index, configletName := range configletNames {
		var configletKey string
		var found bool

		if configletKey, found = configletMap[configletName]; !found {
			return nil, errors.Errorf("getConfigletKeys: Invalid configlet name [%s]",
				configletName)
		}
		configletKeys[index] = configletKey
	}
	return configletKeys, nil
}