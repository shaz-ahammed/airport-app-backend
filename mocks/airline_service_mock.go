// Code generated by MockGen. DO NOT EDIT.
// Source: airport-app-backend/services (interfaces: IAirlineRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "airport-app-backend/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAirlineRepository is a mock of IAirlineRepository interface.
type MockIAirlineRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAirlineRepositoryMockRecorder
}

// MockIAirlineRepositoryMockRecorder is the mock recorder for MockIAirlineRepository.
type MockIAirlineRepositoryMockRecorder struct {
	mock *MockIAirlineRepository
}

// NewMockIAirlineRepository creates a new mock instance.
func NewMockIAirlineRepository(ctrl *gomock.Controller) *MockIAirlineRepository {
	mock := &MockIAirlineRepository{ctrl: ctrl}
	mock.recorder = &MockIAirlineRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAirlineRepository) EXPECT() *MockIAirlineRepositoryMockRecorder {
	return m.recorder
}

// GetAirline mocks base method.
func (m *MockIAirlineRepository) GetAirline(arg0 int) ([]models.Airlines, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAirline", arg0)
	ret0, _ := ret[0].([]models.Airlines)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAirline indicates an expected call of GetAirline.
func (mr *MockIAirlineRepositoryMockRecorder) GetAirline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAirline", reflect.TypeOf((*MockIAirlineRepository)(nil).GetAirline), arg0)
}
