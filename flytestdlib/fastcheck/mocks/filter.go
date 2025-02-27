// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Filter is an autogenerated mock type for the Filter type
type Filter struct {
	mock.Mock
}

type Filter_Expecter struct {
	mock *mock.Mock
}

func (_m *Filter) EXPECT() *Filter_Expecter {
	return &Filter_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, id
func (_m *Filter) Add(ctx context.Context, id []byte) bool {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, []byte) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Filter_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Filter_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - id []byte
func (_e *Filter_Expecter) Add(ctx interface{}, id interface{}) *Filter_Add_Call {
	return &Filter_Add_Call{Call: _e.mock.On("Add", ctx, id)}
}

func (_c *Filter_Add_Call) Run(run func(ctx context.Context, id []byte)) *Filter_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte))
	})
	return _c
}

func (_c *Filter_Add_Call) Return(evicted bool) *Filter_Add_Call {
	_c.Call.Return(evicted)
	return _c
}

func (_c *Filter_Add_Call) RunAndReturn(run func(context.Context, []byte) bool) *Filter_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Contains provides a mock function with given fields: ctx, id
func (_m *Filter) Contains(ctx context.Context, id []byte) bool {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Contains")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, []byte) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Filter_Contains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Contains'
type Filter_Contains_Call struct {
	*mock.Call
}

// Contains is a helper method to define mock.On call
//   - ctx context.Context
//   - id []byte
func (_e *Filter_Expecter) Contains(ctx interface{}, id interface{}) *Filter_Contains_Call {
	return &Filter_Contains_Call{Call: _e.mock.On("Contains", ctx, id)}
}

func (_c *Filter_Contains_Call) Run(run func(ctx context.Context, id []byte)) *Filter_Contains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte))
	})
	return _c
}

func (_c *Filter_Contains_Call) Return(_a0 bool) *Filter_Contains_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Filter_Contains_Call) RunAndReturn(run func(context.Context, []byte) bool) *Filter_Contains_Call {
	_c.Call.Return(run)
	return _c
}

// NewFilter creates a new instance of Filter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFilter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Filter {
	mock := &Filter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
