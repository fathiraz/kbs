// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "ntf/internal/modules/sms/model"

	mock "github.com/stretchr/testify/mock"

	sms_vendor "ntf/config/sms_vendor"

	testing "testing"
)

// SmsUsecase is an autogenerated mock type for the SmsUsecase type
type SmsUsecase struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx
func (_m *SmsUsecase) GetAll(ctx context.Context) ([]sms_vendor.SmsVendorConfig, error) {
	ret := _m.Called(ctx)

	var r0 []sms_vendor.SmsVendorConfig
	if rf, ok := ret.Get(0).(func(context.Context) []sms_vendor.SmsVendorConfig); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sms_vendor.SmsVendorConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: ctx, data
func (_m *SmsUsecase) Send(ctx context.Context, data *model.SmsSendRequest) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.SmsSendRequest) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Toggle provides a mock function with given fields: ctx, data
func (_m *SmsUsecase) Toggle(ctx context.Context, data *model.SmsToggleRequest) (sms_vendor.SmsVendorConfig, error) {
	ret := _m.Called(ctx, data)

	var r0 sms_vendor.SmsVendorConfig
	if rf, ok := ret.Get(0).(func(context.Context, *model.SmsToggleRequest) sms_vendor.SmsVendorConfig); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(sms_vendor.SmsVendorConfig)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.SmsToggleRequest) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSmsUsecase creates a new instance of SmsUsecase. It also registers a cleanup function to assert the mocks expectations.
func NewSmsUsecase(t testing.TB) *SmsUsecase {
	mock := &SmsUsecase{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
