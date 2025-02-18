// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/weaveworks/flintlock/core/ports (interfaces: MicroVMService,MicroVMRepository,EventService,IDService,ImageService,ReconcileMicroVMsUseCase,NetworkService,MicroVMCommandUseCases,MicroVMQueryUseCases)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/weaveworks/flintlock/core/models"
	ports "github.com/weaveworks/flintlock/core/ports"
)

// MockMicroVMService is a mock of MicroVMService interface.
type MockMicroVMService struct {
	ctrl     *gomock.Controller
	recorder *MockMicroVMServiceMockRecorder
}

// MockMicroVMServiceMockRecorder is the mock recorder for MockMicroVMService.
type MockMicroVMServiceMockRecorder struct {
	mock *MockMicroVMService
}

// NewMockMicroVMService creates a new mock instance.
func NewMockMicroVMService(ctrl *gomock.Controller) *MockMicroVMService {
	mock := &MockMicroVMService{ctrl: ctrl}
	mock.recorder = &MockMicroVMServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMicroVMService) EXPECT() *MockMicroVMServiceMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockMicroVMService) Capabilities() models.Capabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(models.Capabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockMicroVMServiceMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockMicroVMService)(nil).Capabilities))
}

// Create mocks base method.
func (m *MockMicroVMService) Create(arg0 context.Context, arg1 *models.MicroVM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMicroVMServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMicroVMService)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockMicroVMService) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMicroVMServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMicroVMService)(nil).Delete), arg0, arg1)
}

// Pause mocks base method.
func (m *MockMicroVMService) Pause(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pause", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pause indicates an expected call of Pause.
func (mr *MockMicroVMServiceMockRecorder) Pause(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pause", reflect.TypeOf((*MockMicroVMService)(nil).Pause), arg0, arg1)
}

// Resume mocks base method.
func (m *MockMicroVMService) Resume(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resume indicates an expected call of Resume.
func (mr *MockMicroVMServiceMockRecorder) Resume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resume", reflect.TypeOf((*MockMicroVMService)(nil).Resume), arg0, arg1)
}

// Start mocks base method.
func (m *MockMicroVMService) Start(arg0 context.Context, arg1 *models.MicroVM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockMicroVMServiceMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockMicroVMService)(nil).Start), arg0, arg1)
}

// State mocks base method.
func (m *MockMicroVMService) State(arg0 context.Context, arg1 string) (ports.MicroVMState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State", arg0, arg1)
	ret0, _ := ret[0].(ports.MicroVMState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockMicroVMServiceMockRecorder) State(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockMicroVMService)(nil).State), arg0, arg1)
}

// Stop mocks base method.
func (m *MockMicroVMService) Stop(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockMicroVMServiceMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockMicroVMService)(nil).Stop), arg0, arg1)
}

// MockMicroVMRepository is a mock of MicroVMRepository interface.
type MockMicroVMRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMicroVMRepositoryMockRecorder
}

// MockMicroVMRepositoryMockRecorder is the mock recorder for MockMicroVMRepository.
type MockMicroVMRepositoryMockRecorder struct {
	mock *MockMicroVMRepository
}

// NewMockMicroVMRepository creates a new mock instance.
func NewMockMicroVMRepository(ctrl *gomock.Controller) *MockMicroVMRepository {
	mock := &MockMicroVMRepository{ctrl: ctrl}
	mock.recorder = &MockMicroVMRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMicroVMRepository) EXPECT() *MockMicroVMRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockMicroVMRepository) Delete(arg0 context.Context, arg1 *models.MicroVM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMicroVMRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMicroVMRepository)(nil).Delete), arg0, arg1)
}

// Exists mocks base method.
func (m *MockMicroVMRepository) Exists(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockMicroVMRepositoryMockRecorder) Exists(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockMicroVMRepository)(nil).Exists), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockMicroVMRepository) Get(arg0 context.Context, arg1 ports.RepositoryGetOptions) (*models.MicroVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*models.MicroVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMicroVMRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMicroVMRepository)(nil).Get), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockMicroVMRepository) GetAll(arg0 context.Context, arg1 string) ([]*models.MicroVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]*models.MicroVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockMicroVMRepositoryMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMicroVMRepository)(nil).GetAll), arg0, arg1)
}

// ReleaseLease mocks base method.
func (m *MockMicroVMRepository) ReleaseLease(arg0 context.Context, arg1 *models.MicroVM) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseLease", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseLease indicates an expected call of ReleaseLease.
func (mr *MockMicroVMRepositoryMockRecorder) ReleaseLease(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseLease", reflect.TypeOf((*MockMicroVMRepository)(nil).ReleaseLease), arg0, arg1)
}

// Save mocks base method.
func (m *MockMicroVMRepository) Save(arg0 context.Context, arg1 *models.MicroVM) (*models.MicroVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(*models.MicroVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockMicroVMRepositoryMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMicroVMRepository)(nil).Save), arg0, arg1)
}

// MockEventService is a mock of EventService interface.
type MockEventService struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceMockRecorder
}

// MockEventServiceMockRecorder is the mock recorder for MockEventService.
type MockEventServiceMockRecorder struct {
	mock *MockEventService
}

// NewMockEventService creates a new mock instance.
func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
	mock := &MockEventService{ctrl: ctrl}
	mock.recorder = &MockEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockEventService) Publish(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockEventServiceMockRecorder) Publish(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockEventService)(nil).Publish), arg0, arg1, arg2)
}

// Subscribe mocks base method.
func (m *MockEventService) Subscribe(arg0 context.Context) (<-chan *ports.EventEnvelope, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(<-chan *ports.EventEnvelope)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockEventServiceMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockEventService)(nil).Subscribe), arg0)
}

// SubscribeTopic mocks base method.
func (m *MockEventService) SubscribeTopic(arg0 context.Context, arg1 string) (<-chan *ports.EventEnvelope, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeTopic", arg0, arg1)
	ret0, _ := ret[0].(<-chan *ports.EventEnvelope)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// SubscribeTopic indicates an expected call of SubscribeTopic.
func (mr *MockEventServiceMockRecorder) SubscribeTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeTopic", reflect.TypeOf((*MockEventService)(nil).SubscribeTopic), arg0, arg1)
}

// SubscribeTopics mocks base method.
func (m *MockEventService) SubscribeTopics(arg0 context.Context, arg1 []string) (<-chan *ports.EventEnvelope, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeTopics", arg0, arg1)
	ret0, _ := ret[0].(<-chan *ports.EventEnvelope)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// SubscribeTopics indicates an expected call of SubscribeTopics.
func (mr *MockEventServiceMockRecorder) SubscribeTopics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeTopics", reflect.TypeOf((*MockEventService)(nil).SubscribeTopics), arg0, arg1)
}

// MockIDService is a mock of IDService interface.
type MockIDService struct {
	ctrl     *gomock.Controller
	recorder *MockIDServiceMockRecorder
}

// MockIDServiceMockRecorder is the mock recorder for MockIDService.
type MockIDServiceMockRecorder struct {
	mock *MockIDService
}

// NewMockIDService creates a new mock instance.
func NewMockIDService(ctrl *gomock.Controller) *MockIDService {
	mock := &MockIDService{ctrl: ctrl}
	mock.recorder = &MockIDServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDService) EXPECT() *MockIDServiceMockRecorder {
	return m.recorder
}

// GenerateRandom mocks base method.
func (m *MockIDService) GenerateRandom() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRandom")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateRandom indicates an expected call of GenerateRandom.
func (mr *MockIDServiceMockRecorder) GenerateRandom() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRandom", reflect.TypeOf((*MockIDService)(nil).GenerateRandom))
}

// MockImageService is a mock of ImageService interface.
type MockImageService struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceMockRecorder
}

// MockImageServiceMockRecorder is the mock recorder for MockImageService.
type MockImageServiceMockRecorder struct {
	mock *MockImageService
}

// NewMockImageService creates a new mock instance.
func NewMockImageService(ctrl *gomock.Controller) *MockImageService {
	mock := &MockImageService{ctrl: ctrl}
	mock.recorder = &MockImageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageService) EXPECT() *MockImageServiceMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockImageService) Exists(arg0 context.Context, arg1 *ports.ImageSpec) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockImageServiceMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockImageService)(nil).Exists), arg0, arg1)
}

// IsMounted mocks base method.
func (m *MockImageService) IsMounted(arg0 context.Context, arg1 *ports.ImageMountSpec) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMounted", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMounted indicates an expected call of IsMounted.
func (mr *MockImageServiceMockRecorder) IsMounted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMounted", reflect.TypeOf((*MockImageService)(nil).IsMounted), arg0, arg1)
}

// Pull mocks base method.
func (m *MockImageService) Pull(arg0 context.Context, arg1 *ports.ImageSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockImageServiceMockRecorder) Pull(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockImageService)(nil).Pull), arg0, arg1)
}

// PullAndMount mocks base method.
func (m *MockImageService) PullAndMount(arg0 context.Context, arg1 *ports.ImageMountSpec) ([]models.Mount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullAndMount", arg0, arg1)
	ret0, _ := ret[0].([]models.Mount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PullAndMount indicates an expected call of PullAndMount.
func (mr *MockImageServiceMockRecorder) PullAndMount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullAndMount", reflect.TypeOf((*MockImageService)(nil).PullAndMount), arg0, arg1)
}

// MockReconcileMicroVMsUseCase is a mock of ReconcileMicroVMsUseCase interface.
type MockReconcileMicroVMsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockReconcileMicroVMsUseCaseMockRecorder
}

// MockReconcileMicroVMsUseCaseMockRecorder is the mock recorder for MockReconcileMicroVMsUseCase.
type MockReconcileMicroVMsUseCaseMockRecorder struct {
	mock *MockReconcileMicroVMsUseCase
}

// NewMockReconcileMicroVMsUseCase creates a new mock instance.
func NewMockReconcileMicroVMsUseCase(ctrl *gomock.Controller) *MockReconcileMicroVMsUseCase {
	mock := &MockReconcileMicroVMsUseCase{ctrl: ctrl}
	mock.recorder = &MockReconcileMicroVMsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReconcileMicroVMsUseCase) EXPECT() *MockReconcileMicroVMsUseCaseMockRecorder {
	return m.recorder
}

// ReconcileMicroVM mocks base method.
func (m *MockReconcileMicroVMsUseCase) ReconcileMicroVM(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMicroVM", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileMicroVM indicates an expected call of ReconcileMicroVM.
func (mr *MockReconcileMicroVMsUseCaseMockRecorder) ReconcileMicroVM(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMicroVM", reflect.TypeOf((*MockReconcileMicroVMsUseCase)(nil).ReconcileMicroVM), arg0, arg1, arg2)
}

// ResyncMicroVMs mocks base method.
func (m *MockReconcileMicroVMsUseCase) ResyncMicroVMs(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResyncMicroVMs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResyncMicroVMs indicates an expected call of ResyncMicroVMs.
func (mr *MockReconcileMicroVMsUseCaseMockRecorder) ResyncMicroVMs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResyncMicroVMs", reflect.TypeOf((*MockReconcileMicroVMsUseCase)(nil).ResyncMicroVMs), arg0, arg1)
}

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// IfaceCreate mocks base method.
func (m *MockNetworkService) IfaceCreate(arg0 context.Context, arg1 ports.IfaceCreateInput) (*ports.IfaceDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IfaceCreate", arg0, arg1)
	ret0, _ := ret[0].(*ports.IfaceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IfaceCreate indicates an expected call of IfaceCreate.
func (mr *MockNetworkServiceMockRecorder) IfaceCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IfaceCreate", reflect.TypeOf((*MockNetworkService)(nil).IfaceCreate), arg0, arg1)
}

// IfaceDelete mocks base method.
func (m *MockNetworkService) IfaceDelete(arg0 context.Context, arg1 ports.DeleteIfaceInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IfaceDelete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IfaceDelete indicates an expected call of IfaceDelete.
func (mr *MockNetworkServiceMockRecorder) IfaceDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IfaceDelete", reflect.TypeOf((*MockNetworkService)(nil).IfaceDelete), arg0, arg1)
}

// IfaceDetails mocks base method.
func (m *MockNetworkService) IfaceDetails(arg0 context.Context, arg1 string) (*ports.IfaceDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IfaceDetails", arg0, arg1)
	ret0, _ := ret[0].(*ports.IfaceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IfaceDetails indicates an expected call of IfaceDetails.
func (mr *MockNetworkServiceMockRecorder) IfaceDetails(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IfaceDetails", reflect.TypeOf((*MockNetworkService)(nil).IfaceDetails), arg0, arg1)
}

// IfaceExists mocks base method.
func (m *MockNetworkService) IfaceExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IfaceExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IfaceExists indicates an expected call of IfaceExists.
func (mr *MockNetworkServiceMockRecorder) IfaceExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IfaceExists", reflect.TypeOf((*MockNetworkService)(nil).IfaceExists), arg0, arg1)
}

// MockMicroVMCommandUseCases is a mock of MicroVMCommandUseCases interface.
type MockMicroVMCommandUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockMicroVMCommandUseCasesMockRecorder
}

// MockMicroVMCommandUseCasesMockRecorder is the mock recorder for MockMicroVMCommandUseCases.
type MockMicroVMCommandUseCasesMockRecorder struct {
	mock *MockMicroVMCommandUseCases
}

// NewMockMicroVMCommandUseCases creates a new mock instance.
func NewMockMicroVMCommandUseCases(ctrl *gomock.Controller) *MockMicroVMCommandUseCases {
	mock := &MockMicroVMCommandUseCases{ctrl: ctrl}
	mock.recorder = &MockMicroVMCommandUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMicroVMCommandUseCases) EXPECT() *MockMicroVMCommandUseCasesMockRecorder {
	return m.recorder
}

// CreateMicroVM mocks base method.
func (m *MockMicroVMCommandUseCases) CreateMicroVM(arg0 context.Context, arg1 *models.MicroVM) (*models.MicroVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMicroVM", arg0, arg1)
	ret0, _ := ret[0].(*models.MicroVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMicroVM indicates an expected call of CreateMicroVM.
func (mr *MockMicroVMCommandUseCasesMockRecorder) CreateMicroVM(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMicroVM", reflect.TypeOf((*MockMicroVMCommandUseCases)(nil).CreateMicroVM), arg0, arg1)
}

// DeleteMicroVM mocks base method.
func (m *MockMicroVMCommandUseCases) DeleteMicroVM(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMicroVM", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMicroVM indicates an expected call of DeleteMicroVM.
func (mr *MockMicroVMCommandUseCasesMockRecorder) DeleteMicroVM(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMicroVM", reflect.TypeOf((*MockMicroVMCommandUseCases)(nil).DeleteMicroVM), arg0, arg1, arg2)
}

// MockMicroVMQueryUseCases is a mock of MicroVMQueryUseCases interface.
type MockMicroVMQueryUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockMicroVMQueryUseCasesMockRecorder
}

// MockMicroVMQueryUseCasesMockRecorder is the mock recorder for MockMicroVMQueryUseCases.
type MockMicroVMQueryUseCasesMockRecorder struct {
	mock *MockMicroVMQueryUseCases
}

// NewMockMicroVMQueryUseCases creates a new mock instance.
func NewMockMicroVMQueryUseCases(ctrl *gomock.Controller) *MockMicroVMQueryUseCases {
	mock := &MockMicroVMQueryUseCases{ctrl: ctrl}
	mock.recorder = &MockMicroVMQueryUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMicroVMQueryUseCases) EXPECT() *MockMicroVMQueryUseCasesMockRecorder {
	return m.recorder
}

// GetAllMicroVM mocks base method.
func (m *MockMicroVMQueryUseCases) GetAllMicroVM(arg0 context.Context, arg1 string) ([]*models.MicroVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMicroVM", arg0, arg1)
	ret0, _ := ret[0].([]*models.MicroVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMicroVM indicates an expected call of GetAllMicroVM.
func (mr *MockMicroVMQueryUseCasesMockRecorder) GetAllMicroVM(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMicroVM", reflect.TypeOf((*MockMicroVMQueryUseCases)(nil).GetAllMicroVM), arg0, arg1)
}

// GetMicroVM mocks base method.
func (m *MockMicroVMQueryUseCases) GetMicroVM(arg0 context.Context, arg1, arg2 string) (*models.MicroVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMicroVM", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.MicroVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMicroVM indicates an expected call of GetMicroVM.
func (mr *MockMicroVMQueryUseCasesMockRecorder) GetMicroVM(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMicroVM", reflect.TypeOf((*MockMicroVMQueryUseCases)(nil).GetMicroVM), arg0, arg1, arg2)
}
