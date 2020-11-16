package initialize

import (
	"gitee.com/chensyf/container/app/main/models"
	"gitee.com/chensyf/container/global"
)

func SetupTables() {
	db := global.DB
	db.AutoMigrate(&models.Student{})
}