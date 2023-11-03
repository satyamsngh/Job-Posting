package repository

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"job-portal-api/internal/models"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(ctx context.Context, UserDetails models.User) (models.User, error) {
	result := r.DB.Create(&UserDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("could not create the user")
	}
	return UserDetails, nil
}
func (r *Repo) CheckEmail(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error) {
	var u models.User
	tx := r.DB.Where("email = ?", email).First(&u)
	if tx.Error != nil {
		return jwt.RegisteredClaims{}, tx.Error
	}

	// We check if the provided password matches the hashed password in the database.
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	c := jwt.RegisteredClaims{
		Issuer:    "jobportal project",
		Subject:   strconv.FormatUint(uint64(u.ID), 10),
		Audience:  jwt.ClaimStrings{"companies"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	return c, nil

}
