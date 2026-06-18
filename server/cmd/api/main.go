package main

import (
	"log"
	"my-fiber-app/server/internal/database"
	"my-fiber-app/server/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
