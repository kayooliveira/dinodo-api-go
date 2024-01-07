package controller

import (
	"github.com/kayooliveira/dinodo-api-go/internal/infra/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitController() {
	db = config.GetMysql()
}
