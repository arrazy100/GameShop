package controller

import (
	"main/constants"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

type GameProductController struct {
	service services.GameProductService
}

func (repo *GameProductController) NewGameProductController(service services.GameProductService) GameProductController {
	return GameProductController{
		service: service,
	}
}

// FindAllGameProduct godoc
// @Summary find all game products
// @Schemes http
// @Tags game-products
// @Accept json
// @Produce json
// @Success 200 {object} models.GameProductDetails
// @Router /game-products/ [get]
func (repo *GameProductController) FindAll(c *fiber.Ctx) error {
	gameProducts, err := repo.service.FindAll()

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gameProducts)
}

// CreateGameProduct godoc
// @Summary create game product
// @Schemes http
// @Tags game-products
// @Accept json
// @Produce json
// @Param gameProduct body models.GameProductInput true "game product"
// @Success 200 {object} models.GameProductDetails
// @Router /game-products/ [post]
func (repo *GameProductController) Create(c *fiber.Ctx) error {
	// parse input
	var gameProductInput models.GameProductInput

	if err := c.BodyParser(&gameProductInput); err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate input
	error_validation := constants.ValidateStruct(gameProductInput)

	if error_validation != nil {
		return constants.ValidationErrorResponse(c, error_validation)
	}

	// create game product
	gameProduct, err := repo.service.Create(&gameProductInput)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, gameProduct)
}

// SetCategoriesGameProduct godoc
// @Summary set categories for game product
// @Schemes http
// @Tags game-products
// @Accept json
// @Produce json
// @Param gameProduct body models.GameProductSetCategories true "game product"
// @Success 200 {object} models.GameProduct
// @Router /game-products/set-categories [post]
func (repo *GameProductController) SetCategories(c *fiber.Ctx) error {
	// parse input
	var data models.GameProductSetCategories

	if err := c.BodyParser(&data); err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate input
	error_validation := constants.ValidateStruct(data)

	if error_validation != nil {
		return constants.ValidationErrorResponse(c, error_validation)
	}

	// set game product categories
	err := repo.service.SetCategories(&data)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponse(c)
}

// SetCompanyGameProduct godoc
// @Summary set company for game product
// @Schemes http
// @Tags game-products
// @Accept json
// @Produce json
// @Param gameProduct body models.GameProductSetCompany true "game product"
// @Success 200 {object} models.GameProduct
// @Router /game-products/set-company [post]
func (repo *GameProductController) SetCompany(c *fiber.Ctx) error {
	// parse input
	var data models.GameProductSetCompany

	if err := c.BodyParser(&data); err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate input
	error_validation := constants.ValidateStruct(data)

	if error_validation != nil {
		return constants.ValidationErrorResponse(c, error_validation)
	}

	// set game product company
	err := repo.service.SetCompany(&data)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponse(c)
}

// SetLicenseGameProduct godoc
// @Summary set license for game product
// @Schemes http
// @Tags game-products
// @Accept json
// @Produce json
// @Param gameProduct body models.GameProductSetLicense true "game product"
// @Success 200 {object} models.GameProduct
// @Router /game-products/set-license [post]
func (repo *GameProductController) SetLicense(c *fiber.Ctx) error {
	// parse input
	var data models.GameProductSetLicense

	if err := c.BodyParser(&data); err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate input
	error_validation := constants.ValidateStruct(data)

	if error_validation != nil {
		return constants.ValidationErrorResponse(c, error_validation)
	}

	// set game product company
	err := repo.service.SetLicense(&data)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponse(c)
}

// SetPromoGameProduct godoc
// @Summary set promo for game product
// @Schemes http
// @Tags game-products
// @Accept json
// @Produce json
// @Param gameProduct body models.GameProductSetPromo true "game product"
// @Success 200 {object} models.GameProduct
// @Router /game-products/set-promo [post]
func (repo *GameProductController) SetPromo(c *fiber.Ctx) error {
	// parse input
	var data models.GameProductSetPromo

	if err := c.BodyParser(&data); err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate input
	error_validation := constants.ValidateStruct(data)

	if error_validation != nil {
		return constants.ValidationErrorResponse(c, error_validation)
	}

	// set game product company
	err := repo.service.SetPromo(&data)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponse(c)
}

func (repo *GameProductController) Route(rg fiber.Router) {
	router := rg.Group("/game-products")
	router.Get("/", repo.FindAll)
	router.Post("/", repo.Create)
	router.Post("/set-categories", repo.SetCategories)
	router.Post("/set-company", repo.SetCompany)
	router.Post("/set-license", repo.SetLicense)
	router.Post("/set-promo", repo.SetPromo)
}
