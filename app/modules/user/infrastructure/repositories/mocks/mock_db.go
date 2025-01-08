// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/modules/user/infrastructure/repositories/interface.go
//
// Generated by this command:
//
//	mockgen -source=./app/modules/user/infrastructure/repositories/interface.go -destination=./app/modules/user/infrastructure/repositories/mocks/mock_db.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	usertypes "github.com/Black-And-White-Club/tcr-bot/app/modules/user/domain/types"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDB is a mock of UserDB interface.
type MockUserDB struct {
	ctrl     *gomock.Controller
	recorder *MockUserDBMockRecorder
	isgomock struct{}
}

// MockUserDBMockRecorder is the mock recorder for MockUserDB.
type MockUserDBMockRecorder struct {
	mock *MockUserDB
}

// NewMockUserDB creates a new mock instance.
func NewMockUserDB(ctrl *gomock.Controller) *MockUserDB {
	mock := &MockUserDB{ctrl: ctrl}
	mock.recorder = &MockUserDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDB) EXPECT() *MockUserDBMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserDB) CreateUser(ctx context.Context, user usertypes.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserDBMockRecorder) CreateUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserDB)(nil).CreateUser), ctx, user)
}

// GetUserByDiscordID mocks base method.
func (m *MockUserDB) GetUserByDiscordID(ctx context.Context, discordID usertypes.DiscordID) (usertypes.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByDiscordID", ctx, discordID)
	ret0, _ := ret[0].(usertypes.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByDiscordID indicates an expected call of GetUserByDiscordID.
func (mr *MockUserDBMockRecorder) GetUserByDiscordID(ctx, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByDiscordID", reflect.TypeOf((*MockUserDB)(nil).GetUserByDiscordID), ctx, discordID)
}

// GetUserRole mocks base method.
func (m *MockUserDB) GetUserRole(ctx context.Context, discordID usertypes.DiscordID) (usertypes.UserRoleEnum, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRole", ctx, discordID)
	ret0, _ := ret[0].(usertypes.UserRoleEnum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRole indicates an expected call of GetUserRole.
func (mr *MockUserDBMockRecorder) GetUserRole(ctx, discordID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRole", reflect.TypeOf((*MockUserDB)(nil).GetUserRole), ctx, discordID)
}

// UpdateUserRole mocks base method.
func (m *MockUserDB) UpdateUserRole(ctx context.Context, discordID usertypes.DiscordID, newRole usertypes.UserRoleEnum) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserRole", ctx, discordID, newRole)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserRole indicates an expected call of UpdateUserRole.
func (mr *MockUserDBMockRecorder) UpdateUserRole(ctx, discordID, newRole any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserRole", reflect.TypeOf((*MockUserDB)(nil).UpdateUserRole), ctx, discordID, newRole)
}
