// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/attenberger/quic-go (interfaces: Unpacker)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	wire "github.com/attenberger/quic-go/internal/wire"
	gomock "github.com/golang/mock/gomock"
)

// MockUnpacker is a mock of Unpacker interface
type MockUnpacker struct {
	ctrl     *gomock.Controller
	recorder *MockUnpackerMockRecorder
}

// MockUnpackerMockRecorder is the mock recorder for MockUnpacker
type MockUnpackerMockRecorder struct {
	mock *MockUnpacker
}

// NewMockUnpacker creates a new mock instance
func NewMockUnpacker(ctrl *gomock.Controller) *MockUnpacker {
	mock := &MockUnpacker{ctrl: ctrl}
	mock.recorder = &MockUnpackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnpacker) EXPECT() *MockUnpackerMockRecorder {
	return m.recorder
}

// Unpack mocks base method
func (m *MockUnpacker) Unpack(arg0 *wire.Header, arg1 []byte) (*unpackedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unpack", arg0, arg1)
	ret0, _ := ret[0].(*unpackedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unpack indicates an expected call of Unpack
func (mr *MockUnpackerMockRecorder) Unpack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unpack", reflect.TypeOf((*MockUnpacker)(nil).Unpack), arg0, arg1)
}
