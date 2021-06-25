// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/internal/printer (interfaces: ObjectInterface)

// Package fake is a generated GoMock package.
package fake

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"

	printer "github.com/vmware-tanzu/octant/internal/printer"
	action "github.com/vmware-tanzu/octant/pkg/action"
	component "github.com/vmware-tanzu/octant/pkg/view/component"
)

// MockObjectInterface is a mock of ObjectInterface interface.
type MockObjectInterface struct {
	ctrl     *gomock.Controller
	recorder *MockObjectInterfaceMockRecorder
}

// MockObjectInterfaceMockRecorder is the mock recorder for MockObjectInterface.
type MockObjectInterfaceMockRecorder struct {
	mock *MockObjectInterface
}

// NewMockObjectInterface creates a new mock instance.
func NewMockObjectInterface(ctrl *gomock.Controller) *MockObjectInterface {
	mock := &MockObjectInterface{ctrl: ctrl}
	mock.recorder = &MockObjectInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectInterface) EXPECT() *MockObjectInterfaceMockRecorder {
	return m.recorder
}

// AddButton mocks base method.
func (m *MockObjectInterface) AddButton(arg0 string, arg1 action.Payload, arg2 ...component.ButtonOption) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddButton", varargs...)
}

// AddButton indicates an expected call of AddButton.
func (mr *MockObjectInterfaceMockRecorder) AddButton(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddButton", reflect.TypeOf((*MockObjectInterface)(nil).AddButton), varargs...)
}

// EnableEvents mocks base method.
func (m *MockObjectInterface) EnableEvents() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnableEvents")
}

// EnableEvents indicates an expected call of EnableEvents.
func (mr *MockObjectInterfaceMockRecorder) EnableEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableEvents", reflect.TypeOf((*MockObjectInterface)(nil).EnableEvents))
}

// EnableJobTemplate mocks base method.
func (m *MockObjectInterface) EnableJobTemplate(arg0 v1beta1.JobTemplateSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnableJobTemplate", arg0)
}

// EnableJobTemplate indicates an expected call of EnableJobTemplate.
func (mr *MockObjectInterfaceMockRecorder) EnableJobTemplate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableJobTemplate", reflect.TypeOf((*MockObjectInterface)(nil).EnableJobTemplate), arg0)
}

// EnablePodTemplate mocks base method.
func (m *MockObjectInterface) EnablePodTemplate(arg0 v1.PodTemplateSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnablePodTemplate", arg0)
}

// EnablePodTemplate indicates an expected call of EnablePodTemplate.
func (mr *MockObjectInterfaceMockRecorder) EnablePodTemplate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePodTemplate", reflect.TypeOf((*MockObjectInterface)(nil).EnablePodTemplate), arg0)
}

// RegisterConfig mocks base method.
func (m *MockObjectInterface) RegisterConfig(arg0 *component.Summary) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterConfig", arg0)
}

// RegisterConfig indicates an expected call of RegisterConfig.
func (mr *MockObjectInterfaceMockRecorder) RegisterConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterConfig", reflect.TypeOf((*MockObjectInterface)(nil).RegisterConfig), arg0)
}

// RegisterItems mocks base method.
func (m *MockObjectInterface) RegisterItems(arg0 ...printer.ItemDescriptor) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RegisterItems", varargs...)
}

// RegisterItems indicates an expected call of RegisterItems.
func (mr *MockObjectInterfaceMockRecorder) RegisterItems(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterItems", reflect.TypeOf((*MockObjectInterface)(nil).RegisterItems), arg0...)
}

// RegisterSummary mocks base method.
func (m *MockObjectInterface) RegisterSummary(arg0 *component.Summary) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterSummary", arg0)
}

// RegisterSummary indicates an expected call of RegisterSummary.
func (mr *MockObjectInterfaceMockRecorder) RegisterSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterSummary", reflect.TypeOf((*MockObjectInterface)(nil).RegisterSummary), arg0)
}
