// Code generated by MockGen. DO NOT EDIT.
// Source: ldap_connection_factory.go
package authentication

import (
	tls "crypto/tls"
	reflect "reflect"

	ldap "github.com/go-ldap/ldap/v3"
	gomock "github.com/golang/mock/gomock"
)

// MockLDAPConnection is a mock of LDAPConnection interface
type MockLDAPConnection struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConnectionMockRecorder
}

// MockLDAPConnectionMockRecorder is the mock recorder for MockLDAPConnection
type MockLDAPConnectionMockRecorder struct {
	mock *MockLDAPConnection
}

// NewMockLDAPConnection creates a new mock instance
func NewMockLDAPConnection(ctrl *gomock.Controller) *MockLDAPConnection {
	mock := &MockLDAPConnection{ctrl: ctrl}
	mock.recorder = &MockLDAPConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLDAPConnection) EXPECT() *MockLDAPConnectionMockRecorder {
	return m.recorder
}

// Bind mocks base method
func (m *MockLDAPConnection) Bind(username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind
func (mr *MockLDAPConnectionMockRecorder) Bind(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockLDAPConnection)(nil).Bind), username, password)
}

// Close mocks base method
func (m *MockLDAPConnection) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockLDAPConnectionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLDAPConnection)(nil).Close))
}

// Search mocks base method
func (m *MockLDAPConnection) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", searchRequest)
	ret0, _ := ret[0].(*ldap.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockLDAPConnectionMockRecorder) Search(searchRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockLDAPConnection)(nil).Search), searchRequest)
}

// Modify mocks base method
func (m *MockLDAPConnection) Modify(modifyRequest *ldap.ModifyRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modify", modifyRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Modify indicates an expected call of Modify
func (mr *MockLDAPConnectionMockRecorder) Modify(modifyRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockLDAPConnection)(nil).Modify), modifyRequest)
}

// StartTLS mocks base method
func (m *MockLDAPConnection) StartTLS(config *tls.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTLS", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartTLS indicates an expected call of StartTLS
func (mr *MockLDAPConnectionMockRecorder) StartTLS(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTLS", reflect.TypeOf((*MockLDAPConnection)(nil).StartTLS), config)
}

// MockLDAPConnectionFactory is a mock of LDAPConnectionFactory interface
type MockLDAPConnectionFactory struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPConnectionFactoryMockRecorder
}

// MockLDAPConnectionFactoryMockRecorder is the mock recorder for MockLDAPConnectionFactory
type MockLDAPConnectionFactoryMockRecorder struct {
	mock *MockLDAPConnectionFactory
}

// NewMockLDAPConnectionFactory creates a new mock instance
func NewMockLDAPConnectionFactory(ctrl *gomock.Controller) *MockLDAPConnectionFactory {
	mock := &MockLDAPConnectionFactory{ctrl: ctrl}
	mock.recorder = &MockLDAPConnectionFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLDAPConnectionFactory) EXPECT() *MockLDAPConnectionFactoryMockRecorder {
	return m.recorder
}

// DialURL mocks base method
func (m *MockLDAPConnectionFactory) DialURL(addr string, opts ldap.DialOpt) (LDAPConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialURL", addr, opts)
	ret0, _ := ret[0].(LDAPConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DialURL indicates an expected call of DialURL
func (mr *MockLDAPConnectionFactoryMockRecorder) DialURL(addr, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialURL", reflect.TypeOf((*MockLDAPConnectionFactory)(nil).DialURL), addr, opts)
}
