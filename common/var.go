package common

import (
	"gin-vben-admin/conf"
	"github.com/casbin/casbin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)
var (
	CONFIG  conf.Config
	VP     *viper.Viper
	LOG    *zap.Logger
	DB     *gorm.DB
	CASBIN *casbin.Enforcer
	CACHE   *redis.Client
)

