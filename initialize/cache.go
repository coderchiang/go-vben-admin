package initialize

import (
	"gin-vben-admin/common"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)


func InitCache(){
	switch common.CONFIG.System.CacheType {
	case "redis":
		InitRedis()
	default:
		InitRedis()
	}

}

func InitRedis(){
	r:=common.CONFIG.Redis
	common.CACHE = redis.NewClient(&redis.Options{
		Addr:    r.Addr, // use default Addr
		Password: r.Password,               // no password set
		DB:       r.DB,                // use default DB
	})

	pong, err := common.CACHE .Ping().Result()
	if err != nil {
		common.LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		common.LOG.Info("redis connect ping response:", zap.String("pong",pong))
	}

}
