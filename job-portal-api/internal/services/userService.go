package services

import (
	"context"
	"errors"
	"fmt"
	"job-portal-api/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Store) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {

	// We hash the user's password for storage in the database.
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}

	u := models.User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashedPass),
	}

	user, err := s.UserRepo.CreateUser(ctx, u)
	if err != nil {
		return models.User{}, err

	}
	return user, nil
}

func (s *Store) Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
	error) {

	claims, err := s.UserRepo.CheckEmail(ctx, email, password)
	if err != nil {
		return jwt.RegisteredClaims{}, errors.New("not able to generate claims")
	}
	return claims, nil
}
