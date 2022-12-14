// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	types "play3/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetAllDishes mocks base method.
func (m *MockRepositoryInterface) GetAllDishes() []types.Dishes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDishes")
	ret0, _ := ret[0].([]types.Dishes)
	return ret0
}

// GetAllDishes indicates an expected call of GetAllDishes.
func (mr *MockRepositoryInterfaceMockRecorder) GetAllDishes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDishes", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAllDishes))
}
