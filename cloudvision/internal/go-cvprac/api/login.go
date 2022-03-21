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

// LoginResp contains the response from a login request
type LoginResp struct {
	Roles              []Role        `json:"roles"`
	AuthorizationType  string        `json:"authorizationType"`
	PermissionList     []Permissions `json:"permissionList"`
	SessionID          string        `json:"sessionId"`
	AuthenticationType string        `json:"authenticationType"`
	User               User          `json:"user"`
	Username           string        `json:"username"`

	ErrorResponse
}

// Role represents CVP role
type Role struct {
	ClassID        int      `json:"classId"`
	ModuleList     []Module `json:"moduleList"`
	CreatedBy      string   `json:"createdBy"`
	FactoryID      int      `json:"factoryId"`
	Name           string   `json:"name"`
	ModuleListSize int      `json:"moduleListSize"`
	Description    string   `json:"description"`
	ID             int      `json:"id"`
	CreatedOn      int64    `json:"createdOn"`
	Key            string   `json:"key"`
}

// Permissions represents CVP permission
type Permissions struct {
	Mode      string `json:"mode"`
	FactoryID int    `json:"factoryId"`
	Name      string `json:"name"`
	ID        int    `json:"id"`
}

// Module ..
type Module struct {
	Mode      string `json:"mode"`
	FactoryID int    `json:"factoryId"`
	Name      string `json:"name"`
	ID        int    `json:"id"`
}

// User represents CVP User
type User struct {
	FirstName     string      `json:"firstName"`
	LastName      string      `json:"lastName"`
	Password      interface{} `json:"password"`
	UserStatus    string      `json:"userStatus"`
	CurrentStatus string      `json:"currentStatus"`
	FactoryID     int         `json:"factoryId"`
	ContactNumber interface{} `json:"contactNumber"`
	LastAccessed  int64       `json:"lastAccessed"`
	UserType      string      `json:"userType"`
	ID            int         `json:"id"`
	UserID        string      `json:"userId"`
	Email         string      `json:"email"`
}

// Login perform loging and save off cookies
func (c *CvpRestAPI) Login(username, password string) (*LoginResp, error) {
	var resp LoginResp

	auth := "{\"userId\":\"" + username + "\", \"password\":\"" + password + "\"}"

	rawResp, err := c.client.Post("/login/authenticate.do", nil, auth)
	if err != nil {
		return nil, errors.Errorf("Login: %s", err)
	}

	if err = json.Unmarshal(rawResp, &resp); err != nil {
		return nil, errors.Errorf("Login: %s Payload:\n%s", err, rawResp)
	}

	if err := resp.Error(); err != nil {
		return nil, errors.Errorf("Login: %s", err)
	}

	return &resp, nil
}

// Logout perform logout
func (c *CvpRestAPI) Logout() error {
	var resp ErrorResponse

	rawResp, err := c.client.Post("/login/logout.do", nil, nil)
	if err != nil {
		return errors.Errorf("Logout: %s", err)
	}

	if err = json.Unmarshal(rawResp, &resp); err != nil {
		return errors.Errorf("Logout: %s Payload:\n%s", err, rawResp)
	}

	if err := resp.Error(); err != nil {
		return errors.Errorf("Logout: %s", err)
	}
	return nil
}
