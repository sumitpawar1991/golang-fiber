package router

import (
	"my-fiber-app/server/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/blog", controller.BlogList)
	app.Get("/blog/:id", controller.BlogDetail)
	app.Post("/blog", controller.BlogCreate)
	app.Put("/blog/:id", controller.BlogUpdate)
	app.Delete("/blog/:id", controller.BlogDelete)
}
