// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onflow/flow-go/network (interfaces: Network)

// Package mocknetwork is a generated GoMock package.
package mocknetwork

import (
	gomock "github.com/golang/mock/gomock"
	go_datastore "github.com/ipfs/go-datastore"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"
	network "github.com/onflow/flow-go/network"
	reflect "reflect"
)

// MockNetwork is a mock of Network interface
type MockNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkMockRecorder
}

// MockNetworkMockRecorder is the mock recorder for MockNetwork
type MockNetworkMockRecorder struct {
	mock *MockNetwork
}

// NewMockNetwork creates a new mock instance
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// Done mocks base method
func (m *MockNetwork) Done() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Done indicates an expected call of Done
func (mr *MockNetworkMockRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockNetwork)(nil).Done))
}

// Ready mocks base method
func (m *MockNetwork) Ready() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Ready indicates an expected call of Ready
func (mr *MockNetworkMockRecorder) Ready() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockNetwork)(nil).Ready))
}

// Register mocks base method
func (m *MockNetwork) Register(arg0 network.Channel, arg1 network.MessageProcessor) (network.Conduit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(network.Conduit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockNetworkMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockNetwork)(nil).Register), arg0, arg1)
}

// RegisterBlobService mocks base method
func (m *MockNetwork) RegisterBlobService(arg0 network.Channel, arg1 go_datastore.Batching) (network.BlobService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterBlobService", arg0, arg1)
	ret0, _ := ret[0].(network.BlobService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterBlobService indicates an expected call of RegisterBlobService
func (mr *MockNetworkMockRecorder) RegisterBlobService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterBlobService", reflect.TypeOf((*MockNetwork)(nil).RegisterBlobService), arg0, arg1)
}

// RegisterPingService mocks base method
func (m *MockNetwork) RegisterPingService(arg0 protocol.ID, arg1 network.PingInfoProvider) (network.PingService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterPingService", arg0, arg1)
	ret0, _ := ret[0].(network.PingService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterPingService indicates an expected call of RegisterPingService
func (mr *MockNetworkMockRecorder) RegisterPingService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPingService", reflect.TypeOf((*MockNetwork)(nil).RegisterPingService), arg0, arg1)
}

// Start mocks base method
func (m *MockNetwork) Start(arg0 irrecoverable.SignalerContext) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start
func (mr *MockNetworkMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockNetwork)(nil).Start), arg0)
}
