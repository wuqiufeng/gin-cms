package routers

import (
	"cmsgo/pkg/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(router *gin.RouterGroup) (r gin.IRouter) {
	authController := &controllers.AuthController{}

	baseRouter := router.Group("base")
	{
		baseRouter.GET("healthcheck", controllers.Healthy)
		baseRouter.GET("/login", authController.JwtAuthLogin)
	}
	return baseRouter
}
