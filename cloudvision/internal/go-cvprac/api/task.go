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

// CvpTask represents a task
type CvpTask struct {
	TemplateID                 string          `json:"templateId"`
	CurrentTaskType            string          `json:"currentTaskType"`
	TaskStatus                 string          `json:"taskStatus"`
	CurrentTaskName            string          `json:"currentTaskName"`
	ExecutedBy                 string          `json:"executedBy"`
	WorkOrderUserDefinedStatus string          `json:"workOrderUserDefinedStatus"`
	WorkOrderID                string          `json:"workOrderId"`
	WorkOrderState             string          `json:"workOrderState"`
	CreatedBy                  string          `json:"createdBy"`
	ExecutedOnInLongFormat     int64           `json:"executedOnInLongFormat"`
	WorkOrderDetails           WorkOrderDetail `json:"workOrderDetails"`
	CreatedOnInLongFormat      int64           `json:"createdOnInLongFormat"`
	Data                       WorkData        `json:"data"`
	Description                string          `json:"description"`
	Note                       string          `json:"note"`
	Name                       string          `json:"name"`
	ID                         int             `json:"id"`

	ErrorResponse
}

// WorkOrderDetail associated with a task
type WorkOrderDetail struct {
	NetElementID       string `json:"netElementId"`
	NetElementHostName string `json:"netElementHostName"`
	IPAddress          string `json:"ipAddress"`
	FactoryID          int    `json:"factoryId"`
	SerialNumber       string `json:"serialNumber"`
	ID                 int    `json:"id"`
}

// WorkData associated with a task
type WorkData struct {
	WorkFlowAction           string   `json:"WORKFLOW_ACTION"`
	CurrentparentContainerID string   `json:"currentparentContainerId"`
	View                     string   `json:"VIEW"`
	NewparentContainerID     string   `json:"newparentContainerId"`
	NetElementID             string   `json:"NETELEMENT_ID"`
	IsConfigPushNeeded       string   `json:"IS_CONFIG_PUSH_NEEDED"`
	IgnoreConfigletList      []string `json:"ignoreConfigletList"`
}

// CvpTaskList represents a task list returned from a request
type CvpTaskList struct {
	Total int       `json:"total"`
	Data  []CvpTask `json:"data"`

	ErrorResponse
}

// LogData log data associated with a task
type LogData struct {
	ClassID              int    `json:"classId"`
	DateTimeInLongFormat int64  `json:"dateTimeInLongFormat"`
	LogDetails           string `json:"logDetails"`
	WorkOrderID          string `json:"workOrderId"`
	FactoryID            int    `json:"factoryId"`
	ObjectName           string `json:"objectName"`
	UserName             string `json:"userName"`
	Key                  string `json:"key"`
	ID                   int    `json:"id"`
}

// CvpLogList represents the log list for a specific task
type CvpLogList struct {
	Total int       `json:"total"`
	Data  []LogData `json:"data"`

	ErrorResponse
}

// GetTaskByID returns the current Task for the specified taskID.
func (c CvpRestAPI) GetTaskByID(taskID int) (*CvpTask, error) {
	var info CvpTask

	query := &url.Values{"taskId": {strconv.Itoa(taskID)}}

	resp, err := c.client.Get("/cvpservice/task/getTaskById.do", query)
	if err != nil {
		return nil, errors.Errorf("GetTaskByID: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetTaskByID: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetTaskByID: %s", err)
	}

	return &info, nil
}

// GetTasks returns the current CVP Tasks that match the provided string
// and within the provided start/end range.
func (c CvpRestAPI) GetTasks(queryStr string, start int, end int) ([]CvpTask, error) {
	var info CvpTaskList
	query := &url.Values{
		"queryparam": {queryStr},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/workflow/getTasks.do", query)
	if err != nil {
		return nil, errors.Errorf("GetTasks: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetTasks: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetTasks: %s", err)
	}

	return info.Data, nil
}

// GetTaskByStatus returns a list of all tasks with the given status.
func (c CvpRestAPI) GetTaskByStatus(status string) ([]CvpTask, error) {
	return c.GetTasks(status, 0, 0)
}

// GetAllTasks returns a list of all the tasks.
func (c CvpRestAPI) GetAllTasks() ([]CvpTask, error) {
	return c.GetTasks("", 0, 0)
}

// GetLogs returns the log entries for the task with the specified taskID and
// within the provide start/end range.
func (c CvpRestAPI) GetLogs(taskID int, start int, end int) ([]LogData, error) {
	var info CvpLogList
	query := &url.Values{
		"id":         {strconv.Itoa(taskID)},
		"queryparam": {""},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/task/getLogsById.do", query)
	if err != nil {
		return nil, errors.Errorf("GetLogs: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetLogs: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetLogs: %s", err)
	}

	return info.Data, nil
}

// GetLogsByID returns the log entries for the task with the specified taskID.
func (c CvpRestAPI) GetLogsByID(taskID int) ([]LogData, error) {
	return c.GetLogs(taskID, 0, 0)
}

// AddNoteToTask adds a note to the task represented by taskID
func (c CvpRestAPI) AddNoteToTask(taskID int, note string) error {
	var info ErrorResponse

	data := map[string]string{
		"workOrderId": strconv.Itoa(taskID),
		"note":        note,
	}
	resp, err := c.client.Post("/cvpservice/task/addNoteToTask.do", nil, data)
	if err != nil {
		return errors.Errorf("AddNoteToTask: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("AddNoteToTask: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("AddNoteToTask: %s", err)
	}
	return nil
}

// ExecuteTask executes a task given the taskID.
func (c CvpRestAPI) ExecuteTask(taskID int) error {
	return c.ExecuteTasks([]int{taskID})
}

// ExecuteTasks executes a task given the taskID.
func (c CvpRestAPI) ExecuteTasks(taskID []int) error {
	var info ErrorResponse
	var taskIDs []string

	for _, task := range taskID {
		taskIDs = append(taskIDs, strconv.Itoa(task))
	}

	data := map[string][]string{
		"data": taskIDs,
	}
	resp, err := c.client.Post("/cvpservice/workflow/executeTask.do", nil, data)
	if err != nil {
		return errors.Errorf("ExecuteTask: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("ExecuteTask: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("ExecuteTask: %s", err)
	}
	return nil
}

// CancelTask cancels the task given the taskID
func (c CvpRestAPI) CancelTask(taskID int) error {
	return c.CancelTasks([]int{taskID})
}

// CancelTasks cancels the list of taskIDs
func (c CvpRestAPI) CancelTasks(taskID []int) error {
	var info ErrorResponse
	var taskIDs []string

	for _, task := range taskID {
		taskIDs = append(taskIDs, strconv.Itoa(task))
	}

	data := map[string][]string{
		"data": taskIDs,
	}
	resp, err := c.client.Post("/cvpservice/task/cancelTask.do", nil, data)
	if err != nil {
		return errors.Errorf("CancelTask: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return errors.Errorf("CancelTask: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return errors.Errorf("CancelTask: %s", err)
	}
	return nil
}
