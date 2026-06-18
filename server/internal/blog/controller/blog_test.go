package controller

import (
	"my-fiber-app/server/controller"

	"github.com/gofiber/fiber/v3"
)

func setupApp() *fiber.App {
	app := fiber.New()

	app.Post("/blogs", controller.BlogCreate)

	return app
}
