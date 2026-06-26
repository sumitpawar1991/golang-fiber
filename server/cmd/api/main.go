package main

import (
	"log"
	"my-fiber-app/server/internal/config"
	"my-fiber-app/server/internal/database/postgres"
)

func main() {

	config.Load()

	//sdatabase.Connect()

	cfg := config.Load()

	db, err := postgres.Connect(cfg.Postgress)

	if err != nil {
		log.Fatal(err)
	}

}
