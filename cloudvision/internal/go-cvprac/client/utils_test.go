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

import "testing"

func TestCvpRac_HostIterNil_UnitTest(t *testing.T) {
	iter, err := NewHostIterator(nil)
	assert(t, err != nil, "Expected error.")
	assert(t, iter == nil, "Expected Nil.")
}

func TestCvpRac_HostIterEmpty_UnitTest(t *testing.T) {
	iter, err := NewHostIterator([]string{})
	assert(t, err != nil, "Expected error.")
	assert(t, iter == nil, "Expected Nil.")
}

func TestCvpRac_HostIterValid_UnitTest(t *testing.T) {
	iter, err := NewHostIterator([]string{"host1"})
	assert(t, err == nil, "No error expected. Got: %v", err)
	assert(t, iter != nil, "No Iterator created.")
}

func TestCvpRac_HostIterNext_UnitTest(t *testing.T) {
	var status bool
	var err error

	hosts := []string{"host1", "host2", "host3"}

	iter, err := NewHostIterator(hosts)
	assert(t, err == nil, "No error expected. Got: %v", err)
	assert(t, iter != nil, "No Iterator created.")

	status = iter.Next()
	assert(t, status, "Expected true, Got false")
	val := iter.Value()
	assert(t, val == "host1", "Expected host1, Got: %s", val)

	status = iter.Next()
	assert(t, status, "Expected true, Got false")
	val = iter.Value()
	assert(t, val == "host2", "Expected host1, Got: %s", val)

	status = iter.Next()
	assert(t, status, "Expected true, Got false")
	val = iter.Value()
	assert(t, val == "host3", "Expected host1, Got: %s", val)

	status = iter.Next()
	assert(t, status == false, "Expected false, Got true")

}

func TestCvpRac_HostIterCycle_UnitTest(t *testing.T) {
	var err error

	hosts := []string{"host1", "host2", "host3"}

	iter, err := NewHostIterator(hosts)
	assert(t, err == nil, "No error expected. Got: %v", err)
	assert(t, iter != nil, "No Iterator created.")

	val := iter.Cycle()
	assert(t, val == "host1", "Expected host1, Got: %s", val)

	val = iter.Cycle()
	assert(t, val == "host2", "Expected host1, Got: %s", val)

	val = iter.Cycle()
	assert(t, val == "host3", "Expected host1, Got: %s", val)

	val = iter.Cycle()
	assert(t, val == "host1", "Expected host1, Got: %s", val)

	val = iter.Cycle()
	assert(t, val == "host2", "Expected host1, Got: %s", val)

	val = iter.Cycle()
	assert(t, val == "host3", "Expected host1, Got: %s", val)
}
