// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/service/LanguageService.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	dto "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
)

// MockLanguageService is a mock of LanguageService interface.
type MockLanguageService struct {
	ctrl     *gomock.Controller
	recorder *MockLanguageServiceMockRecorder
}

// MockLanguageServiceMockRecorder is the mock recorder for MockLanguageService.
type MockLanguageServiceMockRecorder struct {
	mock *MockLanguageService
}

// NewMockLanguageService creates a new mock instance.
func NewMockLanguageService(ctrl *gomock.Controller) *MockLanguageService {
	mock := &MockLanguageService{ctrl: ctrl}
	mock.recorder = &MockLanguageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLanguageService) EXPECT() *MockLanguageServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLanguageService) Create(request dto.LanguageDto) (domain.Language, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", request)
	ret0, _ := ret[0].(domain.Language)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLanguageServiceMockRecorder) Create(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLanguageService)(nil).Create), request)
}

// FindAll mocks base method.
func (m *MockLanguageService) FindAll() (*[]domain.Language, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].(*[]domain.Language)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockLanguageServiceMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockLanguageService)(nil).FindAll))
}

// Update mocks base method.
func (m *MockLanguageService) Update(request domain.Language) (*domain.Language, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", request)
	ret0, _ := ret[0].(*domain.Language)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockLanguageServiceMockRecorder) Update(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLanguageService)(nil).Update), request)
}
