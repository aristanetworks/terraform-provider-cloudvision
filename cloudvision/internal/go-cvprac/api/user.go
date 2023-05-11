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
	"strings"

	"github.com/pkg/errors"
)

var (
	defaultUser = "cvpadmin"
	successMsg  = "success"
)

// UserList represents a list of users and the roles
type UserList struct {
	Total int                 `json:"total"`
	Users []User              `json:"users"`
	Roles map[string][]string `json:"roles"`

	ErrorResponse
}

// SingleUser represents one user and roles associated with that user
type SingleUser struct {
	UserData User     `json:"user"`
	Roles    []string `json:"roles"`

	ErrorResponse
}

// GetAllUsers returns all the existing users in CVP
func (c CvpRestAPI) GetAllUsers(start, end int) (*UserList, error) {
	var users UserList

	query := &url.Values{
		"queryparam": {""},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/cvpservice/user/getUsers.do", query)
	if err != nil {
		return nil, errors.Errorf("GetAllUsers: %s", err)
	}

	if err = json.Unmarshal(resp, &users); err != nil {
		return nil, errors.Errorf("GetAllUsers: %s Payload:\n%s", err, resp)
	}

	if err := users.Error(); err != nil {
		// Entity does not exist
		if users.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetAllUsers: %s", err)
	}
	return &users, nil
}

// GetUser returns the user with ID 'userID' and the roles associated with user
func (c CvpRestAPI) GetUser(userID string) (*SingleUser, error) {
	var info SingleUser

	query := &url.Values{"userId": {userID}}

	resp, err := c.client.Get("/cvpservice/user/getUser.do", query)
	if err != nil {
		return nil, errors.Errorf("GetUser: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetUser: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		return nil, errors.Errorf("GetUser: %s", info.String())
	}
	return &info, nil
}

// AddUser adds 'user' and returns error if any
func (c CvpRestAPI) AddUser(user *SingleUser) error {
	if user == nil {
		return errors.New("AddUser: can not add nil user")
	}
	resp, err := c.client.Post("/cvpservice/user/addUser.do", nil, user)
	if err != nil {
		return errors.Errorf("AddUser: %s", err)
	}
	var addedUser *SingleUser
	if err = json.Unmarshal(resp, &addedUser); err != nil {
		return errors.Errorf("AddUser: JSON unmarshal error: \n%v", err)
	}
	if err = addedUser.Error(); err != nil {
		var retErr error
		if addedUser.ErrorCode == USER_ALREADY_EXISTS ||
			addedUser.ErrorCode == DATA_ALREADY_EXISTS {
			retErr = errors.Errorf("AddUser: user '%s' already exists", addedUser.UserData.UserID)
		} else {
			retErr = errors.Errorf("AddUser: %s", addedUser.String())
		}
		return retErr
	}
	return nil
}

// DeleteUsers deletes specified users
func (c CvpRestAPI) DeleteUsers(userIds []string) error {
	if len(userIds) == 0 {
		return errors.New("DeleteUsers: no user specified for deletion")
	}
	resp, err := c.client.Post("/cvpservice/user/deleteUsers.do", nil, userIds)
	if err != nil {
		return errors.Errorf("DeleteUsers: %s", err)
	}
	var msg struct {
		ResponseMessage string `json:"data"`
		ErrorResponse
	}
	if err = json.Unmarshal(resp, &msg); err != nil {
		return errors.Errorf("DeleteUsers: JSON unmarshal error: \n%v", err)
	}
	var retErr error
	if err = msg.Error(); err != nil {
		switch msg.ErrorCode {
		case SUPERUSER_DELETE_ATTEMPT:
			retErr = errors.Errorf("DeleteUsers: cannot delete superuser '%s'", defaultUser)
		case INVALID_USER:
			retErr = errors.Errorf("DeleteUsers: one of the users in %v does not exist", userIds)
		default:
			retErr = errors.Errorf("DeleteUsers: Unexpected error: %v", err)
		}
	} else {
		lowerCaseResp := strings.ToLower(msg.ResponseMessage)
		if !strings.Contains(lowerCaseResp, successMsg) {
			retErr = errors.New("DeleteUsers: Successful deletion response not found")
		}
	}
	return retErr
}

// UpdateUser updates 'user' having userObj
func (c CvpRestAPI) UpdateUser(user string, userObj *SingleUser) error {
	param := &url.Values{"userId": {user}}
	resp, err := c.client.Post("/cvpservice/user/updateUser.do", param, userObj)
	if err != nil {
		return errors.Errorf("UpdateUser: %v", err)
	}
	var msg struct {
		ResponseMessage string `json:"data"`
		ErrorResponse
	}
	if err = json.Unmarshal(resp, &msg); err != nil {
		return errors.Errorf("UpdateUsers: JSON unmarshal error: \n%v", err)
	}
	var retErr error
	if err = msg.Error(); err != nil {
		if msg.ErrorCode == SUPERUSER_EDIT_ATTEMPT {
			retErr = errors.Errorf("UpdateUsers: can not edit super user '%s'", defaultUser)
		} else {
			retErr = errors.Errorf("UpdateUsers: %v", err)
		}
	} else {
		lowerCaseResp := strings.ToLower(msg.ResponseMessage)
		if !strings.Contains(lowerCaseResp, successMsg) {
			retErr = errors.New("UpdateUsers: Successful updation response not found")
		}
	}
	return retErr
}
