package main

import (
	"fmt"
	"gitee.com/chensyf/container/global"
	"gitee.com/chensyf/container/initialize"
	"log"
)


// @title 容器化Student
// @version 1.0
// @description 容器化Student
// @termsOfService https://github.com/swaggo/gin-swagger
func main() {
	initialize.SetupConfig("config/config.ini")
	// mysql
	initialize.SetupMysql()

	// redis
	initialize.SetupRedis()

	// tables
	initialize.SetupTables()

	// routers
	router := initialize.SetupRouter()

	err := router.Run(fmt.Sprintf(":%d", global.CONFIG.Server.HttpPort))
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
