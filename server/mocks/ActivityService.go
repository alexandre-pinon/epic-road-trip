// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "github.com/alexandre-pinon/epic-road-trip/model"
	mock "github.com/stretchr/testify/mock"
)

// ActivityService is an autogenerated mock type for the ActivityService type
type ActivityService struct {
	mock.Mock
}

// Enjoy provides a mock function with given fields: position
func (_m *ActivityService) Enjoy(position string) ([]*model.Activity, error) {
	ret := _m.Called(position)

	var r0 []*model.Activity
	if rf, ok := ret.Get(0).(func(string) []*model.Activity); ok {
		r0 = rf(position)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Activity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(position)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
