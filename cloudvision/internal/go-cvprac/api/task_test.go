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
	"errors"
	"testing"
)

func Test_CvpGetTaskByIDRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetTaskByID: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetTaskByID(5)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpGetTaskByIDJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetTaskByID(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetTaskByIDEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetTaskByID(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetTaskByIDReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetTaskByID(5); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetTaskByIDValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetTaskByID(1)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

////////////////////

func Test_CvpGetTaskByStatusRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetTasks: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetTaskByStatus("Pending")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpGetTaskByStatusJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetTaskByStatus("Pending"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetTaskByStatusEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetTaskByStatus("Pending"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetTaskByStatusReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetTaskByStatus("Pending"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetTaskByStatusValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetTaskByStatus("Pending")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

//////////////////////

func Test_CvpGetAllTasksRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetTasks: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllTasks()
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpGetAllTasksJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllTasks(); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllTasksEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllTasks(); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllTasksReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllTasks(); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetAllTasksValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllTasks()
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

////////////////////////

func Test_CvpGetLogsByIDRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetLogs: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetLogsByID(5)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpGetLogsByIDJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetLogsByID(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetLogsByIDEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetLogsByID(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetLogsByIDReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetLogsByID(5); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetLogsByIDValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetLogsByID(1)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

///////////////////////

func Test_CvpAddNoteToTaskRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("AddNoteToTask: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	err := api.AddNoteToTask(5, "note string")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpAddNoteToTaskJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if err := api.AddNoteToTask(5, "note string"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpAddNoteToTaskEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if err := api.AddNoteToTask(5, "note string"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpAddNoteToTaskReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if err := api.AddNoteToTask(5, "note string"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpAddNoteToTaskValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	err := api.AddNoteToTask(5, "note string")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

//////////////////////

func Test_CvpExecuteTaskRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("ExecuteTask: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	err := api.ExecuteTask(5)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpExecuteTaskJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if err := api.ExecuteTask(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpExecuteTaskEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if err := api.ExecuteTask(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpExecuteTaskReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if err := api.ExecuteTask(5); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpExecuteTaskValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	err := api.ExecuteTask(1)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

/////////////////////

func Test_CvpCancelTaskRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("CancelTask: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	err := api.CancelTask(5)
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpCancelTaskJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if err := api.CancelTask(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpCancelTaskEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if err := api.CancelTask(5); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpCancelTaskReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if err := api.CancelTask(5); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpCancelTaskValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	err := api.CancelTask(1)
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}
