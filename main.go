package main

import (
	"container/global"
	"container/initialize"
	"fmt"
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

	// tables
	initialize.SetupTables()

	// routers
	router := initialize.SetupRouter()

	err := router.Run(fmt.Sprintf(":%d", global.CONFIG.Server.HttpPort))
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
