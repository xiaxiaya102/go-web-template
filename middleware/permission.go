package middleware

import (
	"go-web-template/global"
	"go-web-template/pkg/handle"
	"go-web-template/pkg/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// PermissionMiddleware 权限验证中间件
func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusForbidden, handle.Error("用户未认证"))
			c.Abort()
			return
		}

		// 检查用户权限
		if !hasPermission(userID.(uint), requiredPermission) {
			c.JSON(http.StatusForbidden, handle.Error("权限不足"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// RoleMiddleware 角色验证中间件
func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusForbidden, handle.Error("用户未认证"))
			c.Abort()
			return
		}

		// 检查用户角色
		if !hasRole(userID.(uint), requiredRoles...) {
			c.JSON(http.StatusForbidden, handle.Error("角色权限不足"))
			c.Abort()
			return
		}

		c.Next()
	}
}

// SuperAdminMiddleware 超级管理员验证中间件
func SuperAdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware("super_admin")
}

// AdminMiddleware 管理员验证中间件（包括超级管理员和普通管理员）
func AdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware("super_admin", "admin")
}

// hasPermission 检查用户是否具有指定权限
func hasPermission(userID uint, permissionCode string) bool {
	// 超级管理员拥有所有权限
	if isSuperAdmin(userID) {
		return true
	}

	// 查询用户权限
	var count int64
	global.DB.Table("users").
		Select("1").
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN role_permissions ON user_roles.role_id = role_permissions.role_id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("users.id = ? AND permissions.code = ? AND permissions.status = 1", userID, permissionCode).
		Count(&count)

	return count > 0
}

// hasRole 检查用户是否具有指定角色
func hasRole(userID uint, roleCodes ...string) bool {
	if len(roleCodes) == 0 {
		return false
	}

	var count int64
	global.DB.Table("users").
		Select("1").
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("users.id = ? AND roles.code IN ? AND roles.status = 1", userID, roleCodes).
		Count(&count)

	return count > 0
}

// isSuperAdmin 检查用户是否为超级管理员
func isSuperAdmin(userID uint) bool {
	return hasRole(userID, "super_admin")
}

// GetUserPermissions 获取用户所有权限
func GetUserPermissions(userID uint) []string {
	var permissions []string

	// 超级管理员拥有所有权限
	if isSuperAdmin(userID) {
		global.DB.Model(&model.Permission{}).
			Where("status = 1").
			Pluck("code", &permissions)
		return permissions
	}

	// 查询用户权限
	global.DB.Table("permissions").
		Select("DISTINCT permissions.code").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ? AND permissions.status = 1", userID).
		Pluck("code", &permissions)

	return permissions
}

// GetUserRoles 获取用户所有角色
func GetUserRoles(userID uint) []string {
	var roles []string

	global.DB.Table("roles").
		Select("roles.code").
		Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ? AND roles.status = 1", userID).
		Pluck("code", &roles)

	return roles
}

// CheckAPIPermission 检查API权限
func CheckAPIPermission(userID uint, method, path string) bool {
	// 超级管理员拥有所有权限
	if isSuperAdmin(userID) {
		return true
	}

	// 构建权限代码
	permissionCode := buildAPIPermissionCode(method, path)

	// 检查权限
	return hasPermission(userID, permissionCode)
}

// buildAPIPermissionCode 构建API权限代码
func buildAPIPermissionCode(method, path string) string {
	// 移除API前缀
	path = strings.TrimPrefix(path, "/api/")

	// 替换路径分隔符
	path = strings.ReplaceAll(path, "/", ":")

	// 构建权限代码：method:module:action
	return strings.ToLower(method) + ":" + path
}

// RequirePermission 权限装饰器函数
func RequirePermission(permission string) gin.HandlerFunc {
	return PermissionMiddleware(permission)
}

// RequireRole 角色装饰器函数
func RequireRole(roles ...string) gin.HandlerFunc {
	return RoleMiddleware(roles...)
}

// RequireSuperAdmin 超级管理员装饰器函数
func RequireSuperAdmin() gin.HandlerFunc {
	return SuperAdminMiddleware()
}

// RequireAdmin 管理员装饰器函数
func RequireAdmin() gin.HandlerFunc {
	return AdminMiddleware()
}
