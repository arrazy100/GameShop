package services

import (
	"main/models"

	"gorm.io/gorm"
)

type GameLicenseService struct {
	db *gorm.DB
}

func (repo *GameLicenseService) NewGameLicenseService(db *gorm.DB) GameLicenseService {
	return GameLicenseService{
		db: db,
	}
}

func (repo *GameLicenseService) FindAll() ([]models.GameLicenseDetails, error) {
	var gameLicenses []models.GameLicense
	var gameLicensesDetails []models.GameLicenseDetails

	// get all game licenses
	err := repo.db.Model(&models.GameLicense{}).Find(&gameLicenses).Error

	for _, gameLicense := range gameLicenses {
		gameLicensesDetails = append(gameLicensesDetails, models.GameLicenseDetails{
			LicenseID:   gameLicense.LicenseID.String(),
			Name:        gameLicense.Name,
			Description: gameLicense.Description,
		})
	}

	if err != nil {
		return nil, err
	}

	return gameLicensesDetails, nil
}

func (repo *GameLicenseService) Create(gameLicense *models.GameLicenseInput) (*models.GameLicenseDetails, error) {
	var gameLicenseDetails models.GameLicenseDetails

	// create game license
	err := repo.db.Model(&models.GameLicense{}).Create(&models.GameLicense{
		Name:        gameLicense.Name,
		Description: gameLicense.Description,
	}).Scan(&gameLicenseDetails).Error

	if err != nil {
		return nil, err
	}

	return &gameLicenseDetails, nil
}
