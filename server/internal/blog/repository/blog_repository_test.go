package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestBlogRepository_ListBlogs(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Failed to create sql mock:%v", err)
	}

	defer sqlDB.Close()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("failed to open GORM:%v", err)
	}

	columns := []string{"id", "title", "content", "created_at", "updated_at"}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "blogs"`)).
		WillReturnRows(sqlmock.NewRows(columns))

	repo := NewBlogRepository(db)

	records, err := repo.ListBlogs()

	if err != nil {
		t.Fatalf("expected no error,got %v", err)
	}

	if records == nil {
		t.Fatal("expected empty slice , got nil")
	}

	if len(records) != 0 {
		t.Fatalf("expected 0 records , got %d", len(records))
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled SQL expections: %v", err)
	}
}
