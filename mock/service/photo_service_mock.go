// Code generated by MockGen. DO NOT EDIT.
// Source: service/photo-service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/vinix997/Final-Project-Hacktiv8/entity"
)

// MockPhotoService is a mock of PhotoService interface.
type MockPhotoService struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoServiceMockRecorder
}

// MockPhotoServiceMockRecorder is the mock recorder for MockPhotoService.
type MockPhotoServiceMockRecorder struct {
	mock *MockPhotoService
}

// NewMockPhotoService creates a new mock instance.
func NewMockPhotoService(ctrl *gomock.Controller) *MockPhotoService {
	mock := &MockPhotoService{ctrl: ctrl}
	mock.recorder = &MockPhotoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoService) EXPECT() *MockPhotoServiceMockRecorder {
	return m.recorder
}

// CreatePhoto mocks base method.
func (m *MockPhotoService) CreatePhoto(ctx context.Context, photo *entity.Photo) (*entity.Photo, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoto", ctx, photo)
	ret0, _ := ret[0].(*entity.Photo)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// CreatePhoto indicates an expected call of CreatePhoto.
func (mr *MockPhotoServiceMockRecorder) CreatePhoto(ctx, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoto", reflect.TypeOf((*MockPhotoService)(nil).CreatePhoto), ctx, photo)
}

// DeletePhoto mocks base method.
func (m *MockPhotoService) DeletePhoto(ctx context.Context, id int) (string, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoto", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// DeletePhoto indicates an expected call of DeletePhoto.
func (mr *MockPhotoServiceMockRecorder) DeletePhoto(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoto", reflect.TypeOf((*MockPhotoService)(nil).DeletePhoto), ctx, id)
}

// GetPhoto mocks base method.
func (m *MockPhotoService) GetPhoto(ctx context.Context) ([]entity.PhotoResponse, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhoto", ctx)
	ret0, _ := ret[0].([]entity.PhotoResponse)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GetPhoto indicates an expected call of GetPhoto.
func (mr *MockPhotoServiceMockRecorder) GetPhoto(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoto", reflect.TypeOf((*MockPhotoService)(nil).GetPhoto), ctx)
}

// UpdatePhoto mocks base method.
func (m *MockPhotoService) UpdatePhoto(ctx context.Context, photo *entity.Photo, id int) (*entity.Photo, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", ctx, photo, id)
	ret0, _ := ret[0].(*entity.Photo)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockPhotoServiceMockRecorder) UpdatePhoto(ctx, photo, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockPhotoService)(nil).UpdatePhoto), ctx, photo, id)
}

// MockPhotoRepository is a mock of PhotoRepository interface.
type MockPhotoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPhotoRepositoryMockRecorder
}

// MockPhotoRepositoryMockRecorder is the mock recorder for MockPhotoRepository.
type MockPhotoRepositoryMockRecorder struct {
	mock *MockPhotoRepository
}

// NewMockPhotoRepository creates a new mock instance.
func NewMockPhotoRepository(ctrl *gomock.Controller) *MockPhotoRepository {
	mock := &MockPhotoRepository{ctrl: ctrl}
	mock.recorder = &MockPhotoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhotoRepository) EXPECT() *MockPhotoRepositoryMockRecorder {
	return m.recorder
}

// CreatePhoto mocks base method.
func (m *MockPhotoRepository) CreatePhoto(ctx context.Context, photo *entity.Photo) (*entity.Photo, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoto", ctx, photo)
	ret0, _ := ret[0].(*entity.Photo)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// CreatePhoto indicates an expected call of CreatePhoto.
func (mr *MockPhotoRepositoryMockRecorder) CreatePhoto(ctx, photo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).CreatePhoto), ctx, photo)
}

// DeletePhoto mocks base method.
func (m *MockPhotoRepository) DeletePhoto(ctx context.Context, id int) (string, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoto", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// DeletePhoto indicates an expected call of DeletePhoto.
func (mr *MockPhotoRepositoryMockRecorder) DeletePhoto(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).DeletePhoto), ctx, id)
}

// GetPhoto mocks base method.
func (m *MockPhotoRepository) GetPhoto(ctx context.Context) ([]entity.PhotoResponse, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhoto", ctx)
	ret0, _ := ret[0].([]entity.PhotoResponse)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GetPhoto indicates an expected call of GetPhoto.
func (mr *MockPhotoRepositoryMockRecorder) GetPhoto(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoto", reflect.TypeOf((*MockPhotoRepository)(nil).GetPhoto), ctx)
}

// UpdatePhoto mocks base method.
func (m *MockPhotoRepository) UpdatePhoto(ctx context.Context, photo *entity.Photo, id int) (*entity.Photo, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", ctx, photo, id)
	ret0, _ := ret[0].(*entity.Photo)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockPhotoRepositoryMockRecorder) UpdatePhoto(ctx, photo, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockPhotoRepository)(nil).UpdatePhoto), ctx, photo, id)
}
