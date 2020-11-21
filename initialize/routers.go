package initialize

import (
	"gitee.com/chensyf/container/app/main/controllers"
	_ "gitee.com/chensyf/container/docs"
	"gitee.com/chensyf/container/global"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	if global.CONFIG.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	api := router.Group("")

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	controllers.StudentController.Router(api)

	return router
}