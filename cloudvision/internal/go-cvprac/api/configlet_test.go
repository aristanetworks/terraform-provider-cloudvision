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

func Test_CvpGetConfigletByNameRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetConfigletByName: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetConfigletByName("configletname")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: [%v] Got: [%v]", expectedErr, err)
	}
}

func Test_CvpGetConfigletByNameJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetConfigletByName("configletname"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetConfigletByNameEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetConfigletByName("configletname"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetConfigletByNameReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetConfigletByName("configletname"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetConfigletByNameValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetConfigletByName("configletname")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

///////////////////// GetAllConfigletHistory

func Test_CvpGetAllConfigletHistoryRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetConfigletHistory: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllConfigletHistory("configletname")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpGetAllConfigletHistoryJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllConfigletHistory("configletname"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllConfigletHistoryEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllConfigletHistory("configletname"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpGetAllConfigletHistoryReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetAllConfigletHistory("configletname"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpGetAllConfigletHistoryValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.GetAllConfigletHistory("configletname")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

//////////////////// AddConfiglet

func Test_CvpAddConfigletRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("AddConfiglet: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.AddConfiglet("configletname", "config string")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpAddConfigletJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.AddConfiglet("configletname", "config string"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpAddConfigletEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.AddConfiglet("configletname", "config string"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpAddConfigletReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if _, err := api.AddConfiglet("configletname", "config string"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpAddConfigletValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	_, err := api.AddConfiglet("configletname", "config string")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

///////////////////// DeleteConfiglet

func Test_CvpDeleteConfigletRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("DeleteConfiglet: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	err := api.DeleteConfiglet("configletname", "key")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpDeleteConfigletJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if err := api.DeleteConfiglet("configletname", "key"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpDeleteConfigletEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if err := api.DeleteConfiglet("configletname", "key"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpDeleteConfigletReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if err := api.DeleteConfiglet("configletname", "key"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpDeleteConfigletValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	err := api.DeleteConfiglet("configletname", "key")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}

////////////////////// UpdateConfiglet

func Test_CvpUpdateConfigletRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("UpdateConfiglet: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	err := api.UpdateConfiglet("config", "name", "key")
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpUpdateConfigletJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if err := api.UpdateConfiglet("config", "name", "key"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpUpdateConfigletEmptyJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("", nil)
	api := NewCvpRestAPI(client)
	if err := api.UpdateConfiglet("config", "name", "key"); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpUpdateConfigletReturnError_UnitTest(t *testing.T) {
	respStr := `{"errorCode": "112498",
  				 "errorMessage": "Unauthorized User"}`

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)
	if err := api.UpdateConfiglet("config", "name", "key"); err == nil {
		t.Fatal("Error should be returned")
	}
}

func Test_CvpUpdateConfigletValid_UnitTest(t *testing.T) {
	client := NewMockClient("{}", nil)
	api := NewCvpRestAPI(client)

	err := api.UpdateConfiglet("config", "name", "key")
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
}
