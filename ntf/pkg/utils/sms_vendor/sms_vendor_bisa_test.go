package sms_vendor

import (
	"context"
	"github.com/stretchr/testify/mock"
	"ntf/config/sms_vendor"
	mocks_http_client "ntf/mocks/pkg/utils/http_client"
	"ntf/pkg/utils/http_client"
	"reflect"
	"testing"
)

func TestNewSmsVendorBisa(t *testing.T) {
	type args struct {
		config sms_vendor.SmsVendorConfig
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success Case",
			args: args{
				config: sms_vendor.SmsVendorConfig{
					Timeout: "30s",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSmsVendorBisa(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSmsVendorBisa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_smsVendorBisaImplementation_Send(t *testing.T) {
	type fields struct {
		name       string
		url        string
		endpoint   string
		token      string
		isDefault  bool
		isActive   bool
		timeout    string
		httpClient http_client.HttpClient
	}
	type args struct {
		ctx     context.Context
		payload []byte
	}

	// set data
	httpClient := &mocks_http_client.HttpClient{}
	httpClient.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult interface{}
		wantErr    error
	}{
		{
			name: "Success Case",
			fields: fields{
				httpClient: httpClient,
			},
			args:       args{},
			wantResult: SmsVendorBisaResponse{},
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &smsVendorBisaImplementation{
				name:       tt.fields.name,
				url:        tt.fields.url,
				endpoint:   tt.fields.endpoint,
				token:      tt.fields.token,
				isDefault:  tt.fields.isDefault,
				isActive:   tt.fields.isActive,
				timeout:    tt.fields.timeout,
				httpClient: tt.fields.httpClient,
			}
			gotResult, err := f.Send(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Send() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_smsVendorBisaImplementation_Toggle(t *testing.T) {
	type fields struct {
		name       string
		url        string
		endpoint   string
		token      string
		isDefault  bool
		isActive   bool
		timeout    string
		httpClient http_client.HttpClient
	}
	type args struct {
		ctx      context.Context
		isActive bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:    "Success Case",
			fields:  fields{},
			args:    args{},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &smsVendorBisaImplementation{
				name:       tt.fields.name,
				url:        tt.fields.url,
				endpoint:   tt.fields.endpoint,
				token:      tt.fields.token,
				isDefault:  tt.fields.isDefault,
				isActive:   tt.fields.isActive,
				timeout:    tt.fields.timeout,
				httpClient: tt.fields.httpClient,
			}
			if err := f.Toggle(tt.args.ctx, tt.args.isActive); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Toggle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_smsVendorBisaImplementation_Data(t *testing.T) {
	type fields struct {
		name       string
		url        string
		endpoint   string
		token      string
		isDefault  bool
		isActive   bool
		timeout    string
		httpClient http_client.HttpClient
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult sms_vendor.SmsVendorConfig
	}{
		{
			name:       "Success Case",
			fields:     fields{},
			wantResult: sms_vendor.SmsVendorConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &smsVendorBisaImplementation{
				name:       tt.fields.name,
				url:        tt.fields.url,
				endpoint:   tt.fields.endpoint,
				token:      tt.fields.token,
				isDefault:  tt.fields.isDefault,
				isActive:   tt.fields.isActive,
				timeout:    tt.fields.timeout,
				httpClient: tt.fields.httpClient,
			}
			if gotResult := f.Data(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Data() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
