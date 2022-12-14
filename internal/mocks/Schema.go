// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	schema "github.com/jsightapi/jsight-schema-core"
	mock "github.com/stretchr/testify/mock"
)

// Schema is an autogenerated mock type for the Schema type
type Schema struct {
	mock.Mock
}

// AddRule provides a mock function with given fields: name, _a1
func (_m *Schema) AddRule(name string, _a1 schema.Rule) error {
	ret := _m.Called(name, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, schema.Rule) error); ok {
		r0 = rf(name, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddType provides a mock function with given fields: name, _a1
func (_m *Schema) AddType(name string, _a1 schema.Schema) error {
	ret := _m.Called(name, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, schema.Schema) error); ok {
		r0 = rf(name, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Check provides a mock function with given fields:
func (_m *Schema) Check() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Example provides a mock function with given fields:
func (_m *Schema) Example() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAST provides a mock function with given fields:
func (_m *Schema) GetAST() (schema.ASTNode, error) {
	ret := _m.Called()

	var r0 schema.ASTNode
	if rf, ok := ret.Get(0).(func() schema.ASTNode); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(schema.ASTNode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Len provides a mock function with given fields:
func (_m *Schema) Len() (uint, error) {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UsedUserTypes provides a mock function with given fields:
func (_m *Schema) UsedUserTypes() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewSchemaT interface {
	mock.TestingT
	Cleanup(func())
}

// NewSchema creates a new instance of Schema. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSchema(t NewSchemaT) *Schema {
	mock := &Schema{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
