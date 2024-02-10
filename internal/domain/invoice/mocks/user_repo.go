// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luchojuarez/call-sorter/internal/domain/invoice (interfaces: UserRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "github.com/luchojuarez/call-sorter/internal/infrastructure/user"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// GetByPhoneNumber mocks base method.
func (m *MockUserRepository) GetByPhoneNumber(arg0 context.Context, arg1 string) (user.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPhoneNumber", arg0, arg1)
	ret0, _ := ret[0].(user.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPhoneNumber indicates an expected call of GetByPhoneNumber.
func (mr *MockUserRepositoryMockRecorder) GetByPhoneNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPhoneNumber", reflect.TypeOf((*MockUserRepository)(nil).GetByPhoneNumber), arg0, arg1)
}
