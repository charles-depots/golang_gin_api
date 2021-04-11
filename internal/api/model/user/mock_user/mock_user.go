// Package mock_pet is a generated GoMock package.
package mock_user

import (
	"github.com/golang/mock/gomock"
	user "golang-gin-api/internal/api/model/user/dbinit"
	"reflect"
)

// MockUserDbInterface is a mock of PetDbInterface interface
type MockUserDbInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPetDbInterfaceMockRecorder
}

// MockPetDbInterfaceMockRecorder is the mock recorder for MockUserDbInterface
type MockPetDbInterfaceMockRecorder struct {
	mock *MockUserDbInterface
}

// NewMockPetDbInterface creates a new mock instance
func NewMockPetDbInterface(ctrl *gomock.Controller) *MockUserDbInterface {
	mock := &MockUserDbInterface{ctrl: ctrl}
	mock.recorder = &MockPetDbInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDbInterface) EXPECT() *MockPetDbInterfaceMockRecorder {
	return m.recorder
}

// MockOwnerDbInterface is a mock of OwnerDbInterface interface
type MockOwnerDbInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOwnerDbInterfaceMockRecorder
}

// MockOwnerDbInterfaceMockRecorder is the mock recorder for MockOwnerDbInterface
type MockOwnerDbInterfaceMockRecorder struct {
	mock *MockOwnerDbInterface
}

// NewMockOwnerDbInterface creates a new mock instance
func NewMockOwnerDbInterface(ctrl *gomock.Controller) *MockOwnerDbInterface {
	mock := &MockOwnerDbInterface{ctrl: ctrl}
	mock.recorder = &MockOwnerDbInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOwnerDbInterface) EXPECT() *MockOwnerDbInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockOwnerDbInterface) Create(arg0 *user.User) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOwnerDbInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOwnerDbInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockOwnerDbInterface) Delete(arg0 *user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockOwnerDbInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOwnerDbInterface)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockOwnerDbInterface) Get(arg0 string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockOwnerDbInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOwnerDbInterface)(nil).Get), arg0)
}

// List mocks base method
func (m *MockOwnerDbInterface) List(arg0 *user.User, arg1, arg2 int) ([]*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockOwnerDbInterfaceMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOwnerDbInterface)(nil).List), arg0, arg1, arg2)
}

// Update mocks base method
func (m *MockOwnerDbInterface) Update(arg0 *user.User) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockOwnerDbInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOwnerDbInterface)(nil).Update), arg0)
}
