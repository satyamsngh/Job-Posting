package models

import (
	"gorm.io/gorm"
)

type Companies struct {
	gorm.Model
	CompanyName string `json:"company_name"`
	FoundedYear int    `json:"founded_year"`
	Location    string `json:"location"`
	UserId      uint   `json:"user_id"`
	Address     string `json:"address"`
	Jobs        []Job  `json:"jobs,omitempty" gorm:"foreignKey:CompanyID"`
}

type NewComapanies struct {
	CompanyName string `json:"company_name" validate:"required"`
	FoundedYear int    `json:"founded_year" validate:"required,number"`
	Location    string `json:"location" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Jobs        []Job  `json:"jobs"`
}

type Job struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	CompanyID   uint   `json:"company_id"`
}
