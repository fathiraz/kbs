package sms_vendor

import (
	"context"
	"ntf/config/sms_vendor"
)

type (

	// SmsVendorHolder
	SmsVendorHolder interface {
		Get(name string) (result SmsVendor)
		GetAll() (result []sms_vendor.SmsVendorConfig)
		ActiveVendor() (result SmsVendor)
	}

	// SmsVendor interfaces
	SmsVendor interface {
		Send(ctx context.Context, payload []byte) (result interface{}, err error)
		Toggle(ctx context.Context, isActive bool) (err error)
		Data() (result sms_vendor.SmsVendorConfig)
	}
)
