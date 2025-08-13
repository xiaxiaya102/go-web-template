package initialization

import (
	"feishuReboot/config"
	"feishuReboot/database"
	"feishuReboot/global"
	"feishuReboot/logger"
)

func InitBase() {
	// 加载配置
	config.LoadConfig()

	// 日志处理
	logger.InitLogging(global.System.Log.Path, "run.log", global.System.Log.Level)

	// 初始化sqlite
	database.InitDB()

}
