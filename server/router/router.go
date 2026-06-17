package router

import (
	"my-fiber-app/server/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/blog", controller.BlogList)
	app.Get("/blog/show/:id", controller.BlogDetail)
	app.Post("/blog", controller.BlogCreate)
	app.Put("/blog/edit/:id", controller.BlogUpdate)
	app.Delete("/blog/delete/:id", controller.BlogDelete)
}
