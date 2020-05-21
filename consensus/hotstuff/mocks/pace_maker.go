// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/dapperlabs/flow-go/consensus/hotstuff/model"
import time "time"

// PaceMaker is an autogenerated mock type for the PaceMaker type
type PaceMaker struct {
	mock.Mock
}

// BlockRateDelay provides a mock function with given fields:
func (_m *PaceMaker) BlockRateDelay() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// CurView provides a mock function with given fields:
func (_m *PaceMaker) CurView() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// OnTimeout provides a mock function with given fields:
func (_m *PaceMaker) OnTimeout() *model.NewViewEvent {
	ret := _m.Called()

	var r0 *model.NewViewEvent
	if rf, ok := ret.Get(0).(func() *model.NewViewEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NewViewEvent)
		}
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *PaceMaker) Start() {
	_m.Called()
}

// TimeoutChannel provides a mock function with given fields:
func (_m *PaceMaker) TimeoutChannel() <-chan time.Time {
	ret := _m.Called()

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func() <-chan time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// UpdateCurViewWithBlock provides a mock function with given fields: block, isLeaderForNextView
func (_m *PaceMaker) UpdateCurViewWithBlock(block *model.Block, isLeaderForNextView bool) (*model.NewViewEvent, bool) {
	ret := _m.Called(block, isLeaderForNextView)

	var r0 *model.NewViewEvent
	if rf, ok := ret.Get(0).(func(*model.Block, bool) *model.NewViewEvent); ok {
		r0 = rf(block, isLeaderForNextView)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NewViewEvent)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*model.Block, bool) bool); ok {
		r1 = rf(block, isLeaderForNextView)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// UpdateCurViewWithQC provides a mock function with given fields: qc
func (_m *PaceMaker) UpdateCurViewWithQC(qc *model.QuorumCertificate) (*model.NewViewEvent, bool) {
	ret := _m.Called(qc)

	var r0 *model.NewViewEvent
	if rf, ok := ret.Get(0).(func(*model.QuorumCertificate) *model.NewViewEvent); ok {
		r0 = rf(qc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NewViewEvent)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(*model.QuorumCertificate) bool); ok {
		r1 = rf(qc)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}
