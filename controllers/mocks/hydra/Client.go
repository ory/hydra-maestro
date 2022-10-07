// Copyright © 2022 Ory Corp

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	hydra "github.com/ory/hydra-maester/hydra"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// DeleteOAuth2Client provides a mock function with given fields: id
func (_m *Client) DeleteOAuth2Client(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOAuth2Client provides a mock function with given fields: id
func (_m *Client) GetOAuth2Client(id string) (*hydra.OAuth2ClientJSON, bool, error) {
	ret := _m.Called(id)

	var r0 *hydra.OAuth2ClientJSON
	if rf, ok := ret.Get(0).(func(string) *hydra.OAuth2ClientJSON); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*hydra.OAuth2ClientJSON)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListOAuth2Client provides a mock function with given fields:
func (_m *Client) ListOAuth2Client() ([]*hydra.OAuth2ClientJSON, error) {
	ret := _m.Called()

	var r0 []*hydra.OAuth2ClientJSON
	if rf, ok := ret.Get(0).(func() []*hydra.OAuth2ClientJSON); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*hydra.OAuth2ClientJSON)
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

// PostOAuth2Client provides a mock function with given fields: o
func (_m *Client) PostOAuth2Client(o *hydra.OAuth2ClientJSON) (*hydra.OAuth2ClientJSON, error) {
	ret := _m.Called(o)

	var r0 *hydra.OAuth2ClientJSON
	if rf, ok := ret.Get(0).(func(*hydra.OAuth2ClientJSON) *hydra.OAuth2ClientJSON); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*hydra.OAuth2ClientJSON)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*hydra.OAuth2ClientJSON) error); ok {
		r1 = rf(o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutOAuth2Client provides a mock function with given fields: o
func (_m *Client) PutOAuth2Client(o *hydra.OAuth2ClientJSON) (*hydra.OAuth2ClientJSON, error) {
	ret := _m.Called(o)

	var r0 *hydra.OAuth2ClientJSON
	if rf, ok := ret.Get(0).(func(*hydra.OAuth2ClientJSON) *hydra.OAuth2ClientJSON); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*hydra.OAuth2ClientJSON)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*hydra.OAuth2ClientJSON) error); ok {
		r1 = rf(o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
