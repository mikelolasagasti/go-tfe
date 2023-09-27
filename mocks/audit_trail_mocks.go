// Code generated by MockGen. DO NOT EDIT.
// Source: audit_trail.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/hashicorp/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockAuditTrails is a mock of AuditTrails interface.
type MockAuditTrails struct {
	ctrl     *gomock.Controller
	recorder *MockAuditTrailsMockRecorder
}

// MockAuditTrailsMockRecorder is the mock recorder for MockAuditTrails.
type MockAuditTrailsMockRecorder struct {
	mock *MockAuditTrails
}

// NewMockAuditTrails creates a new mock instance.
func NewMockAuditTrails(ctrl *gomock.Controller) *MockAuditTrails {
	mock := &MockAuditTrails{ctrl: ctrl}
	mock.recorder = &MockAuditTrailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditTrails) EXPECT() *MockAuditTrailsMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockAuditTrails) List(ctx context.Context, options *tfe.AuditTrailListOptions) (*tfe.AuditTrailList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.AuditTrailList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAuditTrailsMockRecorder) List(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuditTrails)(nil).List), ctx, options)
}
