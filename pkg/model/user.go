package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserID     string    `gorm:"type:varchar(50);unique;not null" json:"user_id"`  // 用户ID
	Username   string    `gorm:"type:varchar(50);unique;not null" json:"username"` // 用户名
	Password   string    `gorm:"type:varchar(255);not null" json:"-"`              // 密码，不返回给前端
	Email      string    `gorm:"type:varchar(100);unique" json:"email"`            // 邮箱
	Phone      string    `gorm:"type:varchar(20);unique" json:"phone"`             // 手机号
	Avatar     string    `gorm:"type:varchar(255)" json:"avatar"`                  // 头像
	Status     int       `gorm:"default:1" json:"status"`                          // 状态：0-禁用，1-启用
	LastLogin  time.Time `json:"last_login"`                                       // 最后登录时间
	LoginCount int       `gorm:"default:0" json:"login_count"`                     // 登录次数
	Remark     string    `gorm:"type:text" json:"remark"`                          // 备注
}

// Role 角色模型
type Role struct {
	BaseModel
	Name        string `gorm:"type:varchar(50);unique;not null" json:"name"` // 角色名称
	Code        string `gorm:"type:varchar(50);unique;not null" json:"code"` // 角色编码
	Description string `gorm:"type:text" json:"description"`                 // 角色描述
	Status      int    `gorm:"default:1" json:"status"`                      // 状态：0-禁用，1-启用
	Sort        int    `gorm:"default:0" json:"sort"`                        // 排序
}

// Permission 权限模型
type Permission struct {
	BaseModel
	Name        string `gorm:"type:varchar(50);unique;not null" json:"name"` // 权限名称
	Code        string `gorm:"type:varchar(50);unique;not null" json:"code"` // 权限编码
	Type        string `gorm:"type:varchar(20);not null" json:"type"`        // 权限类型：menu-菜单，button-按钮，api-接口
	ParentID    uint   `json:"parent_id"`                                    // 父级ID
	Path        string `gorm:"type:varchar(255)" json:"path"`                // 路由路径
	Component   string `gorm:"type:varchar(255)" json:"component"`           // 组件路径
	Icon        string `gorm:"type:varchar(50)" json:"icon"`                 // 图标
	Sort        int    `gorm:"default:0" json:"sort"`                        // 排序
	Status      int    `gorm:"default:1" json:"status"`                      // 状态：0-禁用，1-启用
	Description string `gorm:"type:text" json:"description"`                 // 权限描述
}

// UserRole 用户角色关联表
type UserRole struct {
	BaseModel
	UserID uint `gorm:"not null" json:"user_id"` // 用户ID
	RoleID uint `gorm:"not null" json:"role_id"` // 角色ID
}

// RolePermission 角色权限关联表
type RolePermission struct {
	BaseModel
	RoleID       uint `gorm:"not null" json:"role_id"`       // 角色ID
	PermissionID uint `gorm:"not null" json:"permission_id"` // 权限ID
}

// Menu 菜单模型（权限的一种特殊类型）
type Menu struct {
	BaseModel
	Name      string `gorm:"type:varchar(50);not null" json:"name"` // 菜单名称
	ParentID  uint   `json:"parent_id"`                             // 父级菜单ID
	Path      string `gorm:"type:varchar(255)" json:"path"`         // 路由路径
	Component string `gorm:"type:varchar(255)" json:"component"`    // 组件路径
	Icon      string `gorm:"type:varchar(50)" json:"icon"`          // 图标
	Sort      int    `gorm:"default:0" json:"sort"`                 // 排序
	Status    int    `gorm:"default:1" json:"status"`               // 状态：0-隐藏，1-显示
	Type      int    `gorm:"default:1" json:"type"`                 // 类型：1-菜单，2-按钮
	Remark    string `gorm:"type:text" json:"remark"`               // 备注
}

// OperationLog 操作日志模型
type OperationLog struct {
	BaseModel
	UserID        uint   `gorm:"index;not null" json:"user_id"`                // 操作用户ID
	Username      string `gorm:"type:varchar(50);not null" json:"username"`    // 操作用户名
	Module        string `gorm:"type:varchar(50);not null" json:"module"`      // 操作模块
	Operation     string `gorm:"type:varchar(100);not null" json:"operation"`  // 操作类型
	Method        string `gorm:"type:varchar(10);not null" json:"method"`      // 请求方法
	URL           string `gorm:"type:varchar(500);not null" json:"url"`        // 请求URL
	IP            string `gorm:"type:varchar(45);not null" json:"ip"`          // 操作IP
	UserAgent     string `gorm:"type:varchar(500)" json:"user_agent"`          // 用户代理
	RequestBody   string `gorm:"type:text" json:"request_body"`                // 请求体
	ResponseBody  string `gorm:"type:text" json:"response_body"`               // 响应体
	Status        int    `gorm:"not null" json:"status"`                       // 响应状态码
	Duration      int64  `gorm:"not null" json:"duration"`                     // 执行时间(毫秒)
	ErrorMessage  string `gorm:"type:text" json:"error_message"`               // 错误信息
	OperationTime string `gorm:"type:datetime;not null" json:"operation_time"` // 操作时间
	Description   string `gorm:"type:varchar(500)" json:"description"`         // 操作描述
}
