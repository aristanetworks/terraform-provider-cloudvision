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

func Test_CvpGetAllDevicesRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetAllDevices: GetInventory: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllDevices()
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: [%v] Got: [%v]", expectedErr, err)
	}
}

func Test_CvpGetAllDevicesJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllDevices(); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllDevicesEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllDevices(); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllDevicesReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllDevices(); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetAllDevicesValid_UnitTest(t *testing.T) {
	client := NewMockClient("[]", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllDevices()
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

func Test_CvpGetDeviceByNameRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetDeviceByName: GetInventory: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetDeviceByName("DeviceName")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: [%v] Got: [%v]", expectedErr, err)
	}
}

func Test_CvpGetDeviceByNameJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetDeviceByName("DeviceName"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetDeviceByNameEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetDeviceByName("DeviceName"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetDeviceByNameReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetDeviceByName("DeviceName"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetDeviceByNameValid_UnitTest(t *testing.T) {
	client := NewMockClient("[]", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetDeviceByName("DeviceName")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}
