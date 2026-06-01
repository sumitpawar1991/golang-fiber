package router

import (
	"my-fiber-app/server/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)
}
