package system

import (
	v1 "go-web-template/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("", userApi.CreateUser)                     // 创建用户
		userRouter.GET("", userApi.GetUserList)                     // 获取用户列表
		userRouter.GET("/:id", userApi.GetUser)                     // 获取用户详情
		userRouter.PUT("/:id", userApi.UpdateUser)                  // 更新用户
		userRouter.DELETE("/:id", userApi.DeleteUser)               // 删除用户
		userRouter.POST("/change-password", userApi.ChangePassword) // 修改密码
	}
}
