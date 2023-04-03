// package router

// import (
// 	"goGinTem/controller"
// 	"goGinTem/docs"

// 	"github.com/gin-gonic/gin"
// 	swaggerFiles "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// )

// func UseRoute(Router *gin.RouterGroup) {
// 	docs.SwaggerInfo.BasePath = "/v1"
// 	UserRouter(Router)
// 	UseBasic(Router)
// }

// func UserRouter(Router *gin.RouterGroup) {
// 	r := Router.Group("user")

// 	{
// 		r.GET("list", func(context *gin.Context) {
// 			context.JSON(200, gin.H{
// 				"message": "pong",
// 			})
// 		})

// 		r.POST("login", controller.PasswordLogin)

// 		r.GET("index", controller.GetIndex)
// 	}
// }

// func UseBasic(Router *gin.RouterGroup) {
// 	r := Router
// 	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// }
package router

import (
	"goGinTem/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login", controller.PasswordLogin)
		// UserRouter.POST("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), controller.GetUserList)
		// UserRouter.POST("uploadUserHeaderImage", controller.PutHeaderImage)
	}
}
