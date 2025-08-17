package database

import (
	"fmt"
	"go-web-template/global"
	"go-web-template/logger"
	"go-web-template/pkg/model"
	"golang.org/x/crypto/bcrypt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

const admin = "admin"

func InitDB() {
	var err error
	var dsn string

	dbConfig := global.System.Database

	// 获取GORM配置（包含日志设置）
	gormConfig := GetGormConfig()

	switch dbConfig.Type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
		DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	case "postgres":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Database, dbConfig.Password, dbConfig.SSLMode)
		DB, err = gorm.Open(postgres.Open(dsn), gormConfig)
	case "sqlite":
		DB, err = gorm.Open(sqlite.Dialector{
			DriverName: "sqlite",
			DSN:        dbConfig.Path,
		}, gormConfig)
	default:
		logger.Error("不支持的数据库类型: %s", dbConfig.Type)
		panic("不支持的数据库类型")
	}

	if err != nil {
		logger.Error("数据库连接失败: %v", err)
		panic(err)
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		logger.Error("获取数据库连接失败: %v", err)
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.DB = DB

	// 自动迁移表结构
	err = DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.UserRole{},
		&model.RolePermission{},
		&model.Menu{},
		&model.OperationLog{},
	)

	if err != nil {
		logger.Error("数据库表迁移失败: %v", err)
		panic(err)
	}

	// 创建默认管理员用户
	createDefaultAdmin()

	// 初始化RBAC数据
	InitRBAC()

	logger.Info("数据库初始化完成")
}

// createDefaultAdmin 创建默认管理员用户
func createDefaultAdmin() {
	var count int64
	DB.Model(&model.User{}).Count(&count)

	if count == 0 {
		// 加密默认密码
		hashedPassword, err := hashPassword(global.System.ServerInfo.Password)
		if err != nil {
			logger.Error("密码加密失败: %v", err)
			return
		}

		// 创建默认管理员用户
		adminUser := &model.User{
			UserID:   "admin",
			Username: "admin",
			Password: hashedPassword,
			Status:   1,
		}

		if err := DB.Create(adminUser).Error; err != nil {
			logger.Error("创建默认管理员失败: %v", err)
		} else {
			logger.Info("默认管理员用户创建成功")
		}
	}
}

// hashPassword 加密密码（临时函数，避免循环导入）
func hashPassword(password string) (string, error) {
	// 直接使用bcrypt加密，避免循环导入
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
