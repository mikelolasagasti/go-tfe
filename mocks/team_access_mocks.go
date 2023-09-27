// Code generated by MockGen. DO NOT EDIT.
// Source: team_access.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/hashicorp/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockTeamAccesses is a mock of TeamAccesses interface.
type MockTeamAccesses struct {
	ctrl     *gomock.Controller
	recorder *MockTeamAccessesMockRecorder
}

// MockTeamAccessesMockRecorder is the mock recorder for MockTeamAccesses.
type MockTeamAccessesMockRecorder struct {
	mock *MockTeamAccesses
}

// NewMockTeamAccesses creates a new mock instance.
func NewMockTeamAccesses(ctrl *gomock.Controller) *MockTeamAccesses {
	mock := &MockTeamAccesses{ctrl: ctrl}
	mock.recorder = &MockTeamAccessesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamAccesses) EXPECT() *MockTeamAccessesMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTeamAccesses) Add(ctx context.Context, options tfe.TeamAccessAddOptions) (*tfe.TeamAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, options)
	ret0, _ := ret[0].(*tfe.TeamAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTeamAccessesMockRecorder) Add(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTeamAccesses)(nil).Add), ctx, options)
}

// List mocks base method.
func (m *MockTeamAccesses) List(ctx context.Context, options *tfe.TeamAccessListOptions) (*tfe.TeamAccessList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.TeamAccessList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTeamAccessesMockRecorder) List(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTeamAccesses)(nil).List), ctx, options)
}

// Read mocks base method.
func (m *MockTeamAccesses) Read(ctx context.Context, teamAccessID string) (*tfe.TeamAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, teamAccessID)
	ret0, _ := ret[0].(*tfe.TeamAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTeamAccessesMockRecorder) Read(ctx, teamAccessID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTeamAccesses)(nil).Read), ctx, teamAccessID)
}

// Remove mocks base method.
func (m *MockTeamAccesses) Remove(ctx context.Context, teamAccessID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, teamAccessID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockTeamAccessesMockRecorder) Remove(ctx, teamAccessID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockTeamAccesses)(nil).Remove), ctx, teamAccessID)
}

// Update mocks base method.
func (m *MockTeamAccesses) Update(ctx context.Context, teamAccessID string, options tfe.TeamAccessUpdateOptions) (*tfe.TeamAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, teamAccessID, options)
	ret0, _ := ret[0].(*tfe.TeamAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTeamAccessesMockRecorder) Update(ctx, teamAccessID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTeamAccesses)(nil).Update), ctx, teamAccessID, options)
}
