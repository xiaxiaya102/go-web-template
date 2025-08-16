# Go Web Admin - 后台管理框架

一个基于 Go + Gin + GORM + JWT + RBAC 的现代化后台管理框架，提供完整的用户权限管理功能。

## ✨ 特性

- 🚀 **现代化技术栈**: Go 1.23 + Gin + GORM v2 + JWT + Redis
- 🔐 **完整的认证授权**: JWT Token + 刷新Token机制
- 👥 **RBAC权限模型**: 用户-角色-权限三层权限控制
- 🗄️ **多数据库支持**: MySQL、PostgreSQL、SQLite
- 📚 **API文档**: 集成 Swagger 自动生成API文档
- 🔧 **配置管理**: 支持YAML配置文件热更新
- 📝 **日志系统**: 结构化日志记录
- 🎯 **标准化响应**: 统一的API响应格式
- 🔒 **密码加密**: bcrypt 密码加密存储

## 📁 项目结构

```
go-web-template/
├── api/                    # API接口层
│   └── v1/
│       └── system/         # 系统管理相关接口
├── config/                 # 配置文件
├── database/              # 数据库连接和初始化
├── global/                # 全局变量
├── initialization/        # 系统初始化
├── logger/               # 日志系统
├── middleware/           # 中间件
├── pkg/                  # 核心包
│   ├── dto/             # 数据传输对象
│   ├── handle/          # 响应处理
│   ├── model/           # 数据模型
│   └── utils/           # 工具函数
├── router/              # 路由配置
├── scripts/             # 数据库脚本
├── docs/               # API文档
├── main.go             # 程序入口
├── go.mod              # Go模块文件
└── README.md           # 项目说明
```

## 🚀 快速开始

### 环境要求

- Go 1.23+
- MySQL 8.0+ / PostgreSQL 12+ / SQLite 3+
- Redis 6.0+ (可选)

### 安装部署

1. **克隆项目**
```bash
git clone <repository-url>
cd go-web-template
```

2. **安装依赖**
```bash
go mod tidy
```

3. **配置数据库**

编辑 `config/application-web.yaml` 文件：

```yaml
# MySQL配置示例
database:
  type: mysql
  host: localhost
  port: 3306
  username: root
  password: your_password
  database: go_web_admin
  ssl_mode: disable

# PostgreSQL配置示例
database:
  type: postgres
  host: localhost
  port: 5432
  username: postgres
  password: your_password
  database: go_web_admin
  ssl_mode: disable

# SQLite配置示例
database:
  type: sqlite
  path: data/app.db
```

4. **初始化数据库**

执行对应的SQL脚本：
- MySQL: `scripts/mysql.sql`
- PostgreSQL: `scripts/postgresql.sql`

5. **配置JWT密钥**

修改 `config/application-web.yaml` 中的JWT配置：
```yaml
jwt:
  secret: "your-secret-key-change-in-production"
  expire: 24h
  refresh_ttl: 168h
```

6. **启动服务**
```bash
go run main.go
```

服务将在 `http://localhost:8089` 启动

### 默认账号

- 用户名: `admin`
- 密码: `admin123`

## 📖 API文档

启动服务后，访问 `http://localhost:8089/swagger/index.html` 查看完整的API文档。

### 核心API接口

#### 认证相关
- `POST /api/login` - 用户登录
- `POST /api/refresh` - 刷新Token
- `POST /api/logout` - 用户登出

#### 用户管理
- `GET /api/user` - 获取用户列表
- `POST /api/user` - 创建用户
- `GET /api/user/{id}` - 获取用户详情
- `PUT /api/user/{id}` - 更新用户
- `DELETE /api/user/{id}` - 删除用户
- `POST /api/user/change-password` - 修改密码

### 请求示例

#### 登录
```bash
curl -X POST http://localhost:8089/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "userName": "admin",
    "password": "admin123"
  }'
```

#### 获取用户列表
```bash
curl -X GET http://localhost:8089/api/user \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🔧 配置说明

### 主配置文件 `config/application-web.yaml`

```yaml
# 服务器配置
server:
  port: 8089              # 服务端口
  mode: release           # 运行模式: debug, release
  timeout: 30s            # 请求超时时间

# 日志配置
log:
  path: logs/app          # 日志文件路径
  level: INFO             # 日志级别: DEBUG, INFO, WARN, ERROR

# 数据库配置
database:
  type: mysql             # 数据库类型: mysql, postgres, sqlite
  host: localhost         # 数据库主机
  port: 3306             # 数据库端口
  username: root          # 用户名
  password: ""            # 密码
  database: go_web_admin  # 数据库名
  ssl_mode: disable       # SSL模式

# Redis配置
redis:
  host: localhost         # Redis主机
  port: 6379             # Redis端口
  password: ""            # Redis密码
  database: 0             # Redis数据库

# JWT配置
jwt:
  secret: "your-secret-key"  # JWT密钥
  expire: 24h                # Token过期时间
  refresh_ttl: 168h          # 刷新Token过期时间
```

## 🏗️ 开发指南

### 添加新的API接口

1. **定义DTO结构**
在 `pkg/dto/` 目录下定义请求和响应结构体

2. **创建API处理函数**
在 `api/v1/system/` 目录下创建对应的API文件

3. **添加路由**
在 `router/system/` 目录下添加路由配置

4. **更新Swagger注释**
使用标准的Swagger注释格式

### 数据库模型

所有数据库模型定义在 `pkg/model/` 目录下，使用GORM标签进行字段映射。

### 中间件

- `middleware/auth.go` - JWT认证中间件
- `middleware/block.go` - 请求阻断中间件

### 工具函数

- `pkg/utils/jwt.go` - JWT工具函数
- `pkg/utils/password.go` - 密码加密工具

## 🔒 安全建议

1. **修改默认密码**: 部署前务必修改默认管理员密码
2. **更换JWT密钥**: 使用强随机字符串作为JWT密钥
3. **HTTPS部署**: 生产环境建议使用HTTPS
4. **数据库安全**: 配置数据库访问权限和防火墙
5. **定期更新**: 保持依赖包的及时更新

## 📝 更新日志

### v1.0.0
- 初始版本发布
- 完整的RBAC权限系统
- JWT认证机制
- 多数据库支持
- Swagger API文档

## 🤝 贡献

欢迎提交Issue和Pull Request来帮助改进项目。

## 📄 许可证

本项目采用 MIT 许可证，详情请查看 [LICENSE](LICENSE) 文件。

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 提交 [Issue](../../issues)
- 发送邮件至: [your-email@example.com]

---

⭐ 如果这个项目对你有帮助，请给个Star支持一下！
