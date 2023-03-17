package main

import (
	"github.com/gofiber/fiber/v2"

	"main/configs"
	"main/models"
	"main/services"
)

var (
	gameCategoriesService services.GameCategoriesService
	gameLicenseService    services.GameLicenseService
	gamePromoService      services.GamePromoService
	companyService        services.CompanyService
	gameProductService    services.GameProductService
	server                *fiber.App
	load_configs          configs.Configs
)

func register_models() []interface{} {
	all_models := []interface{}{}

	// append models here
	all_models = append(all_models, &models.GameCategories{})
	all_models = append(all_models, &models.GameLicense{})
	all_models = append(all_models, &models.GamePromo{})
	all_models = append(all_models, &models.Company{})
	// all_models = append(all_models, &models.GameProduct{})

	return all_models
}

func register_services() {
	gameCategoriesService = gameCategoriesService.NewGameCategoriesService(load_configs.DatabaseConnection)
	gameLicenseService = gameLicenseService.NewGameLicenseService(load_configs.DatabaseConnection)
	gamePromoService = gamePromoService.NewGamePromoService(load_configs.DatabaseConnection)
	companyService = companyService.NewCompanyService(load_configs.DatabaseConnection)
	gameProductService = gameProductService.NewGameProductService(load_configs.DatabaseConnection)

}

func init() {
	all_models := register_models()

	load_configs = configs.LoadConfig("configs/app_config.yaml", all_models...)

	register_services()

	server = fiber.New(load_configs.Fiber)
}
