package services

import (
	"context"
	"job-portal-api/internal/models"
)

func (s *Store) CreatCompanies(ctx context.Context, nc models.NewComapanies, UserID uint) (models.Companies, error) {

	com := models.Companies{
		CompanyName: nc.CompanyName,
		FoundedYear: nc.FoundedYear,
		Location:    nc.Location,
		UserId:      UserID,
		Address:     nc.Address,
		Jobs:        nc.Jobs,
	}

	company, err := s.UserRepo.CreateCompany(ctx, com)
	if err != nil {
		return models.Companies{}, err
	}
	return company, nil
}

func (s *Store) ViewCompanies(ctx context.Context, companyID string) ([]models.Companies, error) {
	companies, err := s.UserRepo.ViewCompanies(ctx)
	if err != nil {
		return nil, err
	}
	return companies, nil

}

func (s *Store) ViewCompaniesById(ctx context.Context, companyID uint, userID string) ([]models.Companies, error) {
	company, err := s.UserRepo.ViewCompanyById(ctx, companyID)
	if err != nil {
		return []models.Companies{}, err
	}

	return company, nil
}
func (s *Store) CreateJob(ctx context.Context, job models.Job, userID string) (models.Job, error) {

	job, err := s.UserRepo.CreateJob(ctx, job)
	if err != nil {
		return models.Job{}, err
	}

	return job, nil
}
func (s *Store) ListJobs(ctx context.Context, companyID uint, userid string) ([]models.Job, error) {
	jobs, err := s.UserRepo.ViewJobByCompanyId(ctx, companyID)
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}
func (s *Store) AllJob(ctx context.Context, userId string) ([]models.Job, error) {
	jobs, err := s.UserRepo.FindAllJobs(ctx)
	if err != nil {
		return []models.Job{}, err
	}

	return jobs, nil
}
func (s *Store) JobsByID(ctx context.Context, jobID uint64, userId string) (models.Job, error) {
	job, err := s.UserRepo.ViewJobDetailsById(ctx, jobID)
	if err != nil {
		return models.Job{}, err

	}
	return job, nil
}
