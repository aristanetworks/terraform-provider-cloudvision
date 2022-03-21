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
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ConfigletList represents a list of configlets
type ConfigletList struct {
	Total int         `json:"total"`
	Data  []Configlet `json:"data"`

	ErrorResponse
}

// Configlet represents a Configlet
type Configlet struct {
	IsDefault            string `json:"isDefault"`
	DateTimeInLongFormat int64  `json:"dateTimeInLongFormat"`
	ContainerCount       int    `json:"containerCount"`
	NetElementCount      int    `json:"netElementCount"`
	IsAutoBuilder        string `json:"isAutoBuilder"`
	Reconciled           bool   `json:"reconciled"`
	FactoryID            int    `json:"factoryId"`
	Config               string `json:"config"`
	User                 string `json:"user"`
	Note                 string `json:"note"`
	Name                 string `json:"name"`
	Key                  string `json:"key"`
	ID                   int    `json:"id"`
	Type                 string `json:"type"`

	ErrorResponse
}

func (c Configlet) String() string {
	return c.Name
}

// ConfigletHistoryEntry represents a configlet history entry
type ConfigletHistoryEntry struct {
	ConfigletID                 string `json:"configletId"`
	OldUserID                   string `json:"oldUserId"`
	NewUserID                   string `json:"newUserId"`
	OldConfig                   string `json:"oldConfig"`
	NewConfig                   string `json:"newConfig"`
	OldDate                     string `json:"oldDate"`
	NewDate                     string `json:"newDate"`
	OldDateTimeInLongFormat     int64  `json:"oldDateTimeInLongFormat"`
	UpdatedDateTimeInLongFormat int64  `json:"updatedDateTimeInLongFormat"`
	FactoryID                   int    `json:"factoryId"`
	Key                         string `json:"key"`
	ID                          int    `json:"id"`
}

// ConfigletHistoryList represents a list of ConfigletHistoryEntry's
type ConfigletHistoryList struct {
	Total       int                     `json:"total"`
	HistoryList []ConfigletHistoryEntry `json:"configletHistory"`

	ErrorResponse
}

// ConfigletOpReturn represents the
type ConfigletOpReturn struct {
	Data Configlet `json:"data"`

	ErrorResponse
}

// ConfigletUpdateReturn ...
type ConfigletUpdateReturn struct {
	Data    string   `json:"data"`
	TaskIDs []string `json:"taskIds"`

	ErrorResponse
}

// ConfigletVerifyResp represents
type ConfigletVerifyResp struct {
	ID      string `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  []struct {
		Output   string   `json:"output"`
		Messages []string `json:"messages"`
	} `json:"result"`
	Warnings     []string      `json:"warnings"`
	WarningCount int           `json:"warningCount"`
	Errors       []VerifyError `json:"errors"`
	ErrorCount   int           `json:"errorCount"`
}

// VerifyError represents an error related to verification of a config
type VerifyError struct {
	LineNo string `json:"lineNo"`
	Error  string `json:"error"`
}

// GetConfigletsInfo returns configlet info
func (c CvpRestAPI) GetConfigletsInfo(start int, end int) ([]Configlet, error) {
	var info ConfigletList

	query := &url.Values{
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/configlet/getConfiglets.do", query)
	if err != nil {
		return nil, errors.Wrap(err, "GetConfigletsInfo")
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetConfigletsInfo: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Wrap(err, "GetConfigletsInfo")
	}
	return info.Data, nil
}

// GetConfiglets returns configlet info
func (c CvpRestAPI) GetConfiglets() ([]Configlet, error) {
	return c.GetConfigletsInfo(0, 0)
}

// GetConfigletByName returns the configlet with the specified name
func (c CvpRestAPI) GetConfigletByName(name string) (*Configlet, error) {
	var info Configlet

	query := &url.Values{"name": {name}}

	resp, err := c.client.Get("/cvpservice/configlet/getConfigletByName.do", query)
	if err != nil {
		return nil, errors.Errorf("GetConfigletByName: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetConfigletByName: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		// Entity does not exist
		if info.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetConfigletByName: %s", err)
	}
	return &info, nil
}

// GetConfigletByID returns the configlet with the specified ID
func (c CvpRestAPI) GetConfigletByID(ID string) (*Configlet, error) {
	var info Configlet

	query := &url.Values{"id": {ID}}

	resp, err := c.client.Get("/cvpservice/configlet/getConfigletById.do", query)
	if err != nil {
		return nil, errors.Errorf("GetConfigletByID: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetConfigletByID: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		// Entity does not exist
		if info.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetConfigletByID: %s", err)
	}
	return &info, nil
}

// GetConfigletHistory returns the history for a configlet provided the key, and a range.
func (c CvpRestAPI) GetConfigletHistory(key string, start int,
	end int) (*ConfigletHistoryList, error) {
	var info ConfigletHistoryList

	//queryparam := url.Values{"name": {key},}

	query := &url.Values{
		"configletId": {key},
		"queryparam":  {""},
		"startIndex":  {strconv.Itoa(start)},
		"endIndex":    {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/configlet/getConfigletHistory.do", query)
	if err != nil {
		return nil, errors.Errorf("GetConfigletHistory: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetConfigletHistory: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetConfigletHistory: %s", err)
	}

	return &info, nil
}

// GetAllConfigletHistory returns all the history for a given configlet
func (c CvpRestAPI) GetAllConfigletHistory(key string) (*ConfigletHistoryList, error) {
	return c.GetConfigletHistory(key, 0, 0)
}

// AddConfiglet creates/adds a configlet
func (c CvpRestAPI) AddConfiglet(name string, config string) (*Configlet, error) {
	var info ConfigletOpReturn

	data := map[string]string{
		"name":   name,
		"config": config,
	}

	resp, err := c.client.Post("/cvpservice/configlet/addConfiglet.do", nil, data)
	if err != nil {
		return nil, errors.Errorf("AddConfiglet: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("AddConfiglet: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("AddConfiglet: %s", err)
	}

	return &info.Data, nil
}

// DeleteConfiglet deletes a configlet.
func (c CvpRestAPI) DeleteConfiglet(name string, key string) error {
	var info ErrorResponse

	data := []map[string]string{
		{
			"name": name,
			"key":  key,
		},
	}
	resp, err := c.client.Post("/cvpservice/configlet/deleteConfiglet.do", nil, data)
	if err != nil {
		return errors.Errorf("DeleteConfiglet: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("DeleteConfiglet: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("DeleteConfiglet: %s", err)
	}

	return nil
}

// updateConfiglet updates a configlet.
func (c CvpRestAPI) updateConfiglet(config string, name string, key string,
	waitForTaskIds bool) (*ConfigletUpdateReturn, error) {
	var info ConfigletUpdateReturn

	data := struct {
		Config         string `json:"config"`
		Key            string `json:"key"`
		Name           string `json:"name"`
		WaitForTaskIds bool   `json:"waitForTaskIds,omitempty"`
	}{
		Config:         config,
		Key:            key,
		Name:           name,
		WaitForTaskIds: waitForTaskIds,
	}

	resp, err := c.client.Post("/cvpservice/configlet/updateConfiglet.do", nil, data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Wrapf(err, "updateConfiglet: Body: %s", resp)
	}

	if err := info.Error(); err != nil {
		return nil, err
	}

	return &info, nil
}

// UpdateConfiglet updates a configlet.
func (c CvpRestAPI) UpdateConfiglet(config string, name string, key string) error {
	_, err := c.updateConfiglet(config, name, key, false)
	if err != nil {
		return errors.Errorf("UpdateConfiglet: %s", err)
	}
	return nil
}

// UpdateConfigletWaitForTask updates a configlet and waits for tasks to be returned.
func (c CvpRestAPI) UpdateConfigletWaitForTask(config string, name string, key string) ([]string,
	error) {
	data, err := c.updateConfiglet(config, name, key, true)
	if err != nil {
		return nil, errors.Errorf("UpdateConfigletWaitForTask: %s", err)
	}
	return data.TaskIDs, nil
}

// AddConfigletNote creates/adds a configlet note
func (c CvpRestAPI) AddConfigletNote(key string, note string) error {
	data := map[string]string{
		"key":  key,
		"note": note,
	}

	resp, err := c.client.Post("/cvpservice/configlet/addNoteToConfiglet.do", nil, data)
	if err != nil {
		return errors.Errorf("AddConfigletNote: %s", err)
	}

	info := struct {
		Data string `json:"data"`

		ErrorResponse
	}{}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("AddConfigletNote: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("AddConfigletNote: %s", err)
	}

	return nil
}

// VerifyConfig verifies a configlet config config
func (c CvpRestAPI) VerifyConfig(netElement string, config string) error {
	var info ConfigletVerifyResp
	data := map[string]string{
		"config":       config,
		"netElementId": netElement,
	}

	resp, err := c.client.Post("/cvpservice/configlet/validateConfig.do", nil, data)
	if err != nil {
		return errors.Wrap(err, "VerifyConfig")
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("VerifyConfig: %s Payload:\n%s", err, resp)
	}

	if info.ErrorCount != 0 {
		var msg []string
		for _, e := range info.Errors {
			msg = append(msg, fmt.Sprintf("LineNo:%s - [%s]", e.LineNo, e.Error))
		}
		return errors.Errorf("%s", strings.Join(msg, ", "))
	}
	return nil
}

// SearchConfigletsWithRange search function for configlets.
func (c CvpRestAPI) SearchConfigletsWithRange(searchStr string, start int,
	end int) (*ConfigletList, error) {
	var info ConfigletList

	query := &url.Values{
		"queryparam": {searchStr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/configlet/searchConfiglets.do", query)
	if err != nil {
		return nil, errors.Errorf("SearchConfiglets: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("SearchConfiglets: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("SearchConfiglets: %s", err)
	}

	return &info, nil
}

// SearchConfiglets search function for configlets.
func (c CvpRestAPI) SearchConfiglets(searchStr string) (*ConfigletList, error) {
	return c.SearchConfigletsWithRange(searchStr, 0, 0)
}

// GetAppliedDevices Returns a list of devices to which the named configlet is applied
func (c CvpRestAPI) GetAppliedDevices(configletName string) ([]ObjectInfo, error) {
	return c.GetAppliedDevicesWithRange(configletName, 0, 0)
}

// GetAppliedDevicesWithRange Returns a list of devices to which the named configlet is applied
func (c CvpRestAPI) GetAppliedDevicesWithRange(configletName string, start int,
	end int) ([]ObjectInfo, error) {
	var info GenericReq

	query := &url.Values{
		"configletName": {configletName},
		"startIndex":    {strconv.Itoa(start)},
		"endIndex":      {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/configlet/getAppliedDevices.do", query)
	if err != nil {
		return nil, errors.Errorf("GetAppliedDevices: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetAppliedDevices: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetAppliedDevices: %s", err)
	}

	return info.Data, nil
}
