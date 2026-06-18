package database

import (
	"log"
	"my-fiber-app/server/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {

	postgres_username := os.Getenv("POSTGRES_DB_USERNAME")

	postgres_password := os.Getenv("POSTGRES_DB_USEPASSWORD")
	postgres_dbname := os.Getenv("POSTGRES_DB_DBNAME")

	dsn := "host=localhost user=" + postgres_username + " password=" + postgres_password + " dbname=" + postgres_dbname + " port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Error),
	// })

	var err error

	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("database connection failed")
	}

	log.Println("connection successfully done")

	//DBConn.AutoMigrate(new(model.Blog))

	DBConn.AutoMigrate(&model.Blog{})

}
