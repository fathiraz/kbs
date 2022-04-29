package sms_vendor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"ntf/config/sms_vendor"
	"ntf/pkg/utils/http_client"
)

type (
	smsVendorKitaImplementation struct {
		name       string
		url        string
		endpoint   string
		token      string
		isDefault  bool
		isActive   bool
		timeout    string
		httpClient http_client.HttpClient
	}

	SmsVendorKitaRequest struct {
		Number  string
		Message string
	}

	SmsVendorKitaResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
)

// JSON function to set escape html to false
func (s *SmsVendorKitaRequest) JSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(s)
	return buffer.Bytes(), err
}

// NewSmsVendorKita function to init sms vendor kita
func NewSmsVendorKita(config sms_vendor.SmsVendorConfig) (SmsVendor, error) {
	httpClient, err := http_client.NewHttpClient(config.Timeout)
	if err != nil {
		return nil, err
	}

	// set implementation
	smsImplementation := smsVendorKitaImplementation{
		name:       config.Name,
		url:        config.Url,
		endpoint:   config.Endpoint,
		token:      config.Token,
		isDefault:  config.IsDefault,
		isActive:   config.IsActive,
		timeout:    config.Timeout,
		httpClient: httpClient,
	}

	return &smsImplementation, nil
}

// Send function to send sms
func (f *smsVendorKitaImplementation) Send(ctx context.Context, payload []byte) (result interface{}, err error) {

	var (
		smsResponse  SmsVendorKitaResponse
		headers     = map[string]string{
			echo.HeaderAuthorization: f.token,
			echo.HeaderContentType:   echo.MIMEApplicationJSON,
		}
	)

	// set url
	url := fmt.Sprintf("%s/%s", f.url, f.endpoint)

	// send to sms vendor bisa
	err = f.httpClient.Post(ctx, url, headers, payload, &smsResponse)
	if err != nil {
		return
	}
	// set result
	result = smsResponse

	return result, nil
}

// Toggle function to set toggle is_active
func (f *smsVendorKitaImplementation) Toggle(ctx context.Context, isActive bool) (err error) {
	f.isActive = isActive
	return
}

// Data function to get data
func (f *smsVendorKitaImplementation) Data() (result sms_vendor.SmsVendorConfig) {
	result = sms_vendor.SmsVendorConfig{
		Name:      f.name,
		Url:       f.url,
		Endpoint:  f.endpoint,
		Token:     f.token,
		IsDefault: f.isDefault,
		IsActive:  f.isActive,
		Timeout:   f.timeout,
	}
	return
}
