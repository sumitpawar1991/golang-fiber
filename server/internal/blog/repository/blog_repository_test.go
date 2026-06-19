package repository

import (
	blogModel "my-fiber-app/server/internal/blog/model"
	"my-fiber-app/server/internal/database"
	"my-fiber-app/server/internal/database/postgres"
	"my-fiber-app/server/test/helper"

	"testing"
)

func TestCreateBlog(t *testing.T) {

	cfg := helper.TestPostgresConfig()

	db, err := postgres.Connect(cfg)

	if err != nil {
		t.Fatalf("db connection failed: %v", err)
	}

	err = database.Migrate(db)

	if err != nil {
		t.Fatalf("migration failed: %v", err)
	}

	repo := NewBlogRepository(nil)

	blog := &blogModel.Blog{
		Title: "Test Blog",
		Post:  "This is test content",
	}

	err = repo.Create(blog)

	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}
}
