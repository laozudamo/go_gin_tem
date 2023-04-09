package router

import (
	"goGinTem/controller"
	"goGinTem/middlewares"

	"github.com/gin-gonic/gin"
)

func InitTopicRouter(Router *gin.RouterGroup) {
	Router.POST("topic", middlewares.JWTAuth(), controller.CreatTopic)
	Router.POST("reviewTopic", middlewares.JWTAuth(), controller.ReviewTopic)
}
