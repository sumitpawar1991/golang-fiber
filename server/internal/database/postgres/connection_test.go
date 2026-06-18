package postgres

import (
	"my-fiber-app/server/internal/config"
	"testing"
)

func TestPostgresConnection(t *testing.T) {

	cfg := config.PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "password",
		Database: "golang_test",
		SSLMode:  "disable",
	}

	db, err := Connect(cfg)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if db == nil {
		t.Fatal("expected database connection")
	}
}
