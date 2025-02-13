// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luchojuarez/call-sorter/internal/domain/invoice (interfaces: InvoiceRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	invoice "github.com/luchojuarez/call-sorter/internal/domain/invoice"
)

// MockInvoiceRepository is a mock of InvoiceRepository interface.
type MockInvoiceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInvoiceRepositoryMockRecorder
}

// MockInvoiceRepositoryMockRecorder is the mock recorder for MockInvoiceRepository.
type MockInvoiceRepositoryMockRecorder struct {
	mock *MockInvoiceRepository
}

// NewMockInvoiceRepository creates a new mock instance.
func NewMockInvoiceRepository(ctrl *gomock.Controller) *MockInvoiceRepository {
	mock := &MockInvoiceRepository{ctrl: ctrl}
	mock.recorder = &MockInvoiceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvoiceRepository) EXPECT() *MockInvoiceRepositoryMockRecorder {
	return m.recorder
}

// GetByPhoneAndMonth mocks base method.
func (m *MockInvoiceRepository) GetByPhoneAndMonth(arg0 context.Context, arg1 string, arg2 time.Month) (*invoice.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPhoneAndMonth", arg0, arg1, arg2)
	ret0, _ := ret[0].(*invoice.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPhoneAndMonth indicates an expected call of GetByPhoneAndMonth.
func (mr *MockInvoiceRepositoryMockRecorder) GetByPhoneAndMonth(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPhoneAndMonth", reflect.TypeOf((*MockInvoiceRepository)(nil).GetByPhoneAndMonth), arg0, arg1, arg2)
}

// Save mocks base method.
func (m *MockInvoiceRepository) Save(arg0 context.Context, arg1 invoice.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockInvoiceRepositoryMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockInvoiceRepository)(nil).Save), arg0, arg1)
}
