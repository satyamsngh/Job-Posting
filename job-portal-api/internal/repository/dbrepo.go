package repository

import "job-portal-api/internal/models"

func (r *Repo) AutoMigrate() error {

	err := r.DB.Migrator().AutoMigrate(&models.User{}, &models.Companies{}, &models.Job{})
	if err != nil {
		return err
	}

	err = r.DB.Migrator().AutoMigrate(&models.User{}, &models.Companies{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
