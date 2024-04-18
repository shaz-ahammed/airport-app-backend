// Code generated by MockGen. DO NOT EDIT.
// Source: airport-app-backend/services (interfaces: IGateRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "airport-app-backend/models"
	context "context"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIGateRepository is a mock of IGateRepository interface.
type MockIGateRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIGateRepositoryMockRecorder
}

// MockIGateRepositoryMockRecorder is the mock recorder for MockIGateRepository.
type MockIGateRepositoryMockRecorder struct {
	mock *MockIGateRepository
}

// NewMockIGateRepository creates a new mock instance.
func NewMockIGateRepository(ctrl *gomock.Controller) *MockIGateRepository {
	mock := &MockIGateRepository{ctrl: ctrl}
	mock.recorder = &MockIGateRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGateRepository) EXPECT() *MockIGateRepositoryMockRecorder {
	return m.recorder
}

// GetGateByID mocks base method.
func (m *MockIGateRepository) GetGateByID(arg0 context.Context, arg1 *gin.Context, arg2 string) (*models.Gate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGateByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Gate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGateByID indicates an expected call of GetGateByID.
func (mr *MockIGateRepositoryMockRecorder) GetGateByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGateByID", reflect.TypeOf((*MockIGateRepository)(nil).GetGateByID), arg0, arg1, arg2)
}

// GetGates mocks base method.
func (m *MockIGateRepository) GetGates(arg0, arg1 int, arg2 context.Context, arg3 *gin.Context) ([]models.Gate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGates", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]models.Gate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGates indicates an expected call of GetGates.
func (mr *MockIGateRepositoryMockRecorder) GetGates(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGates", reflect.TypeOf((*MockIGateRepository)(nil).GetGates), arg0, arg1, arg2, arg3)
}
