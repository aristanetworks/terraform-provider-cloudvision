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

	"github.com/pkg/errors"
)

// BuilderInfo represents the builders for a netelemet
type BuilderInfo struct {
	Total           int           `json:"total"`
	BuilderList     []Configlet   `json:"builderList"`
	BuildMapperList []BuilderMaps `json:"buildMapperList"`

	ErrorResponse
}

// BuilderMaps represents
type BuilderMaps struct {
	BuilderName   string `json:"builderName"`
	BuilderID     string `json:"builderId"`
	ContainerID   string `json:"containerId"`
	ContainerName string `json:"containerName"`
}

// GetHierarchicalConfigletBuilders returns the configlet with the specified key
func (c CvpRestAPI) GetHierarchicalConfigletBuilders(container *Container) (*BuilderInfo, error) {
	var info BuilderInfo

	if container == nil {
		return nil, errors.Errorf("GetHierarchicalConfigletBuilders: container nil")
	}

	query := &url.Values{
		"containerId": {container.Key},
		"queryParam":  {},
		"startIndex":  {"0"},
		"endIndex":    {"0"},
	}

	resp, err := c.client.Get("/cvpservice/configlet/getHierarchicalConfigletBuilders.do", query)
	if err != nil {
		return nil, errors.Errorf("GetHierarchicalConfigletBuilders: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetHierarchicalConfigletBuilders: %s Payload:\n%s", err, resp)
	}
	return &info, nil
}

// ConfigletBuilder represents ConfigletBuilder info
type ConfigletBuilder struct {
	IsAssigned bool          `json:"isAssigned"`
	SSLConfig  bool          `json:"sslConfig"`
	Editable   bool          `json:"editable"`
	Name       string        `json:"name"`
	FormList   []interface{} `json:"formList"`
	MainScript struct {
		Data string `json:"data"`
		Key  string `json:"key"`
	} `json:"main_script"`

	ErrorResponse
}

// ConfigletBuilderResp represents the
type ConfigletBuilderResp struct {
	Data ConfigletBuilder `json:"data"`

	ErrorResponse
}

// GetConfigletBuilderByKey returns the configlet with the specified key
func (c CvpRestAPI) GetConfigletBuilderByKey(key string) (*ConfigletBuilder, error) {
	var info ConfigletBuilderResp

	query := &url.Values{
		"type": {},
		"id":   {key},
	}
	resp, err := c.client.Get("/cvpservice/configlet/getConfigletBuilder.do", query)
	if err != nil {
		return nil, errors.Errorf("GetConfigletBuilderByKey: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetConfigletBuilderByKey: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		// Entity does not exist
		if info.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetConfigletBuilderByKey: %s", err)
	}
	return &info.Data, nil
}

// GetConfigletBuilderByName returns the configlet with the specified key
func (c CvpRestAPI) GetConfigletBuilderByName(name string) (*ConfigletBuilder, error) {
	configlet, err := c.GetConfigletByName(name)
	if err != nil {
		return nil, errors.Wrap(err, "GetConfigletBuilderByName")
	}

	builder, err := c.GetConfigletBuilderByKey(configlet.Key)
	return builder, errors.Wrap(err, "GetConfigletBuilderByName")
}

// AutoConfigletResp represents the
type AutoConfigletResp struct {
	Data []ConfigletExecStatus `json:"data"`

	ErrorResponse
}

// ConfigletExecStatus ...
type ConfigletExecStatus struct {
	Status       string    `json:"status"`
	Configlet    Configlet `json:"configlet"`
	NetElementID string    `json:"netElementId"`
	IsExistingGC bool      `json:"isExistingGc"`
	PythonError  *struct {
		ErrorMessage      string `json:"errorMessage"`
		ErrorPoint        string `json:"errorPoint"`
		ErrorPointMessage string `json:"errorPointMessage"`
		ErrorType         string `json:"errorType"`
		LineNumber        string `json:"lineNumber"`
	} `json:"pythonError"`
}

// GenerateAutoConfiglet ...
// If devKeyList is empty, then exec on all devices in container
func (c CvpRestAPI) GenerateAutoConfiglet(devKeyList []string, builderKey string,
	containerKey string, pageType string) ([]ConfigletExecStatus, error) {
	var info AutoConfigletResp

	if pageType != "netelement" && pageType != "container" {
		return nil, errors.Errorf("GenerateAutoConfiglet: pageType must be " +
			"[netelement | container]")
	}

	data := map[string]interface{}{
		"netElementIds":      devKeyList,
		"configletBuilderId": builderKey,
		"containerId":        containerKey,
		"pageType":           pageType,
	}

	resp, err := c.client.Post("/cvpservice/configlet/autoConfigletGenerator.do", nil, data)
	if err != nil {
		return nil, errors.Wrap(err, "GenerateAutoConfiglet")
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GenerateAutoConfiglet: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GenerateAutoConfiglet: %s", err)
	}

	for _, builderStatus := range info.Data {
		pyError := builderStatus.PythonError
		if pyError != nil {
			return nil, errors.Errorf("ErrorMsg [%s] Line [%s]",
				pyError.ErrorMessage, pyError.LineNumber)
		}
	}
	return info.Data, nil
}

// GenerateConfigletForDevice ...
func (c CvpRestAPI) GenerateConfigletForDevice(dev *NetElement, builder *ConfigletBuilder) (*Configlet, error) {
	if dev == nil {
		return nil, errors.Errorf("GenerateConfigletForDevice: dev nil")
	}
	if builder == nil {
		return nil, errors.Errorf("GenerateConfigletForDevice: builder ref nil")
	}
	pageType := "netelement"

	builderConfiglet, err := c.GetConfigletByName(builder.Name)
	if err != nil {
		return nil, errors.Wrap(err, "GenerateConfigletForDevice")
	}

	if len(builder.FormList) != 0 {
		return nil, errors.Errorf("GenerateConfigletForDevice: FormLists not supported")
	}

	builderStatus, err := c.GenerateAutoConfiglet([]string{dev.SystemMacAddress},
		builderConfiglet.Key,
		"", pageType)
	if err != nil {
		return nil, errors.Wrap(err, "GenerateConfigletForDevice")
	}

	configlet := builderStatus[0].Configlet

	return &configlet, nil
}

// GenerateConfigletForContainer ...
func (c CvpRestAPI) GenerateConfigletForContainer(container *Container,
	builder *ConfigletBuilder, devList []NetElement) ([]Configlet, error) {
	if container == nil {
		return nil, errors.Errorf("GenerateConfigletForContainer: container nil")
	}
	if builder == nil {
		return nil, errors.Errorf("GenerateConfigletForContainer: builder ref nil")
	}

	devMacList := make([]string, len(devList))
	for idx, dev := range devList {
		devMacList[idx] = dev.SystemMacAddress
	}

	pageType := "container"

	builderConfiglet, err := c.GetConfigletByName(builder.Name)
	if err != nil {
		return nil, errors.Wrap(err, "GenerateConfigletForContainer")
	}

	if len(builder.FormList) != 0 {
		return nil, errors.Errorf("GenerateConfigletForContainer: FormLists not supported")
	}

	builderStatus, err := c.GenerateAutoConfiglet(devMacList,
		builderConfiglet.Key,
		container.Key, pageType)
	if err != nil {
		return nil, errors.Wrap(err, "GenerateConfigletForContainer")
	}

	var configlets []Configlet
	for _, configletStatus := range builderStatus {
		configlets = append(configlets, configletStatus.Configlet)
	}
	return configlets, nil
}
