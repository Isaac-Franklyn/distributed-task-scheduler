// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/ports/api_port.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAPIService is a mock of APIService interface.
type MockAPIService struct {
	ctrl     *gomock.Controller
	recorder *MockAPIServiceMockRecorder
}

// MockAPIServiceMockRecorder is the mock recorder for MockAPIService.
type MockAPIServiceMockRecorder struct {
	mock *MockAPIService
}

// NewMockAPIService creates a new mock instance.
func NewMockAPIService(ctrl *gomock.Controller) *MockAPIService {
	mock := &MockAPIService{ctrl: ctrl}
	mock.recorder = &MockAPIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIService) EXPECT() *MockAPIServiceMockRecorder {
	return m.recorder
}

// Validate mocks base method.
func (m *MockAPIService) Validate(task *models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockAPIServiceMockRecorder) Validate(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockAPIService)(nil).Validate), task)
}
