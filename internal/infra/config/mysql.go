package config

import (
	"fmt"
	"os"

	"github.com/kayooliveira/dinodo-api-go/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMysql() (*gorm.DB, error) {

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbDatabase := os.Getenv("DATABASE_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("MySQL connection error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&entity.Task{})

	if err != nil {
		fmt.Printf("MySQL automigration error: %v", err)
		return nil, err
	}

	return db, err
}
