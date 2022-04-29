package model

import (
	"ntf/pkg/utils/sms_vendor"
)

type (

	// SmsSendRequest struct to hold data sms send
	SmsSendRequest struct {
		Number  string `json:"number" validate:"required"`
		Message string `json:"message" validate:"required"`
	}

	// SmsToggleRequest struct to hold data sms toggle
	SmsToggleRequest struct {
		Name  string `json:"name" validate:"required"`
		Toggle bool `json:"toggle" validate:"required"`
	}
)

// ToSmsVendorKitaRequest function to set data request to sms vendor kita request
func (s SmsSendRequest) ToSmsVendorKitaRequest() (result sms_vendor.SmsVendorKitaRequest) {
	result = sms_vendor.SmsVendorKitaRequest{
		Number:  s.Number,
		Message: s.Message,
	}
	return
}

// ToSmsVendorBisaRequest function to set data request to sms vendor bisa request
func (s SmsSendRequest) ToSmsVendorBisaRequest() (result sms_vendor.SmsVendorBisaRequest) {
	result = sms_vendor.SmsVendorBisaRequest{
		Number:  s.Number,
		Message: s.Message,
	}
	return
}

// ToSmsVendorKitabisaRequest function to set data request to sms vendor kitabisa request
func (s SmsSendRequest) ToSmsVendorKitabisaRequest() (result sms_vendor.SmsVendorKitabisaRequest) {
	result = sms_vendor.SmsVendorKitabisaRequest{
		Number:  s.Number,
		Message: s.Message,
	}
	return
}
