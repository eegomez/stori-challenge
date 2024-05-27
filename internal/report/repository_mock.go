// Code generated by MockGen. DO NOT EDIT.
// Source: internal/report/repository.go

// Package report is a generated GoMock package.
package report

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// readCsvTransactionFile mocks base method.
func (m *MockRepository) readCsvTransactionFile(ctx context.Context) ([]Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "readCsvTransactionFile", ctx)
	ret0, _ := ret[0].([]Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// readCsvTransactionFile indicates an expected call of readCsvTransactionFile.
func (mr *MockRepositoryMockRecorder) readCsvTransactionFile(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "readCsvTransactionFile", reflect.TypeOf((*MockRepository)(nil).readCsvTransactionFile), ctx)
}