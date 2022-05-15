// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"

	testing "testing"

	time "time"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Authenticator provides a mock function with given fields: ctx
func (_m *AuthService) Authenticator(ctx *gin.Context) (interface{}, error) {
	ret := _m.Called(ctx)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*gin.Context) interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdentityHandler provides a mock function with given fields: ctx
func (_m *AuthService) IdentityHandler(ctx *gin.Context) interface{} {
	ret := _m.Called(ctx)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(*gin.Context) interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// LoginResponse provides a mock function with given fields: ctx, code, token, expire
func (_m *AuthService) LoginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	_m.Called(ctx, code, token, expire)
}

// LogoutResponse provides a mock function with given fields: ctx, code
func (_m *AuthService) LogoutResponse(ctx *gin.Context, code int) {
	_m.Called(ctx, code)
}

// PayloadFunc provides a mock function with given fields: data
func (_m *AuthService) PayloadFunc(data interface{}) jwt.MapClaims {
	ret := _m.Called(data)

	var r0 jwt.MapClaims
	if rf, ok := ret.Get(0).(func(interface{}) jwt.MapClaims); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.MapClaims)
		}
	}

	return r0
}

// RefreshResponse provides a mock function with given fields: ctx, code, message, _a3
func (_m *AuthService) RefreshResponse(ctx *gin.Context, code int, message string, _a3 time.Time) {
	_m.Called(ctx, code, message, _a3)
}

// Unauthorized provides a mock function with given fields: ctx, code, message
func (_m *AuthService) Unauthorized(ctx *gin.Context, code int, message string) {
	_m.Called(ctx, code, message)
}

// NewAuthService creates a new instance of AuthService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthService(t testing.TB) *AuthService {
	mock := &AuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}