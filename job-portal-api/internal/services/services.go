package services

import (
	"context"
	"errors"
	"job-portal-api/internal/models"
	"job-portal-api/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source service.go -destination mockmodels/service_mock.go -package mockmodels

type Service interface {
	CreatCompanies(ctx context.Context, nc models.NewComapanies, UserId uint) (models.Companies, error)
	ViewCompanies(ctx context.Context, companyId string) ([]models.Companies, error)
	ViewCompaniesById(ctx context.Context, companybyid uint, userId string) ([]models.Companies, error)
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	CreateJob(ctx context.Context, newJob models.Job, userId string) (models.Job, error)
	AllJob(ctx context.Context, userId string) ([]models.Job, error)
	ListJobs(ctx context.Context, companyId uint, userId string) ([]models.Job, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
		error)
	JobsByID(ctx context.Context, jobID uint64, userId string) (models.Job, error)
}

type Store struct {
	UserRepo repository.UserRepo
}

func NewStore(userRepo repository.UserRepo) (Service, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Store{
		UserRepo: userRepo,
	}, nil
}
