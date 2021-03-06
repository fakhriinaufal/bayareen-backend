// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import products "bayareen-backend/features/products"

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *Data) Create(data *products.Core) (*products.Core, error) {
	ret := _m.Called(data)

	var r0 *products.Core
	if rf, ok := ret.Get(0).(func(*products.Core) *products.Core); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Data) Delete(id []int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Data) GetAll() []products.Core {
	ret := _m.Called()

	var r0 []products.Core
	if rf, ok := ret.Get(0).(func() []products.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Core)
		}
	}

	return r0
}

// GetById provides a mock function with given fields: id
func (_m *Data) GetById(id int) (*products.Core, error) {
	ret := _m.Called(id)

	var r0 *products.Core
	if rf, ok := ret.Get(0).(func(int) *products.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.Core)
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

// GetByProviderId provides a mock function with given fields: providerId
func (_m *Data) GetByProviderId(providerId int) ([]products.Core, error) {
	ret := _m.Called(providerId)

	var r0 []products.Core
	if rf, ok := ret.Get(0).(func(int) []products.Core); ok {
		r0 = rf(providerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(providerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: data
func (_m *Data) Update(data *products.Core) (*products.Core, error) {
	ret := _m.Called(data)

	var r0 *products.Core
	if rf, ok := ret.Get(0).(func(*products.Core) *products.Core); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*products.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*products.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
