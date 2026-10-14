package repository

import (
	"my-fiber-app/server/internal/blog/model"

	"gorm.io/gorm"
)

// handle db operation
type BlogRepository struct {
	db *gorm.DB
}

// its create a repository with a dababase connection
func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (r *BlogRepository) ListBlogs() ([]model.Blog, error) {

	records := make([]model.Blog, 0)

	if err := r.db.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
