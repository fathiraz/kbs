package model

import (
	"ntf/pkg/utils/sms_vendor"
	"reflect"
	"testing"
)

func TestSmsSendRequest_ToSmsVendorKitaRequest(t *testing.T) {
	type fields struct {
		Number  string
		Message string
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult sms_vendor.SmsVendorKitaRequest
	}{
		{
			name: "Success Case",
			fields: fields{
				Number:  "1",
				Message: "a",
			},
			wantResult: sms_vendor.SmsVendorKitaRequest{
				Number:  "1",
				Message: "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SmsSendRequest{
				Number:  tt.fields.Number,
				Message: tt.fields.Message,
			}
			if gotResult := s.ToSmsVendorKitaRequest(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ToSmsVendorKitaRequest() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSmsSendRequest_ToSmsVendorBisaRequest(t *testing.T) {
	type fields struct {
		Number  string
		Message string
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult sms_vendor.SmsVendorBisaRequest
	}{
		{
			name: "Success Case",
			fields: fields{
				Number:  "1",
				Message: "a",
			},
			wantResult: sms_vendor.SmsVendorBisaRequest{
				Number:  "1",
				Message: "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SmsSendRequest{
				Number:  tt.fields.Number,
				Message: tt.fields.Message,
			}
			if gotResult := s.ToSmsVendorBisaRequest(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ToSmsVendorBisaRequest() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSmsSendRequest_ToSmsVendorKitabisaRequest(t *testing.T) {
	type fields struct {
		Number  string
		Message string
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult sms_vendor.SmsVendorKitabisaRequest
	}{
		{
			name: "Success Case",
			fields: fields{
				Number:  "1",
				Message: "a",
			},
			wantResult: sms_vendor.SmsVendorKitabisaRequest{
				Number:  "1",
				Message: "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SmsSendRequest{
				Number:  tt.fields.Number,
				Message: tt.fields.Message,
			}
			if gotResult := s.ToSmsVendorKitabisaRequest(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ToSmsVendorKitabisaRequest() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
