package database

import (
	blogModel "my-fiber-app/server/internal/blog/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	return db.AutoMigrate(
		&blogModel.Blog{},
	)
}
