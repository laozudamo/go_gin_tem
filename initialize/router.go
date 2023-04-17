package initialize

import (
	"goGinTem/middlewares"
	"goGinTem/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	//路由中间件
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	ApiGroup := Router.Group("/v1/")

	// 设置跨域中间件
	Router.Use(middlewares.Cors())

	// 图片路径 大小限制
	Router.MaxMultipartMemory = 8 << 20 // 8 MiB
	Router.Static("/uploads", "./uploads")

	//路由分组
	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	router.InitTopicRouter(ApiGroup)
	router.InitUpload(ApiGroup)

	return Router
}
