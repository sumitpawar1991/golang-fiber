package repository

import (
	"errors"
	blogModel "my-fiber-app/server/internal/blog/model"

	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (r *BlogRepository) Create(blog *blogModel.Blog) error {

	if r.db == nil {
		return errors.New("database connection is nil")
	}
	return r.db.Create(blog).Error
}
