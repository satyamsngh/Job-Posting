package repository

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"job-portal-api/internal/models"
)

type Repo struct {
	DB *gorm.DB
}

//go:generate mockgen -source=repository.go -destination=mock_repository.go -package=repository
type UserRepo interface {
	CreateUser(ctx context.Context, userData models.User) (models.User, error)
	CheckEmail(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error)

	CreateCompany(ctx context.Context, companyData models.Companies) (models.Companies, error)
	ViewCompanies(ctx context.Context) ([]models.Companies, error)
	ViewCompanyById(ctx context.Context, cid uint) ([]models.Companies, error)

	CreateJob(ctx context.Context, jobData models.Job) (models.Job, error)
	FindJob(ctx context.Context, cid uint64) ([]models.Job, error)
	FindAllJobs(ctx context.Context) ([]models.Job, error)
	ViewJobDetailsById(ctx context.Context, jid uint64) (models.Job, error)
	ViewJobByCompanyId(ctx context.Context, id uint) ([]models.Job, error)
	AutoMigrate() error
}

func NewRepository(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db cannot be null")
	}
	return &Repo{
		DB: db,
	}, nil
}
