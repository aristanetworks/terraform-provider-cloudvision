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
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	gcfg "gopkg.in/gcfg.v1"
)

var configPath = ""

func init() {
	var err error
	configPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

// Config Represents the config file.
type Config struct {
	Nodes map[string]*Node `gcfg:"node"`
}

// Node defines our dut to test against
type Node struct {
	Username string
	Password string
}

// GetNodeIds returns a list of node Ids
func (c Config) GetNodeIds() []string {
	keys := make([]string, len(c.Nodes))
	i := 0
	for k := range c.Nodes {
		keys[i] = k
		i++
	}
	return keys
}

// getUsername
func (n Node) getUsername() string {
	return n.Username
}

// getPassword
func (n Node) getPassword() string {
	return n.Password
}

func (n Node) String() string {
	return "[<Node> Username:" + n.Username + " Passwd:" + n.Password + "]"
}

var config Config

// GetConfig gets the config object
func GetConfig() Config {
	return config
}

// InitConfigFromString creates a config object
func InitConfigFromString(configString string) (Config, error) {
	err := gcfg.ReadStringInto(&config, configString)
	return config, err
}

// GetConfigPath returns the full path to config filename
func GetConfigPath(filename string) string {
	return path.Join(configPath, filename)
}

// LoadConfigFile reads the config file in.
func LoadConfigFile(file string) (Config, error) {
	err := gcfg.ReadFileInto(&config, GetConfigPath(file))
	return config, err
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d: "+msg+"\033[39m\n\n",
			append([]interface{}{filepath.Base(file), line}, v...)...)
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n",
			filepath.Base(file), line, err.Error())
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		tb.Fatalf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n",
			filepath.Base(file), line, exp, act)
	}
}
