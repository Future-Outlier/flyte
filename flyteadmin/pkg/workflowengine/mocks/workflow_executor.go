// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	interfaces "github.com/flyteorg/flyte/flyteadmin/pkg/workflowengine/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// WorkflowExecutor is an autogenerated mock type for the WorkflowExecutor type
type WorkflowExecutor struct {
	mock.Mock
}

type WorkflowExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *WorkflowExecutor) EXPECT() *WorkflowExecutor_Expecter {
	return &WorkflowExecutor_Expecter{mock: &_m.Mock}
}

// Abort provides a mock function with given fields: ctx, data
func (_m *WorkflowExecutor) Abort(ctx context.Context, data interfaces.AbortData) error {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Abort")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.AbortData) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WorkflowExecutor_Abort_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Abort'
type WorkflowExecutor_Abort_Call struct {
	*mock.Call
}

// Abort is a helper method to define mock.On call
//   - ctx context.Context
//   - data interfaces.AbortData
func (_e *WorkflowExecutor_Expecter) Abort(ctx interface{}, data interface{}) *WorkflowExecutor_Abort_Call {
	return &WorkflowExecutor_Abort_Call{Call: _e.mock.On("Abort", ctx, data)}
}

func (_c *WorkflowExecutor_Abort_Call) Run(run func(ctx context.Context, data interfaces.AbortData)) *WorkflowExecutor_Abort_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interfaces.AbortData))
	})
	return _c
}

func (_c *WorkflowExecutor_Abort_Call) Return(_a0 error) *WorkflowExecutor_Abort_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WorkflowExecutor_Abort_Call) RunAndReturn(run func(context.Context, interfaces.AbortData) error) *WorkflowExecutor_Abort_Call {
	_c.Call.Return(run)
	return _c
}

// Execute provides a mock function with given fields: ctx, data
func (_m *WorkflowExecutor) Execute(ctx context.Context, data interfaces.ExecutionData) (interfaces.ExecutionResponse, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 interfaces.ExecutionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.ExecutionData) (interfaces.ExecutionResponse, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.ExecutionData) interfaces.ExecutionResponse); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(interfaces.ExecutionResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interfaces.ExecutionData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WorkflowExecutor_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type WorkflowExecutor_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - data interfaces.ExecutionData
func (_e *WorkflowExecutor_Expecter) Execute(ctx interface{}, data interface{}) *WorkflowExecutor_Execute_Call {
	return &WorkflowExecutor_Execute_Call{Call: _e.mock.On("Execute", ctx, data)}
}

func (_c *WorkflowExecutor_Execute_Call) Run(run func(ctx context.Context, data interfaces.ExecutionData)) *WorkflowExecutor_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interfaces.ExecutionData))
	})
	return _c
}

func (_c *WorkflowExecutor_Execute_Call) Return(_a0 interfaces.ExecutionResponse, _a1 error) *WorkflowExecutor_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WorkflowExecutor_Execute_Call) RunAndReturn(run func(context.Context, interfaces.ExecutionData) (interfaces.ExecutionResponse, error)) *WorkflowExecutor_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with given fields:
func (_m *WorkflowExecutor) ID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WorkflowExecutor_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type WorkflowExecutor_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *WorkflowExecutor_Expecter) ID() *WorkflowExecutor_ID_Call {
	return &WorkflowExecutor_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *WorkflowExecutor_ID_Call) Run(run func()) *WorkflowExecutor_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WorkflowExecutor_ID_Call) Return(_a0 string) *WorkflowExecutor_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WorkflowExecutor_ID_Call) RunAndReturn(run func() string) *WorkflowExecutor_ID_Call {
	_c.Call.Return(run)
	return _c
}

// NewWorkflowExecutor creates a new instance of WorkflowExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWorkflowExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *WorkflowExecutor {
	mock := &WorkflowExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
