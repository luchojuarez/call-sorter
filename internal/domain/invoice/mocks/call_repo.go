// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luchojuarez/call-sorter/internal/domain/invoice (interfaces: CallRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	callservice "github.com/luchojuarez/call-sorter/internal/infrastructure/callservice"
)

// MockCallRepository is a mock of CallRepository interface.
type MockCallRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCallRepositoryMockRecorder
}

// MockCallRepositoryMockRecorder is the mock recorder for MockCallRepository.
type MockCallRepositoryMockRecorder struct {
	mock *MockCallRepository
}

// NewMockCallRepository creates a new mock instance.
func NewMockCallRepository(ctrl *gomock.Controller) *MockCallRepository {
	mock := &MockCallRepository{ctrl: ctrl}
	mock.recorder = &MockCallRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCallRepository) EXPECT() *MockCallRepositoryMockRecorder {
	return m.recorder
}

// FindByPhoneAndMonthAndYear mocks base method.
func (m *MockCallRepository) FindByPhoneAndMonthAndYear(arg0 context.Context, arg1 string, arg2 time.Month, arg3 int) ([]callservice.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPhoneAndMonthAndYear", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]callservice.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPhoneAndMonthAndYear indicates an expected call of FindByPhoneAndMonthAndYear.
func (mr *MockCallRepositoryMockRecorder) FindByPhoneAndMonthAndYear(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPhoneAndMonthAndYear", reflect.TypeOf((*MockCallRepository)(nil).FindByPhoneAndMonthAndYear), arg0, arg1, arg2, arg3)
}
