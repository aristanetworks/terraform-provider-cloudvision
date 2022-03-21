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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	cvprac "github.com/aristanetworks/go-cvprac/v3"
	cvpapi "github.com/aristanetworks/go-cvprac/v3/api"

	"net/http"

	resty "gopkg.in/resty.v1"
)

var headers = map[string]string{
	"Accept":       "application/json",
	"Content-Type": "application/json",
	"User-Agent":   "Golang cvprac/" + cvprac.Version,
}

var (
	// DefaultTimeOut connection timeout default
	DefaultTimeOut = time.Duration(60 * time.Second)
	// DefaultProtocol uses https for default connection
	DefaultProtocol = "https"
	// DefaultHosts set to local host ip
	DefaultHosts = []string{"127.0.0.1"}
	// NumRetryRequests specifies the number for retries to attempt
	NumRetryRequests = 3
)

// UNDEFPORT undefined port
const UNDEFPORT = -1

type authInfo struct {
	Username string
	Password string
}

// CvpClient represents a CVP client api connection
type CvpClient struct {
	cvpapi.ClientInterface
	Hosts     []string
	HostPool  *HostIterator
	Port      int
	Protocol  string
	authInfo  *authInfo
	Timeout   time.Duration
	Transport http.RoundTripper
	Client    *resty.Client
	SessID    string
	url       string
	API       *cvpapi.CvpRestAPI
	Debug     bool
	IsCvaas   bool
	Tenant    string
	Token     string
}

// Option is a Client Option...function that sets a value and returns
type Option func(*CvpClient) error

// Hosts sets the available host ip/names to use for this
// Client
func Hosts(hosts ...string) Option {
	return func(c *CvpClient) error {
		var err error
		if len(hosts) <= 0 {
			return errors.New("Must define at least one host")
		}
		c.Hosts = hosts
		c.HostPool, err = NewHostIterator(c.Hosts)
		if err != nil {
			return err
		}
		return nil
	}
}

// Port sets the port option for this Client
func Port(port int) Option {
	return func(c *CvpClient) error {
		if port <= 0 {
			return errors.Errorf("Invalid port number [%d]", port)
		}
		c.Port = port
		return nil
	}
}

// Protocol sets the protocol option for this Client
func Protocol(proto string) Option {
	return func(c *CvpClient) error {
		switch proto {
		case "http":
		case "https":
		default:
			return errors.Errorf("Invalid protocol [%s]", proto)
		}
		c.Protocol = proto
		return nil
	}
}

// ConnectTimeout sets the connection timeout for this Client
func ConnectTimeout(timeout int) Option {
	return func(c *CvpClient) error {
		if timeout < 0 {
			return errors.New("Timeout (seconds) must be >= 0")
		}
		c.Timeout = time.Duration(timeout) * time.Second
		if c.Client != nil {
			c.Client.SetTimeout(c.Timeout)
		}
		return nil
	}
}

// Transport sets the connection Transport for this Client
func Transport(transport http.RoundTripper) Option {
	return func(c *CvpClient) error {
		if transport != nil {
			c.Transport = transport
			if c.Client != nil {
				c.Client.SetTransport(transport)
			}
		}
		return nil
	}
}

// Debug sets the debug option for this Client
func Debug(enable bool) Option {
	return func(c *CvpClient) error {
		c.Debug = enable
		if c.Client != nil {
			c.Client.SetDebug(c.Debug)
		}
		return nil
	}
}

// Cvaas enables the cvaas support for this Client
func Cvaas(enable bool, tenant string) Option {
	return func(c *CvpClient) error {
		c.IsCvaas = enable
		c.Tenant = tenant
		return nil
	}
}

// SetOption takes one or more option function and applies them in order
func (c *CvpClient) SetOption(options ...Option) error {
	for _, opt := range options {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

// SetHosts sets the hosts
func (c *CvpClient) SetHosts(host ...string) error {
	return c.SetOption(Hosts(host...))
}

// SetPort sets the port for this connection
func (c *CvpClient) SetPort(port int) error {
	return c.SetOption(Port(port))
}

// SetProtocol sets the protocol (i.e. http or https) associated with this
// connection
func (c *CvpClient) SetProtocol(proto string) error {
	return c.SetOption(Protocol(proto))
}

// SetConnectTimeout sets the connection timeout associated with this
// connection
func (c *CvpClient) SetConnectTimeout(timeout int) error {
	return c.SetOption(ConnectTimeout(timeout))
}

// SetTransport sets the connection Transport associated with this
// connection
func (c *CvpClient) SetTransport(transport http.RoundTripper) error {
	return c.SetOption(Transport(transport))
}

// SetDebug enables or disables debugging.
func (c *CvpClient) SetDebug(enable bool) error {
	return c.SetOption(Debug(enable))
}

// SetCvaas adds support for cvaas.
func (c *CvpClient) SetCvaas(enable bool, tenant string) error {
	return c.SetOption(Cvaas(enable, tenant))
}

// NewCvpClient creates a new CVP RESTful Client
func NewCvpClient(options ...Option) (*CvpClient, error) {
	c := &CvpClient{
		Port:     UNDEFPORT,
		Protocol: DefaultProtocol,
		Timeout:  DefaultTimeOut,
		Hosts:    DefaultHosts,
	}

	// Parse Options
	if err := c.SetOption(options...); err != nil {
		return nil, err
	}

	c.initSession(c.Hosts[0])

	c.API = cvpapi.NewCvpRestAPI(c)

	return c, nil
}

// GetPort returns the port this client will use for connectivity
func (c *CvpClient) GetPort() int {
	if c.Port != UNDEFPORT {
		return c.Port
	}

	if c.Protocol == "https" {
		return 443
	}
	return 80
}

// GetSessionID returns the current Session ID
func (c *CvpClient) GetSessionID() string {
	return c.SessID
}

// Connect Login to CVP and get a session ID and cookie.
func (c *CvpClient) Connect(username string, password string) error {
	c.authInfo = &authInfo{username, password}

	return c.createSession(true)
}

// Connect to CVP with a token. Takes the cvpToken parameter as an input for the string.
func (c *CvpClient) ConnectWithToken(cvpToken string) error {
	c.Token = cvpToken
	return c.createSession(true)
}

func (c *CvpClient) createSession(allNodes bool) error {
	var errorMsg []string

	numNodes := len(c.Hosts)
	if !allNodes && numNodes > 1 {
		numNodes--
	}
	for nodeIter := 0; nodeIter < numNodes; nodeIter++ {
		host := c.HostPool.Cycle()

		c.initSession(host)

		if err := c.login(); err != nil {
			tmpMsg := fmt.Sprintf("createSession: Error: %s", err.Error())
			errorMsg = append(errorMsg, tmpMsg)
			continue
		}
		return nil
	}
	return errors.New(strings.Join(errorMsg, "\n"))
}

func (c *CvpClient) initSession(host string) error {
	if host == "" {
		return errors.Errorf("initSession: No host provided")
	}

	c.url = fmt.Sprintf("%s://%s:%d", c.Protocol, host, c.GetPort())
	//if !c.IsCvaas {
	//	c.url = c.url + "/web"
	//} else {
	//	c.url = c.url + "/cvpservice"
	//}

	c.Client = resty.New()

	// Make sure to set transport before SetTLSClientConfig()
	// If Transport is nil, SetTransport() creates a default.
	c.Client.SetTransport(c.Transport)

	if c.Protocol == "https" {
		c.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	// If there is a token with the TokenField then set token within the rest library.
	if c.Token != "" {
		c.Client.SetAuthToken(c.Token)
	}

	c.Client.SetHostURL(c.url)
	c.Client.SetHeaders(headers)
	c.Client.SetTimeout(c.Timeout)
	c.Client.SetDebug(c.Debug)
	return nil
}

func (c *CvpClient) resetSession() error {
	// reset session to the current host we are connected to
	if err := c.initSession(c.HostPool.Value()); err != nil {
		return errors.Wrap(err, "resetSession")
	}

	if err := c.login(); err != nil {
		return errors.Wrap(err, "resetSession")
	}
	return nil
}

func (c *CvpClient) login() error {
	if c.IsCvaas {
		return c.loginCvaas()
	}
	if c.Token != "" { // If a token exists do not use one of the logincvaas or loginonprem and do not create a cookie the auth header is used with the token.
		return nil
	}
	if c.authInfo.Username != "" {
		return c.loginOnPrem()
	} else {
		return nil
	}
}

func (c *CvpClient) loginCvaas() error {
	request := c.Client.R()
	auth := `{"org":"` + c.Tenant + `", "name":"` + c.authInfo.Username +
		`", "password":"` + c.authInfo.Password + `"}`
	resp, err := request.SetBody(auth).Post("/api/v1/oauth?provider=local&next=false")
	if err != nil {
		return errors.Wrap(err, "login")
	}

	if err = checkResponseStatus(resp); err != nil {
		return errors.Wrap(err, "login")
	}

	c.Client.SetCookies(resp.Cookies())
	return nil
}

func (c *CvpClient) loginOnPrem() error {
	var loginResp cvpapi.LoginResp

	c.SessID = ""
	request := c.Client.R()

	auth := "{\"userId\":\"" + c.authInfo.Username +
		"\", \"password\":\"" + c.authInfo.Password + "\"}"

	resp, err := request.SetBody(auth).Post("/login/authenticate.do")
	if err != nil {
		return errors.Wrap(err, "login")
	}

	if err = checkResponse(resp); err != nil {
		return errors.Wrap(err, "login")
	}

	if err = json.Unmarshal(resp.Body(), &loginResp); err != nil {
		return errors.Wrap(err, "login")
	}
	c.SessID = loginResp.SessionID

	c.Client.SetCookies(resp.Cookies())

	return nil
}

func (c *CvpClient) makeRequest(reqType string, url string, params *url.Values,
	data interface{}) ([]byte, error) {
	var err error
	var resp *resty.Response
	var formattedParams map[string]string

	if c.Client == nil {
		return nil, errors.Errorf("makeRequest: No valid session to CVP [%s]", c.url)
	}

	retryCnt := NumRetryRequests

	if params != nil {
		formattedParams, err = parseURLValues(params)
		if err != nil {
			return nil, err
		}
	}

	nodeCnt := len(c.Hosts)
	for nodeCnt > 0 {

		request := c.Client.R()
		request.SetQueryParams(formattedParams)

		// If we've seen an error
		if err != nil {
			// Decrement count as another node will be tried, if there
			// are no more nodes then return error.
			nodeCnt--
			if nodeCnt == 0 {
				return nil, err
			}
			// Not the first time through the loop. Retrying request so
			// create a session to another CVP node...but exclude this one.
			if err := c.createSession(false); err != nil {
				return nil, err
			}
			retryCnt = NumRetryRequests
		}

		// Clear our errors
		err = nil

		// Check reqType
		switch reqType {
		case "GET":
			resp, err = request.Get(url)
		case "POST":
			resp, err = request.SetBody(data).Post(url)
		case "DELETE":
			resp, err = request.SetBody(data).Delete(url)
		default:
			return nil, errors.Errorf("Invalid. Request type [%s] not implemented", reqType)
		}

		if err != nil {
			return nil, err
		}
		// Underlying request issue. Could be getsockopt error (like network not reachable)
		if resp.RawResponse == nil {
			// retry another session
			err = errors.New("RawResponse error")
			continue
		}

		status := resp.StatusCode()

		if status == 301 {
			// retry another session
			err = errors.Errorf("Status [%d]", status)
			continue
		}
		// From 2018.2.0 onwards, a '401' response is returned for
		// an 'Unauthorized' user, unlike previous releases where a
		// CvpError with code '112498' was returned. This might happen
		// if the session expires, which is after 12 hours of inactivity.
		// In this case, the session must be refreshed. Retry same host.
		if status == 401 {
			retryCnt--
			if retryCnt > 0 {
				// reset our session
				if err := c.resetSession(); err != nil {
					// try another session
					err = errors.Wrap(err, "makeRequest")
				}
			} else {
				err = errors.Errorf("Status [%d]", status)
			}
			continue
		}

		// client error
		if status != http.StatusOK {
			// retry another session
			err = errors.Errorf("Status [%d]", status)
			continue
		}
		break
	}
	return resp.Body(), nil
}

// Get implemented as part of cvprac api client interface
func (c *CvpClient) Get(url string, params *url.Values) ([]byte, error) {
	return c.makeRequest("GET", url, params, nil)
}

// Post implemented as part of cvprac api client interface
func (c *CvpClient) Post(url string, params *url.Values, data interface{}) ([]byte, error) {
	return c.makeRequest("POST", url, params, data)
}

// Delete implemented as part of cvprac api client interface
func (c *CvpClient) Delete(url string, params *url.Values, data interface{}) ([]byte, error) {
	return c.makeRequest("DELETE", url, params, data)
}

func parseURLValues(params *url.Values) (map[string]string, error) {
	newMap := make(map[string]string)
	for k, v := range *params {
		if len(v) > 1 {
			return nil, errors.Errorf("Parsing URL Values: Multiple values for "+
				"param %s. Values: %v", k, v)
		}
		newMap[k] = params.Get(k)
	}
	return newMap, nil
}

func checkResponseStatus(resp *resty.Response) error {
	// Underlying request issue. Could be getsockopt error (like network not reachable)
	if resp.RawResponse == nil {
		// retry another session
		return errors.New("checkResponseStatus: RawResponse error")
	}

	status := resp.StatusCode()
	if status != http.StatusOK {
		return errors.Errorf("checkResponseStatus: %s, StatusCode: %d", resp.Status(), status)
	}

	return nil
}

func checkResponse(resp *resty.Response) error {
	if err := checkResponseStatus(resp); err != nil {
		return errors.Wrap(err, "checkResponse")
	}

	var info cvpapi.ErrorResponse
	if err := json.Unmarshal(resp.Body(), &info); err != nil {
		return errors.Wrap(err, "checkResponse")
	}

	if err := info.Error(); err != nil {
		return errors.Wrap(err, "checkResponse")
	}
	return nil
}
