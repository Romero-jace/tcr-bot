// Code generated by MockGen. DO NOT EDIT.
// Source: ./infrastructure/subscribers/interface.go
//
// Generated by this command:
//
//	mockgen -source=./infrastructure/subscribers/interface.go -destination=./infrastructure/subscribers/mocks/mock_subscribers.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSubscribers is a mock of Subscribers interface.
type MockSubscribers struct {
	ctrl     *gomock.Controller
	recorder *MockSubscribersMockRecorder
	isgomock struct{}
}

// MockSubscribersMockRecorder is the mock recorder for MockSubscribers.
type MockSubscribersMockRecorder struct {
	mock *MockSubscribers
}

// NewMockSubscribers creates a new mock instance.
func NewMockSubscribers(ctrl *gomock.Controller) *MockSubscribers {
	mock := &MockSubscribers{ctrl: ctrl}
	mock.recorder = &MockSubscribersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscribers) EXPECT() *MockSubscribersMockRecorder {
	return m.recorder
}

// SubscribeToLeaderboardEvents mocks base method.
func (m *MockSubscribers) SubscribeToLeaderboardEvents(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToLeaderboardEvents", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeToLeaderboardEvents indicates an expected call of SubscribeToLeaderboardEvents.
func (mr *MockSubscribersMockRecorder) SubscribeToLeaderboardEvents(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToLeaderboardEvents", reflect.TypeOf((*MockSubscribers)(nil).SubscribeToLeaderboardEvents), ctx)
}