package config

import "os"

func initRedis(conf *AppConfig) {
	addr := os.Getenv("REDIS_ADDR")
	pass := os.Getenv("REDIS_PASS")

	conf.Redis.Addr = addr
	conf.Redis.Pass = pass
}
