package models

import "github.com/google/uuid"

type GameCategories struct {
	CategoryID  uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string         `gorm:"type:varchar(100);not null"`
	Description string         `gorm:"type:varchar(1000);not null"`
	Games       []*GameProduct `gorm:"many2many:game_product_categories;"`
}

type GameCategoriesInput struct {
	Name        string `validate:"required"`
	Description string `validate:"required"`
	Games       []string
}

type GameCategoriesDetails struct {
	CategoryID  string
	Name        string
	Description string
	Games       []string
}
