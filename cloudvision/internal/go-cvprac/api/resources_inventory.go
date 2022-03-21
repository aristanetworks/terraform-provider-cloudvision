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
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type OnboardDeviceRequest struct {
	Key struct {
		RequestID string `json:"requestId"`
	} `json:"key"`
	HostnameOrIP string `json:"hostnameOrIp"`
	DeviceType   string `json:"device_type"`
}

type OnboardDeviceResponse struct {
	Value struct {
		Key struct {
			RequestID string `json:"requestId"`
		} `json:"key"`
		DeviceID      string `json:"deviceId"`
		Status        string `json:"status"`
		StatusMessage string `json:"statusMessage"`
	} `json:"value"`
	Time time.Time `json:"time"`
}

type DecomDeviceRequest struct {
	Key struct {
		RequestID string `json:"request_id"`
	} `json:"key"`
	DeviceID string `json:"device_id"`
}

type DecomDeviceResponse struct {
	Value struct {
		Key struct {
			RequestID string `json:"requestId"`
		} `json:"key"`
		Status        string `json:"status"`
		StatusMessage string `json:"statusMessage"`
	} `json:"value"`
	Time time.Time `json:"time"`
}

type OnboardStatusResponse struct {
	Value struct {
		Key struct {
			RequestID string `json:"requestId"`
		} `json:"key"`
		DeviceID      string `json:"deviceId"`
		Status        string `json:"status"`
		StatusMessage string `json:"statusMessage"`
	} `json:"value"`
	Time time.Time `json:"time"`
}

type DecomStatusResponse struct {
	Value struct {
		Key struct {
			RequestID string `json:"requestId"`
		} `json:"key"`
		Status        string `json:"status"`
		StatusMessage string `json:"statusMessage"`
	} `json:"value"`
	Time time.Time `json:"time"`
}
type DecinStatusallResponse struct {
	Result struct {
		Value struct {
			Key struct {
				RequestID string `json:"requestId"`
			} `json:"key"`
			Status        string `json:"status"`
			StatusMessage string `json:"statusMessage"`
		} `json:"value"`
		Time time.Time `json:"time"`
		Type string    `json:"type"`
	} `json:"result"`
}

func (c CvpRestAPI) OnboardDevice(deviceIPAddress, devtype string) (*OnboardDeviceResponse, error) {
	//https://aristanetworks.github.io/cloudvision-apis/examples/rest/inventory/
	//device ip address is the hostname or device ip addrress.
	//devtype is eos at the time being but could be another type like wifi.

	//The uuid is going to be completely random but is required for the request.
	id := uuid.NewString()
	//initialize info and device.
	info := &OnboardDeviceResponse{}
	data := &OnboardDeviceRequest{}
	data.Key.RequestID = id
	data.DeviceType = devtype
	data.HostnameOrIP = deviceIPAddress

	resp, err := c.client.Post("/api/resources/inventory/v1/DeviceOnboardingConfig", nil, data)
	if err != nil {
		return nil, errors.Errorf("Issue adding device: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("OnboardDevice: %s Payload:\n%s", err, resp)
	}
	return info, nil
}

func (c CvpRestAPI) DecomDevice(deviceid string) (*DecomDeviceResponse, error) {
	//https://aristanetworks.github.io/cloudvision-apis/examples/rest/inventory/
	//deviceid is going to be the devices serial number as presented in CVP
	//The uuid is going to be completely random but is required for the request.
	id := uuid.NewString()
	data := &DecomDeviceRequest{}
	info := &DecomDeviceResponse{}

	data.Key.RequestID = id
	data.DeviceID = deviceid

	resp, err := c.client.Post("/api/resources/inventory/v1/DeviceDecommissioningConfig", nil, data)
	if err != nil {
		return nil, errors.Errorf("Issue removing device: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("DecomDevice: %s Payload:\n%s", err, resp)
	}
	return info, nil
}

func (c CvpRestAPI) OnboardStatus(requestid string) (string, error) {
	info := OnboardStatusResponse{}

	resp, err := c.client.Get("/api/resources/inventory/v1/DeviceOnboarding?key.requestId="+requestid, nil)
	if err != nil {
		return info.Value.Status, errors.Errorf("Error querying for device status: %s", err)
	}
	if err = json.Unmarshal(resp, &info); err != nil {
		return info.Value.Status, errors.Errorf("DecomDevice: %s Payload:\n%s", err, resp)
	}
	return info.Value.Status, err

}

func (c CvpRestAPI) DecomStatus(requestid string) (string, error) {
	info := DecomStatusResponse{}

	resp, err := c.client.Get("/api/resources/inventory/v1/DeviceDecommissioning?key.requestId="+requestid, nil)
	if err != nil {
		return info.Value.Status, errors.Errorf("Error querying for device status: %s", err)
	}
	if err = json.Unmarshal(resp, &info); err != nil {
		return info.Value.Status, errors.Errorf("DecomDevice: %s Payload:\n%s", err, resp)
	}
	return info.Value.Status, err

}

func (c CvpRestAPI) DecomStatusAll() ([]DecinStatusallResponse, error) {
	info := []DecinStatusallResponse{}

	resp, err := c.client.Get("/api/resources/inventory/v1/DeviceDecommissioning/all", nil)
	if err != nil {
		return info, errors.Errorf("Error querying for device status: %s", err)
	}
	if err = json.Unmarshal(resp, &info); err != nil {
		return info, errors.Errorf("DecomDevice: %s Payload:\n%s", err, resp)
	}
	return info, err
}
