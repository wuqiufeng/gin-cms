package routers

import (
	"cmsgo/pkg/api/controllers"
	"cmsgo/pkg/api/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	var router = gin.Default()

	router.Use(gin.Recovery())

	// 跨域
	router.Use(middleware.Cors())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthcheck", controllers.Healthy)

	// 方便统一添加路由组前缀 多服务器上线使用
	apiGroup := router.Group("")

	InitBaseRouter(apiGroup)
	InitUserRouter(apiGroup)

	return router
}
