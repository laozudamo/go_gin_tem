package main

import (
	"fmt"
	"goGinTem/global"
	"goGinTem/initialize"

	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()

	initialize.InitLogger()

	initialize.InitMysqlDB()

	// 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// initialize.InitRedis()

	Router := initialize.Routers()
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is hello func", zap.String("error", "启动错误!"))
	}
}
