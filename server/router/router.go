package router

import (
	"my-fiber-app/server/controller"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/blog", controller.BlogList)
	app.Get("/blog/show/:id", controller.BlogShow)
	app.Post("/blog", controller.BlogCreate)
	app.Put("/blog/edit/:id", controller.BlogEdit)
	app.Delete("/blog/delete/:id", controller.BlogDelete)
}
