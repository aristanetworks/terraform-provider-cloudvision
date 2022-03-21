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

	"github.com/pkg/errors"
)

// CvpInfo represents Cvp specific information
type CvpInfo struct {
	AppVersion string `json:"appVersion"`
	Version    string `json:"version"`
}

func (i CvpInfo) String() string {
	return "version:" + i.Version + ", appVersion:" + i.AppVersion
}

// GetCvpInfo returns the CvpInfo from the Client connection.
func (c CvpRestAPI) GetCvpInfo() (*CvpInfo, error) {
	var info CvpInfo

	resp, err := c.client.Get("/cvpInfo/getCvpInfo.do", nil)
	if err != nil {
		return nil, errors.Errorf("GetCvpInfo: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetCvpInfo: %s Payload:\n%s", err, resp)
	}

	return &info, nil
}
