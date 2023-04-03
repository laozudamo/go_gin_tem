package global

import (
	"goGinTem/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Settings config.ServerConfig
	Lg       *zap.Logger
	DB       *gorm.DB
	Redis    *redis.Client
	Trans    ut.Translator
)
