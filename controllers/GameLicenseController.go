package controller

import (
	"main/constants"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

type GameLicenseController struct {
	service services.GameLicenseService
}

func (repo *GameLicenseController) NewGameLicenseController(service services.GameLicenseService) GameLicenseController {
	return GameLicenseController{
		service: service,
	}
}

// FindAllGameLicense returns all game licenses
// @Summary Returns all game licenses
// @Description Returns all game licenses
// @Tags game-license
// @Accept json
// @Produce json
// @Success 200 {object} models.GameLicenseDetails
// @Router /game-license/ [get]
func (repo *GameLicenseController) FindAll(c *fiber.Ctx) error {
	gameLicenses, err := repo.service.FindAll()

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gameLicenses)
}

// CreateGameLicense creates a new game license
// @Summary Creates a new game license
// @Tags game-license
// @Accept json
// @Produce json
// @Param input body models.GameLicenseInput true "Game license input"
// @Success 200 {object} models.GameLicenseDetails
// @Router /game-license/ [post]
func (repo *GameLicenseController) Create(c *fiber.Ctx) error {
	// parse input
	var data models.GameLicenseInput

	if err := c.BodyParser(&data); err != nil {
		constants.ErrorResponse(c, err)
	}

	// validate struct
	error_validation := constants.ValidateStruct(data)

	if error_validation != nil {
		constants.ValidationErrorResponse(c, error_validation)
	}

	gameLicense, err := repo.service.Create(&data)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gameLicense)
}

func (repo *GameLicenseController) Route(rg fiber.Router) {
	router := rg.Group("/game-license")
	router.Get("/", repo.FindAll)
	router.Post("/", repo.Create)
}
