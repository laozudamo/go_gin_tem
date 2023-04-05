package router

import (
	"goGinTem/controller"
	"goGinTem/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)
		UserRouter.GET("getUserInfo", middlewares.JWTAuth(), controller.GetUserInfo)
		UserRouter.POST("updateUserInfo", middlewares.JWTAuth(), controller.UpdateUserInfo)
	}
}
