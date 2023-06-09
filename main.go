package main

import (
	"github.com/VINDA-98/Fasthop/bootstrap"
	"github.com/VINDA-98/Fasthop/global"
	"log"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			err := db.Close()
			if err != nil {
				log.Println("")
				return
			}
		}
		if global.App.Redis != nil {
			err := global.App.Redis.Close()
			if err != nil {
				return
			}
		}
	}()

	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 初始化文件系统
	bootstrap.InitializeStorage()

	// 初始化计划任务
	bootstrap.InitializeCron()

	// 启动服务器
	bootstrap.RunServer()
}
