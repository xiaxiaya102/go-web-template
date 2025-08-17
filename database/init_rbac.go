package database

import (
	"go-web-template/global"
	"go-web-template/logger"
	"go-web-template/pkg/model"
)

// InitRBAC 初始化RBAC数据
func InitRBAC() {
	logger.Info("开始初始化RBAC数据...")

	// 初始化权限
	initPermissions()

	// 初始化角色
	initRoles()

	// 初始化角色权限关联
	initRolePermissions()

	// 为默认管理员分配超级管理员角色
	assignSuperAdminRole()

	logger.Info("RBAC数据初始化完成")
}

// initPermissions 初始化权限数据
func initPermissions() {
	permissions := []model.Permission{
		// 用户管理权限
		{Name: "查看用户列表", Code: "user:list", Type: "api", Description: "查看用户列表权限", Status: 1, Sort: 1},
		{Name: "查看用户详情", Code: "user:detail", Type: "api", Description: "查看用户详情权限", Status: 1, Sort: 2},
		{Name: "创建用户", Code: "user:create", Type: "api", Description: "创建用户权限", Status: 1, Sort: 3},
		{Name: "编辑用户", Code: "user:update", Type: "api", Description: "编辑用户权限", Status: 1, Sort: 4},
		{Name: "删除用户", Code: "user:delete", Type: "api", Description: "删除用户权限", Status: 1, Sort: 5},

		// 角色管理权限
		{Name: "查看角色列表", Code: "role:list", Type: "api", Description: "查看角色列表权限", Status: 1, Sort: 11},
		{Name: "查看角色详情", Code: "role:detail", Type: "api", Description: "查看角色详情权限", Status: 1, Sort: 12},
		{Name: "创建角色", Code: "role:create", Type: "api", Description: "创建角色权限", Status: 1, Sort: 13},
		{Name: "编辑角色", Code: "role:update", Type: "api", Description: "编辑角色权限", Status: 1, Sort: 14},
		{Name: "删除角色", Code: "role:delete", Type: "api", Description: "删除角色权限", Status: 1, Sort: 15},

		// 权限管理权限
		{Name: "查看权限列表", Code: "permission:list", Type: "api", Description: "查看权限列表权限", Status: 1, Sort: 21},
		{Name: "查看权限详情", Code: "permission:detail", Type: "api", Description: "查看权限详情权限", Status: 1, Sort: 22},
		{Name: "创建权限", Code: "permission:create", Type: "api", Description: "创建权限权限", Status: 1, Sort: 23},
		{Name: "编辑权限", Code: "permission:update", Type: "api", Description: "编辑权限权限", Status: 1, Sort: 24},
		{Name: "删除权限", Code: "permission:delete", Type: "api", Description: "删除权限权限", Status: 1, Sort: 25},

		// 操作日志权限
		{Name: "查看操作日志", Code: "operation-logs:list", Type: "api", Description: "查看操作日志权限", Status: 1, Sort: 31},
		{Name: "查看日志详情", Code: "operation-logs:detail", Type: "api", Description: "查看日志详情权限", Status: 1, Sort: 32},
		{Name: "删除操作日志", Code: "operation-logs:delete", Type: "api", Description: "删除操作日志权限", Status: 1, Sort: 33},
		{Name: "清空操作日志", Code: "operation-logs:clear", Type: "api", Description: "清空操作日志权限", Status: 1, Sort: 34},
		{Name: "查看日志统计", Code: "operation-logs:stats", Type: "api", Description: "查看日志统计权限", Status: 1, Sort: 35},

		// 系统管理权限
		{Name: "系统设置", Code: "system:setting", Type: "api", Description: "系统设置权限", Status: 1, Sort: 41},
		{Name: "系统监控", Code: "system:monitor", Type: "api", Description: "系统监控权限", Status: 1, Sort: 42},

		// 菜单权限
		{Name: "用户管理菜单", Code: "menu:user", Type: "menu", Path: "/user", Description: "用户管理菜单权限", Status: 1, Sort: 101},
		{Name: "角色管理菜单", Code: "menu:role", Type: "menu", Path: "/role", Description: "角色管理菜单权限", Status: 1, Sort: 102},
		{Name: "权限管理菜单", Code: "menu:permission", Type: "menu", Path: "/permission", Description: "权限管理菜单权限", Status: 1, Sort: 103},
		{Name: "操作日志菜单", Code: "menu:operation-logs", Type: "menu", Path: "/operation-logs", Description: "操作日志菜单权限", Status: 1, Sort: 104},
		{Name: "系统管理菜单", Code: "menu:system", Type: "menu", Path: "/system", Description: "系统管理菜单权限", Status: 1, Sort: 105},
	}

	for _, permission := range permissions {
		var count int64
		global.DB.Model(&model.Permission{}).Where("code = ?", permission.Code).Count(&count)
		if count == 0 {
			if err := global.DB.Create(&permission).Error; err != nil {
				logger.Error("创建权限失败: %s, 错误: %v", permission.Code, err)
			} else {
				logger.Info("创建权限成功: %s", permission.Code)
			}
		}
	}
}

// initRoles 初始化角色数据
func initRoles() {
	roles := []model.Role{
		{Name: "超级管理员", Code: "super_admin", Description: "系统超级管理员，拥有所有权限", Status: 1, Sort: 1},
		{Name: "管理员", Code: "admin", Description: "系统管理员，拥有大部分管理权限", Status: 1, Sort: 2},
		{Name: "普通用户", Code: "user", Description: "普通用户，拥有基本查看权限", Status: 1, Sort: 3},
		{Name: "访客", Code: "guest", Description: "访客用户，只有查看权限", Status: 1, Sort: 4},
	}

	for _, role := range roles {
		var count int64
		global.DB.Model(&model.Role{}).Where("code = ?", role.Code).Count(&count)
		if count == 0 {
			if err := global.DB.Create(&role).Error; err != nil {
				logger.Error("创建角色失败: %s, 错误: %v", role.Code, err)
			} else {
				logger.Info("创建角色成功: %s", role.Code)
			}
		}
	}
}

// initRolePermissions 初始化角色权限关联
func initRolePermissions() {
	// 管理员角色权限（除了超级管理员专有权限外的所有权限）
	adminPermissions := []string{
		"user:list", "user:detail", "user:create", "user:update",
		"role:list", "role:detail",
		"permission:list", "permission:detail",
		"operation-logs:list", "operation-logs:detail", "operation-logs:stats",
		"menu:user", "menu:role", "menu:permission", "menu:operation-logs",
	}

	// 普通用户权限（基本查看权限）
	userPermissions := []string{
		"user:list", "user:detail",
		"role:list", "role:detail",
		"permission:list", "permission:detail",
		"menu:user", "menu:role", "menu:permission",
	}

	// 访客权限（只读权限）
	guestPermissions := []string{
		"user:list",
		"role:list",
		"permission:list",
	}

	// 分配管理员权限
	assignRolePermissions("admin", adminPermissions)

	// 分配普通用户权限
	assignRolePermissions("user", userPermissions)

	// 分配访客权限
	assignRolePermissions("guest", guestPermissions)
}

// assignRolePermissions 为角色分配权限
func assignRolePermissions(roleCode string, permissionCodes []string) {
	var role model.Role
	if err := global.DB.Where("code = ?", roleCode).First(&role).Error; err != nil {
		logger.Error("角色不存在: %s", roleCode)
		return
	}

	for _, permissionCode := range permissionCodes {
		var permission model.Permission
		if err := global.DB.Where("code = ?", permissionCode).First(&permission).Error; err != nil {
			logger.Error("权限不存在: %s", permissionCode)
			continue
		}

		// 检查关联是否已存在
		var count int64
		global.DB.Model(&model.RolePermission{}).
			Where("role_id = ? AND permission_id = ?", role.ID, permission.ID).
			Count(&count)

		if count == 0 {
			rolePermission := &model.RolePermission{
				RoleID:       role.ID,
				PermissionID: permission.ID,
			}
			if err := global.DB.Create(rolePermission).Error; err != nil {
				logger.Error("分配权限失败: 角色=%s, 权限=%s, 错误=%v", roleCode, permissionCode, err)
			}
		}
	}

	logger.Info("角色权限分配完成: %s", roleCode)
}

// assignSuperAdminRole 为默认管理员分配超级管理员角色
func assignSuperAdminRole() {
	var adminUser model.User
	if err := global.DB.Where("username = ?", "admin").First(&adminUser).Error; err != nil {
		logger.Error("默认管理员用户不存在")
		return
	}

	var superAdminRole model.Role
	if err := global.DB.Where("code = ?", "super_admin").First(&superAdminRole).Error; err != nil {
		logger.Error("超级管理员角色不存在")
		return
	}

	// 检查是否已分配角色
	var count int64
	global.DB.Model(&model.UserRole{}).
		Where("user_id = ? AND role_id = ?", adminUser.ID, superAdminRole.ID).
		Count(&count)

	if count == 0 {
		userRole := &model.UserRole{
			UserID: adminUser.ID,
			RoleID: superAdminRole.ID,
		}
		if err := global.DB.Create(userRole).Error; err != nil {
			logger.Error("分配超级管理员角色失败: %v", err)
		} else {
			logger.Info("默认管理员已分配超级管理员角色")
		}
	}
}
