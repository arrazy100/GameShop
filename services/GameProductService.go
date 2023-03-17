package services

import (
	"main/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameProductService struct {
	db *gorm.DB
}

func (repo *GameProductService) NewGameProductService(db *gorm.DB) GameProductService {
	return GameProductService{
		db: db,
	}
}

func (repo *GameProductService) FindAll() ([]models.GameProductDetails, error) {
	var gameProducts []models.GameProduct
	var gameProductsDetails []models.GameProductDetails

	// get all game products
	err := repo.db.Model(&models.GameProduct{}).
		Preload("Categories").
		Preload("CompanyID").
		Preload("LicenseID").
		Preload("PromoID").
		Find(&gameProducts).Error

	for _, gameProduct := range gameProducts {
		categories := make([]string, len(gameProduct.Categories))

		for i, category := range gameProduct.Categories {
			categories[i] = category.Name
		}

		// parse company
		var company string
		if gameProduct.CompanyID.CompanyID != uuid.Nil {
			company = gameProduct.CompanyID.Name
		}

		// parse license
		var license string
		if gameProduct.LicenseID.LicenseID != uuid.Nil {
			license = gameProduct.LicenseID.Name
		}

		// parse promo
		var promo string
		if gameProduct.PromoID.PromoID != uuid.Nil {
			promo = gameProduct.PromoID.Name
		}

		gameProductsDetails = append(gameProductsDetails, models.GameProductDetails{
			ProductID:   gameProduct.ProductID.String(),
			Name:        gameProduct.Name,
			Description: gameProduct.Description,
			Price:       gameProduct.Price,
			Categories:  categories,
			CompanyName: company,
			LicenseName: license,
			PromoName:   promo,
		})
	}

	if err != nil {
		return nil, err
	}

	return gameProductsDetails, nil
}

func (repo *GameProductService) Create(gameProduct *models.GameProductInput) (*models.GameProductDetails, error) {
	// convert input to model
	newGameProduct := models.GameProduct{
		Name:        gameProduct.Name,
		Description: gameProduct.Description,
		Price:       gameProduct.Price,
	}

	// parse categories
	for _, category := range gameProduct.Categories {
		var gameCategory models.GameCategories
		repo.db.Where("name = ?", category).First(&gameCategory)
		if gameCategory.Name != "" {
			newGameProduct.Categories = append(newGameProduct.Categories, &gameCategory)
		}
	}

	// parse license
	var gameLicense models.GameLicense
	repo.db.Where("license_id = ?", gameProduct.LicenseID).First(&gameLicense)
	newGameProduct.LicenseID = gameLicense

	// parse promo
	var gamePromo models.GamePromo
	repo.db.Where("promo_id = ?", gameProduct.PromoID).First(&gamePromo)
	newGameProduct.PromoID = gamePromo

	// parse company
	var company models.Company
	repo.db.Where("company_id = ?", gameProduct.CompanyID).First(&company)
	newGameProduct.CompanyID = company

	// save to db
	err := repo.db.Create(&newGameProduct).Error

	if err != nil {
		return nil, err
	}

	// convert categories to []string
	categories := make([]string, len(newGameProduct.Categories))

	for i, category := range newGameProduct.Categories {
		categories[i] = category.Name
	}

	// convert newGameProduct to GameProductDetails
	gameProductDetails := models.GameProductDetails{
		ProductID:   newGameProduct.ProductID.String(),
		Name:        newGameProduct.Name,
		Description: newGameProduct.Description,
		Price:       newGameProduct.Price,
		Categories:  categories,
		CompanyName: company.Name,
		LicenseName: gameLicense.Name,
		PromoName:   gamePromo.Name,
	}

	return &gameProductDetails, nil
}

func (repo *GameProductService) SetCategories(data *models.GameProductSetCategories) error {
	// get game product
	var gameProduct models.GameProduct
	repo.db.Where("product_id = ?", data.ProductID).First(&gameProduct)

	// parse categories
	for _, category := range data.Categories {
		var gameCategory models.GameCategories
		err := repo.db.Where("name = ?", category).First(&gameCategory).Error

		if err != nil {
			continue
		}

		gameProduct.Categories = append(gameProduct.Categories, &gameCategory)
		gameCategory.Games = append(gameCategory.Games, &gameProduct)

		// save to db
		err = repo.db.Save(&gameCategory).Error

		if err != nil {
			return err
		}
	}

	// save to db
	err := repo.db.Save(&gameProduct).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GameProductService) SetCompany(data *models.GameProductSetCompany) error {
	// get game product
	var gameProduct models.GameProduct
	repo.db.Where("product_id = ?", data.ProductID).First(&gameProduct)

	// get company
	var company models.Company
	repo.db.Where("company_id = ?", data.CompanyID).First(&company)

	// set company
	gameProduct.CompanyID = company

	// save to db
	err := repo.db.Save(&gameProduct).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GameProductService) SetLicense(data *models.GameProductSetLicense) error {
	// get game product
	var gameProduct models.GameProduct
	repo.db.Where("product_id = ?", data.ProductID).First(&gameProduct)

	// get license
	var license models.GameLicense
	repo.db.Where("license_id = ?", data.LicenseID).First(&license)

	// set license
	gameProduct.LicenseID = license

	// save to db
	err := repo.db.Save(&gameProduct).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *GameProductService) SetPromo(data *models.GameProductSetPromo) error {
	// get game product
	var gameProduct models.GameProduct
	repo.db.Where("product_id = ?", data.ProductID).First(&gameProduct)

	// get promo
	var promo models.GamePromo
	repo.db.Where("promo_id = ?", data.PromoID).First(&promo)

	// set promo
	gameProduct.PromoID = promo

	// save to db
	err := repo.db.Save(&gameProduct).Error

	if err != nil {
		return err
	}

	return nil
}
