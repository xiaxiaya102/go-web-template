# 操作日志系统使用指南

## 📖 概述

操作日志系统用于记录用户在系统中的所有操作行为，包括登录、查询、新增、修改、删除等操作。这对于系统审计、安全监控、问题排查等场景非常重要。

## 🗃️ 数据库表结构

### OperationLog 操作日志表

| 字段名 | 类型 | 说明 | 索引 |
|--------|------|------|------|
| id | uint | 主键ID | PRIMARY |
| user_id | uint | 操作用户ID | INDEX |
| username | varchar(50) | 操作用户名 | - |
| module | varchar(50) | 操作模块 | - |
| operation | varchar(100) | 操作类型 | - |
| method | varchar(10) | 请求方法 | - |
| url | varchar(500) | 请求URL | - |
| ip | varchar(45) | 操作IP | - |
| user_agent | varchar(500) | 用户代理 | - |
| request_body | text | 请求体 | - |
| response_body | text | 响应体 | - |
| status | int | 响应状态码 | - |
| duration | bigint | 执行时间(毫秒) | - |
| error_message | text | 错误信息 | - |
| operation_time | datetime | 操作时间 | - |
| description | varchar(500) | 操作描述 | - |
| created_at | datetime | 创建时间 | - |
| updated_at | datetime | 更新时间 | - |

## 🔧 功能特性

### 1. 自动记录
- ✅ **请求信息**: 自动记录HTTP方法、URL、请求体
- ✅ **响应信息**: 自动记录响应状态码、响应体、执行时间
- ✅ **用户信息**: 自动记录操作用户ID、用户名
- ✅ **环境信息**: 自动记录客户端IP、User-Agent
- ✅ **操作分类**: 自动识别操作模块和操作类型

### 2. 安全保护
- 🔒 **敏感数据过滤**: 自动过滤密码、token等敏感信息
- 📏 **数据长度限制**: 限制请求体和响应体长度，避免日志过大
- 🚀 **异步保存**: 异步保存日志，不影响业务响应性能

### 3. 智能分析
- 📊 **操作分类**: 根据HTTP方法和URL自动分类操作类型
- 🏷️ **模块识别**: 根据URL路径自动识别操作模块
- 📝 **描述生成**: 自动生成操作描述

## 📊 API接口

### 1. 获取操作日志列表
```http
GET /api/operation-logs?page=1&pageSize=10&username=admin&module=user
```

**查询参数**:
- `page`: 页码 (默认: 1)
- `pageSize`: 每页数量 (默认: 10)
- `username`: 用户名筛选
- `module`: 模块筛选
- `operation`: 操作类型筛选
- `ip`: IP地址筛选
- `startTime`: 开始时间 (格式: 2006-01-02)
- `endTime`: 结束时间 (格式: 2006-01-02)

### 2. 获取操作日志详情
```http
GET /api/operation-logs/{id}
```

### 3. 删除操作日志
```http
DELETE /api/operation-logs/{id}
```

### 4. 批量删除操作日志
```http
POST /api/operation-logs/batch-delete
Content-Type: application/json

{
  "ids": [1, 2, 3]
}
```

### 5. 清空操作日志
```http
POST /api/operation-logs/clear
Content-Type: application/json

{
  "startTime": "2025-01-01",
  "endTime": "2025-01-31"
}
```

或者按天数清空:
```json
{
  "days": 30
}
```

### 6. 获取操作日志统计
```http
GET /api/operation-logs/stats?days=7
```

## 🎯 操作类型映射

| HTTP方法 | URL特征 | 操作类型 |
|----------|---------|----------|
| GET | /list, /page | 查询列表 |
| GET | 其他 | 查询详情 |
| POST | /login | 用户登录 |
| POST | 其他 | 新增 |
| PUT/PATCH | - | 修改 |
| DELETE | - | 删除 |

## 🏷️ 模块识别规则

根据URL路径自动识别模块:
- `/api/user/*` → user模块
- `/api/role/*` → role模块
- `/api/permission/*` → permission模块
- `/api/operation-logs/*` → operation-logs模块

## 🔒 安全特性

### 1. 敏感数据过滤
自动过滤以下敏感字段:
- `password`
- `token`
- `secret`
- `key`

### 2. 数据长度限制
- 请求体和响应体最大长度: 2000字符
- 超出部分会被截断并添加"..."标识

### 3. IP获取优先级
1. `X-Forwarded-For` 头部
2. `X-Real-IP` 头部
3. `RemoteAddr`

## 📈 统计功能

### 1. 总体统计
- 指定时间范围内的总操作次数

### 2. 模块统计
- 按模块统计操作次数
- 按操作次数降序排列

### 3. 操作类型统计
- 按操作类型统计次数
- 按操作次数降序排列

### 4. 用户统计
- 按用户统计操作次数
- 排除匿名用户
- 显示前10名活跃用户

### 5. 日期统计
- 按日期统计操作次数
- 用于生成操作趋势图

## 🛠️ 配置说明

### 1. 启用操作日志
操作日志中间件已自动添加到需要认证的路由组中:

```go
PrivateGroup.Use(middleware.OperationLogMiddleware())
```

### 2. 自定义配置
可以在中间件中自定义以下配置:
- 敏感字段列表
- 数据长度限制
- 模块识别规则
- 操作类型映射

### 3. 性能优化
- 使用异步保存，不影响业务响应
- 可以考虑使用消息队列进一步优化
- 定期清理历史日志数据

## 📋 最佳实践

### 1. 日志保留策略
- 建议保留3-6个月的操作日志
- 定期清理过期日志，避免数据库过大
- 重要操作日志可以备份到其他存储

### 2. 监控告警
- 监控异常操作（如大量删除操作）
- 监控登录失败次数
- 监控来自异常IP的操作

### 3. 权限控制
- 只有管理员可以查看操作日志
- 普通用户只能查看自己的操作记录
- 敏感操作需要额外审批

### 4. 数据分析
- 定期分析用户操作习惯
- 识别系统使用热点
- 优化用户体验

## 🔍 故障排查

### 1. 日志记录失败
检查以下问题:
- 数据库连接是否正常
- 操作日志表是否存在
- 磁盘空间是否充足

### 2. 性能问题
- 检查日志表大小
- 考虑添加索引
- 定期清理历史数据

### 3. 数据不准确
- 检查中间件是否正确配置
- 验证用户信息获取逻辑
- 检查IP获取逻辑

## 🚀 扩展功能

### 1. 实时监控
- 集成WebSocket实现实时日志推送
- 添加操作日志仪表板

### 2. 高级分析
- 用户行为分析
- 操作路径分析
- 异常检测

### 3. 集成第三方
- 集成ELK进行日志分析
- 集成监控系统进行告警
- 集成审计系统进行合规检查
