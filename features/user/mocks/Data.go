// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	user "bayareen-backend/features/user"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *Data) Create(data user.UserCore) (user.UserCore, error) {
	ret := _m.Called(data)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Data) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Data) GetAll() []user.UserCore {
	ret := _m.Called()

	var r0 []user.UserCore
	if rf, ok := ret.Get(0).(func() []user.UserCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.UserCore)
		}
	}

	return r0
}

// GetByEmail provides a mock function with given fields: email
func (_m *Data) GetByEmail(email string) (user.UserCore, error) {
	ret := _m.Called(email)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(string) user.UserCore); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Data) GetById(id int) (user.UserCore, error) {
	ret := _m.Called(id)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(int) user.UserCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: _a0
func (_m *Data) Login(_a0 user.UserCore) (user.UserCore, error) {
	ret := _m.Called(_a0)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: data
func (_m *Data) Update(data user.UserCore) (user.UserCore, error) {
	ret := _m.Called(data)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
