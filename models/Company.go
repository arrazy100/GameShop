package models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	CompanyID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `gorm:"type:varchar(100);not null"`
	State     string    `gorm:"type:varchar(100);not null"`
	City      string    `gorm:"type:varchar(100);not null"`
	Country   string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
	IsActive  bool      `gorm:"type:boolean;not null"`
}

type CompanyInput struct {
	Name    string `validate:"required"`
	State   string `validate:"required"`
	City    string `validate:"required"`
	Country string `validate:"required"`
}

type CompanyDetails struct {
	CompanyID uuid.UUID `json:"company_id"`
	Name      string    `json:"name"`
	State     string    `json:"state"`
	City      string    `json:"city"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}
