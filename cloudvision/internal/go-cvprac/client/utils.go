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

package client

import "github.com/pkg/errors"

const initVal = -1

// HostIterator implements an itorator for a list of hostnames/ips
type HostIterator struct {
	current int
	hosts   []string
}

// Value returns the value for the current entry
func (h *HostIterator) Value() string {
	return h.hosts[h.current]
}

// Next returns true if there is another element
// to iterate. False if we reach the end of the list
func (h *HostIterator) Next() bool {
	h.current++
	if h.current >= len(h.hosts) {
		h.current = initVal
		return false
	}
	return true
}

// Cycle returns the next host in the list. If we've exceeded the
// length of the list, then circle back to the first.
func (h *HostIterator) Cycle() string {
	h.current++
	if h.current >= len(h.hosts) {
		h.current = 0
	}
	return h.hosts[h.current]
}

// NewHostIterator inits an itorator for host list
func NewHostIterator(hosts []string) (*HostIterator, error) {
	if len(hosts) == 0 {
		return nil, errors.New("Can not iterate over empty list")
	}
	return &HostIterator{hosts: hosts, current: initVal}, nil
}
