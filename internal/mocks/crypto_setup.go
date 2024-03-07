// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go/internal/handshake (interfaces: CryptoSetup)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package mocks -destination crypto_setup_tmp.go github.com/quic-go/quic-go/internal/handshake CryptoSetup
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	handshake "github.com/danielpfeifer02/quic-go-no-crypto/internal/handshake"
	protocol "github.com/danielpfeifer02/quic-go-no-crypto/internal/protocol"
	gomock "go.uber.org/mock/gomock"
)

// MockCryptoSetup is a mock of CryptoSetup interface.
type MockCryptoSetup struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoSetupMockRecorder
}

// MockCryptoSetupMockRecorder is the mock recorder for MockCryptoSetup.
type MockCryptoSetupMockRecorder struct {
	mock *MockCryptoSetup
}

// NewMockCryptoSetup creates a new mock instance.
func NewMockCryptoSetup(ctrl *gomock.Controller) *MockCryptoSetup {
	mock := &MockCryptoSetup{ctrl: ctrl}
	mock.recorder = &MockCryptoSetupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoSetup) EXPECT() *MockCryptoSetupMockRecorder {
	return m.recorder
}

// ChangeConnectionID mocks base method.
func (m *MockCryptoSetup) ChangeConnectionID(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ChangeConnectionID", arg0)
}

// ChangeConnectionID indicates an expected call of ChangeConnectionID.
func (mr *MockCryptoSetupMockRecorder) ChangeConnectionID(arg0 any) *CryptoSetupChangeConnectionIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeConnectionID", reflect.TypeOf((*MockCryptoSetup)(nil).ChangeConnectionID), arg0)
	return &CryptoSetupChangeConnectionIDCall{Call: call}
}

// CryptoSetupChangeConnectionIDCall wrap *gomock.Call
type CryptoSetupChangeConnectionIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupChangeConnectionIDCall) Return() *CryptoSetupChangeConnectionIDCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupChangeConnectionIDCall) Do(f func(protocol.ConnectionID)) *CryptoSetupChangeConnectionIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupChangeConnectionIDCall) DoAndReturn(f func(protocol.ConnectionID)) *CryptoSetupChangeConnectionIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockCryptoSetup) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCryptoSetupMockRecorder) Close() *CryptoSetupCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCryptoSetup)(nil).Close))
	return &CryptoSetupCloseCall{Call: call}
}

// CryptoSetupCloseCall wrap *gomock.Call
type CryptoSetupCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupCloseCall) Return(arg0 error) *CryptoSetupCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupCloseCall) Do(f func() error) *CryptoSetupCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupCloseCall) DoAndReturn(f func() error) *CryptoSetupCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ConnectionState mocks base method.
func (m *MockCryptoSetup) ConnectionState() handshake.ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(handshake.ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockCryptoSetupMockRecorder) ConnectionState() *CryptoSetupConnectionStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockCryptoSetup)(nil).ConnectionState))
	return &CryptoSetupConnectionStateCall{Call: call}
}

// CryptoSetupConnectionStateCall wrap *gomock.Call
type CryptoSetupConnectionStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupConnectionStateCall) Return(arg0 handshake.ConnectionState) *CryptoSetupConnectionStateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupConnectionStateCall) Do(f func() handshake.ConnectionState) *CryptoSetupConnectionStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupConnectionStateCall) DoAndReturn(f func() handshake.ConnectionState) *CryptoSetupConnectionStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DiscardInitialKeys mocks base method.
func (m *MockCryptoSetup) DiscardInitialKeys() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DiscardInitialKeys")
}

// DiscardInitialKeys indicates an expected call of DiscardInitialKeys.
func (mr *MockCryptoSetupMockRecorder) DiscardInitialKeys() *CryptoSetupDiscardInitialKeysCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardInitialKeys", reflect.TypeOf((*MockCryptoSetup)(nil).DiscardInitialKeys))
	return &CryptoSetupDiscardInitialKeysCall{Call: call}
}

// CryptoSetupDiscardInitialKeysCall wrap *gomock.Call
type CryptoSetupDiscardInitialKeysCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupDiscardInitialKeysCall) Return() *CryptoSetupDiscardInitialKeysCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupDiscardInitialKeysCall) Do(f func()) *CryptoSetupDiscardInitialKeysCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupDiscardInitialKeysCall) DoAndReturn(f func()) *CryptoSetupDiscardInitialKeysCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get0RTTOpener mocks base method.
func (m *MockCryptoSetup) Get0RTTOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTOpener indicates an expected call of Get0RTTOpener.
func (mr *MockCryptoSetupMockRecorder) Get0RTTOpener() *CryptoSetupGet0RTTOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTOpener", reflect.TypeOf((*MockCryptoSetup)(nil).Get0RTTOpener))
	return &CryptoSetupGet0RTTOpenerCall{Call: call}
}

// CryptoSetupGet0RTTOpenerCall wrap *gomock.Call
type CryptoSetupGet0RTTOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGet0RTTOpenerCall) Return(arg0 handshake.LongHeaderOpener, arg1 error) *CryptoSetupGet0RTTOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGet0RTTOpenerCall) Do(f func() (handshake.LongHeaderOpener, error)) *CryptoSetupGet0RTTOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGet0RTTOpenerCall) DoAndReturn(f func() (handshake.LongHeaderOpener, error)) *CryptoSetupGet0RTTOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get0RTTSealer mocks base method.
func (m *MockCryptoSetup) Get0RTTSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTSealer indicates an expected call of Get0RTTSealer.
func (mr *MockCryptoSetupMockRecorder) Get0RTTSealer() *CryptoSetupGet0RTTSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTSealer", reflect.TypeOf((*MockCryptoSetup)(nil).Get0RTTSealer))
	return &CryptoSetupGet0RTTSealerCall{Call: call}
}

// CryptoSetupGet0RTTSealerCall wrap *gomock.Call
type CryptoSetupGet0RTTSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGet0RTTSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *CryptoSetupGet0RTTSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGet0RTTSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *CryptoSetupGet0RTTSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGet0RTTSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *CryptoSetupGet0RTTSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get1RTTOpener mocks base method.
func (m *MockCryptoSetup) Get1RTTOpener() (handshake.ShortHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTOpener")
	ret0, _ := ret[0].(handshake.ShortHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTOpener indicates an expected call of Get1RTTOpener.
func (mr *MockCryptoSetupMockRecorder) Get1RTTOpener() *CryptoSetupGet1RTTOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTOpener", reflect.TypeOf((*MockCryptoSetup)(nil).Get1RTTOpener))
	return &CryptoSetupGet1RTTOpenerCall{Call: call}
}

// CryptoSetupGet1RTTOpenerCall wrap *gomock.Call
type CryptoSetupGet1RTTOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGet1RTTOpenerCall) Return(arg0 handshake.ShortHeaderOpener, arg1 error) *CryptoSetupGet1RTTOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGet1RTTOpenerCall) Do(f func() (handshake.ShortHeaderOpener, error)) *CryptoSetupGet1RTTOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGet1RTTOpenerCall) DoAndReturn(f func() (handshake.ShortHeaderOpener, error)) *CryptoSetupGet1RTTOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get1RTTSealer mocks base method.
func (m *MockCryptoSetup) Get1RTTSealer() (handshake.ShortHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTSealer")
	ret0, _ := ret[0].(handshake.ShortHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTSealer indicates an expected call of Get1RTTSealer.
func (mr *MockCryptoSetupMockRecorder) Get1RTTSealer() *CryptoSetupGet1RTTSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTSealer", reflect.TypeOf((*MockCryptoSetup)(nil).Get1RTTSealer))
	return &CryptoSetupGet1RTTSealerCall{Call: call}
}

// CryptoSetupGet1RTTSealerCall wrap *gomock.Call
type CryptoSetupGet1RTTSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGet1RTTSealerCall) Return(arg0 handshake.ShortHeaderSealer, arg1 error) *CryptoSetupGet1RTTSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGet1RTTSealerCall) Do(f func() (handshake.ShortHeaderSealer, error)) *CryptoSetupGet1RTTSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGet1RTTSealerCall) DoAndReturn(f func() (handshake.ShortHeaderSealer, error)) *CryptoSetupGet1RTTSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetHandshakeOpener mocks base method.
func (m *MockCryptoSetup) GetHandshakeOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeOpener indicates an expected call of GetHandshakeOpener.
func (mr *MockCryptoSetupMockRecorder) GetHandshakeOpener() *CryptoSetupGetHandshakeOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeOpener", reflect.TypeOf((*MockCryptoSetup)(nil).GetHandshakeOpener))
	return &CryptoSetupGetHandshakeOpenerCall{Call: call}
}

// CryptoSetupGetHandshakeOpenerCall wrap *gomock.Call
type CryptoSetupGetHandshakeOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGetHandshakeOpenerCall) Return(arg0 handshake.LongHeaderOpener, arg1 error) *CryptoSetupGetHandshakeOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGetHandshakeOpenerCall) Do(f func() (handshake.LongHeaderOpener, error)) *CryptoSetupGetHandshakeOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGetHandshakeOpenerCall) DoAndReturn(f func() (handshake.LongHeaderOpener, error)) *CryptoSetupGetHandshakeOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetHandshakeSealer mocks base method.
func (m *MockCryptoSetup) GetHandshakeSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeSealer indicates an expected call of GetHandshakeSealer.
func (mr *MockCryptoSetupMockRecorder) GetHandshakeSealer() *CryptoSetupGetHandshakeSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeSealer", reflect.TypeOf((*MockCryptoSetup)(nil).GetHandshakeSealer))
	return &CryptoSetupGetHandshakeSealerCall{Call: call}
}

// CryptoSetupGetHandshakeSealerCall wrap *gomock.Call
type CryptoSetupGetHandshakeSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGetHandshakeSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *CryptoSetupGetHandshakeSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGetHandshakeSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *CryptoSetupGetHandshakeSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGetHandshakeSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *CryptoSetupGetHandshakeSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetInitialOpener mocks base method.
func (m *MockCryptoSetup) GetInitialOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialOpener indicates an expected call of GetInitialOpener.
func (mr *MockCryptoSetupMockRecorder) GetInitialOpener() *CryptoSetupGetInitialOpenerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialOpener", reflect.TypeOf((*MockCryptoSetup)(nil).GetInitialOpener))
	return &CryptoSetupGetInitialOpenerCall{Call: call}
}

// CryptoSetupGetInitialOpenerCall wrap *gomock.Call
type CryptoSetupGetInitialOpenerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGetInitialOpenerCall) Return(arg0 handshake.LongHeaderOpener, arg1 error) *CryptoSetupGetInitialOpenerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGetInitialOpenerCall) Do(f func() (handshake.LongHeaderOpener, error)) *CryptoSetupGetInitialOpenerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGetInitialOpenerCall) DoAndReturn(f func() (handshake.LongHeaderOpener, error)) *CryptoSetupGetInitialOpenerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetInitialSealer mocks base method.
func (m *MockCryptoSetup) GetInitialSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialSealer indicates an expected call of GetInitialSealer.
func (mr *MockCryptoSetupMockRecorder) GetInitialSealer() *CryptoSetupGetInitialSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialSealer", reflect.TypeOf((*MockCryptoSetup)(nil).GetInitialSealer))
	return &CryptoSetupGetInitialSealerCall{Call: call}
}

// CryptoSetupGetInitialSealerCall wrap *gomock.Call
type CryptoSetupGetInitialSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGetInitialSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *CryptoSetupGetInitialSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGetInitialSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *CryptoSetupGetInitialSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGetInitialSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *CryptoSetupGetInitialSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSessionTicket mocks base method.
func (m *MockCryptoSetup) GetSessionTicket() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionTicket")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionTicket indicates an expected call of GetSessionTicket.
func (mr *MockCryptoSetupMockRecorder) GetSessionTicket() *CryptoSetupGetSessionTicketCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionTicket", reflect.TypeOf((*MockCryptoSetup)(nil).GetSessionTicket))
	return &CryptoSetupGetSessionTicketCall{Call: call}
}

// CryptoSetupGetSessionTicketCall wrap *gomock.Call
type CryptoSetupGetSessionTicketCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupGetSessionTicketCall) Return(arg0 []byte, arg1 error) *CryptoSetupGetSessionTicketCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupGetSessionTicketCall) Do(f func() ([]byte, error)) *CryptoSetupGetSessionTicketCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupGetSessionTicketCall) DoAndReturn(f func() ([]byte, error)) *CryptoSetupGetSessionTicketCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandleMessage mocks base method.
func (m *MockCryptoSetup) HandleMessage(arg0 []byte, arg1 protocol.EncryptionLevel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockCryptoSetupMockRecorder) HandleMessage(arg0, arg1 any) *CryptoSetupHandleMessageCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockCryptoSetup)(nil).HandleMessage), arg0, arg1)
	return &CryptoSetupHandleMessageCall{Call: call}
}

// CryptoSetupHandleMessageCall wrap *gomock.Call
type CryptoSetupHandleMessageCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupHandleMessageCall) Return(arg0 error) *CryptoSetupHandleMessageCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupHandleMessageCall) Do(f func([]byte, protocol.EncryptionLevel) error) *CryptoSetupHandleMessageCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupHandleMessageCall) DoAndReturn(f func([]byte, protocol.EncryptionLevel) error) *CryptoSetupHandleMessageCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NextEvent mocks base method.
func (m *MockCryptoSetup) NextEvent() handshake.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextEvent")
	ret0, _ := ret[0].(handshake.Event)
	return ret0
}

// NextEvent indicates an expected call of NextEvent.
func (mr *MockCryptoSetupMockRecorder) NextEvent() *CryptoSetupNextEventCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextEvent", reflect.TypeOf((*MockCryptoSetup)(nil).NextEvent))
	return &CryptoSetupNextEventCall{Call: call}
}

// CryptoSetupNextEventCall wrap *gomock.Call
type CryptoSetupNextEventCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupNextEventCall) Return(arg0 handshake.Event) *CryptoSetupNextEventCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupNextEventCall) Do(f func() handshake.Event) *CryptoSetupNextEventCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupNextEventCall) DoAndReturn(f func() handshake.Event) *CryptoSetupNextEventCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetHandshakeConfirmed mocks base method.
func (m *MockCryptoSetup) SetHandshakeConfirmed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHandshakeConfirmed")
}

// SetHandshakeConfirmed indicates an expected call of SetHandshakeConfirmed.
func (mr *MockCryptoSetupMockRecorder) SetHandshakeConfirmed() *CryptoSetupSetHandshakeConfirmedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHandshakeConfirmed", reflect.TypeOf((*MockCryptoSetup)(nil).SetHandshakeConfirmed))
	return &CryptoSetupSetHandshakeConfirmedCall{Call: call}
}

// CryptoSetupSetHandshakeConfirmedCall wrap *gomock.Call
type CryptoSetupSetHandshakeConfirmedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupSetHandshakeConfirmedCall) Return() *CryptoSetupSetHandshakeConfirmedCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupSetHandshakeConfirmedCall) Do(f func()) *CryptoSetupSetHandshakeConfirmedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupSetHandshakeConfirmedCall) DoAndReturn(f func()) *CryptoSetupSetHandshakeConfirmedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetLargest1RTTAcked mocks base method.
func (m *MockCryptoSetup) SetLargest1RTTAcked(arg0 protocol.PacketNumber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLargest1RTTAcked", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLargest1RTTAcked indicates an expected call of SetLargest1RTTAcked.
func (mr *MockCryptoSetupMockRecorder) SetLargest1RTTAcked(arg0 any) *CryptoSetupSetLargest1RTTAckedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLargest1RTTAcked", reflect.TypeOf((*MockCryptoSetup)(nil).SetLargest1RTTAcked), arg0)
	return &CryptoSetupSetLargest1RTTAckedCall{Call: call}
}

// CryptoSetupSetLargest1RTTAckedCall wrap *gomock.Call
type CryptoSetupSetLargest1RTTAckedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupSetLargest1RTTAckedCall) Return(arg0 error) *CryptoSetupSetLargest1RTTAckedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupSetLargest1RTTAckedCall) Do(f func(protocol.PacketNumber) error) *CryptoSetupSetLargest1RTTAckedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupSetLargest1RTTAckedCall) DoAndReturn(f func(protocol.PacketNumber) error) *CryptoSetupSetLargest1RTTAckedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// StartHandshake mocks base method.
func (m *MockCryptoSetup) StartHandshake() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartHandshake")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartHandshake indicates an expected call of StartHandshake.
func (mr *MockCryptoSetupMockRecorder) StartHandshake() *CryptoSetupStartHandshakeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartHandshake", reflect.TypeOf((*MockCryptoSetup)(nil).StartHandshake))
	return &CryptoSetupStartHandshakeCall{Call: call}
}

// CryptoSetupStartHandshakeCall wrap *gomock.Call
type CryptoSetupStartHandshakeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoSetupStartHandshakeCall) Return(arg0 error) *CryptoSetupStartHandshakeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoSetupStartHandshakeCall) Do(f func() error) *CryptoSetupStartHandshakeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoSetupStartHandshakeCall) DoAndReturn(f func() error) *CryptoSetupStartHandshakeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
