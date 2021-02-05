package svc

import (
	"cron-service/api/internal/config"
	"github.com/robfig/cron"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

var c = cron.New()

func getCronApp() *cron.Cron  {

	return c
}

func getRedisClient(config config.Config) *redis.Redis {
	return redis.NewRedis(config.RedisUri,redis.ClusterType)
}