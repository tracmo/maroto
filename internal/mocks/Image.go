// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	internal "github.com/tracmo/maroto/internal"
	consts "github.com/tracmo/maroto/pkg/consts"

	mock "github.com/stretchr/testify/mock"

	props "github.com/tracmo/maroto/pkg/props"
)

// Image is an autogenerated mock type for the Image type
type Image struct {
	mock.Mock
}

// AddFromBase64 provides a mock function with given fields: stringBase64, cell, prop, extension
func (_m *Image) AddFromBase64(stringBase64 string, cell internal.Cell, prop props.Rect, extension consts.Extension) error {
	ret := _m.Called(stringBase64, cell, prop, extension)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, internal.Cell, props.Rect, consts.Extension) error); ok {
		r0 = rf(stringBase64, cell, prop, extension)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddFromFile provides a mock function with given fields: path, cell, prop
func (_m *Image) AddFromFile(path string, cell internal.Cell, prop props.Rect) error {
	ret := _m.Called(path, cell, prop)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, internal.Cell, props.Rect) error); ok {
		r0 = rf(path, cell, prop)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewImage interface {
	mock.TestingT
	Cleanup(func())
}

// NewImage creates a new instance of Image. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewImage(t mockConstructorTestingTNewImage) *Image {
	mock := &Image{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
