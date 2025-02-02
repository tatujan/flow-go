// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockinsecure

import (
	insecure "github.com/onflow/flow-go/insecure"
	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"

	mock "github.com/stretchr/testify/mock"
)

// AttackOrchestrator is an autogenerated mock type for the AttackOrchestrator type
type AttackOrchestrator struct {
	mock.Mock
}

// Done provides a mock function with given fields:
func (_m *AttackOrchestrator) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// HandleEventFromCorruptedNode provides a mock function with given fields: _a0
func (_m *AttackOrchestrator) HandleEventFromCorruptedNode(_a0 *insecure.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*insecure.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *AttackOrchestrator) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *AttackOrchestrator) Start(_a0 irrecoverable.SignalerContext) {
	_m.Called(_a0)
}
