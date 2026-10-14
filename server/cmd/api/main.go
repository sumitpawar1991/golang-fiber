package main

import (
	"log"
	"my-fiber-app/server/internal/config"
	"my-fiber-app/server/internal/database/postgres"
	"my-fiber-app/server/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {

	//config.Load()

	//sdatabase.Connect()

	cfg := config.Load()

	db, err := postgres.Connect(cfg.Postgres)

	if err != nil {
		log.Fatal(err)
	}

	postgressqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	if err := postgressqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("postgres sql connected successfully")

	// app := fiber.New()

	// log.Fatal(app.Listen(":3000"))
	// // app.Get("/", func(c fiber.Ctx) error {
	// // 	return c.SendString("Fiber working!")
	// // })

	// app.Listen(":3000")

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
