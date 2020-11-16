package initialize

import (
	"container/app/main/models"
	"container/global"
)

func SetupTables() {
	db := global.DB
	db.AutoMigrate(&models.Student{})
}