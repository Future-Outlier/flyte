// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MutableBranchNodeStatus is an autogenerated mock type for the MutableBranchNodeStatus type
type MutableBranchNodeStatus struct {
	mock.Mock
}

type MutableBranchNodeStatus_Expecter struct {
	mock *mock.Mock
}

func (_m *MutableBranchNodeStatus) EXPECT() *MutableBranchNodeStatus_Expecter {
	return &MutableBranchNodeStatus_Expecter{mock: &_m.Mock}
}

// GetFinalizedNode provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) GetFinalizedNode() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFinalizedNode")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MutableBranchNodeStatus_GetFinalizedNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFinalizedNode'
type MutableBranchNodeStatus_GetFinalizedNode_Call struct {
	*mock.Call
}

// GetFinalizedNode is a helper method to define mock.On call
func (_e *MutableBranchNodeStatus_Expecter) GetFinalizedNode() *MutableBranchNodeStatus_GetFinalizedNode_Call {
	return &MutableBranchNodeStatus_GetFinalizedNode_Call{Call: _e.mock.On("GetFinalizedNode")}
}

func (_c *MutableBranchNodeStatus_GetFinalizedNode_Call) Run(run func()) *MutableBranchNodeStatus_GetFinalizedNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableBranchNodeStatus_GetFinalizedNode_Call) Return(_a0 *string) *MutableBranchNodeStatus_GetFinalizedNode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableBranchNodeStatus_GetFinalizedNode_Call) RunAndReturn(run func() *string) *MutableBranchNodeStatus_GetFinalizedNode_Call {
	_c.Call.Return(run)
	return _c
}

// GetPhase provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) GetPhase() v1alpha1.BranchNodePhase {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPhase")
	}

	var r0 v1alpha1.BranchNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.BranchNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.BranchNodePhase)
	}

	return r0
}

// MutableBranchNodeStatus_GetPhase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPhase'
type MutableBranchNodeStatus_GetPhase_Call struct {
	*mock.Call
}

// GetPhase is a helper method to define mock.On call
func (_e *MutableBranchNodeStatus_Expecter) GetPhase() *MutableBranchNodeStatus_GetPhase_Call {
	return &MutableBranchNodeStatus_GetPhase_Call{Call: _e.mock.On("GetPhase")}
}

func (_c *MutableBranchNodeStatus_GetPhase_Call) Run(run func()) *MutableBranchNodeStatus_GetPhase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableBranchNodeStatus_GetPhase_Call) Return(_a0 v1alpha1.BranchNodePhase) *MutableBranchNodeStatus_GetPhase_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableBranchNodeStatus_GetPhase_Call) RunAndReturn(run func() v1alpha1.BranchNodePhase) *MutableBranchNodeStatus_GetPhase_Call {
	_c.Call.Return(run)
	return _c
}

// IsDirty provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) IsDirty() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsDirty")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MutableBranchNodeStatus_IsDirty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDirty'
type MutableBranchNodeStatus_IsDirty_Call struct {
	*mock.Call
}

// IsDirty is a helper method to define mock.On call
func (_e *MutableBranchNodeStatus_Expecter) IsDirty() *MutableBranchNodeStatus_IsDirty_Call {
	return &MutableBranchNodeStatus_IsDirty_Call{Call: _e.mock.On("IsDirty")}
}

func (_c *MutableBranchNodeStatus_IsDirty_Call) Run(run func()) *MutableBranchNodeStatus_IsDirty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableBranchNodeStatus_IsDirty_Call) Return(_a0 bool) *MutableBranchNodeStatus_IsDirty_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MutableBranchNodeStatus_IsDirty_Call) RunAndReturn(run func() bool) *MutableBranchNodeStatus_IsDirty_Call {
	_c.Call.Return(run)
	return _c
}

// SetBranchNodeError provides a mock function with given fields:
func (_m *MutableBranchNodeStatus) SetBranchNodeError() {
	_m.Called()
}

// MutableBranchNodeStatus_SetBranchNodeError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBranchNodeError'
type MutableBranchNodeStatus_SetBranchNodeError_Call struct {
	*mock.Call
}

// SetBranchNodeError is a helper method to define mock.On call
func (_e *MutableBranchNodeStatus_Expecter) SetBranchNodeError() *MutableBranchNodeStatus_SetBranchNodeError_Call {
	return &MutableBranchNodeStatus_SetBranchNodeError_Call{Call: _e.mock.On("SetBranchNodeError")}
}

func (_c *MutableBranchNodeStatus_SetBranchNodeError_Call) Run(run func()) *MutableBranchNodeStatus_SetBranchNodeError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MutableBranchNodeStatus_SetBranchNodeError_Call) Return() *MutableBranchNodeStatus_SetBranchNodeError_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableBranchNodeStatus_SetBranchNodeError_Call) RunAndReturn(run func()) *MutableBranchNodeStatus_SetBranchNodeError_Call {
	_c.Call.Return(run)
	return _c
}

// SetBranchNodeSuccess provides a mock function with given fields: id
func (_m *MutableBranchNodeStatus) SetBranchNodeSuccess(id string) {
	_m.Called(id)
}

// MutableBranchNodeStatus_SetBranchNodeSuccess_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBranchNodeSuccess'
type MutableBranchNodeStatus_SetBranchNodeSuccess_Call struct {
	*mock.Call
}

// SetBranchNodeSuccess is a helper method to define mock.On call
//   - id string
func (_e *MutableBranchNodeStatus_Expecter) SetBranchNodeSuccess(id interface{}) *MutableBranchNodeStatus_SetBranchNodeSuccess_Call {
	return &MutableBranchNodeStatus_SetBranchNodeSuccess_Call{Call: _e.mock.On("SetBranchNodeSuccess", id)}
}

func (_c *MutableBranchNodeStatus_SetBranchNodeSuccess_Call) Run(run func(id string)) *MutableBranchNodeStatus_SetBranchNodeSuccess_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MutableBranchNodeStatus_SetBranchNodeSuccess_Call) Return() *MutableBranchNodeStatus_SetBranchNodeSuccess_Call {
	_c.Call.Return()
	return _c
}

func (_c *MutableBranchNodeStatus_SetBranchNodeSuccess_Call) RunAndReturn(run func(string)) *MutableBranchNodeStatus_SetBranchNodeSuccess_Call {
	_c.Call.Return(run)
	return _c
}

// NewMutableBranchNodeStatus creates a new instance of MutableBranchNodeStatus. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMutableBranchNodeStatus(t interface {
	mock.TestingT
	Cleanup(func())
}) *MutableBranchNodeStatus {
	mock := &MutableBranchNodeStatus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
