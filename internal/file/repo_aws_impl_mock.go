package file

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockS3ClientInterface is a mock of S3ClientInterface interface
type MockS3ClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientInterfaceMockRecorder
}

// MockS3ClientInterfaceMockRecorder is the mock recorder for MockS3ClientInterface
type MockS3ClientInterfaceMockRecorder struct {
	mock *MockS3ClientInterface
}

// NewMockS3ClientInterface creates a new mock instance
func NewMockS3ClientInterface(ctrl *gomock.Controller) *MockS3ClientInterface {
	mock := &MockS3ClientInterface{ctrl: ctrl}
	mock.recorder = &MockS3ClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3ClientInterface) EXPECT() *MockS3ClientInterfaceMockRecorder {
	return m.recorder
}

// GetObject mocks base method
func (m *MockS3ClientInterface) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockS3ClientInterfaceMockRecorder) GetObject(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3ClientInterface)(nil).GetObject), varargs...)
}
