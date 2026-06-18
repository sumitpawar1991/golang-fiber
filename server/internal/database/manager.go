package database

import (
	"my-fiber-app/server/internal/config"
	"my-fiber-app/server/internal/database/postgres"

	"gorm.io/gorm"
)

type Manager struct {
	Postgres *gorm.DB
}

func New(cfg *config.Config) (*Manager, error) {
	postgresDB, err := postgres.Connect(cfg.Postgres)

	if err != nil {
		return nil, err
	}

	return &Manager{
		Postgres: postgresDB,
	}, nil
}
