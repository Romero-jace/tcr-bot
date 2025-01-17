// Code generated by MockGen. DO NOT EDIT.
// Source: app/modules/user/application/interface.go
//
// Generated by this command:
//
//	mockgen -source=app/modules/user/application/interface.go -destination=app/modules/user/application/mocks/mock_service.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	events "github.com/Black-And-White-Club/tcr-bot/app/modules/user/domain/events"
	usertypes "github.com/Black-And-White-Club/tcr-bot/app/modules/user/domain/types"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockService) GetUser(ctx context.Context, discordID usertypes.DiscordID) (usertypes.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, discordID)
	ret0, _ := ret[0].(usertypes.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(ctx, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), ctx, discordID)
}

// GetUserRole mocks base method.
func (m *MockService) GetUserRole(ctx context.Context, discordID usertypes.DiscordID) (usertypes.UserRoleEnum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", ctx, discordID)
	ret0, _ := ret[0].(usertypes.UserRoleEnum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockServiceMockRecorder) GetUserRole(ctx, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockService)(nil).GetUserRole), ctx, discordID)
}

// OnUserRoleUpdateRequest mocks base method.
func (m *MockService) OnUserRoleUpdateRequest(ctx context.Context, req events.UserRoleUpdateRequestPayload) (*events.UserRoleUpdateResponsePayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnUserRoleUpdateRequest", ctx, req)
	ret0, _ := ret[0].(*events.UserRoleUpdateResponsePayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnUserRoleUpdateRequest indicates an expected call of OnUserRoleUpdateRequest.
func (mr *MockServiceMockRecorder) OnUserRoleUpdateRequest(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUserRoleUpdateRequest", reflect.TypeOf((*MockService)(nil).OnUserRoleUpdateRequest), ctx, req)
}

// OnUserSignupRequest mocks base method.
func (m *MockService) OnUserSignupRequest(ctx context.Context, req events.UserSignupRequestPayload) (*events.UserSignupResponsePayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnUserSignupRequest", ctx, req)
	ret0, _ := ret[0].(*events.UserSignupResponsePayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnUserSignupRequest indicates an expected call of OnUserSignupRequest.
func (mr *MockServiceMockRecorder) OnUserSignupRequest(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUserSignupRequest", reflect.TypeOf((*MockService)(nil).OnUserSignupRequest), ctx, req)
}
