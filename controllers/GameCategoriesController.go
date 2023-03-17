package controller

import (
	"main/constants"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

type GameCategoriesController struct {
	service services.GameCategoriesService
}

func (repo *GameCategoriesController) NewGameCategoriesController(service services.GameCategoriesService) GameCategoriesController {
	return GameCategoriesController{
		service: service,
	}
}

// GetAllGameCategories godoc
// @Summary find all game categories
// @Schemes http
// @Tags game-categories
// @Accept json
// @Produce json
// @Success 200 {object} models.GameCategories
// @Router /game-categories/ [get]
func (repo *GameCategoriesController) FindAll(c *fiber.Ctx) error {
	gameCategories, err := repo.service.FindAll()

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gameCategories)
}

// CreateGameCategory godoc
// @Summary create game category
// @Schemes http
// @Tags game-categories
// @Accept json
// @Produce json
// @Param gameCategory body models.GameCategoriesInput true "Game Category"
// @Success 200 {object} models.GameCategories
// @Router /game-categories/ [post]
func (repo *GameCategoriesController) Create(c *fiber.Ctx) error {
	// parse input
	var gameCategoryInput models.GameCategoriesInput

	if err := c.BodyParser(&gameCategoryInput); err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate input
	error_validation := constants.ValidateStruct(gameCategoryInput)

	if error_validation != nil {
		return constants.ValidationErrorResponse(c, error_validation)
	}

	// create game category
	gameCategory, err := repo.service.Create(&gameCategoryInput)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gameCategory)
}

func (repo *GameCategoriesController) Route(rg fiber.Router) {
	router := rg.Group("/game-categories")
	router.Get("/", repo.FindAll)
	router.Post("/", repo.Create)
}
