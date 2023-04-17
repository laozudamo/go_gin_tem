package router

import (
	"goGinTem/controller"
	"goGinTem/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUpload(Router *gin.RouterGroup) {
	Router.POST("upload", middlewares.JWTAuth(), controller.Upload)
}
