package initialize

import (
	"goGinTem/config"
	"goGinTem/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()

	// 文件的路径如何设置
	v.SetConfigFile("./settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	serverConfig := config.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	global.Settings = serverConfig

}
