// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/flyteorg/flyte/datacatalog/pkg/repositories/models"
)

// PartitionRepo is an autogenerated mock type for the PartitionRepo type
type PartitionRepo struct {
	mock.Mock
}

type PartitionRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *PartitionRepo) EXPECT() *PartitionRepo_Expecter {
	return &PartitionRepo_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, in
func (_m *PartitionRepo) Create(ctx context.Context, in models.Partition) error {
	ret := _m.Called(ctx, in)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Partition) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PartitionRepo_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type PartitionRepo_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - in models.Partition
func (_e *PartitionRepo_Expecter) Create(ctx interface{}, in interface{}) *PartitionRepo_Create_Call {
	return &PartitionRepo_Create_Call{Call: _e.mock.On("Create", ctx, in)}
}

func (_c *PartitionRepo_Create_Call) Run(run func(ctx context.Context, in models.Partition)) *PartitionRepo_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Partition))
	})
	return _c
}

func (_c *PartitionRepo_Create_Call) Return(_a0 error) *PartitionRepo_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PartitionRepo_Create_Call) RunAndReturn(run func(context.Context, models.Partition) error) *PartitionRepo_Create_Call {
	_c.Call.Return(run)
	return _c
}

// NewPartitionRepo creates a new instance of PartitionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPartitionRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *PartitionRepo {
	mock := &PartitionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
