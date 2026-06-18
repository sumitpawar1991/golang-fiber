package main

import (
	"my-fiber-app/server/internal/config"
	"my-fiber-app/server/internal/database"
)

func main() {

	config.Load()

	database.Connect()

}
