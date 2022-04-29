// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	configsms_vendor "ntf/config/sms_vendor"

	mock "github.com/stretchr/testify/mock"

	sms_vendor "ntf/pkg/utils/sms_vendor"

	testing "testing"
)

// SmsVendorHolder is an autogenerated mock type for the SmsVendorHolder type
type SmsVendorHolder struct {
	mock.Mock
}

// ActiveVendor provides a mock function with given fields:
func (_m *SmsVendorHolder) ActiveVendor() sms_vendor.SmsVendor {
	ret := _m.Called()

	var r0 sms_vendor.SmsVendor
	if rf, ok := ret.Get(0).(func() sms_vendor.SmsVendor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sms_vendor.SmsVendor)
		}
	}

	return r0
}

// Get provides a mock function with given fields: name
func (_m *SmsVendorHolder) Get(name string) sms_vendor.SmsVendor {
	ret := _m.Called(name)

	var r0 sms_vendor.SmsVendor
	if rf, ok := ret.Get(0).(func(string) sms_vendor.SmsVendor); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sms_vendor.SmsVendor)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *SmsVendorHolder) GetAll() []configsms_vendor.SmsVendorConfig {
	ret := _m.Called()

	var r0 []configsms_vendor.SmsVendorConfig
	if rf, ok := ret.Get(0).(func() []configsms_vendor.SmsVendorConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]configsms_vendor.SmsVendorConfig)
		}
	}

	return r0
}

// NewSmsVendorHolder creates a new instance of SmsVendorHolder. It also registers a cleanup function to assert the mocks expectations.
func NewSmsVendorHolder(t testing.TB) *SmsVendorHolder {
	mock := &SmsVendorHolder{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
