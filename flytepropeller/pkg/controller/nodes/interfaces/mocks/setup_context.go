// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	promutils "github.com/flyteorg/flytestdlib/promutils"
	mock "github.com/stretchr/testify/mock"
)

// SetupContext is an autogenerated mock type for the SetupContext type
type SetupContext struct {
	mock.Mock
}

type SetupContext_EnqueueOwner struct {
	*mock.Call
}

func (_m SetupContext_EnqueueOwner) Return(_a0 func(string)) *SetupContext_EnqueueOwner {
	return &SetupContext_EnqueueOwner{Call: _m.Call.Return(_a0)}
}

func (_m *SetupContext) OnEnqueueOwner() *SetupContext_EnqueueOwner {
	c_call := _m.On("EnqueueOwner")
	return &SetupContext_EnqueueOwner{Call: c_call}
}

func (_m *SetupContext) OnEnqueueOwnerMatch(matchers ...interface{}) *SetupContext_EnqueueOwner {
	c_call := _m.On("EnqueueOwner", matchers...)
	return &SetupContext_EnqueueOwner{Call: c_call}
}

// EnqueueOwner provides a mock function with given fields:
func (_m *SetupContext) EnqueueOwner() func(string) {
	ret := _m.Called()

	var r0 func(string)
	if rf, ok := ret.Get(0).(func() func(string)); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(string))
		}
	}

	return r0
}

type SetupContext_MetricsScope struct {
	*mock.Call
}

func (_m SetupContext_MetricsScope) Return(_a0 promutils.Scope) *SetupContext_MetricsScope {
	return &SetupContext_MetricsScope{Call: _m.Call.Return(_a0)}
}

func (_m *SetupContext) OnMetricsScope() *SetupContext_MetricsScope {
	c_call := _m.On("MetricsScope")
	return &SetupContext_MetricsScope{Call: c_call}
}

func (_m *SetupContext) OnMetricsScopeMatch(matchers ...interface{}) *SetupContext_MetricsScope {
	c_call := _m.On("MetricsScope", matchers...)
	return &SetupContext_MetricsScope{Call: c_call}
}

// MetricsScope provides a mock function with given fields:
func (_m *SetupContext) MetricsScope() promutils.Scope {
	ret := _m.Called()

	var r0 promutils.Scope
	if rf, ok := ret.Get(0).(func() promutils.Scope); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(promutils.Scope)
		}
	}

	return r0
}

type SetupContext_OwnerKind struct {
	*mock.Call
}

func (_m SetupContext_OwnerKind) Return(_a0 string) *SetupContext_OwnerKind {
	return &SetupContext_OwnerKind{Call: _m.Call.Return(_a0)}
}

func (_m *SetupContext) OnOwnerKind() *SetupContext_OwnerKind {
	c_call := _m.On("OwnerKind")
	return &SetupContext_OwnerKind{Call: c_call}
}

func (_m *SetupContext) OnOwnerKindMatch(matchers ...interface{}) *SetupContext_OwnerKind {
	c_call := _m.On("OwnerKind", matchers...)
	return &SetupContext_OwnerKind{Call: c_call}
}

// OwnerKind provides a mock function with given fields:
func (_m *SetupContext) OwnerKind() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
