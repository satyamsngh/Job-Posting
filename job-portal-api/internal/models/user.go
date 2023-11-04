package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"-"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
