package models

import (
	"time"

	"github.com/google/uuid"
)

type GamePromo struct {
	PromoID     uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:varchar(1000);not null"`
	Discount    float64   `gorm:"type:decimal(10,2);not null"`
	StartDate   time.Time `gorm:"type:timestamp;not null"`
	EndDate     time.Time `gorm:"type:timestamp;not null"`
	IsActive    bool      `gorm:"type:boolean;not null"`
	CompanyID   Company   `gorm:"foreignkey:CompanyID"`
}

type GamePromoInput struct {
	Name        string    `validate:"required"`
	Description string    `validate:"required"`
	Discount    float64   `validate:"required"`
	StartDate   time.Time `validate:"required"`
	EndDate     time.Time `validate:"required"`
	IsActive    bool      `validate:"required"`
	CompanyID   uuid.UUID `validate:"required"`
}

type GamePromoDetails struct {
	PromoID     string    `json:"promo_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Discount    float64   `json:"discount"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	IsActive    bool      `json:"is_active"`
	CompanyName string    `json:"company_name"`
}
