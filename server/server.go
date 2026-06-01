package main

import (
	"log"
	"my-fiber-app/server/database"
	"my-fiber-app/server/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error in loading .env file")
	}
	database.ConnectDB()

}

func main() {

	postgresDb, err := database.DBConn.DB()

	if err != nil {
		panic("error in postgresdb connection")
	}

	defer postgresDb.Close()

	app := fiber.New()

	app.Use(logger.New())
	router.SetupRoutes(app)
	// app.Get("/", func(c fiber.Ctx) error {

	// 	return c.JSON(fiber.Map{"message": "welcome to my first blog application"})
	// 	//return c.SendString("Hello world")
	// })

	app.Listen(":8000")
}
