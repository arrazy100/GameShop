package services

import (
	"main/models"

	"gorm.io/gorm"
)

type GameCategoriesService struct {
	db *gorm.DB
}

func (repo *GameCategoriesService) NewGameCategoriesService(db *gorm.DB) GameCategoriesService {
	return GameCategoriesService{
		db: db,
	}
}

func (repo *GameCategoriesService) FindAll() ([]models.GameCategoriesDetails, error) {
	var gameCategories []models.GameCategories
	var gameCategoriesDetails []models.GameCategoriesDetails

	// get all game categories
	err := repo.db.Model(&models.GameCategories{}).
		Preload("Games").
		Find(&gameCategories).Error

	if err != nil {
		return nil, err
	}

	// convert to details model
	for _, gameCategories := range gameCategories {
		var games []string

		for _, gameProduct := range gameCategories.Games {
			// get game
			var game models.GameProduct
			repo.db.Where("product_id = ?", gameProduct.ProductID).First(&game)

			if game.Name != "" {
				games = append(games, gameProduct.Name)
			}
		}

		gameCategoriesDetails = append(gameCategoriesDetails, models.GameCategoriesDetails{
			CategoryID:  gameCategories.CategoryID.String(),
			Name:        gameCategories.Name,
			Description: gameCategories.Description,
			Games:       games,
		})
	}

	return gameCategoriesDetails, nil
}

func (repo *GameCategoriesService) Create(gameCategory *models.GameCategoriesInput) (*models.GameCategories, error) {
	// convert input to model
	newGameCategory := models.GameCategories{
		Name:        gameCategory.Name,
		Description: gameCategory.Description,
	}

	// parse games
	for _, game := range gameCategory.Games {
		var gameProduct models.GameProduct
		repo.db.Where("name = ?", game).First(&gameProduct)
		newGameCategory.Games = append(newGameCategory.Games, &gameProduct)
	}

	// save game category
	err := repo.db.Create(&newGameCategory).Error

	if err != nil {
		return nil, err
	}

	return &newGameCategory, nil
}
