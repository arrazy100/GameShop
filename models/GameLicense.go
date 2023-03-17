package models

import "github.com/google/uuid"

type GameLicense struct {
	LicenseID   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:varchar(1000);not null"`
}

type GameLicenseInput struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

type GameLicenseDetails struct {
	LicenseID   string `json:"license_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
