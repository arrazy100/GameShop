package services

import (
	"main/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GamePromoService struct {
	db *gorm.DB
}

func (repo *GamePromoService) NewGamePromoService(db *gorm.DB) GamePromoService {
	return GamePromoService{
		db: db,
	}
}

func (repo *GamePromoService) FindAll() ([]models.GamePromoDetails, error) {
	var gamePromos []models.GamePromo
	var gamePromosDetails []models.GamePromoDetails

	// get all game promos
	err := repo.db.Model(&models.GamePromo{}).
		Preload("CompanyID").
		Find(&gamePromos).Error

	if err != nil {
		return nil, err
	}

	for _, gamePromo := range gamePromos {
		// parse company
		var company string
		if gamePromo.CompanyID.CompanyID != uuid.Nil {
			company = gamePromo.CompanyID.Name
		}

		gamePromosDetails = append(gamePromosDetails, models.GamePromoDetails{
			PromoID:     gamePromo.PromoID.String(),
			Name:        gamePromo.Name,
			StartDate:   gamePromo.StartDate,
			EndDate:     gamePromo.EndDate,
			Discount:    gamePromo.Discount,
			CompanyName: company,
		})
	}

	return gamePromosDetails, nil
}

func (repo *GamePromoService) Create(gamePromoInput models.GamePromoInput) (*models.GamePromo, error) {
	var gamePromo models.GamePromo

	// get company by id
	var company models.Company
	repo.db.Where("company_id = ?", gamePromoInput.CompanyID).First(&company)

	// create game promo
	gamePromo = models.GamePromo{
		Name:        gamePromoInput.Name,
		Description: gamePromoInput.Description,
		Discount:    gamePromoInput.Discount,
		StartDate:   gamePromoInput.StartDate,
		EndDate:     gamePromoInput.EndDate,
		IsActive:    gamePromoInput.IsActive,
		CompanyID:   company,
	}

	err := repo.db.Model(&models.GamePromo{}).Create(&gamePromo).Error

	if err != nil {
		return nil, err
	}

	return &gamePromo, nil
}
