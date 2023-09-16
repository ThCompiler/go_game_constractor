// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene (interfaces: MessageMatcher)

// Package mock_scene is a generated GoMock package.
package mock_scene

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessageMatcher is a mock of MessageMatcher interface.
type MockMessageMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMatcherMockRecorder
}

// MockMessageMatcherMockRecorder is the mock recorder for MockMessageMatcher.
type MockMessageMatcherMockRecorder struct {
	mock *MockMessageMatcher
}

// NewMockMessageMatcher creates a new mock instance.
func NewMockMessageMatcher(ctrl *gomock.Controller) *MockMessageMatcher {
	mock := &MockMessageMatcher{ctrl: ctrl}
	mock.recorder = &MockMessageMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageMatcher) EXPECT() *MockMessageMatcherMockRecorder {
	return m.recorder
}

// GetMatchedName mocks base method.
func (m *MockMessageMatcher) GetMatchedName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchedName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMatchedName indicates an expected call of GetMatchedName.
func (mr *MockMessageMatcherMockRecorder) GetMatchedName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchedName", reflect.TypeOf((*MockMessageMatcher)(nil).GetMatchedName))
}

// Match mocks base method.
func (m *MockMessageMatcher) Match(arg0 string) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// Match indicates an expected call of Match.
func (mr *MockMessageMatcherMockRecorder) Match(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockMessageMatcher)(nil).Match), arg0)
}