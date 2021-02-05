package svc

import (
	"cron-service/api/internal/config"
	"github.com/robfig/cron"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Cron   *cron.Cron
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cron: getCronApp(),
		Redis: getRedisClient(c),
	}
}
