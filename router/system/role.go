package system

import (
	v1 "go-web-template/api/v1"
	"go-web-template/middleware"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role")
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.GET("", middleware.RequirePermission("role:list"), roleApi.GetRoleList)         // 获取角色列表
		roleRouter.GET("/:id", middleware.RequirePermission("role:detail"), roleApi.GetRole)       // 获取角色详情
		roleRouter.POST("", middleware.RequirePermission("role:create"), roleApi.CreateRole)       // 创建角色
		roleRouter.PUT("/:id", middleware.RequirePermission("role:update"), roleApi.UpdateRole)    // 更新角色
		roleRouter.DELETE("/:id", middleware.RequirePermission("role:delete"), roleApi.DeleteRole) // 删除角色
	}
}
