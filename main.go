package main

import (
	"fmt"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"

	"go.uber.org/zap"
)

func main() {
	// 1. viper 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("viper init fail, err:%v\n", err)
		return
	}
	// 2. zap 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("zap init fail, err:%v\n", err)
		return
	}
	zap.L().Debug("logger init success...")
	defer zap.L().Sync()
	// 3. gorm 初始化 MySQL 连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("mysql init fail, err:%v\n", err)
		return
	}
	// 4. go-redis 初始化 Redis 连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init fail, err:%v\n", err)
		return
	}
	defer redis.Close()
	// 5. gin 注册路由
	r := routes.Setup()
	// 6. 启动服务
	r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
}
