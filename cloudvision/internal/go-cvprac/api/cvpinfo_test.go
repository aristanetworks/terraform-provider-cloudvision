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
	"errors"
	"reflect"
	"testing"
)

func Test_CvpInfoRetError_UnitTest(t *testing.T) {
	clientErr := errors.New("Client error")
	expectedErr := errors.New("GetCvpInfo: Client error")

	client := NewMockClient("", clientErr)
	api := NewCvpRestAPI(client)

	_, err := api.GetCvpInfo()
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Expected Client error: %v Got: %v", expectedErr, err)
	}
}

func Test_CvpInfoJsonError_UnitTest(t *testing.T) {
	client := NewMockClient("{", nil)
	api := NewCvpRestAPI(client)
	if _, err := api.GetCvpInfo(); err == nil {
		t.Fatal("JSON unmarshal error should be returned")
	}
}

func Test_CvpInfoValid_UnitTest(t *testing.T) {
	expectedResp := &CvpInfo{}

	respStr := `{"appVersion": "Phase_2_Sprint_34_HF09",
  			  "version": "2017.1.0.1"}`

	json.Unmarshal([]byte(respStr), expectedResp)

	client := NewMockClient(respStr, nil)
	api := NewCvpRestAPI(client)

	info, err := api.GetCvpInfo()
	if err != nil {
		t.Fatalf("Valid case failed with error: %v", err)
	}
	if !reflect.DeepEqual(expectedResp, info) {
		t.Fatalf("Expected: %v Got: %v", expectedResp, info)
	}
}
