// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	catalog "github.com/flyteorg/flyteplugins/go/tasks/pluginmachinery/catalog"
	mock "github.com/stretchr/testify/mock"
)

// UploadFuture is an autogenerated mock type for the UploadFuture type
type UploadFuture struct {
	mock.Mock
}

type UploadFuture_GetResponseError struct {
	*mock.Call
}

func (_m UploadFuture_GetResponseError) Return(_a0 error) *UploadFuture_GetResponseError {
	return &UploadFuture_GetResponseError{Call: _m.Call.Return(_a0)}
}

func (_m *UploadFuture) OnGetResponseError() *UploadFuture_GetResponseError {
	c_call := _m.On("GetResponseError")
	return &UploadFuture_GetResponseError{Call: c_call}
}

func (_m *UploadFuture) OnGetResponseErrorMatch(matchers ...interface{}) *UploadFuture_GetResponseError {
	c_call := _m.On("GetResponseError", matchers...)
	return &UploadFuture_GetResponseError{Call: c_call}
}

// GetResponseError provides a mock function with given fields:
func (_m *UploadFuture) GetResponseError() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type UploadFuture_GetResponseStatus struct {
	*mock.Call
}

func (_m UploadFuture_GetResponseStatus) Return(_a0 catalog.ResponseStatus) *UploadFuture_GetResponseStatus {
	return &UploadFuture_GetResponseStatus{Call: _m.Call.Return(_a0)}
}

func (_m *UploadFuture) OnGetResponseStatus() *UploadFuture_GetResponseStatus {
	c_call := _m.On("GetResponseStatus")
	return &UploadFuture_GetResponseStatus{Call: c_call}
}

func (_m *UploadFuture) OnGetResponseStatusMatch(matchers ...interface{}) *UploadFuture_GetResponseStatus {
	c_call := _m.On("GetResponseStatus", matchers...)
	return &UploadFuture_GetResponseStatus{Call: c_call}
}

// GetResponseStatus provides a mock function with given fields:
func (_m *UploadFuture) GetResponseStatus() catalog.ResponseStatus {
	ret := _m.Called()

	var r0 catalog.ResponseStatus
	if rf, ok := ret.Get(0).(func() catalog.ResponseStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(catalog.ResponseStatus)
	}

	return r0
}

// OnReady provides a mock function with given fields: handler
func (_m *UploadFuture) OnReady(handler catalog.ReadyHandler) {
	_m.Called(handler)
}
