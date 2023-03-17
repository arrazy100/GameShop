package main

import (
	"log"
	controller "main/controllers"
	"strconv"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "main/docs"
)

var (
	gameCategoriesController controller.GameCategoriesController
	gameLicenseController    controller.GameLicenseController
	gamePromoController      controller.GamePromoController
	companyController        controller.CompanyController
	gameProductController    controller.GameProductController
	router                   fiber.Router
)

func register_controllers() {
	gameCategoriesController = gameCategoriesController.NewGameCategoriesController(gameCategoriesService)
	gameCategoriesController.Route(router)
	gameLicenseController = gameLicenseController.NewGameLicenseController(gameLicenseService)
	gameLicenseController.Route(router)
	gamePromoController = gamePromoController.NewGamePromoController(gamePromoService)
	gamePromoController.Route(router)
	companyController = companyController.NewCompanyController(companyService)
	companyController.Route(router)
	gameProductController = gameProductController.NewGameProductController(gameProductService)
	gameProductController.Route(router)

}

func startFiberServer() {
	server.Get("/swagger/*", fiberSwagger.WrapHandler)

	router = server.Group("/api/v1")

	register_controllers()

	log.Fatal(server.Listen(":" + strconv.Itoa(load_configs.AppPort)))
}
