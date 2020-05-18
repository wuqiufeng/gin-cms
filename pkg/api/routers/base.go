package routers

import (
	"cmsgo/pkg/api/controllers"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(router *gin.RouterGroup) (r gin.IRouter) {
	baseRouter := router.Group("base")
	{
		baseRouter.GET("healthcheck", controllers.Healthy)
	}
	return baseRouter
}
