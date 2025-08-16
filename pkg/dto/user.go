package dto

import "time"

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	UserID   string `json:"user_id" binding:"required"`  // 用户ID
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	Email    string `json:"email"`                       // 邮箱
	Phone    string `json:"phone"`                       // 手机号
	Avatar   string `json:"avatar"`                      // 头像
	Status   int    `json:"status"`                      // 状态：0-禁用，1-启用
	Remark   string `json:"remark"`                      // 备注
	RoleIDs  []uint `json:"role_ids"`                    // 角色ID列表
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	ID       uint   `json:"id" binding:"required"`       // 用户ID
	UserID   string `json:"user_id" binding:"required"`  // 用户ID
	Username string `json:"username" binding:"required"` // 用户名
	Email    string `json:"email"`                       // 邮箱
	Phone    string `json:"phone"`                       // 手机号
	Avatar   string `json:"avatar"`                      // 头像
	Status   int    `json:"status"`                      // 状态：0-禁用，1-启用
	Remark   string `json:"remark"`                      // 备注
	RoleIDs  []uint `json:"role_ids"`                    // 角色ID列表
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"` // 旧密码
	NewPassword string `json:"new_password" binding:"required"` // 新密码
}

// UserResponse 用户响应
type UserResponse struct {
	ID         uint      `json:"id"`          // 用户ID
	UserID     string    `json:"user_id"`     // 用户ID
	Username   string    `json:"username"`    // 用户名
	Email      string    `json:"email"`       // 邮箱
	Phone      string    `json:"phone"`       // 手机号
	Avatar     string    `json:"avatar"`      // 头像
	Status     int       `json:"status"`      // 状态：0-禁用，1-启用
	LastLogin  time.Time `json:"last_login"`  // 最后登录时间
	LoginCount int       `json:"login_count"` // 登录次数
	Remark     string    `json:"remark"`      // 备注
	CreatedAt  time.Time `json:"created_at"`  // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`  // 更新时间
	Roles      []string  `json:"roles"`       // 角色列表
}

// UserListRequest 用户列表请求
type UserListRequest struct {
	Page     int    `form:"page" binding:"min=1"`      // 页码
	PageSize int    `form:"page_size" binding:"min=1"` // 每页数量
	Username string `form:"username"`                  // 用户名（模糊搜索）
	Status   *int   `form:"status"`                    // 状态筛选
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	List  []UserResponse `json:"list"`  // 用户列表
	Total int64          `json:"total"` // 总数
}

// CreateRoleRequest 创建角色请求
type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"` // 角色名称
	Code        string `json:"code" binding:"required"` // 角色编码
	Description string `json:"description"`             // 角色描述
	Status      int    `json:"status"`                  // 状态：0-禁用，1-启用
	Sort        int    `json:"sort"`                    // 排序
}

// UpdateRoleRequest 更新角色请求
type UpdateRoleRequest struct {
	ID          uint   `json:"id" binding:"required"`   // 角色ID
	Name        string `json:"name" binding:"required"` // 角色名称
	Code        string `json:"code" binding:"required"` // 角色编码
	Description string `json:"description"`             // 角色描述
	Status      int    `json:"status"`                  // 状态：0-禁用，1-启用
	Sort        int    `json:"sort"`                    // 排序
}

// RoleResponse 角色响应
type RoleResponse struct {
	ID          uint      `json:"id"`          // 角色ID
	Name        string    `json:"name"`        // 角色名称
	Code        string    `json:"code"`        // 角色编码
	Description string    `json:"description"` // 角色描述
	Status      int       `json:"status"`      // 状态：0-禁用，1-启用
	Sort        int       `json:"sort"`        // 排序
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`  // 更新时间
}

// AssignRolePermissionsRequest 分配角色权限请求
type AssignRolePermissionsRequest struct {
	RoleID        uint   `json:"role_id" binding:"required"`        // 角色ID
	PermissionIDs []uint `json:"permission_ids" binding:"required"` // 权限ID列表
}

// CreatePermissionRequest 创建权限请求
type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required"` // 权限名称
	Code        string `json:"code" binding:"required"` // 权限编码
	Type        string `json:"type" binding:"required"` // 权限类型：menu-菜单，button-按钮，api-接口
	ParentID    uint   `json:"parent_id"`               // 父级ID
	Path        string `json:"path"`                    // 路由路径
	Component   string `json:"component"`               // 组件路径
	Icon        string `json:"icon"`                    // 图标
	Sort        int    `json:"sort"`                    // 排序
	Status      int    `json:"status"`                  // 状态：0-禁用，1-启用
	Description string `json:"description"`             // 权限描述
}

// PermissionResponse 权限响应
type PermissionResponse struct {
	ID          uint                 `json:"id"`                 // 权限ID
	Name        string               `json:"name"`               // 权限名称
	Code        string               `json:"code"`               // 权限编码
	Type        string               `json:"type"`               // 权限类型
	ParentID    uint                 `json:"parent_id"`          // 父级ID
	Path        string               `json:"path"`               // 路由路径
	Component   string               `json:"component"`          // 组件路径
	Icon        string               `json:"icon"`               // 图标
	Sort        int                  `json:"sort"`               // 排序
	Status      int                  `json:"status"`             // 状态
	Description string               `json:"description"`        // 权限描述
	CreatedAt   time.Time            `json:"created_at"`         // 创建时间
	UpdatedAt   time.Time            `json:"updated_at"`         // 更新时间
	Children    []PermissionResponse `json:"children,omitempty"` // 子权限
}

// MenuResponse 菜单响应
type MenuResponse struct {
	ID        uint           `json:"id"`                 // 菜单ID
	Name      string         `json:"name"`               // 菜单名称
	ParentID  uint           `json:"parent_id"`          // 父级菜单ID
	Path      string         `json:"path"`               // 路由路径
	Component string         `json:"component"`          // 组件路径
	Icon      string         `json:"icon"`               // 图标
	Sort      int            `json:"sort"`               // 排序
	Status    int            `json:"status"`             // 状态
	Type      int            `json:"type"`               // 类型
	Remark    string         `json:"remark"`             // 备注
	Children  []MenuResponse `json:"children,omitempty"` // 子菜单
}
