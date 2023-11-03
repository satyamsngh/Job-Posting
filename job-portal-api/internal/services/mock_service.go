package services

import (
	"context"
	models "job-portal-api/internal/models"
	reflect "reflect"

	v5 "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

type MockServiceMockRecorder struct {
	mock *MockService
}

func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

func (m *MockService) Authenticate(ctx context.Context, email, password string) (v5.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, email, password)
	ret0, _ := ret[0].(v5.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) Authenticate(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockService)(nil).Authenticate), ctx, email, password)
}

func (m *MockService) AutoMigrate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoMigrate")
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockServiceMockRecorder) AutoMigrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockService)(nil).AutoMigrate))
}

func (m *MockService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, nu)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) CreateUser(ctx, nu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), ctx, nu)
}

func (m *MockService) CreatCompanies(ctx context.Context, ni models.NewComapanies, userId uint) (models.Companies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompanies", ctx, ni, userId)
	ret0, _ := ret[0].(models.Companies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) CreatCompanies(ctx, ni, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompanies", reflect.TypeOf((*MockService)(nil).CreatCompanies), ctx, ni, userId)
}

func (m *MockService) ViewCompanies(ctx context.Context, companyId string) ([]models.Companies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompanies", ctx, companyId)
	ret0, _ := ret[0].([]models.Companies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewCompanies(ctx, companyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompanies", reflect.TypeOf((*MockService)(nil).ViewCompanies), ctx, companyId)
}

func (m *MockService) ViewCompaniesById(ctx context.Context, companyById uint, userId string) ([]models.Companies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "api/companies/:companyID", ctx, companyById, userId)
	ret0, _ := ret[0].([]models.Companies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ViewCompaniesById(ctx, companyById, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "api/companies/:companyID", reflect.TypeOf((*MockService)(nil).ViewCompaniesById), ctx, companyById, userId)
}

func (m *MockService) CreateJob(ctx context.Context, newJob models.Job, userId string) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, newJob, userId)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) CreateJob(ctx, newJob, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockService)(nil).CreateJob), ctx, newJob, userId)
}

func (m *MockService) AllJob(ctx context.Context, userId string) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllJob", ctx, userId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) AllJob(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllJob", reflect.TypeOf((*MockService)(nil).AllJob), ctx, userId)
}

func (m *MockService) ListJobs(ctx context.Context, companyId uint, userId string) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "api/companies/:companyID/list-jobs", ctx, companyId, userId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) ListJobs(ctx, companyId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "api/companies/:companyID/list-jobs", reflect.TypeOf((*MockService)(nil).ListJobs), ctx, companyId, userId)
}

func (m *MockService) JobsById(ctx context.Context, jobById uint, userId string) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "api/jobs/:jobID", ctx, jobById, userId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServiceMockRecorder) JobsById(ctx, jobById, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "api/jobs/:jobID", reflect.TypeOf((*MockService)(nil).JobsById), ctx, jobById, userId)
}
