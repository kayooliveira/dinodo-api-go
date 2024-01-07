package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitConfig() error {
	initEnv()

	var err error
	db, err = initMysql()

	if err != nil {
		logger.Err("Error initializing MySQL")
		return err
	}

	return nil
}

func GetMysql() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
