# GORM日志功能使用指南

## 📖 概述

GORM提供了强大的日志功能，可以帮助开发者调试SQL查询、监控数据库性能、排查问题。本项目已经集成了GORM日志功能。

## 🔧 配置方式

### 1. 日志级别配置

在 `config/application-web.yaml` 中设置日志级别：

```yaml
log:
  path: logs/app
  level: DEBUG  # DEBUG, INFO, WARN, ERROR
```

### 2. GORM日志级别映射

| 应用日志级别 | GORM日志级别 | 显示内容 |
|-------------|-------------|----------|
| DEBUG | Info | 显示所有SQL语句、参数、执行时间 |
| INFO | Warn | 只显示慢查询和错误 |
| WARN | Error | 只显示错误 |
| ERROR | Silent | 静默模式，不显示任何SQL日志 |

## 📊 日志输出示例

### 1. 普通查询日志
```
[2025-08-17 09:15:30] [0.45ms] [rows:1] SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
```

### 2. 慢查询警告
```
[2025-08-17 09:15:31] [1.23s] [rows:100] SLOW SQL >= 1s
SELECT * FROM `users` WHERE status = 1
```

### 3. 错误日志
```
[2025-08-17 09:15:32] [0.12ms] [rows:0] 
Error 1054: Unknown column 'non_existent_column' in 'where clause'
SELECT * FROM `users` WHERE non_existent_column = 'value'
```

### 4. 事务日志
```
[2025-08-17 09:15:33] [0.05ms] [rows:0] BEGIN
[2025-08-17 09:15:33] [0.12ms] [rows:1] INSERT INTO `users` (...)
[2025-08-17 09:15:33] [0.03ms] [rows:0] ROLLBACK
```

## 🛠️ 使用方法

### 1. 开启SQL日志（开发环境）

修改配置文件：
```yaml
log:
  level: DEBUG  # 显示所有SQL
```

### 2. 生产环境配置

```yaml
log:
  level: WARN   # 只显示错误和慢查询
```

### 3. 代码中临时开启日志

```go
// 临时开启详细日志
db := global.DB.Debug()
db.Where("username = ?", "admin").First(&user)

// 或者为单个查询开启
global.DB.Debug().Find(&users)
```

## 📈 性能监控功能

### 1. 慢查询检测

默认配置：超过1秒的查询会被标记为慢查询

```go
// 在 database/gorm_logger_config.go 中配置
SlowThreshold: time.Second,  // 可以调整阈值
```

### 2. 查询统计

每个查询都会显示：
- 执行时间
- 影响行数
- SQL语句和参数

## 🔍 调试技巧

### 1. 查看生成的SQL

```go
// 方法1：使用Debug()
global.DB.Debug().Where("status = ?", 1).Find(&users)

// 方法2：使用ToSQL()（需要额外配置）
sql := global.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
    return tx.Where("status = ?", 1).Find(&users)
})
fmt.Println("Generated SQL:", sql)
```

### 2. 分析查询性能

```go
// 使用EXPLAIN分析查询计划
var result []map[string]interface{}
global.DB.Raw("EXPLAIN SELECT * FROM users WHERE status = ?", 1).Scan(&result)
```

### 3. 监控数据库连接

```go
sqlDB, err := global.DB.DB()
if err == nil {
    stats := sqlDB.Stats()
    fmt.Printf("Open connections: %d\n", stats.OpenConnections)
    fmt.Printf("In use: %d\n", stats.InUse)
    fmt.Printf("Idle: %d\n", stats.Idle)
}
```

## 📝 自定义日志器

### 1. 文件日志

```go
// 将SQL日志写入文件
logFile, _ := os.OpenFile("sql.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
newLogger := gormLogger.New(
    log.New(logFile, "\r\n", log.LstdFlags),
    gormLogger.Config{
        SlowThreshold: time.Second,
        LogLevel:      gormLogger.Info,
        Colorful:      false, // 文件日志不需要颜色
    },
)
```

### 2. 结构化日志

```go
// 集成到现有的日志系统
type StructuredLogger struct {
    gormLogger.Interface
}

func (l *StructuredLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
    sql, rows := fc()
    elapsed := time.Since(begin)
    
    // 使用结构化日志
    logger.Info("SQL executed", map[string]interface{}{
        "sql":      sql,
        "duration": elapsed,
        "rows":     rows,
        "error":    err,
    })
}
```

## 🚀 最佳实践

### 1. 环境区分

```go
func getLogLevel() gormLogger.LogLevel {
    if os.Getenv("ENV") == "production" {
        return gormLogger.Error  // 生产环境只记录错误
    }
    return gormLogger.Info      // 开发环境记录所有SQL
}
```

### 2. 敏感信息过滤

```go
// 过滤敏感参数
type SafeLogger struct {
    gormLogger.Interface
}

func (l *SafeLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
    sql, rows := fc()
    // 过滤密码等敏感信息
    safeSql := strings.ReplaceAll(sql, "password", "***")
    // ... 记录安全的SQL
}
```

### 3. 性能监控集成

```go
// 集成到监控系统
func (l *MonitorLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
    sql, rows := fc()
    elapsed := time.Since(begin)
    
    // 发送指标到监控系统
    metrics.RecordSQLDuration(elapsed)
    if elapsed > time.Second {
        metrics.IncrementSlowQuery()
    }
    if err != nil {
        metrics.IncrementSQLError()
    }
}
```

## 🔧 运行演示

要查看GORM日志功能，可以运行演示程序：

```bash
# 运行GORM日志演示
go run examples/gorm_logging_demo.go

# 或者修改配置文件中的日志级别为DEBUG，然后启动应用
# 所有数据库操作都会显示SQL日志
go run main.go
```

## 📋 常见问题

### Q: 如何只显示慢查询？
A: 设置日志级别为 `WARN`，只有超过阈值的查询和错误会被记录。

### Q: 如何关闭SQL日志？
A: 设置日志级别为 `ERROR` 或在GORM配置中使用 `gormLogger.Silent`。

### Q: 如何自定义慢查询阈值？
A: 在 `database/gorm_logger_config.go` 中修改 `SlowThreshold` 值。

### Q: 生产环境建议的配置？
A: 建议使用 `WARN` 级别，记录慢查询和错误，便于性能监控和问题排查。
