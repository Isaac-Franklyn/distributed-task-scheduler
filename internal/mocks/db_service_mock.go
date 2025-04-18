// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/ports/db_port.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/Isaac-Franklyn/distributed-task-scheduler/internal/domain/models"
	gomock "github.com/golang/mock/gomock"
)

// MockDbService is a mock of DbService interface.
type MockDbService struct {
	ctrl     *gomock.Controller
	recorder *MockDbServiceMockRecorder
}

// MockDbServiceMockRecorder is the mock recorder for MockDbService.
type MockDbServiceMockRecorder struct {
	mock *MockDbService
}

// NewMockDbService creates a new mock instance.
func NewMockDbService(ctrl *gomock.Controller) *MockDbService {
	mock := &MockDbService{ctrl: ctrl}
	mock.recorder = &MockDbServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDbService) EXPECT() *MockDbServiceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDbService) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDbServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDbService)(nil).Close))
}

// CreateTaskTable mocks base method.
func (m *MockDbService) CreateTaskTable() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateTaskTable")
}

// CreateTaskTable indicates an expected call of CreateTaskTable.
func (mr *MockDbServiceMockRecorder) CreateTaskTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTaskTable", reflect.TypeOf((*MockDbService)(nil).CreateTaskTable))
}

// SaveTaskToDb mocks base method.
func (m *MockDbService) SaveTaskToDb(task *models.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTaskToDb", task)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTaskToDb indicates an expected call of SaveTaskToDb.
func (mr *MockDbServiceMockRecorder) SaveTaskToDb(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTaskToDb", reflect.TypeOf((*MockDbService)(nil).SaveTaskToDb), task)
}
