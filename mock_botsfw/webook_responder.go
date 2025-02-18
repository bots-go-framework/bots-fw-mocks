// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bots-go-framework/bots-fw/botsfw (interfaces: WebhookResponder)
//
// Generated by this command:
//
//	mockgen github.com/bots-go-framework/bots-fw/botsfw WebhookResponder
//

// Package mock_botsfw is a generated GoMock package.
package mock_botsfw

import (
	context "context"
	reflect "reflect"

	botsfw "github.com/bots-go-framework/bots-fw/botsfw"
	gomock "go.uber.org/mock/gomock"
)

// MockWebhookResponder is a mock of WebhookResponder interface.
type MockWebhookResponder struct {
	ctrl     *gomock.Controller
	recorder *MockWebhookResponderMockRecorder
	isgomock struct{}
}

// MockWebhookResponderMockRecorder is the mock recorder for MockWebhookResponder.
type MockWebhookResponderMockRecorder struct {
	mock *MockWebhookResponder
}

// NewMockWebhookResponder creates a new mock instance.
func NewMockWebhookResponder(ctrl *gomock.Controller) *MockWebhookResponder {
	mock := &MockWebhookResponder{ctrl: ctrl}
	mock.recorder = &MockWebhookResponderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebhookResponder) EXPECT() *MockWebhookResponderMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MockWebhookResponder) DeleteMessage(c context.Context, messageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", c, messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockWebhookResponderMockRecorder) DeleteMessage(c, messageID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockWebhookResponder)(nil).DeleteMessage), c, messageID)
}

// SendMessage mocks base method.
func (m_2 *MockWebhookResponder) SendMessage(c context.Context, m botsfw.MessageFromBot, channel botsfw.BotAPISendMessageChannel) (botsfw.OnMessageSentResponse, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMessage", c, m, channel)
	ret0, _ := ret[0].(botsfw.OnMessageSentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockWebhookResponderMockRecorder) SendMessage(c, m, channel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockWebhookResponder)(nil).SendMessage), c, m, channel)
}
