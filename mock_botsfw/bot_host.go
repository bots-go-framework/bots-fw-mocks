// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bots-go-framework/bots-fw/botsfw (interfaces: BotHost)
//
// Generated by this command:
//
//	mockgen github.com/bots-go-framework/bots-fw/botsfw BotHost
//

// Package mock_botsfw is a generated GoMock package.
package mock_botsfw

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBotHost is a mock of BotHost interface.
type MockBotHost struct {
	ctrl     *gomock.Controller
	recorder *MockBotHostMockRecorder
	isgomock struct{}
}

// MockBotHostMockRecorder is the mock recorder for MockBotHost.
type MockBotHostMockRecorder struct {
	mock *MockBotHost
}

// NewMockBotHost creates a new mock instance.
func NewMockBotHost(ctrl *gomock.Controller) *MockBotHost {
	mock := &MockBotHost{ctrl: ctrl}
	mock.recorder = &MockBotHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBotHost) EXPECT() *MockBotHostMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockBotHost) Context(r *http.Request) context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context", r)
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockBotHostMockRecorder) Context(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockBotHost)(nil).Context), r)
}

// GetHTTPClient mocks base method.
func (m *MockBotHost) GetHTTPClient(c context.Context) *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPClient", c)
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// GetHTTPClient indicates an expected call of GetHTTPClient.
func (mr *MockBotHostMockRecorder) GetHTTPClient(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPClient", reflect.TypeOf((*MockBotHost)(nil).GetHTTPClient), c)
}
