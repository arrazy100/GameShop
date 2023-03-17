package controller

import (
	"main/constants"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

type GamePromoController struct {
	service services.GamePromoService
}

func (repo *GamePromoController) NewGamePromoController(service services.GamePromoService) GamePromoController {
	return GamePromoController{
		service: service,
	}
}

// FindAllGamePromo godoc
// @Summary Find all game promo
// @Tags game-promo
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GamePromoDetails
// @Router /game-promo/ [get]
func (repo *GamePromoController) FindAll(c *fiber.Ctx) error {
	gamePromos, err := repo.service.FindAll()

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gamePromos)
}

// CreateGamePromo godoc
// @Summary Create game promo
// @Tags game-promo
// @Accept  json
// @Produce  json
// @Param gamePromo body models.GamePromoInput true "game promo"
// @Success 200 {object} models.GamePromoDetails
// @Router /game-promo/ [post]
func (repo *GamePromoController) Create(c *fiber.Ctx) error {
	// parse input
	var data models.GamePromoInput

	if err := c.BodyParser(&data); err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate input
	error_validation := constants.ValidateStruct(data)

	if error_validation != nil {
		return constants.ValidationErrorResponse(c, error_validation)
	}

	gamePromo, err := repo.service.Create(data)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gamePromo)
}

func (repo *GamePromoController) Route(rg fiber.Router) {
	router := rg.Group("/game-promo")
	router.Get("/", repo.FindAll)
	router.Post("/", repo.Create)
}
