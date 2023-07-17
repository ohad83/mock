// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/mock/gomock (interfaces: Matcher)
//
// Generated by this command:
//
//	mockgen -destination internal/mock_gomock/mock_matcher.go go.uber.org/mock/gomock Matcher
//
// Package mock_gomock is a generated GoMock package.
package mock_gomock

import (
	reflect "reflect"

	gomock "github.com/ohad83/mock/gomock"
)

// MockMatcher is a mock of Matcher interface.
type MockMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMockRecorder
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher.
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance.
func NewMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// Matches mocks base method.
func (m *MockMatcher) Matches(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Matches", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Matches indicates an expected call of Matches.
func (mr *MockMatcherMockRecorder) Matches(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Matches", reflect.TypeOf((*MockMatcher)(nil).Matches), arg0)
}

// String mocks base method.
func (m *MockMatcher) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockMatcherMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockMatcher)(nil).String))
}
