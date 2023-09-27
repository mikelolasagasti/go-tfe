// Code generated by MockGen. DO NOT EDIT.
// Source: apply.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	tfe "github.com/hashicorp/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockApplies is a mock of Applies interface.
type MockApplies struct {
	ctrl     *gomock.Controller
	recorder *MockAppliesMockRecorder
}

// MockAppliesMockRecorder is the mock recorder for MockApplies.
type MockAppliesMockRecorder struct {
	mock *MockApplies
}

// NewMockApplies creates a new mock instance.
func NewMockApplies(ctrl *gomock.Controller) *MockApplies {
	mock := &MockApplies{ctrl: ctrl}
	mock.recorder = &MockAppliesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplies) EXPECT() *MockAppliesMockRecorder {
	return m.recorder
}

// Logs mocks base method.
func (m *MockApplies) Logs(ctx context.Context, applyID string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", ctx, applyID)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockAppliesMockRecorder) Logs(ctx, applyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockApplies)(nil).Logs), ctx, applyID)
}

// Read mocks base method.
func (m *MockApplies) Read(ctx context.Context, applyID string) (*tfe.Apply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, applyID)
	ret0, _ := ret[0].(*tfe.Apply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockAppliesMockRecorder) Read(ctx, applyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockApplies)(nil).Read), ctx, applyID)
}
