package http_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"ntf/config/environment"
	"time"
)

type (

	// HttpClient
	HttpClient interface {
		Post(ctx context.Context, url string, headers map[string]string, body []byte, response interface{}) (err error)
	}

	httpClient struct {
		client *resty.Client
	}
)

// NewHttpClient function
func NewHttpClient(timeout string) (*httpClient, error) {

	var (
		httpClient  = &httpClient{}
		httpTimeout time.Duration
		client      *resty.Client
		err         error
	)

	// set custom http client
	httpTimeout, err = time.ParseDuration(timeout)
	if err != nil {
		return nil, err
	}

	// set resty client
	client = resty.New()
	client.SetTimeout(httpTimeout)

	// set debug
	if environment.GetEnv().AppDebug {
		client.SetDebug(true)
	}

	// set client
	httpClient.client = client

	return httpClient, nil
}

// Post function
func (h *httpClient) Post(ctx context.Context, url string, headers map[string]string, body []byte, response interface{}) (err error) {

	var (
		client       = h.client
		httpResponse *resty.Response
	)

	// set header
	if headers != nil {
		for key, value := range headers {
			client.SetHeader(key, value)
		}
	}

	// post request
	httpResponse, err = client.
		R().
		EnableTrace().
		SetContext(ctx).
		SetBody(body).
		Post(url)
	if err != nil {
		return
	}

	// check response code
	if httpResponse.StatusCode() >= http.StatusBadRequest {
		err = fmt.Errorf(string(httpResponse.Body()))
		return
	}

	// parse request body
	err = json.Unmarshal(httpResponse.Body(), &response)
	if err != nil {
		return
	}

	return
}
