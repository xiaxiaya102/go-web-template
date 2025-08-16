package initialization

import (
	"go-web-template/config"
	"go-web-template/database"
	"go-web-template/global"
	"go-web-template/logger"
)

func InitBase() {
	// 加载配置
	config.LoadConfig()

	// 日志处理
	logger.InitLogging(global.System.Log.Path, "run.log", global.System.Log.Level)

	// 初始化数据库
	database.InitDB()

	// 初始化Redis
	database.InitRedis()

}
