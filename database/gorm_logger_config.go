package database

import (
	"context"
	"fmt"
	"go-web-template/global"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// GetGormConfig 获取GORM配置，包含日志设置
func GetGormConfig() *gorm.Config {
	// 根据环境和配置决定日志级别
	logLevel := getLogLevel()

	// 创建自定义日志器
	newLogger := gormLogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second, // 慢查询阈值
			LogLevel:                  logLevel,    // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound错误
			Colorful:                  true,        // 彩色输出
		},
	)

	return &gorm.Config{
		Logger: newLogger,
		// 其他配置...
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
}

// getLogLevel 根据配置获取日志级别
func getLogLevel() gormLogger.LogLevel {
	// 根据全局配置的日志级别决定GORM日志级别
	switch global.System.Log.Level {
	case "DEBUG":
		return gormLogger.Info // 显示所有SQL
	case "INFO":
		return gormLogger.Warn // 只显示慢查询和错误
	case "WARN":
		return gormLogger.Error // 只显示错误
	case "ERROR":
		return gormLogger.Silent // 静默模式
	default:
		return gormLogger.Warn
	}
}

// 自定义日志器示例 - 可以将SQL日志写入文件
type CustomLogger struct {
	gormLogger.Interface
}

func NewCustomLogger() gormLogger.Interface {
	return &CustomLogger{
		Interface: gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Second,
				LogLevel:      gormLogger.Info,
				Colorful:      true,
			},
		),
	}
}

// Trace 自定义SQL跟踪
func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// 可以在这里添加自定义逻辑，比如写入文件、发送到监控系统等
	sql, rows := fc()
	elapsed := time.Since(begin)

	// 记录到自定义日志
	fmt.Printf("[GORM] [%v] [rows:%d] %s\n", elapsed, rows, sql)

	// 调用原始的Trace方法
	l.Interface.Trace(ctx, begin, fc, err)
}
