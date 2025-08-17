# MySQL数据库启动问题修复指南

## 问题描述

在使用MySQL数据库启动项目时，遇到以下错误：
```
Error 1170 (42000): BLOB/TEXT column 'type' used in key specification without a key length
```

## 问题原因

这个错误是因为MySQL对TEXT/LONGTEXT字段作为索引时需要指定长度。在GORM中，当字段类型为`string`且没有明确指定数据库字段类型时，GORM会将其映射为`longtext`类型，而MySQL要求TEXT类型字段在创建索引时必须指定长度。

## 解决方案

### 1. 修改模型定义

在所有模型的字符串字段上明确指定MySQL字段类型：

```go
// 修改前
Type string `gorm:"not null" json:"type"`

// 修改后  
Type string `gorm:"type:varchar(20);not null" json:"type"`
```

### 2. 已修复的字段

以下字段已经修复：

**User模型**:
- UserID: `varchar(50)`
- Username: `varchar(50)`
- Password: `varchar(255)`
- Email: `varchar(100)`
- Phone: `varchar(20)`
- Avatar: `varchar(255)`
- Remark: `text`

**Role模型**:
- Name: `varchar(50)`
- Code: `varchar(50)`
- Description: `text`

**Permission模型**:
- Name: `varchar(50)`
- Code: `varchar(50)`
- Type: `varchar(20)` ← 主要问题字段
- Path: `varchar(255)`
- Component: `varchar(255)`
- Icon: `varchar(50)`
- Description: `text`

**Menu模型**:
- Name: `varchar(50)`
- Path: `varchar(255)`
- Component: `varchar(255)`
- Icon: `varchar(50)`
- Remark: `text`

### 3. 数据库处理方式

#### 方式一：删除现有表（推荐）
如果是开发环境且数据不重要，可以删除现有表让GORM重新创建：

```sql
DROP TABLE IF EXISTS `permissions`;
DROP TABLE IF EXISTS `role_permissions`;
DROP TABLE IF EXISTS `user_roles`;
DROP TABLE IF EXISTS `menus`;
DROP TABLE IF EXISTS `roles`;
DROP TABLE IF EXISTS `users`;
```

#### 方式二：修改现有表结构
如果需要保留数据，可以执行 `scripts/fix_mysql_schema.sql` 脚本来修改现有表结构。

### 4. 启动项目

修复后，重新启动项目：
```bash
go run main.go
```

## 预防措施

1. **明确指定字段类型**: 在定义GORM模型时，对于字符串字段明确指定数据库字段类型
2. **使用合适的长度**: 根据实际需求设置合适的字段长度
3. **测试多种数据库**: 在不同数据库上测试模型定义的兼容性

## 相关文档

- [GORM字段标签文档](https://gorm.io/docs/models.html#Fields-Tags)
- [MySQL数据类型文档](https://dev.mysql.com/doc/refman/8.0/en/data-types.html)
