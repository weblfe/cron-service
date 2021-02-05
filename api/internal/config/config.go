package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	RedisUri string `json:"redis_uri"`
	RedisType string `json:"redis_type"`
	DatabaseUri string `json:"database_uri"`
}
