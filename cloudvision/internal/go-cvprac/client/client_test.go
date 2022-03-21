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

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"sync/atomic"
	"testing"
	"time"
)

var debugFlag = flag.Bool("debug", false, "Enable debug")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestCvpRac_ClientOptions_UnitTest(t *testing.T) {
	hosts := []string{"host1", "host2", "host3"}

	_, err := NewCvpClient(Protocol("http"))
	ok(t, err)
	_, err = NewCvpClient(Protocol("https"))
	ok(t, err)
	_, err = NewCvpClient(Protocol("bogus"))
	assert(t, err != nil, "Invalid protocol should return error")

	_, err = NewCvpClient(ConnectTimeout(0))
	ok(t, err)
	_, err = NewCvpClient(ConnectTimeout(15))
	ok(t, err)
	_, err = NewCvpClient(ConnectTimeout(-1))
	assert(t, err != nil, "Invalid timeout should return error")

	_, err = NewCvpClient(Port(9000))
	ok(t, err)
	_, err = NewCvpClient(Port(-1))
	assert(t, err != nil, "Invalid port should return error")

	_, err = NewCvpClient(Hosts(hosts...))
	ok(t, err)
	_, err = NewCvpClient(Hosts(nil...))
	assert(t, err != nil, "Nil host list should return error")
	_, err = NewCvpClient(Hosts([]string{}...))
	assert(t, err != nil, "Empty host list should return error")

	_, err = NewCvpClient(
		Protocol("http"),
		Hosts(hosts...),
		Port(9000),
		ConnectTimeout(19),
		Debug(*debugFlag))
	ok(t, err)

	client, err := NewCvpClient()
	ok(t, err)
	err = client.SetProtocol("http")
	ok(t, err)
	err = client.SetProtocol("https")
	ok(t, err)
	err = client.SetProtocol("bogus")
	assert(t, err != nil, "Invalid protocol should return error")

	err = client.SetConnectTimeout(0)
	ok(t, err)
	err = client.SetConnectTimeout(15)
	ok(t, err)
	err = client.SetConnectTimeout(-1)
	assert(t, err != nil, "Invalid timeout should return error")

	err = client.SetPort(9000)
	ok(t, err)
	err = client.SetPort(-1)
	assert(t, err != nil, "Invalid port should return error")

	err = client.SetHosts(hosts...)
	ok(t, err)
	err = client.SetHosts(nil...)
	assert(t, err != nil, "Nil host list should return error")
	err = client.SetHosts([]string{}...)
	assert(t, err != nil, "Empty host list should return error")

	err = client.SetDebug(true)
	ok(t, err)
}

func TestCvpRac_Client_UnitTest(t *testing.T) {
	ts := createServer(t)
	defer ts.Close()

	host, port, err := parseURL(ts.URL)
	if err != nil {
		t.Fatalf("Parsing test server URL: %s", err)
	}

	hosts := []string{host}

	cvpClient, _ := NewCvpClient(
		Protocol("http"),
		Hosts(hosts...),
		Port(port),
		Debug(*debugFlag))

	err = cvpClient.Connect("cvpadmin", "cvp123")
	ok(t, err)
}

func TestCvpRac_ClientLoginFailRetryPass_UnitTest(t *testing.T) {
	ts1 := createServer(t)
	defer ts1.Close()

	host, port, err := parseURL(ts1.URL)
	if err != nil {
		t.Fatalf("Parsing test server URL: %s", err)
	}

	hosts := []string{host, host, host}

	cvpClient, _ := NewCvpClient(
		Protocol("http"),
		Hosts(hosts...),
		Port(port),
		Debug(*debugFlag))

	err = cvpClient.Connect("deny2", "cvp123")
	ok(t, err)
}

func TestCvpRac_ClientLoginFailRetryFail_UnitTest(t *testing.T) {
	ts1 := createServer(t)
	defer ts1.Close()

	host, port, err := parseURL(ts1.URL)
	if err != nil {
		t.Fatalf("Parsing test server URL: %s", err)
	}

	hosts := []string{host, host, host}

	cvpClient, _ := NewCvpClient(
		Protocol("http"),
		Hosts(hosts...),
		Port(port),
		Debug(*debugFlag))

	err = cvpClient.Connect("denyAll", "cvp123")
	assert(t, err != nil, "Connect returned no error.")
}

func TestCvpRac_ClientStatusBadRequest_UnitTest(t *testing.T) {
	ts := createServer(t)
	defer ts.Close()

	host, port, err := parseURL(ts.URL)
	if err != nil {
		t.Fatalf("Parsing test server URL: %s", err)
	}

	hosts := []string{host}

	cvpClient, _ := NewCvpClient(
		Protocol("http"),
		Hosts(hosts...),
		Port(port),
		Debug(*debugFlag))

	err = cvpClient.Connect("cvpadmin", "cvp123")
	ok(t, err)

	_, err = cvpClient.Post("/StatusBadRequest-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusBadRequest(400)")
	assert(t, err.Error() == "Status [400]", "Got: %s", err)
}

func TestCvpRac_ClientRetrySingleHost_UnitTest(t *testing.T) {
	ts1 := createServer(t)
	defer ts1.Close()

	host, port, err := parseURL(ts1.URL)
	if err != nil {
		t.Fatalf("Parsing test server URL: %s", err)
	}

	hosts := []string{host}

	cvpClient, _ := NewCvpClient(
		Protocol("http"),
		Hosts(hosts...),
		Port(port),
		Debug(*debugFlag))

	err = cvpClient.Connect("cvpadmin", "cvp123")
	ok(t, err)

	_, err = cvpClient.Post("/retrycount-test", nil, nil)
	ok(t, err)

	_, err = cvpClient.Post("/StatusMovedPermanently-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusMovedPermanently(301)")

	_, err = cvpClient.Post("/StatusBadRequest-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusBadRequest(400)")
	assert(t, err.Error() == "Status [400]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusUnauthorized-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusUnauthorized(401)")
	assert(t, err.Error() == "Status [401]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusForbidden-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusForbidden(403)")
	assert(t, err.Error() == "Status [403]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusNotFound-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusNotFound(404)")
	assert(t, err.Error() == "Status [404]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusInternalServerError-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusInternalServerError(500).")
	assert(t, err.Error() == "Status [500]", "Got: %s", err)
}

func TestCvpRac_ClientRetry_MultiHost_UnitTest(t *testing.T) {
	ts1 := createServer(t)
	defer ts1.Close()

	host, port, err := parseURL(ts1.URL)
	if err != nil {
		t.Fatalf("Parsing test server URL: %s", err)
	}

	hosts := []string{host, host, host}

	cvpClient, _ := NewCvpClient(
		Protocol("http"),
		Hosts(hosts...),
		Port(port),
		Debug(*debugFlag))

	err = cvpClient.Connect("cvpadmin", "cvp123")
	ok(t, err)

	_, err = cvpClient.Post("/retrycount-test", nil, nil)
	ok(t, err)

	_, err = cvpClient.Post("/StatusMovedPermanently-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusMovedPermanently(301)")

	_, err = cvpClient.Post("/StatusBadRequest-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusBadRequest(400)")
	assert(t, err.Error() == "Status [400]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusUnauthorized-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusUnauthorized(401)")
	assert(t, err.Error() == "Status [401]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusForbidden-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusForbidden(403)")
	assert(t, err.Error() == "Status [403]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusNotFound-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusNotFound(404)")
	assert(t, err.Error() == "Status [404]", "Got: %s", err)

	_, err = cvpClient.Post("/StatusInternalServerError-test", nil, nil)
	assert(t, err != nil, "POST returned no error when it should have "+
		"returned StatusInternalServerError(500).")
	assert(t, err.Error() == "Status [500]", "Got: %s", err)
}

func createServer(t *testing.T) *httptest.Server {
	var attempt int32

	ts := createTestServer(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Method: %v", r.Method)
		t.Logf("Path: %v", r.URL.Path)
		t.Logf("RawQuery: %v", r.URL.RawQuery)
		t.Logf("Content-Type: %v", r.Header.Get("Content-Type"))

		if r.Method == "POST" {
			if r.URL.Path == "/web/login/authenticate.do" {

				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)

				var creds map[string]string
				data := json.NewDecoder(r.Body)
				if err := data.Decode(&creds); err != nil {
					t.Errorf("createServer: %s", err)
				}

				if creds["userId"] == "deny2" {
					if attp == 2 {
						w.WriteHeader(http.StatusOK)
						fmt.Fprintf(w, `{ "message": "Accepted", "attempt": %d }`, attp)
					} else {
						w.WriteHeader(http.StatusUnauthorized)
						fmt.Fprintf(w, `{ "id": "StatusUnauthorized", "message": "nope. retry", `+
							`"attempt": %d }`, attp)
					}
				} else if creds["userId"] == "denyAll" {
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintf(w, `{ "id": "StatusUnauthorized", "message": "nope. retry", `+
						`"attempt": %d }`, attp)
				} else {
					w.WriteHeader(http.StatusOK)
					fmt.Fprintf(w, `{ "message": "Accepted", "attempt": %d }`, attp)
				}
			} else if r.URL.Path == "/web/retrycount-test" {
				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)
				if attp <= 3 {
					time.Sleep(time.Second * 6)
				}
				fmt.Fprintf(w, `{ "message": "ClientRetry", "attempt": %d }`, attp)
			} else if r.URL.Path == "/web/StatusMovedPermanently-test" {
				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)
				w.WriteHeader(http.StatusMovedPermanently)
				fmt.Fprintf(w, `{ "id": "StatusMovedPermanently", "message": "moved", `+
					`"attempt": %d }`, attp)
			} else if r.URL.Path == "/web/StatusBadRequest-test" {
				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, `{ "id": "bad_request", "message": "bad", `+
					`"attempt": %d }`, attp)
			} else if r.URL.Path == "/web/StatusUnauthorized-test" {
				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, `{ "id": "StatusUnauthorized", "message": "nope", `+
					`"attempt": %d }`, attp)
			} else if r.URL.Path == "/web/StatusForbidden-test" {
				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, `{ "id": "StatusUnauthorized", "message": "nope", `+
					`"attempt": %d }`, attp)
				_, _ = w.Write([]byte(`{ "id": "StatusForbidden", "message": "nope" }`))
			} else if r.URL.Path == "/web/StatusNotFound-test" {
				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, `{ "id": "StatusUnauthorized", "message": "nope", `+
					`"attempt": %d }`, attp)
				_, _ = w.Write([]byte(`{ "id": "StatusNotFound", "message": "not found" }`))

			} else if r.URL.Path == "/web/StatusInternalServerError-test" {
				attp := atomic.AddInt32(&attempt, 1)
				t.Logf("Attempt: %d", attp)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{ "id": "StatusUnauthorized", "message": "nope", `+
					`"attempt": %d }`, attp)
				_, _ = w.Write([]byte(
					`{ "id": "StatusInternalServerError", "message": "server error" }`))

			} else {
				attp := atomic.AddInt32(&attempt, 1)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{ "id": "StatusUnauthorized", "message": "nope", `+
					`"attempt": %d }`, attp)
				_, _ = w.Write([]byte(`{ "message": "Bad" }`))
			}
		}
	})

	return ts
}

func createTestServer(fn func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(fn))
}

func parseURL(urlStr string) (string, int, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return "", -1, fmt.Errorf("Parsing test server url: %s", err)
	}
	port, err := strconv.Atoi(url.Port())
	if err != nil {
		return "", -1, fmt.Errorf("Parsing test server url '%s' for port. %s", url.Host, err)
	}
	return url.Hostname(), port, nil
}
