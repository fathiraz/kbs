package sms_vendor

import (
	"ntf/config/sms_vendor"
	"ntf/pkg/shared/constant"
)

type (
	smsVendorHolderImplementation struct {
		smsVendors map[string]SmsVendor
	}
)

// NewSmsVendorHolder function to set sms vendor holder
func NewSmsVendorHolder(smsVendorConfig map[string]sms_vendor.SmsVendorConfig) (result SmsVendorHolder, err error) {

	// set sms vendor implementation
	var holder = smsVendorHolderImplementation{
		smsVendors: make(map[string]SmsVendor),
	}

	// set implementation sms vendor
	for _, config := range smsVendorConfig {
		var implementation SmsVendor

		// set implementation
		implementation, err = implementationSmsVendorByName(config)

		// append to holder
		holder.smsVendors[config.Name] = implementation
	}

	result = &holder

	return
}

// Get function to get sms vendor by name
func (s smsVendorHolderImplementation) Get(name string) (result SmsVendor) {
	return s.smsVendors[name]
}

// GetAll function to get all sms vendor
func (s smsVendorHolderImplementation) GetAll() (result []sms_vendor.SmsVendorConfig) {
	for _, smsVendor := range s.smsVendors {
		result = append(result, smsVendor.Data())
	}
	return
}

// ActiveVendor function to get active sms vendor
func (s smsVendorHolderImplementation) ActiveVendor() (result SmsVendor) {
	var (
		defaultVendor SmsVendor
	)

	// loop to check each sms vendor
	for _, smsVendor := range s.smsVendors {

		// set default vendor
		if smsVendor.Data().IsDefault {
			defaultVendor = smsVendor
		}

		// set sms vendor is active
		if smsVendor.Data().IsActive {
			result = smsVendor
		}
	}

	// set result from default vendor if no active vendor found
	if result == nil {
		result = defaultVendor
	}

	return
}

func implementationSmsVendorByName(config sms_vendor.SmsVendorConfig) (implementation SmsVendor, err error) {

	// set implementation based on sms vendor name
	switch config.Name {
	case constant.SmsVendorNameKita:
		implementation, err = NewSmsVendorKita(config)
	case constant.SmsVendorNameBisa:
		implementation, err = NewSmsVendorBisa(config)
	case constant.SmsVendorNameKitabisa:
		implementation, err = NewSmsVendorKitabisa(config)
	}

	if err != nil {
		return
	}
	return
}
