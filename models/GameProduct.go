package models

import "github.com/google/uuid"

type GameProduct struct {
	ProductID   uuid.UUID         `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string            `gorm:"type:varchar(100);not null"`
	Description string            `gorm:"type:varchar(1000);not null"`
	Price       float64           `gorm:"type:decimal(10,2);not null"`
	Categories  []*GameCategories `gorm:"many2many:game_product_categories;"`
	LicenseID   GameLicense       `gorm:"foreignkey:LicenseID"`
	PromoID     GamePromo         `gorm:"foreignkey:PromoID"`
	CompanyID   Company           `gorm:"foreignkey:CompanyID"`
}

type GameProductInput struct {
	Name        string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float64 `validate:"required"`
	Categories  []string
	LicenseID   string
	PromoID     string
	CompanyID   string
}

type GameProductDetails struct {
	ProductID   string
	Name        string
	Description string
	Price       float64
	Categories  []string
	LicenseName string
	PromoName   string
	CompanyName string
}

type GameProductSetCategories struct {
	ProductID  string   `validate:"required"`
	Categories []string `validate:"required"`
}

type GameProductSetCompany struct {
	ProductID string `validate:"required"`
	CompanyID string `validate:"required"`
}

type GameProductSetLicense struct {
	ProductID string `validate:"required"`
	LicenseID string `validate:"required"`
}

type GameProductSetPromo struct {
	ProductID string `validate:"required"`
	PromoID   string `validate:"required"`
}
