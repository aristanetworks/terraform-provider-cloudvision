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

// RoleList list of roles
type RoleList struct {
	Total int    `json:"total"`
	Roles []Role `json:"Roles"`

	ErrorResponse
}

// SingleRole represents one role
type SingleRole struct {
	RoleData Role `json:"role"`

	ErrorResponse
}

// GetAllRoles returns all the existing roles in CVP
func (c CvpRestAPI) GetAllRoles(start, end int) (*RoleList, error) {
	var roles RoleList

	query := &url.Values{
		"queryparam": {""},
		"startIndex": {strconv.Itoa(start)},
		"endIndex":   {strconv.Itoa(end)},
	}

	resp, err := c.client.Get("/role/getRoles.do", query)
	if err != nil {
		return nil, errors.Errorf("GetAllRoles: %s", err)
	}

	if err = json.Unmarshal(resp, &roles); err != nil {
		return nil, errors.Errorf("GetAllRoles: %s Payload:\n%s", err, resp)
	}

	if err := roles.Error(); err != nil {
		// Entity does not exist
		if roles.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetAllRoles: %s", err)
	}
	return &roles, nil
}

// GetRoleByID returns the role with ID 'roleID'
func (c CvpRestAPI) GetRoleByID(roleID string) (*SingleRole, error) {
	var info SingleRole

	query := &url.Values{"roleId": {roleID}}

	resp, err := c.client.Get("/role/getRole.do", query)
	if err != nil {
		return nil, errors.Errorf("GetRoleByID: %s", err)
	}

	if err = json.Unmarshal(resp, &info); err != nil {
		return nil, errors.Errorf("GetRoleByID: %s Payload:\n%s", err, resp)
	}

	if err := info.Error(); err != nil {
		// Entity does not exist
		if info.ErrorCode == "132801" {
			return nil, nil
		}
		return nil, errors.Errorf("GetRoleByID: %s", err)
	}
	return &info, nil
}

// GetRoleByName returns a role having name 'roleName'
// by querying for all the roles and returning one that matches
func (c CvpRestAPI) GetRoleByName(roleName string) (*SingleRole, error) {
	allRoles, err := c.GetAllRoles(0, 0)
	if err != nil {
		return nil, errors.Errorf("GetRoleByName: [%s]", err)
	}
	if allRoles != nil {
		for _, role := range allRoles.Roles {
			if role.Name == roleName {
				return &SingleRole{RoleData: role}, nil
			}
		}
	}
	return nil, errors.Errorf("GetRoleByName: could not find a role with role name- %s", roleName)
}

// AddRole adds a custom role
func (c CvpRestAPI) AddRole(role *SingleRole) (*SingleRole, error) {
	if role == nil {
		return nil, errors.New("AddRole: can not add nil role")
	}

	resp, err := c.client.Post("/role/createRole.do", nil, role.RoleData)
	if err != nil {
		return nil, errors.Errorf("AddRole: Error: [%v]", err)
	}
	var returnedRole SingleRole
	if err = json.Unmarshal(resp, &returnedRole); err != nil {
		return nil, errors.Errorf("AddRole: unmarshal error- [%v] \nin response- [%v]", err, resp)
	}
	var retErr error
	if err = returnedRole.Error(); err != nil {
		if returnedRole.ErrorCode == ROLE_ALREADY_EXISTS {
			retErr = errors.Errorf("AddRole: Role with key '%s' already exists", role.RoleData.Key)
		} else {
			retErr = errors.Errorf("AddRole: %s", err)
		}
	}
	if retErr == nil {
		return &returnedRole, nil
	}
	return nil, retErr
}

// DeleteRoles deletes the roles with specified keys
func (c CvpRestAPI) DeleteRoles(roleIds []string) error {
	if len(roleIds) == 0 {
		return errors.New("DeleteRoles: empty roleId list")
	}
	resp, err := c.client.Post("/role/deleteRoles.do", nil, roleIds)
	if err != nil {
		return errors.Errorf("DeleteRoles: Error: [%v]", err)
	}
	var msg struct {
		ResponseMessage string `json:"data"`
		ErrorResponse
	}
	if err := json.Unmarshal(resp, &msg); err != nil {
		return errors.Errorf("DeleteRoles: unmarshal error - [%v] \nin response - [%v]", err, resp)
	}
	var retErr error
	if err := msg.Error(); err != nil {
		switch msg.ErrorCode {
		case DEFAULT_ROLE_DELETE:
			retErr = errors.Errorf("DeleteRoles: can not delete default role: [%s]", roleIds)
		case INVALID_ROLE, ENTITY_DOES_NOT_EXIST:
			retErr = errors.Errorf("DeleteRoles: one of the role in [%s] does not exist", roleIds)
		default:
			retErr = errors.Errorf("DeleteRoles: Unexpected error: %v", err)
		}
	} else {
		lowerCaseResp := strings.ToLower(msg.ResponseMessage)
		if !strings.Contains(lowerCaseResp, successMsg) {
			retErr = errors.Errorf("DeleteRoles: Successful deletion response not found in [%s]",
				lowerCaseResp)
		}
	}
	return retErr
}

// UpdateRole updates the given role and also the user permissions
func (c CvpRestAPI) UpdateRole(role *SingleRole) error {
	if role == nil {
		return errors.New("UpdateRole: can not update a nil role")
	}
	resp, err := c.client.Post("/role/updateRole.do", nil, role.RoleData)
	if err != nil {
		return errors.Errorf("UpdateRole: Error: [%v]", err)
	}
	var msg struct {
		ResponseMessage string `json:"data"`
		ErrorResponse
	}
	if err := json.Unmarshal(resp, &msg); err != nil {
		return errors.Errorf("UpdateRole: unmarshal error - [%v] \nin response - [%v]", err, resp)
	}
	if err = msg.Error(); err != nil {
		return errors.Errorf("UpdateRole: Unexpected error - [%v]", err)
	}
	lowerCaseResp := strings.ToLower(msg.ResponseMessage)
	if !strings.Contains(lowerCaseResp, successMsg) {
		return errors.Errorf("UpdateRole: Successful update response not found in [%s]",
			lowerCaseResp)
	}
	return nil
}
