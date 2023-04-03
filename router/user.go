package router

import (
	"goGinTem/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)
		// UserRouter.POST("register", controller.Resgeter)
	}
}
