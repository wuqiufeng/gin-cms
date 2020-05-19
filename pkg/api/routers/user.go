package routers

import (
	"cmsgo/pkg/api/controllers"
	"cmsgo/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userController := &controllers.UserController{}
	UserRouter := Router.Group("user").Use(middleware.JWTAuth())
	{
		UserRouter.POST("changePassword", userController.ResetPassword)
	}

}
