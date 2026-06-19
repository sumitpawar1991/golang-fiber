package helper

import (
	"my-fiber-app/server/internal/config"
)

func TestPostgresConfig() config.PostgresConfig {
	return config.PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Database: "golang_test",
		SSLMode:  "disable",
	}
}
