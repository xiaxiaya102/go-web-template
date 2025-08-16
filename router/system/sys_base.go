package system

import (
	v1 "go-web-template/api/v1"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("refresh", baseApi.RefreshToken)
	}
	return baseRouter
}

// InitAuthBaseRouter 初始化需要认证的基础路由
func (s *BaseRouter) InitAuthBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("logout", baseApi.Logout)
	}
	return baseRouter
}
