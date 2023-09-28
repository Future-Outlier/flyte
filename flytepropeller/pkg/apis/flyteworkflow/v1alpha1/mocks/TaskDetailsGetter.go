// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// TaskDetailsGetter is an autogenerated mock type for the TaskDetailsGetter type
type TaskDetailsGetter struct {
	mock.Mock
}

type TaskDetailsGetter_GetTask struct {
	*mock.Call
}

func (_m TaskDetailsGetter_GetTask) Return(_a0 v1alpha1.ExecutableTask, _a1 error) *TaskDetailsGetter_GetTask {
	return &TaskDetailsGetter_GetTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TaskDetailsGetter) OnGetTask(id string) *TaskDetailsGetter_GetTask {
	c_call := _m.On("GetTask", id)
	return &TaskDetailsGetter_GetTask{Call: c_call}
}

func (_m *TaskDetailsGetter) OnGetTaskMatch(matchers ...interface{}) *TaskDetailsGetter_GetTask {
	c_call := _m.On("GetTask", matchers...)
	return &TaskDetailsGetter_GetTask{Call: c_call}
}

// GetTask provides a mock function with given fields: id
func (_m *TaskDetailsGetter) GetTask(id string) (v1alpha1.ExecutableTask, error) {
	ret := _m.Called(id)

	var r0 v1alpha1.ExecutableTask
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableTask); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
