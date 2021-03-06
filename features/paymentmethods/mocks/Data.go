// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	paymentmethods "bayareen-backend/features/paymentmethods"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *Data) Create(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	ret := _m.Called(data)

	var r0 *paymentmethods.Core
	if rf, ok := ret.Get(0).(func(*paymentmethods.Core) *paymentmethods.Core); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentmethods.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*paymentmethods.Core) error); ok {
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
func (_m *Data) GetAll() []paymentmethods.Core {
	ret := _m.Called()

	var r0 []paymentmethods.Core
	if rf, ok := ret.Get(0).(func() []paymentmethods.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]paymentmethods.Core)
		}
	}

	return r0
}

// GetById provides a mock function with given fields: id
func (_m *Data) GetById(id int) (*paymentmethods.Core, error) {
	ret := _m.Called(id)

	var r0 *paymentmethods.Core
	if rf, ok := ret.Get(0).(func(int) *paymentmethods.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentmethods.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: method, channel
func (_m *Data) GetByName(method string, channel string) (int, error) {
	ret := _m.Called(method, channel)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(method, channel)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(method, channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: data
func (_m *Data) Update(data *paymentmethods.Core) (*paymentmethods.Core, error) {
	ret := _m.Called(data)

	var r0 *paymentmethods.Core
	if rf, ok := ret.Get(0).(func(*paymentmethods.Core) *paymentmethods.Core); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paymentmethods.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*paymentmethods.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
