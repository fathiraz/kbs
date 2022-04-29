// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// HttpClient is an autogenerated mock type for the HttpClient type
type HttpClient struct {
	mock.Mock
}

// Post provides a mock function with given fields: ctx, url, headers, body, response
func (_m *HttpClient) Post(ctx context.Context, url string, headers map[string]string, body []byte, response interface{}) error {
	ret := _m.Called(ctx, url, headers, body, response)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string, []byte, interface{}) error); ok {
		r0 = rf(ctx, url, headers, body, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewHttpClient creates a new instance of HttpClient. It also registers a cleanup function to assert the mocks expectations.
func NewHttpClient(t testing.TB) *HttpClient {
	mock := &HttpClient{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}