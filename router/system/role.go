package system

import (
	v1 "go-web-template/api/v1"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role")
	roleApi := v1.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.POST("", roleApi.CreateRole)       // 创建角色
		roleRouter.GET("", roleApi.GetRoleList)       // 获取角色列表
		roleRouter.GET("/:id", roleApi.GetRole)       // 获取角色详情
		roleRouter.PUT("/:id", roleApi.UpdateRole)    // 更新角色
		roleRouter.DELETE("/:id", roleApi.DeleteRole) // 删除角色
	}
}
