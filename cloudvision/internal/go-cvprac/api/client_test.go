package cvpapi

import (
	"crypto/tls"
	"fmt"
	"net/url"
	"time"

	resty "gopkg.in/resty.v1"
)

// MockClient is made up of a mock response or error
type MockClient struct {
	response []byte
	err      error
}

// NewMockClient creates a MockClient with a response/error
func NewMockClient(response string, err error) *MockClient {
	return &MockClient{
		response: []byte(response),
		err:      err,
	}
}

// Get satisfies the api ClientInterface for Get operation
func (c *MockClient) Get(url string, params *url.Values) ([]byte, error) {
	return c.response, c.err
}

// Post satisfies the api ClientInterface for Post operation
func (c *MockClient) Post(url string, params *url.Values, data interface{}) ([]byte, error) {
	return c.response, c.err
}

// Delete satisfies the api ClientInterface for Delete operation
func (c *MockClient) Delete(url string, params *url.Values, data interface{}) ([]byte, error) {
	return c.response, c.err
}

// RealClient is a simple client implementing the cvpapi ClientInterface
type RealClient struct {
	ClientInterface
	Timeout time.Duration
	Client  *resty.Client
	Headers map[string]string
	url     string
	API     *CvpRestAPI
	URL     string
	Debug   bool
}

// NewRealClient creates a new client
func NewRealClient(host string, proto string, port int) *RealClient {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"User-Agent":   "Golang cvpractest",
	}
	c := &RealClient{
		Headers: headers,
	}
	c.URL = fmt.Sprintf("%s://%s:%d/web", proto, host, port)

	c.Client = resty.New()
	c.Client.SetHeaders(headers)
	c.Client.SetHostURL(c.URL)
	c.Client.SetTimeout(45 * time.Second)

	if proto == "https" {
		c.Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return c
}

// Get satisfies the api ClientInterface for Get operation
func (c *RealClient) Get(url string, params *url.Values) ([]byte, error) {
	var err error
	var resp *resty.Response

	request := c.Client.R()

	if params != nil {
		formatedParams, err := parseURLValues(params)
		if err != nil {
			return nil, err
		}
		request.SetQueryParams(formatedParams)
	}

	resp, err = request.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil

}

// Post satisfies the api ClientInterface for Post operation
func (c *RealClient) Post(url string, params *url.Values, data interface{}) ([]byte, error) {
	var err error
	var resp *resty.Response

	request := c.Client.R()

	if params != nil {
		formatedParams, err := parseURLValues(params)
		if err != nil {
			return nil, err
		}
		request.SetQueryParams(formatedParams)
	}

	resp, err = request.SetBody(data).Post(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func parseURLValues(params *url.Values) (map[string]string, error) {
	newMap := make(map[string]string)
	for k, v := range *params {
		if len(v) > 1 {
			return nil, fmt.Errorf("Parsing URL Values: Multiple values for param %s. Values: %v",
				k, v)
		}
		newMap[k] = params.Get(k)
	}
	return newMap, nil
}
