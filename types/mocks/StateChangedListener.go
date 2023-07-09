// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	types "github.com/getamis/alice/types"
	mock "github.com/stretchr/testify/mock"
)

// StateChangedListener is an autogenerated mock type for the StateChangedListener type
type StateChangedListener struct {
	mock.Mock
}

// OnStateChanged provides a mock function with given fields: oldState, newState
func (_m *StateChangedListener) OnStateChanged(oldState types.MainState, newState types.MainState) {
	_m.Called(oldState, newState)
}

// NewStateChangedListener creates a new instance of StateChangedListener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateChangedListener(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateChangedListener {
	mock := &StateChangedListener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
