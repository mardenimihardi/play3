// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	types "play3/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecaseInterface is a mock of UsecaseInterface interface.
type MockUsecaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseInterfaceMockRecorder
}

// MockUsecaseInterfaceMockRecorder is the mock recorder for MockUsecaseInterface.
type MockUsecaseInterfaceMockRecorder struct {
	mock *MockUsecaseInterface
}

// NewMockUsecaseInterface creates a new mock instance.
func NewMockUsecaseInterface(ctrl *gomock.Controller) *MockUsecaseInterface {
	mock := &MockUsecaseInterface{ctrl: ctrl}
	mock.recorder = &MockUsecaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseInterface) EXPECT() *MockUsecaseInterfaceMockRecorder {
	return m.recorder
}

// GetAllDishes mocks base method.
func (m *MockUsecaseInterface) GetAllDishes() []types.Dishes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDishes")
	ret0, _ := ret[0].([]types.Dishes)
	return ret0
}

// GetAllDishes indicates an expected call of GetAllDishes.
func (mr *MockUsecaseInterfaceMockRecorder) GetAllDishes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDishes", reflect.TypeOf((*MockUsecaseInterface)(nil).GetAllDishes))
}