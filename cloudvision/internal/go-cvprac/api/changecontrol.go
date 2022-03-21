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

	"github.com/pkg/errors"
)

// ChangeControlTask represents a task
type ChangeControlTask struct {
	ContainerName              string `json:"containerName"`
	CurrentTaskName            string `json:"currentTaskName"`
	Description                string `json:"description"`
	CreatedOnInLongFormat      int64  `json:"createdOnInLongFormat"`
	WorkOrderID                string `json:"workOrderId"`
	NetElementHostName         string `json:"netElementHostName"`
	Note                       string `json:"note"`
	NetElementID               string `json:"netElementId"`
	CreatedBy                  string `json:"createdBy"`
	TemplateID                 string `json:"templateId"`
	ExecutedOnInLongFormat     int64  `json:"executedOnInLongFormat"`
	ExecutedBy                 string `json:"executedBy"`
	WorkOrderUserDefinedStatus string `json:"workOrderUserDefinedStatus"`
	IPAddress                  string `json:"ipAddress"`
	Model                      string `json:"model"`
	WorkOrderState             string `json:"workOrderState"`
	CcID                       string `json:"ccId"`

	ErrorResponse
}

// ChangeControlTaskInfo represents task info to be used for applying a task
// to an existing or new change control.
type ChangeControlTaskInfo struct {
	TaskID              string `json:"taskId"`
	TaskOrder           int    `json:"taskOrder"`
	SnapshotTemplateKey string `json:"snapshotTemplateKey"`
	ClonedCcID          string `json:"clonedCcId"`
}

// ChangeControl represents a CVP Change Control object
type ChangeControl struct {
	CreatedBy             string `json:"createdBy"`
	FactoryID             int    `json:"factoryId"`
	TaskEndTime           int64  `json:"taskEndTime"`
	DeviceCount           int    `json:"deviceCount"`
	StopOnError           bool   `json:"stopOnError"`
	ID                    int    `json:"id"`
	ScheduledTimestamp    int64  `json:"scheduledTimestamp"`
	CountryID             string `json:"countryId"`
	CcID                  string `json:"ccId"`
	PreSnapshotEndTime    int64  `json:"preSnapshotEndTime"`
	TaskCount             int    `json:"taskCount"`
	TaskStartTime         int64  `json:"taskStartTime"`
	ScheduledBy           string `json:"scheduledBy"`
	ScheduledByPassword   string `json:"scheduledByPassword"`
	Type                  string `json:"type"`
	Status                string `json:"status"`
	PreSnapshotStartTime  int64  `json:"preSnapshotStartTime"`
	PostSnapshotEndTime   int64  `json:"postSnapshotEndTime"`
	CreatedTimestamp      int64  `json:"createdTimestamp"`
	Key                   string `json:"key"`
	CcName                string `json:"ccName"`
	ExecutedTimestamp     int64  `json:"executedTimestamp"`
	ExecutedBy            string `json:"executedBy"`
	ClassID               int    `json:"classId"`
	ContainerName         string `json:"containerName"`
	StopOnErrorStatus     string `json:"stopOnErrorStatus"`
	Notes                 string `json:"notes"`
	DateTime              string `json:"dateTime"`
	PostSnapshotStartTime int64  `json:"postSnapshotStartTime"`
	ContainerKey          string `json:"containerKey"`
	TimeZone              string `json:"timeZone"`

	ErrorResponse
}

// ChangeControlList represents a return data for getting a list of all change controls
type ChangeControlList struct {
	Total int             `json:"total"`
	Data  []ChangeControl `json:"data"`

	ErrorResponse
}

// ChangeControlTaskList represents a return data for getting a list of tasks available
// for change control
type ChangeControlTaskList struct {
	Total int                 `json:"total"`
	Data  []ChangeControlTask `json:"data"`

	ErrorResponse
}

// AddOrUpdateChangeControlResp is the response returned for addOrUpdateChangeControl API call
type AddOrUpdateChangeControlResp struct {
	Data string `json:"data"`
	CcID string `json:"ccId"`

	ErrorResponse
}

// AddNotesToChangeControlResp is the response returned for addNotesToChangeControl API call
type AddNotesToChangeControlResp struct {
	Data string `json:"data"`

	ErrorResponse
}

// GetChangeControls returns a list of ChangeControls.
//
// Failed search returns empty
// {
//   "total": 0,
//   "data": []
// }
func (c CvpRestAPI) GetChangeControls(
	querystr string, start int, end int) ([]ChangeControl, error) {
	var changeControlInfo ChangeControlList
	query := &url.Values{
		"queryparam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/changeControl/getChangeControls.do",
		query)
	if err != nil {
		return nil, errors.Errorf("GetChangeControls: %s",
			err)
	}

	if err = json.Unmarshal(resp, &changeControlInfo); err != nil {
		return nil, errors.Errorf("GetChangeControls: %s Payload:\n%s",
			err, resp)

	}

	if err := changeControlInfo.Error(); err != nil {
		return nil, errors.Errorf("GetChangeControls: %s",
			err)
	}

	return changeControlInfo.Data, nil
}

// GetChangeControlAvailableTasks returns a list of ChangeControlTask's.
//
// Failed search returns empty
// {
//   "total": 0,
//   "data": []
// }
func (c CvpRestAPI) GetChangeControlAvailableTasks(
	querystr string, start int, end int) ([]ChangeControlTask, error) {
	var availableTaskInfo ChangeControlTaskList
	query := &url.Values{
		"queryparam": {querystr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/changeControl/getTasksByStatus.do",
		query)
	if err != nil {
		return nil, errors.Errorf("GetChangeControlAvailableTasks: %s",
			err)
	}

	if err = json.Unmarshal(resp, &availableTaskInfo); err != nil {
		return nil, errors.Errorf("GetChangeControlAvailableTasks: %s Payload:\n%s",
			err, resp)

	}

	if err := availableTaskInfo.Error(); err != nil {
		return nil, errors.Errorf("GetChangeControlAvailableTasks: %s",
			err)
	}

	return availableTaskInfo.Data, nil
}

// CreateChangeControl adds a note to the Change Control represented by ccID
func (c CvpRestAPI) CreateChangeControl(ccName, timeZone, countryID, dateTime, snapshotTemplateKey,
	changeControlType, stopOnError string, tasks []ChangeControlTaskInfo) (string, error) {
	var info AddOrUpdateChangeControlResp

	data := map[string]interface{}{
		"timeZone":            timeZone,
		"countryId":           countryID,
		"dateTime":            dateTime,
		"ccName":              ccName,
		"snapshotTemplateKey": snapshotTemplateKey,
		"type":                changeControlType,
		"stopOnError":         stopOnError,
		"deletedTaskIds":      []string{},
		"changeControlTasks":  tasks,
	}
	resp, err := c.client.Post("/cvpservice/changeControl/addOrUpdateChangeControl.do", nil, data)
	if err != nil {
		return "", errors.Errorf("CreateChangeControl: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return "", errors.Errorf("CreateChangeControl: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return "", errors.Errorf("CreateChangeControl: %s", err)
	}
	return info.CcID, nil
}

// AddNotesToChangeControl adds a note to the Change Control represented by ccID
func (c CvpRestAPI) AddNotesToChangeControl(ccID int, notes string) error {
	var info AddNotesToChangeControlResp

	data := map[string]string{
		"ccId":  strconv.Itoa(ccID),
		"notes": notes,
	}
	resp, err := c.client.Post("/cvpservice/changeControl/addNotesToChangeControl.do", nil, data)
	if err != nil {
		return errors.Errorf("AddNotesToChangeControl: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("AddNotesToChangeControl: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("AddNotesToChangeControl: %s", err)
	}
	return nil
}
