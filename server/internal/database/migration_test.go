package database

import (
	blogModel "my-fiber-app/server/internal/blog/model"
	"my-fiber-app/server/internal/database/postgres"
	"my-fiber-app/server/test/helper"
	"testing"
)

func TestBlogMigration(t *testing.T) {

	//cfg := config.Load()

	// cfg := config.PostgresConfig{
	// 	Host:     "localhost",
	// 	Port:     "5432",
	// 	User:     "postgres",
	// 	Password: "password",
	// 	Database: "golang_test",
	// 	SSLMode:  "disable",
	// }

	cfg := helper.TestPostgresConfig()

	db, err := postgres.Connect(cfg)

	if err != nil {
		t.Fatalf("postgres database connection failed :%v", err)
	}

	err = Migrate(db)

	if err != nil {
		t.Fatalf("Migration failed :%v", err)
	}

	if !db.Migrator().HasTable(&blogModel.Blog{}) {
		t.Fatal("table is not created")
	}
}
