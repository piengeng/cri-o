// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/containers/storage (interfaces: Store)

// Package containerstoragemock is a generated GoMock package.
package containerstoragemock

import (
	storage "github.com/containers/storage"
	graphdriver "github.com/containers/storage/drivers"
	archive "github.com/containers/storage/pkg/archive"
	idtools "github.com/containers/storage/pkg/idtools"
	lockfile "github.com/containers/storage/pkg/lockfile"
	gomock "github.com/golang/mock/gomock"
	digest "github.com/opencontainers/go-digest"
	io "io"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// ApplyDiff mocks base method
func (m *MockStore) ApplyDiff(arg0 string, arg1 io.Reader) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyDiff", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyDiff indicates an expected call of ApplyDiff
func (mr *MockStoreMockRecorder) ApplyDiff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyDiff", reflect.TypeOf((*MockStore)(nil).ApplyDiff), arg0, arg1)
}

// Changes mocks base method
func (m *MockStore) Changes(arg0, arg1 string) ([]archive.Change, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes", arg0, arg1)
	ret0, _ := ret[0].([]archive.Change)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Changes indicates an expected call of Changes
func (mr *MockStoreMockRecorder) Changes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockStore)(nil).Changes), arg0, arg1)
}

// Container mocks base method
func (m *MockStore) Container(arg0 string) (*storage.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Container", arg0)
	ret0, _ := ret[0].(*storage.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Container indicates an expected call of Container
func (mr *MockStoreMockRecorder) Container(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Container", reflect.TypeOf((*MockStore)(nil).Container), arg0)
}

// ContainerBigData mocks base method
func (m *MockStore) ContainerBigData(arg0, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerBigData", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerBigData indicates an expected call of ContainerBigData
func (mr *MockStoreMockRecorder) ContainerBigData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerBigData", reflect.TypeOf((*MockStore)(nil).ContainerBigData), arg0, arg1)
}

// ContainerBigDataDigest mocks base method
func (m *MockStore) ContainerBigDataDigest(arg0, arg1 string) (digest.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerBigDataDigest", arg0, arg1)
	ret0, _ := ret[0].(digest.Digest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerBigDataDigest indicates an expected call of ContainerBigDataDigest
func (mr *MockStoreMockRecorder) ContainerBigDataDigest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerBigDataDigest", reflect.TypeOf((*MockStore)(nil).ContainerBigDataDigest), arg0, arg1)
}

// ContainerBigDataSize mocks base method
func (m *MockStore) ContainerBigDataSize(arg0, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerBigDataSize", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerBigDataSize indicates an expected call of ContainerBigDataSize
func (mr *MockStoreMockRecorder) ContainerBigDataSize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerBigDataSize", reflect.TypeOf((*MockStore)(nil).ContainerBigDataSize), arg0, arg1)
}

// ContainerByLayer mocks base method
func (m *MockStore) ContainerByLayer(arg0 string) (*storage.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerByLayer", arg0)
	ret0, _ := ret[0].(*storage.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerByLayer indicates an expected call of ContainerByLayer
func (mr *MockStoreMockRecorder) ContainerByLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerByLayer", reflect.TypeOf((*MockStore)(nil).ContainerByLayer), arg0)
}

// ContainerDirectory mocks base method
func (m *MockStore) ContainerDirectory(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerDirectory", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerDirectory indicates an expected call of ContainerDirectory
func (mr *MockStoreMockRecorder) ContainerDirectory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerDirectory", reflect.TypeOf((*MockStore)(nil).ContainerDirectory), arg0)
}

// ContainerParentOwners mocks base method
func (m *MockStore) ContainerParentOwners(arg0 string) ([]int, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerParentOwners", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ContainerParentOwners indicates an expected call of ContainerParentOwners
func (mr *MockStoreMockRecorder) ContainerParentOwners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerParentOwners", reflect.TypeOf((*MockStore)(nil).ContainerParentOwners), arg0)
}

// ContainerRunDirectory mocks base method
func (m *MockStore) ContainerRunDirectory(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerRunDirectory", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerRunDirectory indicates an expected call of ContainerRunDirectory
func (mr *MockStoreMockRecorder) ContainerRunDirectory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerRunDirectory", reflect.TypeOf((*MockStore)(nil).ContainerRunDirectory), arg0)
}

// ContainerSize mocks base method
func (m *MockStore) ContainerSize(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerSize", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerSize indicates an expected call of ContainerSize
func (mr *MockStoreMockRecorder) ContainerSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerSize", reflect.TypeOf((*MockStore)(nil).ContainerSize), arg0)
}

// Containers mocks base method
func (m *MockStore) Containers() ([]storage.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Containers")
	ret0, _ := ret[0].([]storage.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Containers indicates an expected call of Containers
func (mr *MockStoreMockRecorder) Containers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Containers", reflect.TypeOf((*MockStore)(nil).Containers))
}

// CreateContainer mocks base method
func (m *MockStore) CreateContainer(arg0 string, arg1 []string, arg2, arg3, arg4 string, arg5 *storage.ContainerOptions) (*storage.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*storage.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockStoreMockRecorder) CreateContainer(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockStore)(nil).CreateContainer), arg0, arg1, arg2, arg3, arg4, arg5)
}

// CreateImage mocks base method
func (m *MockStore) CreateImage(arg0 string, arg1 []string, arg2, arg3 string, arg4 *storage.ImageOptions) (*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImage", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateImage indicates an expected call of CreateImage
func (mr *MockStoreMockRecorder) CreateImage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImage", reflect.TypeOf((*MockStore)(nil).CreateImage), arg0, arg1, arg2, arg3, arg4)
}

// CreateLayer mocks base method
func (m *MockStore) CreateLayer(arg0, arg1 string, arg2 []string, arg3 string, arg4 bool, arg5 *storage.LayerOptions) (*storage.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLayer", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*storage.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLayer indicates an expected call of CreateLayer
func (mr *MockStoreMockRecorder) CreateLayer(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLayer", reflect.TypeOf((*MockStore)(nil).CreateLayer), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Delete mocks base method
func (m *MockStore) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), arg0)
}

// DeleteContainer mocks base method
func (m *MockStore) DeleteContainer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContainer indicates an expected call of DeleteContainer
func (mr *MockStoreMockRecorder) DeleteContainer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContainer", reflect.TypeOf((*MockStore)(nil).DeleteContainer), arg0)
}

// DeleteImage mocks base method
func (m *MockStore) DeleteImage(arg0 string, arg1 bool) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteImage indicates an expected call of DeleteImage
func (mr *MockStoreMockRecorder) DeleteImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockStore)(nil).DeleteImage), arg0, arg1)
}

// DeleteLayer mocks base method
func (m *MockStore) DeleteLayer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLayer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLayer indicates an expected call of DeleteLayer
func (mr *MockStoreMockRecorder) DeleteLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLayer", reflect.TypeOf((*MockStore)(nil).DeleteLayer), arg0)
}

// Diff mocks base method
func (m *MockStore) Diff(arg0, arg1 string, arg2 *storage.DiffOptions) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Diff", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Diff indicates an expected call of Diff
func (mr *MockStoreMockRecorder) Diff(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Diff", reflect.TypeOf((*MockStore)(nil).Diff), arg0, arg1, arg2)
}

// DiffSize mocks base method
func (m *MockStore) DiffSize(arg0, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiffSize", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiffSize indicates an expected call of DiffSize
func (mr *MockStoreMockRecorder) DiffSize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiffSize", reflect.TypeOf((*MockStore)(nil).DiffSize), arg0, arg1)
}

// Exists mocks base method
func (m *MockStore) Exists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockStoreMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockStore)(nil).Exists), arg0)
}

// Free mocks base method
func (m *MockStore) Free() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Free")
}

// Free indicates an expected call of Free
func (mr *MockStoreMockRecorder) Free() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockStore)(nil).Free))
}

// FromContainerDirectory mocks base method
func (m *MockStore) FromContainerDirectory(arg0, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromContainerDirectory", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromContainerDirectory indicates an expected call of FromContainerDirectory
func (mr *MockStoreMockRecorder) FromContainerDirectory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromContainerDirectory", reflect.TypeOf((*MockStore)(nil).FromContainerDirectory), arg0, arg1)
}

// FromContainerRunDirectory mocks base method
func (m *MockStore) FromContainerRunDirectory(arg0, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FromContainerRunDirectory", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FromContainerRunDirectory indicates an expected call of FromContainerRunDirectory
func (mr *MockStoreMockRecorder) FromContainerRunDirectory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FromContainerRunDirectory", reflect.TypeOf((*MockStore)(nil).FromContainerRunDirectory), arg0, arg1)
}

// GIDMap mocks base method
func (m *MockStore) GIDMap() []idtools.IDMap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GIDMap")
	ret0, _ := ret[0].([]idtools.IDMap)
	return ret0
}

// GIDMap indicates an expected call of GIDMap
func (mr *MockStoreMockRecorder) GIDMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GIDMap", reflect.TypeOf((*MockStore)(nil).GIDMap))
}

// GetDigestLock mocks base method
func (m *MockStore) GetDigestLock(arg0 digest.Digest) (lockfile.Locker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDigestLock", arg0)
	ret0, _ := ret[0].(lockfile.Locker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDigestLock indicates an expected call of GetDigestLock
func (mr *MockStoreMockRecorder) GetDigestLock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDigestLock", reflect.TypeOf((*MockStore)(nil).GetDigestLock), arg0)
}

// GraphDriver mocks base method
func (m *MockStore) GraphDriver() (graphdriver.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GraphDriver")
	ret0, _ := ret[0].(graphdriver.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GraphDriver indicates an expected call of GraphDriver
func (mr *MockStoreMockRecorder) GraphDriver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GraphDriver", reflect.TypeOf((*MockStore)(nil).GraphDriver))
}

// GraphDriverName mocks base method
func (m *MockStore) GraphDriverName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GraphDriverName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GraphDriverName indicates an expected call of GraphDriverName
func (mr *MockStoreMockRecorder) GraphDriverName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GraphDriverName", reflect.TypeOf((*MockStore)(nil).GraphDriverName))
}

// GraphOptions mocks base method
func (m *MockStore) GraphOptions() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GraphOptions")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GraphOptions indicates an expected call of GraphOptions
func (mr *MockStoreMockRecorder) GraphOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GraphOptions", reflect.TypeOf((*MockStore)(nil).GraphOptions))
}

// GraphRoot mocks base method
func (m *MockStore) GraphRoot() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GraphRoot")
	ret0, _ := ret[0].(string)
	return ret0
}

// GraphRoot indicates an expected call of GraphRoot
func (mr *MockStoreMockRecorder) GraphRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GraphRoot", reflect.TypeOf((*MockStore)(nil).GraphRoot))
}

// Image mocks base method
func (m *MockStore) Image(arg0 string) (*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Image", arg0)
	ret0, _ := ret[0].(*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Image indicates an expected call of Image
func (mr *MockStoreMockRecorder) Image(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Image", reflect.TypeOf((*MockStore)(nil).Image), arg0)
}

// ImageBigData mocks base method
func (m *MockStore) ImageBigData(arg0, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageBigData", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageBigData indicates an expected call of ImageBigData
func (mr *MockStoreMockRecorder) ImageBigData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageBigData", reflect.TypeOf((*MockStore)(nil).ImageBigData), arg0, arg1)
}

// ImageBigDataDigest mocks base method
func (m *MockStore) ImageBigDataDigest(arg0, arg1 string) (digest.Digest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageBigDataDigest", arg0, arg1)
	ret0, _ := ret[0].(digest.Digest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageBigDataDigest indicates an expected call of ImageBigDataDigest
func (mr *MockStoreMockRecorder) ImageBigDataDigest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageBigDataDigest", reflect.TypeOf((*MockStore)(nil).ImageBigDataDigest), arg0, arg1)
}

// ImageBigDataSize mocks base method
func (m *MockStore) ImageBigDataSize(arg0, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageBigDataSize", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageBigDataSize indicates an expected call of ImageBigDataSize
func (mr *MockStoreMockRecorder) ImageBigDataSize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageBigDataSize", reflect.TypeOf((*MockStore)(nil).ImageBigDataSize), arg0, arg1)
}

// ImageSize mocks base method
func (m *MockStore) ImageSize(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageSize", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImageSize indicates an expected call of ImageSize
func (mr *MockStoreMockRecorder) ImageSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageSize", reflect.TypeOf((*MockStore)(nil).ImageSize), arg0)
}

// Images mocks base method
func (m *MockStore) Images() ([]storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Images")
	ret0, _ := ret[0].([]storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Images indicates an expected call of Images
func (mr *MockStoreMockRecorder) Images() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Images", reflect.TypeOf((*MockStore)(nil).Images))
}

// ImagesByDigest mocks base method
func (m *MockStore) ImagesByDigest(arg0 digest.Digest) ([]*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagesByDigest", arg0)
	ret0, _ := ret[0].([]*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImagesByDigest indicates an expected call of ImagesByDigest
func (mr *MockStoreMockRecorder) ImagesByDigest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagesByDigest", reflect.TypeOf((*MockStore)(nil).ImagesByDigest), arg0)
}

// ImagesByTopLayer mocks base method
func (m *MockStore) ImagesByTopLayer(arg0 string) ([]*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagesByTopLayer", arg0)
	ret0, _ := ret[0].([]*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImagesByTopLayer indicates an expected call of ImagesByTopLayer
func (mr *MockStoreMockRecorder) ImagesByTopLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagesByTopLayer", reflect.TypeOf((*MockStore)(nil).ImagesByTopLayer), arg0)
}

// Layer mocks base method
func (m *MockStore) Layer(arg0 string) (*storage.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Layer", arg0)
	ret0, _ := ret[0].(*storage.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Layer indicates an expected call of Layer
func (mr *MockStoreMockRecorder) Layer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Layer", reflect.TypeOf((*MockStore)(nil).Layer), arg0)
}

// LayerParentOwners mocks base method
func (m *MockStore) LayerParentOwners(arg0 string) ([]int, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerParentOwners", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LayerParentOwners indicates an expected call of LayerParentOwners
func (mr *MockStoreMockRecorder) LayerParentOwners(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerParentOwners", reflect.TypeOf((*MockStore)(nil).LayerParentOwners), arg0)
}

// LayerSize mocks base method
func (m *MockStore) LayerSize(arg0 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayerSize", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayerSize indicates an expected call of LayerSize
func (mr *MockStoreMockRecorder) LayerSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayerSize", reflect.TypeOf((*MockStore)(nil).LayerSize), arg0)
}

// Layers mocks base method
func (m *MockStore) Layers() ([]storage.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Layers")
	ret0, _ := ret[0].([]storage.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Layers indicates an expected call of Layers
func (mr *MockStoreMockRecorder) Layers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Layers", reflect.TypeOf((*MockStore)(nil).Layers))
}

// LayersByCompressedDigest mocks base method
func (m *MockStore) LayersByCompressedDigest(arg0 digest.Digest) ([]storage.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayersByCompressedDigest", arg0)
	ret0, _ := ret[0].([]storage.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayersByCompressedDigest indicates an expected call of LayersByCompressedDigest
func (mr *MockStoreMockRecorder) LayersByCompressedDigest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayersByCompressedDigest", reflect.TypeOf((*MockStore)(nil).LayersByCompressedDigest), arg0)
}

// LayersByUncompressedDigest mocks base method
func (m *MockStore) LayersByUncompressedDigest(arg0 digest.Digest) ([]storage.Layer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LayersByUncompressedDigest", arg0)
	ret0, _ := ret[0].([]storage.Layer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LayersByUncompressedDigest indicates an expected call of LayersByUncompressedDigest
func (mr *MockStoreMockRecorder) LayersByUncompressedDigest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LayersByUncompressedDigest", reflect.TypeOf((*MockStore)(nil).LayersByUncompressedDigest), arg0)
}

// ListContainerBigData mocks base method
func (m *MockStore) ListContainerBigData(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainerBigData", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainerBigData indicates an expected call of ListContainerBigData
func (mr *MockStoreMockRecorder) ListContainerBigData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainerBigData", reflect.TypeOf((*MockStore)(nil).ListContainerBigData), arg0)
}

// ListImageBigData mocks base method
func (m *MockStore) ListImageBigData(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImageBigData", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImageBigData indicates an expected call of ListImageBigData
func (mr *MockStoreMockRecorder) ListImageBigData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImageBigData", reflect.TypeOf((*MockStore)(nil).ListImageBigData), arg0)
}

// Lookup mocks base method
func (m *MockStore) Lookup(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup
func (mr *MockStoreMockRecorder) Lookup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockStore)(nil).Lookup), arg0)
}

// Metadata mocks base method
func (m *MockStore) Metadata(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Metadata indicates an expected call of Metadata
func (mr *MockStoreMockRecorder) Metadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockStore)(nil).Metadata), arg0)
}

// Mount mocks base method
func (m *MockStore) Mount(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mount indicates an expected call of Mount
func (mr *MockStoreMockRecorder) Mount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockStore)(nil).Mount), arg0, arg1)
}

// MountImage mocks base method
func (m *MockStore) MountImage(arg0 string, arg1 []string, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MountImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MountImage indicates an expected call of MountImage
func (mr *MockStoreMockRecorder) MountImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MountImage", reflect.TypeOf((*MockStore)(nil).MountImage), arg0, arg1, arg2)
}

// Mounted mocks base method
func (m *MockStore) Mounted(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mounted", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mounted indicates an expected call of Mounted
func (mr *MockStoreMockRecorder) Mounted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mounted", reflect.TypeOf((*MockStore)(nil).Mounted), arg0)
}

// Names mocks base method
func (m *MockStore) Names(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Names", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Names indicates an expected call of Names
func (mr *MockStoreMockRecorder) Names(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Names", reflect.TypeOf((*MockStore)(nil).Names), arg0)
}

// PutLayer mocks base method
func (m *MockStore) PutLayer(arg0, arg1 string, arg2 []string, arg3 string, arg4 bool, arg5 *storage.LayerOptions, arg6 io.Reader) (*storage.Layer, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutLayer", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(*storage.Layer)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PutLayer indicates an expected call of PutLayer
func (mr *MockStoreMockRecorder) PutLayer(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutLayer", reflect.TypeOf((*MockStore)(nil).PutLayer), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// RunRoot mocks base method
func (m *MockStore) RunRoot() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunRoot")
	ret0, _ := ret[0].(string)
	return ret0
}

// RunRoot indicates an expected call of RunRoot
func (mr *MockStoreMockRecorder) RunRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunRoot", reflect.TypeOf((*MockStore)(nil).RunRoot))
}

// SetContainerBigData mocks base method
func (m *MockStore) SetContainerBigData(arg0, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContainerBigData", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContainerBigData indicates an expected call of SetContainerBigData
func (mr *MockStoreMockRecorder) SetContainerBigData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContainerBigData", reflect.TypeOf((*MockStore)(nil).SetContainerBigData), arg0, arg1, arg2)
}

// SetContainerDirectoryFile mocks base method
func (m *MockStore) SetContainerDirectoryFile(arg0, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContainerDirectoryFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContainerDirectoryFile indicates an expected call of SetContainerDirectoryFile
func (mr *MockStoreMockRecorder) SetContainerDirectoryFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContainerDirectoryFile", reflect.TypeOf((*MockStore)(nil).SetContainerDirectoryFile), arg0, arg1, arg2)
}

// SetContainerRunDirectoryFile mocks base method
func (m *MockStore) SetContainerRunDirectoryFile(arg0, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContainerRunDirectoryFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetContainerRunDirectoryFile indicates an expected call of SetContainerRunDirectoryFile
func (mr *MockStoreMockRecorder) SetContainerRunDirectoryFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContainerRunDirectoryFile", reflect.TypeOf((*MockStore)(nil).SetContainerRunDirectoryFile), arg0, arg1, arg2)
}

// SetImageBigData mocks base method
func (m *MockStore) SetImageBigData(arg0, arg1 string, arg2 []byte, arg3 func([]byte) (digest.Digest, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetImageBigData", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetImageBigData indicates an expected call of SetImageBigData
func (mr *MockStoreMockRecorder) SetImageBigData(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetImageBigData", reflect.TypeOf((*MockStore)(nil).SetImageBigData), arg0, arg1, arg2, arg3)
}

// SetMetadata mocks base method
func (m *MockStore) SetMetadata(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMetadata", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMetadata indicates an expected call of SetMetadata
func (mr *MockStoreMockRecorder) SetMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetadata", reflect.TypeOf((*MockStore)(nil).SetMetadata), arg0, arg1)
}

// SetNames mocks base method
func (m *MockStore) SetNames(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNames", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetNames indicates an expected call of SetNames
func (mr *MockStoreMockRecorder) SetNames(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNames", reflect.TypeOf((*MockStore)(nil).SetNames), arg0, arg1)
}

// Shutdown mocks base method
func (m *MockStore) Shutdown(arg0 bool) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockStoreMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockStore)(nil).Shutdown), arg0)
}

// Status mocks base method
func (m *MockStore) Status() ([][2]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].([][2]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockStoreMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockStore)(nil).Status))
}

// UIDMap mocks base method
func (m *MockStore) UIDMap() []idtools.IDMap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UIDMap")
	ret0, _ := ret[0].([]idtools.IDMap)
	return ret0
}

// UIDMap indicates an expected call of UIDMap
func (mr *MockStoreMockRecorder) UIDMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UIDMap", reflect.TypeOf((*MockStore)(nil).UIDMap))
}

// Unmount mocks base method
func (m *MockStore) Unmount(arg0 string, arg1 bool) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmount", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unmount indicates an expected call of Unmount
func (mr *MockStoreMockRecorder) Unmount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmount", reflect.TypeOf((*MockStore)(nil).Unmount), arg0, arg1)
}

// UnmountImage mocks base method
func (m *MockStore) UnmountImage(arg0 string, arg1 bool) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmountImage", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnmountImage indicates an expected call of UnmountImage
func (mr *MockStoreMockRecorder) UnmountImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmountImage", reflect.TypeOf((*MockStore)(nil).UnmountImage), arg0, arg1)
}

// Version mocks base method
func (m *MockStore) Version() ([][2]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].([][2]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockStoreMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockStore)(nil).Version))
}

// Wipe mocks base method
func (m *MockStore) Wipe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wipe")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wipe indicates an expected call of Wipe
func (mr *MockStoreMockRecorder) Wipe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wipe", reflect.TypeOf((*MockStore)(nil).Wipe))
}
