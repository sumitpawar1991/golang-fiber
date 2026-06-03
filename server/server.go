package main

import (
	"log"
	"my-fiber-app/server/database"
	"my-fiber-app/server/router"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
		},
		AllowCredentials: false,
	}))

	app.Use(logger.New())
	router.SetupRoutes(app)
	// app.Get("/", func(c fiber.Ctx) error {

	// 	return c.JSON(fiber.Map{"message": "welcome to my first blog application"})
	// 	//return c.SendString("Hello world")
	// })

	app.Listen(":8000")
}
