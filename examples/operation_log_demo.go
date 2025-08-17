package main

import (
	"fmt"
	"go-web-template/config"
	"go-web-template/database"
	"go-web-template/global"
	"go-web-template/logger"
	"go-web-template/pkg/model"
	"time"
)

func main() {
	fmt.Println("=== 操作日志系统演示 ===")

	// 初始化配置和数据库
	config.LoadConfig()
	logger.InitLogging(global.System.Log.Path, "operation_log_demo.log", "DEBUG")
	database.InitDB()

	fmt.Println("\n1. 创建测试操作日志")
	createTestLogs()

	fmt.Println("\n2. 查询操作日志")
	queryLogs()

	fmt.Println("\n3. 统计操作日志")
	statsLogs()

	fmt.Println("\n=== 演示完成 ===")
}

// createTestLogs 创建测试操作日志
func createTestLogs() {
	testLogs := []model.OperationLog{
		{
			UserID:        1,
			Username:      "admin",
			Module:        "user",
			Operation:     "查询列表",
			Method:        "GET",
			URL:           "/api/users?page=1&pageSize=10",
			IP:            "192.168.1.100",
			UserAgent:     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
			RequestBody:   "",
			ResponseBody:  `{"errCode":0,"errMsg":"success","data":{"list":[],"total":0}}`,
			Status:        200,
			Duration:      45,
			ErrorMessage:  "",
			OperationTime: time.Now().Format("2006-01-02 15:04:05"),
			Description:   "查询用户列表",
		},
		{
			UserID:        1,
			Username:      "admin",
			Module:        "user",
			Operation:     "新增",
			Method:        "POST",
			URL:           "/api/users",
			IP:            "192.168.1.100",
			UserAgent:     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
			RequestBody:   `{"username":"testuser","password":"***","email":"test@example.com"}`,
			ResponseBody:  `{"errCode":0,"errMsg":"success","data":{"id":2}}`,
			Status:        200,
			Duration:      120,
			ErrorMessage:  "",
			OperationTime: time.Now().Add(-1 * time.Hour).Format("2006-01-02 15:04:05"),
			Description:   "新增用户数据",
		},
		{
			UserID:        1,
			Username:      "admin",
			Module:        "role",
			Operation:     "修改",
			Method:        "PUT",
			URL:           "/api/roles/1",
			IP:            "192.168.1.100",
			UserAgent:     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
			RequestBody:   `{"name":"管理员","description":"系统管理员角色"}`,
			ResponseBody:  `{"errCode":0,"errMsg":"success"}`,
			Status:        200,
			Duration:      80,
			ErrorMessage:  "",
			OperationTime: time.Now().Add(-2 * time.Hour).Format("2006-01-02 15:04:05"),
			Description:   "修改角色数据",
		},
		{
			UserID:        2,
			Username:      "testuser",
			Module:        "system",
			Operation:     "用户登录",
			Method:        "POST",
			URL:           "/api/login",
			IP:            "192.168.1.101",
			UserAgent:     "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36",
			RequestBody:   `{"userName":"testuser","password":"***"}`,
			ResponseBody:  `{"errCode":0,"errMsg":"success","data":{"token":"eyJ..."}}`,
			Status:        200,
			Duration:      200,
			ErrorMessage:  "",
			OperationTime: time.Now().Add(-3 * time.Hour).Format("2006-01-02 15:04:05"),
			Description:   "用户登录",
		},
		{
			UserID:        0,
			Username:      "anonymous",
			Module:        "system",
			Operation:     "用户登录",
			Method:        "POST",
			URL:           "/api/login",
			IP:            "192.168.1.102",
			UserAgent:     "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15",
			RequestBody:   `{"userName":"wronguser","password":"***"}`,
			ResponseBody:  `{"errCode":-1,"errMsg":"用户名或密码错误"}`,
			Status:        401,
			Duration:      50,
			ErrorMessage:  "用户名或密码错误",
			OperationTime: time.Now().Add(-4 * time.Hour).Format("2006-01-02 15:04:05"),
			Description:   "用户登录失败",
		},
	}

	for i, log := range testLogs {
		if err := global.DB.Create(&log).Error; err != nil {
			fmt.Printf("创建测试日志 %d 失败: %v\n", i+1, err)
		} else {
			fmt.Printf("✅ 创建测试日志 %d: %s %s %s\n", i+1, log.Username, log.Operation, log.Module)
		}
	}
}

// queryLogs 查询操作日志
func queryLogs() {
	// 查询最近的操作日志
	var logs []model.OperationLog
	if err := global.DB.Order("created_at DESC").Limit(5).Find(&logs).Error; err != nil {
		fmt.Printf("查询操作日志失败: %v\n", err)
		return
	}

	fmt.Printf("最近5条操作日志:\n")
	for _, log := range logs {
		fmt.Printf("- [%s] %s %s %s (%dms) %s\n",
			log.OperationTime, log.Username, log.Method, log.Module, log.Duration, log.Description)
	}

	// 按用户查询
	var userLogs []model.OperationLog
	if err := global.DB.Where("username = ?", "admin").Order("created_at DESC").Limit(3).Find(&userLogs).Error; err != nil {
		fmt.Printf("查询用户日志失败: %v\n", err)
		return
	}

	fmt.Printf("\nadmin用户的操作日志:\n")
	for _, log := range userLogs {
		fmt.Printf("- [%s] %s %s (%d) %s\n",
			log.OperationTime, log.Method, log.URL, log.Status, log.Description)
	}
}

// statsLogs 统计操作日志
func statsLogs() {
	// 总数统计
	var totalCount int64
	global.DB.Model(&model.OperationLog{}).Count(&totalCount)
	fmt.Printf("操作日志总数: %d\n", totalCount)

	// 按模块统计
	type ModuleStats struct {
		Module string `json:"module"`
		Count  int64  `json:"count"`
	}
	var moduleStats []ModuleStats
	global.DB.Model(&model.OperationLog{}).
		Select("module, COUNT(*) as count").
		Group("module").
		Order("count DESC").
		Scan(&moduleStats)

	fmt.Printf("\n按模块统计:\n")
	for _, stat := range moduleStats {
		fmt.Printf("- %s: %d次\n", stat.Module, stat.Count)
	}

	// 按操作类型统计
	type OperationStats struct {
		Operation string `json:"operation"`
		Count     int64  `json:"count"`
	}
	var operationStats []OperationStats
	global.DB.Model(&model.OperationLog{}).
		Select("operation, COUNT(*) as count").
		Group("operation").
		Order("count DESC").
		Scan(&operationStats)

	fmt.Printf("\n按操作类型统计:\n")
	for _, stat := range operationStats {
		fmt.Printf("- %s: %d次\n", stat.Operation, stat.Count)
	}

	// 按用户统计
	type UserStats struct {
		Username string `json:"username"`
		Count    int64  `json:"count"`
	}
	var userStats []UserStats
	global.DB.Model(&model.OperationLog{}).
		Select("username, COUNT(*) as count").
		Where("username != 'anonymous'").
		Group("username").
		Order("count DESC").
		Scan(&userStats)

	fmt.Printf("\n按用户统计:\n")
	for _, stat := range userStats {
		fmt.Printf("- %s: %d次\n", stat.Username, stat.Count)
	}

	// 错误统计
	var errorCount int64
	global.DB.Model(&model.OperationLog{}).Where("status >= 400").Count(&errorCount)
	fmt.Printf("\n错误操作数: %d\n", errorCount)

	// 平均响应时间
	type AvgDuration struct {
		AvgDuration float64 `json:"avg_duration"`
	}
	var avgDuration AvgDuration
	global.DB.Model(&model.OperationLog{}).
		Select("AVG(duration) as avg_duration").
		Scan(&avgDuration)
	fmt.Printf("平均响应时间: %.2fms\n", avgDuration.AvgDuration)
}
