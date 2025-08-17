package main

import (
	"fmt"
	"go-web-template/config"
	"go-web-template/database"
	"go-web-template/global"
	"go-web-template/logger"
	"go-web-template/pkg/model"
)

func main() {
	fmt.Println("=== GORM日志功能演示 ===")

	// 初始化配置和数据库
	config.LoadConfig()
	logger.InitLogging(global.System.Log.Path, "gorm_demo.log", "DEBUG") // 设置为DEBUG级别
	database.InitDB()

	fmt.Println("\n1. 基础查询 - 会显示SQL语句")
	var users []model.User
	result := global.DB.Find(&users)
	fmt.Printf("查询结果: 找到 %d 个用户\n", result.RowsAffected)

	fmt.Println("\n2. 条件查询 - 会显示带参数的SQL")
	var user model.User
	global.DB.Where("username = ?", "admin").First(&user)
	fmt.Printf("找到用户: %s\n", user.Username)

	fmt.Println("\n3. 复杂查询 - 会显示JOIN等复杂SQL")
	var userWithRoles []struct {
		model.User
		RoleName string `json:"role_name"`
	}
	global.DB.Table("users").
		Select("users.*, roles.name as role_name").
		Joins("LEFT JOIN user_roles ON users.id = user_roles.user_id").
		Joins("LEFT JOIN roles ON user_roles.role_id = roles.id").
		Where("users.status = ?", 1).
		Find(&userWithRoles)

	fmt.Println("\n4. 慢查询演示 - 会触发慢查询警告")
	// 这个查询会被标记为慢查询（如果超过1秒）
	global.DB.Raw("SELECT SLEEP(0.1)").Scan(&struct{}{})

	fmt.Println("\n5. 错误查询 - 会显示错误日志")
	global.DB.Where("non_existent_column = ?", "value").Find(&users)

	fmt.Println("\n6. 事务操作 - 会显示BEGIN/COMMIT/ROLLBACK")
	tx := global.DB.Begin()
	tx.Create(&model.User{
		UserID:   "test_user",
		Username: "test",
		Password: "test_password",
		Status:   1,
	})
	tx.Rollback() // 回滚，不实际创建

	fmt.Println("\n=== 演示完成 ===")
}
