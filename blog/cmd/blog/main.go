package main

import (
	"chalet/blog/internal/config"
	"chalet/blog/server"
	"chalet/infra/mysql"
	"chalet/infra/redis"
)

func init() {
	config.Init()
	redis.Init(config.AppConfig.Redis)
	mysql.Init(config.AppConfig.Mysql)
}

func main() {
	server.Run()
}
