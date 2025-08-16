package system

import (
	v1 "go-web-template/api/v1"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

func (p *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	permissionRouter := Router.Group("permission")
	permissionApi := v1.ApiGroupApp.SystemApiGroup.PermissionApi
	{
		permissionRouter.POST("", permissionApi.CreatePermission)       // 创建权限
		permissionRouter.GET("", permissionApi.GetPermissionList)       // 获取权限列表
		permissionRouter.GET("/:id", permissionApi.GetPermission)       // 获取权限详情
		permissionRouter.PUT("/:id", permissionApi.UpdatePermission)    // 更新权限
		permissionRouter.DELETE("/:id", permissionApi.DeletePermission) // 删除权限
	}
}
