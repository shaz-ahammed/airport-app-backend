// Code generated by MockGen. DO NOT EDIT.
// Source: airport-app-backend/services (interfaces: IHealthRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "airport-app-backend/models"
	context "context"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIHealthRepository is a mock of IHealthRepository interface.
type MockIHealthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIHealthRepositoryMockRecorder
}

// MockIHealthRepositoryMockRecorder is the mock recorder for MockIHealthRepository.
type MockIHealthRepositoryMockRecorder struct {
	mock *MockIHealthRepository
}

// NewMockIHealthRepository creates a new mock instance.
func NewMockIHealthRepository(ctrl *gomock.Controller) *MockIHealthRepository {
	mock := &MockIHealthRepository{ctrl: ctrl}
	mock.recorder = &MockIHealthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHealthRepository) EXPECT() *MockIHealthRepositoryMockRecorder {
	return m.recorder
}

// GetAppHealth mocks base method.
func (m *MockIHealthRepository) GetAppHealth(arg0 context.Context, arg1 *gin.Context) models.AppHealth {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppHealth", arg0, arg1)
	ret0, _ := ret[0].(models.AppHealth)
	return ret0
}

// GetAppHealth indicates an expected call of GetAppHealth.
func (mr *MockIHealthRepositoryMockRecorder) GetAppHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppHealth", reflect.TypeOf((*MockIHealthRepository)(nil).GetAppHealth), arg0, arg1)
}
