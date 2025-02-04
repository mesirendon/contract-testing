// Code generated by mockery v2.51.1. DO NOT EDIT.

package handler

import (
	model "github.com/mesirendon/contract-testing/provider/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// mockUserGetter is an autogenerated mock type for the userGetter type
type mockUserGetter struct {
	mock.Mock
}

type mockUserGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *mockUserGetter) EXPECT() *mockUserGetter_Expecter {
	return &mockUserGetter_Expecter{mock: &_m.Mock}
}

// GetUser provides a mock function with given fields: id
func (_m *mockUserGetter) GetUser(id string) (model.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) model.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockUserGetter_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type mockUserGetter_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - id string
func (_e *mockUserGetter_Expecter) GetUser(id interface{}) *mockUserGetter_GetUser_Call {
	return &mockUserGetter_GetUser_Call{Call: _e.mock.On("GetUser", id)}
}

func (_c *mockUserGetter_GetUser_Call) Run(run func(id string)) *mockUserGetter_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *mockUserGetter_GetUser_Call) Return(_a0 model.User, _a1 error) *mockUserGetter_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockUserGetter_GetUser_Call) RunAndReturn(run func(string) (model.User, error)) *mockUserGetter_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// newMockUserGetter creates a new instance of mockUserGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockUserGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockUserGetter {
	mock := &mockUserGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
