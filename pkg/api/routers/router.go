package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	var router = gin.Default()

	//TODO
	//middleware

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.GET("/healthcheck", controllers.Healthy)

	apiGroup := router.Group("")

	InitBaseRouter(apiGroup)

	return router
}
