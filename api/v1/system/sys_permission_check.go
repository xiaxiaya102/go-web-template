package system

import (
	"fmt"
	"go-web-template/middleware"
	"go-web-template/pkg/handle"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermissionCheckApi struct{}

// getUserID 安全地获取用户ID
func getUserID(c *gin.Context) (uint, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, fmt.Errorf("用户未认证")
	}

	// 添加调试信息
	fmt.Printf("DEBUG: userID type: %T, value: %v\n", userID, userID)

	switch v := userID.(type) {
	case uint:
		return v, nil
	case int:
		return uint(v), nil
	case int64:
		return uint(v), nil
	case float64:
		// JSON反序列化时数字可能变成float64
		return uint(v), nil
	case string:
		if id, err := strconv.ParseUint(v, 10, 32); err == nil {
			return uint(id), nil
		}
		return 0, fmt.Errorf("用户ID格式错误: %s", v)
	default:
		return 0, fmt.Errorf("用户ID格式错误: %T, value: %v", v, v)
	}
}

// GetUserPermissions 获取当前用户权限
// @Summary 获取当前用户权限
// @Description 获取当前登录用户的所有权限列表
// @Tags 权限检查
// @Accept json
// @Produce json
// @Success 200 {object} handle.Response{data=[]string} "获取成功"
// @Security BearerAuth
// @Router /user/permissions [get]
func (p *PermissionCheckApi) GetUserPermissions(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, handle.Error(err.Error()))
		return
	}

	permissions := middleware.GetUserPermissions(userID)
	c.JSON(http.StatusOK, handle.Success(permissions))
}

// GetUserRoles 获取当前用户角色
// @Summary 获取当前用户角色
// @Description 获取当前登录用户的所有角色列表
// @Tags 权限检查
// @Accept json
// @Produce json
// @Success 200 {object} handle.Response{data=[]string} "获取成功"
// @Security BearerAuth
// @Router /user/roles [get]
func (p *PermissionCheckApi) GetUserRoles(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, handle.Error(err.Error()))
		return
	}

	roles := middleware.GetUserRoles(userID)
	c.JSON(http.StatusOK, handle.Success(roles))
}

// CheckPermission 检查权限
// @Summary 检查权限
// @Description 检查当前用户是否具有指定权限
// @Tags 权限检查
// @Accept json
// @Produce json
// @Param permission query string true "权限代码"
// @Success 200 {object} handle.Response{data=bool} "检查结果"
// @Security BearerAuth
// @Router /user/check-permission [get]
func (p *PermissionCheckApi) CheckPermission(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, handle.Error(err.Error()))
		return
	}

	permission := c.Query("permission")
	if permission == "" {
		c.JSON(http.StatusOK, handle.Error("权限代码不能为空"))
		return
	}

	// 直接检查权限代码，而不是API路径
	hasPermission := false
	permissions := middleware.GetUserPermissions(userID)
	for _, p := range permissions {
		if p == permission {
			hasPermission = true
			break
		}
	}
	c.JSON(http.StatusOK, handle.Success(hasPermission))
}

// CheckRole 检查角色
// @Summary 检查角色
// @Description 检查当前用户是否具有指定角色
// @Tags 权限检查
// @Accept json
// @Produce json
// @Param role query string true "角色代码"
// @Success 200 {object} handle.Response{data=bool} "检查结果"
// @Security BearerAuth
// @Router /user/check-role [get]
func (p *PermissionCheckApi) CheckRole(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, handle.Error(err.Error()))
		return
	}

	role := c.Query("role")
	if role == "" {
		c.JSON(http.StatusOK, handle.Error("角色代码不能为空"))
		return
	}

	hasRole := middleware.GetUserRoles(userID)
	for _, userRole := range hasRole {
		if userRole == role {
			c.JSON(http.StatusOK, handle.Success(true))
			return
		}
	}

	c.JSON(http.StatusOK, handle.Success(false))
}

// GetUserInfo 获取当前用户信息（包含权限和角色）
// @Summary 获取当前用户信息
// @Description 获取当前登录用户的详细信息，包括权限和角色
// @Tags 权限检查
// @Accept json
// @Produce json
// @Success 200 {object} handle.Response{data=map[string]interface{}} "获取成功"
// @Security BearerAuth
// @Router /user/info [get]
func (p *PermissionCheckApi) GetUserInfo(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusOK, handle.Error(err.Error()))
		return
	}

	username, _ := c.Get("username")

	permissions := middleware.GetUserPermissions(userID)
	roles := middleware.GetUserRoles(userID)

	userInfo := map[string]interface{}{
		"user_id":     userID,
		"username":    username,
		"permissions": permissions,
		"roles":       roles,
		"is_super_admin": func() bool {
			for _, role := range roles {
				if role == "super_admin" {
					return true
				}
			}
			return false
		}(),
	}

	c.JSON(http.StatusOK, handle.Success(userInfo))
}
