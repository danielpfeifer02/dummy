// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: CryptoDataHandler)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_crypto_data_handler_test.go github.com/quic-go/quic-go CryptoDataHandler
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	handshake "github.com/danielpfeifer02/quic-go-no-crypto/internal/handshake"
	protocol "github.com/danielpfeifer02/quic-go-no-crypto/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockCryptoDataHandler is a mock of CryptoDataHandler interface.
type MockCryptoDataHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoDataHandlerMockRecorder
}

// MockCryptoDataHandlerMockRecorder is the mock recorder for MockCryptoDataHandler.
type MockCryptoDataHandlerMockRecorder struct {
	mock *MockCryptoDataHandler
}

// NewMockCryptoDataHandler creates a new mock instance.
func NewMockCryptoDataHandler(ctrl *gomock.Controller) *MockCryptoDataHandler {
	mock := &MockCryptoDataHandler{ctrl: ctrl}
	mock.recorder = &MockCryptoDataHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoDataHandler) EXPECT() *MockCryptoDataHandlerMockRecorder {
	return m.recorder
}

// HandleMessage mocks base method.
func (m *MockCryptoDataHandler) HandleMessage(arg0 []byte, arg1 protocol.EncryptionLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockCryptoDataHandlerMockRecorder) HandleMessage(arg0, arg1 any) *CryptoDataHandlerHandleMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockCryptoDataHandler)(nil).HandleMessage), arg0, arg1)
	return &CryptoDataHandlerHandleMessageCall{Call: call}
}

// CryptoDataHandlerHandleMessageCall wrap *gomock.Call
type CryptoDataHandlerHandleMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoDataHandlerHandleMessageCall) Return(arg0 error) *CryptoDataHandlerHandleMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoDataHandlerHandleMessageCall) Do(f func([]byte, protocol.EncryptionLevel) error) *CryptoDataHandlerHandleMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoDataHandlerHandleMessageCall) DoAndReturn(f func([]byte, protocol.EncryptionLevel) error) *CryptoDataHandlerHandleMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NextEvent mocks base method.
func (m *MockCryptoDataHandler) NextEvent() handshake.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextEvent")
	ret0, _ := ret[0].(handshake.Event)
	return ret0
}

// NextEvent indicates an expected call of NextEvent.
func (mr *MockCryptoDataHandlerMockRecorder) NextEvent() *CryptoDataHandlerNextEventCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextEvent", reflect.TypeOf((*MockCryptoDataHandler)(nil).NextEvent))
	return &CryptoDataHandlerNextEventCall{Call: call}
}

// CryptoDataHandlerNextEventCall wrap *gomock.Call
type CryptoDataHandlerNextEventCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoDataHandlerNextEventCall) Return(arg0 handshake.Event) *CryptoDataHandlerNextEventCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoDataHandlerNextEventCall) Do(f func() handshake.Event) *CryptoDataHandlerNextEventCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoDataHandlerNextEventCall) DoAndReturn(f func() handshake.Event) *CryptoDataHandlerNextEventCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
