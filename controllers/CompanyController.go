package controller

import (
	"main/constants"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	service services.CompanyService
}

func (repo *CompanyController) NewCompanyController(service services.CompanyService) CompanyController {
	return CompanyController{
		service: service,
	}
}

// FindAllCompany godoc
// @Summary find all companies
// @Tags companies
// @Accept  json
// @Produce  json
// @Success 200 {object} models.CompanyDetails
// @Router /companies/ [get]
func (repo *CompanyController) FindAll(c *fiber.Ctx) error {
	companies, err := repo.service.FindAll()

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, companies)
}

// CreateCompany godoc
// @Summary create a new company
// @Tags companies
// @Accept  json
// @Produce  json
// @Param company body models.CompanyInput true "Company"
// @Success 200 {object} models.CompanyDetails
// @Router /companies/ [post]
func (repo *CompanyController) Create(c *fiber.Ctx) error {
	var companyInput models.CompanyInput

	err := c.BodyParser(&companyInput)
	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	// validate struct
	error_validations := constants.ValidateStruct(&companyInput)
	if err != nil {
		return constants.ValidationErrorResponse(c, error_validations)
	}

	company, err := repo.service.Create(companyInput)

	if err != nil {
		return constants.ErrorResponse(c, err)
	}

	return constants.SuccessResponseWithData(c, company)
}

func (repo *CompanyController) Route(rg fiber.Router) {
	router := rg.Group("/companies")
	router.Get("/", repo.FindAll)
	router.Post("/", repo.Create)
}
