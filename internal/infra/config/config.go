package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitConfig() error {
	initEnv()

	var err error
	db, err = initMysql()

	if err != nil {
		fmt.Println("Error initializing MySQL")
		return err
	}

	return nil
}

func GetMysql() *gorm.DB {
	return db
}
