// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	model "github.com/alexandre-pinon/epic-road-trip/model"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// GoogleService is an autogenerated mock type for the GoogleService type
type GoogleService struct {
	mock.Mock
}

// Drink provides a mock function with given fields: url, position, constraint
func (_m *GoogleService) Drink(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {
	ret := _m.Called(url, position, constraint)

	var r0 *[]model.ActivityResult
	if rf, ok := ret.Get(0).(func(string, model.Location, model.Constraints) *[]model.ActivityResult); ok {
		r0 = rf(url, position, constraint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.ActivityResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.Location, model.Constraints) error); ok {
		r1 = rf(url, position, constraint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Eat provides a mock function with given fields: url, position, constraint
func (_m *GoogleService) Eat(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {
	ret := _m.Called(url, position, constraint)

	var r0 *[]model.ActivityResult
	if rf, ok := ret.Get(0).(func(string, model.Location, model.Constraints) *[]model.ActivityResult); ok {
		r0 = rf(url, position, constraint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.ActivityResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.Location, model.Constraints) error); ok {
		r1 = rf(url, position, constraint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Enjoy provides a mock function with given fields: url, position, constraint
func (_m *GoogleService) Enjoy(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {
	ret := _m.Called(url, position, constraint)

	var r0 *[]model.ActivityResult
	if rf, ok := ret.Get(0).(func(string, model.Location, model.Constraints) *[]model.ActivityResult); ok {
		r0 = rf(url, position, constraint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.ActivityResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.Location, model.Constraints) error); ok {
		r1 = rf(url, position, constraint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GeoCoding provides a mock function with given fields: url, position
func (_m *GoogleService) GeoCoding(url string, position string) (*model.Location, error) {
	ret := _m.Called(url, position)

	var r0 *model.Location
	if rf, ok := ret.Get(0).(func(string, string) *model.Location); ok {
		r0 = rf(url, position)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(url, position)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDirections provides a mock function with given fields: url, directionsFormData
func (_m *GoogleService) GetDirections(url string, directionsFormData *model.DirectionsFormData) (*[]model.Itinerary, error) {
	ret := _m.Called(url, directionsFormData)

	var r0 *[]model.Itinerary
	if rf, ok := ret.Get(0).(func(string, *model.DirectionsFormData) *[]model.Itinerary); ok {
		r0 = rf(url, directionsFormData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.Itinerary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *model.DirectionsFormData) error); ok {
		r1 = rf(url, directionsFormData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sleep provides a mock function with given fields: url, position, constraint
func (_m *GoogleService) Sleep(url string, position model.Location, constraint model.Constraints) (*[]model.ActivityResult, error) {
	ret := _m.Called(url, position, constraint)

	var r0 *[]model.ActivityResult
	if rf, ok := ret.Get(0).(func(string, model.Location, model.Constraints) *[]model.ActivityResult); ok {
		r0 = rf(url, position, constraint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.ActivityResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.Location, model.Constraints) error); ok {
		r1 = rf(url, position, constraint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGoogleService creates a new instance of GoogleService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewGoogleService(t testing.TB) *GoogleService {
	mock := &GoogleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
