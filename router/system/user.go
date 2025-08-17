package system

import (
	v1 "go-web-template/api/v1"
	"go-web-template/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	permissionCheckApi := v1.ApiGroupApp.SystemApiGroup.PermissionCheckApi
	{
		// 权限检查相关接口（所有登录用户都可访问）
		userRouter.GET("/permissions", permissionCheckApi.GetUserPermissions)   // 获取用户权限
		userRouter.GET("/roles", permissionCheckApi.GetUserRoles)               // 获取用户角色
		userRouter.GET("/info", permissionCheckApi.GetUserInfo)                 // 获取用户信息
		userRouter.GET("/check-permission", permissionCheckApi.CheckPermission) // 检查权限
		userRouter.GET("/check-role", permissionCheckApi.CheckRole)             // 检查角色

		// 用户管理接口（需要相应权限）
		userRouter.GET("", middleware.RequirePermission("user:list"), userApi.GetUserList)         // 获取用户列表
		userRouter.GET("/:id", middleware.RequirePermission("user:detail"), userApi.GetUser)       // 获取用户详情
		userRouter.POST("", middleware.RequirePermission("user:create"), userApi.CreateUser)       // 创建用户
		userRouter.PUT("/:id", middleware.RequirePermission("user:update"), userApi.UpdateUser)    // 更新用户
		userRouter.DELETE("/:id", middleware.RequirePermission("user:delete"), userApi.DeleteUser) // 删除用户
		userRouter.POST("/change-password", userApi.ChangePassword)                                // 修改密码（自己的密码）
	}
}
