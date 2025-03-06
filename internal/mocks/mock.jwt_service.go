// Code generated by mockery v2.53.0. DO NOT EDIT.

package mocks

import (
	jwt "gin-framework-boilerplate/pkg/jwt"

	mock "github.com/stretchr/testify/mock"
)

// JWTService is an autogenerated mock type for the JWTService type
type JWTService struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: userId, role, email
func (_m *JWTService) GenerateToken(userId string, role string, email string) (string, error) {
	ret := _m.Called(userId, role, email)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(userId, role, email)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(userId, role, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(userId, role, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseToken provides a mock function with given fields: tokenString
func (_m *JWTService) ParseToken(tokenString string) (jwt.JwtCustomClaim, error) {
	ret := _m.Called(tokenString)

	if len(ret) == 0 {
		panic("no return value specified for ParseToken")
	}

	var r0 jwt.JwtCustomClaim
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (jwt.JwtCustomClaim, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) jwt.JwtCustomClaim); ok {
		r0 = rf(tokenString)
	} else {
		r0 = ret.Get(0).(jwt.JwtCustomClaim)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTService creates a new instance of JWTService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTService(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTService {
	mock := &JWTService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
