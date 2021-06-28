package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"kitutils/base/config"
	"kitutils/base/tool"
	"time"
)

func initRedis() {
	if len(config.GetRedisConfig().GetIP())==0{
		tool.NewLogger().Warn("None Redis Init")
		return
	}
	redisDb = redis.NewClient(
		&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()),
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			Password:     config.GetRedisConfig().GetPass(),
			PoolSize:     config.GetRedisConfig().GetMaxOpen(),
		},
	)
	err = redisDb.Ping().Err()
	if nil != err {
		tool.NewLogger().Error("ping redis err:", zap.Error(err))
		panic(err)
	}
	tool.NewLogger().Debug("redis success : " + fmt.Sprintf("%s:%s", config.GetRedisConfig().GetIP(), config.GetRedisConfig().GetPort()))

}

func closeRedis() {
	if redisDb != nil {
		_ = redisDb.Close()
	}
}
