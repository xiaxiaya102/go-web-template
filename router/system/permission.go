package system

import (
	v1 "go-web-template/api/v1"
	"go-web-template/middleware"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

func (p *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	permissionRouter := Router.Group("permission")
	permissionApi := v1.ApiGroupApp.SystemApiGroup.PermissionApi
	{
		permissionRouter.GET("", middleware.RequirePermission("permission:list"), permissionApi.GetPermissionList)   // 获取权限列表
		permissionRouter.GET("/:id", middleware.RequirePermission("permission:detail"), permissionApi.GetPermission) // 获取权限详情
		permissionRouter.POST("", middleware.RequireSuperAdmin(), permissionApi.CreatePermission)                    // 创建权限（仅超级管理员）
		permissionRouter.PUT("/:id", middleware.RequireSuperAdmin(), permissionApi.UpdatePermission)                 // 更新权限（仅超级管理员）
		permissionRouter.DELETE("/:id", middleware.RequireSuperAdmin(), permissionApi.DeletePermission)              // 删除权限（仅超级管理员）
	}
}
