// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tus/tusd/pkg/gcsstore (interfaces: GCSReader,GCSAPI)

// Package gcsstore_test is a generated GoMock package.
package gcsstore_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	gcsstore "github.com/tus/tusd/pkg/gcsstore"
	io "io"
	reflect "reflect"
)

// MockGCSReader is a mock of GCSReader interface
type MockGCSReader struct {
	ctrl     *gomock.Controller
	recorder *MockGCSReaderMockRecorder
}

// MockGCSReaderMockRecorder is the mock recorder for MockGCSReader
type MockGCSReaderMockRecorder struct {
	mock *MockGCSReader
}

// NewMockGCSReader creates a new mock instance
func NewMockGCSReader(ctrl *gomock.Controller) *MockGCSReader {
	mock := &MockGCSReader{ctrl: ctrl}
	mock.recorder = &MockGCSReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGCSReader) EXPECT() *MockGCSReaderMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockGCSReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockGCSReaderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGCSReader)(nil).Close))
}

// ContentType mocks base method
func (m *MockGCSReader) ContentType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentType indicates an expected call of ContentType
func (mr *MockGCSReaderMockRecorder) ContentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentType", reflect.TypeOf((*MockGCSReader)(nil).ContentType))
}

// Read mocks base method
func (m *MockGCSReader) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockGCSReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockGCSReader)(nil).Read), arg0)
}

// Remain mocks base method
func (m *MockGCSReader) Remain() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remain")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Remain indicates an expected call of Remain
func (mr *MockGCSReaderMockRecorder) Remain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remain", reflect.TypeOf((*MockGCSReader)(nil).Remain))
}

// Size mocks base method
func (m *MockGCSReader) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockGCSReaderMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockGCSReader)(nil).Size))
}

// MockGCSAPI is a mock of GCSAPI interface
type MockGCSAPI struct {
	ctrl     *gomock.Controller
	recorder *MockGCSAPIMockRecorder
}

// MockGCSAPIMockRecorder is the mock recorder for MockGCSAPI
type MockGCSAPIMockRecorder struct {
	mock *MockGCSAPI
}

// NewMockGCSAPI creates a new mock instance
func NewMockGCSAPI(ctrl *gomock.Controller) *MockGCSAPI {
	mock := &MockGCSAPI{ctrl: ctrl}
	mock.recorder = &MockGCSAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGCSAPI) EXPECT() *MockGCSAPIMockRecorder {
	return m.recorder
}

// ComposeObjects mocks base method
func (m *MockGCSAPI) ComposeObjects(arg0 context.Context, arg1 gcsstore.GCSComposeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComposeObjects", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComposeObjects indicates an expected call of ComposeObjects
func (mr *MockGCSAPIMockRecorder) ComposeObjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComposeObjects", reflect.TypeOf((*MockGCSAPI)(nil).ComposeObjects), arg0, arg1)
}

// DeleteObject mocks base method
func (m *MockGCSAPI) DeleteObject(arg0 context.Context, arg1 gcsstore.GCSObjectParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockGCSAPIMockRecorder) DeleteObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockGCSAPI)(nil).DeleteObject), arg0, arg1)
}

// DeleteObjectsWithFilter mocks base method
func (m *MockGCSAPI) DeleteObjectsWithFilter(arg0 context.Context, arg1 gcsstore.GCSFilterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectsWithFilter", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjectsWithFilter indicates an expected call of DeleteObjectsWithFilter
func (mr *MockGCSAPIMockRecorder) DeleteObjectsWithFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectsWithFilter", reflect.TypeOf((*MockGCSAPI)(nil).DeleteObjectsWithFilter), arg0, arg1)
}

// FilterObjects mocks base method
func (m *MockGCSAPI) FilterObjects(arg0 context.Context, arg1 gcsstore.GCSFilterParams) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterObjects", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterObjects indicates an expected call of FilterObjects
func (mr *MockGCSAPIMockRecorder) FilterObjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterObjects", reflect.TypeOf((*MockGCSAPI)(nil).FilterObjects), arg0, arg1)
}

// GetObjectSize mocks base method
func (m *MockGCSAPI) GetObjectSize(arg0 context.Context, arg1 gcsstore.GCSObjectParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectSize", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectSize indicates an expected call of GetObjectSize
func (mr *MockGCSAPIMockRecorder) GetObjectSize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectSize", reflect.TypeOf((*MockGCSAPI)(nil).GetObjectSize), arg0, arg1)
}

// ReadObject mocks base method
func (m *MockGCSAPI) ReadObject(arg0 context.Context, arg1 gcsstore.GCSObjectParams) (gcsstore.GCSReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadObject", arg0, arg1)
	ret0, _ := ret[0].(gcsstore.GCSReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadObject indicates an expected call of ReadObject
func (mr *MockGCSAPIMockRecorder) ReadObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadObject", reflect.TypeOf((*MockGCSAPI)(nil).ReadObject), arg0, arg1)
}

// SetObjectMetadata mocks base method
func (m *MockGCSAPI) SetObjectMetadata(arg0 context.Context, arg1 gcsstore.GCSObjectParams, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetObjectMetadata", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetObjectMetadata indicates an expected call of SetObjectMetadata
func (mr *MockGCSAPIMockRecorder) SetObjectMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetObjectMetadata", reflect.TypeOf((*MockGCSAPI)(nil).SetObjectMetadata), arg0, arg1, arg2)
}

// WriteObject mocks base method
func (m *MockGCSAPI) WriteObject(arg0 context.Context, arg1 gcsstore.GCSObjectParams, arg2 io.Reader) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteObject indicates an expected call of WriteObject
func (mr *MockGCSAPIMockRecorder) WriteObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteObject", reflect.TypeOf((*MockGCSAPI)(nil).WriteObject), arg0, arg1, arg2)
}