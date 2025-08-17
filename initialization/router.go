package initialization

import (
	"go-web-template/logger"
	"go-web-template/middleware"
	"go-web-template/router"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-web-template/docs" // 这里会自动生成
)

// 初始化总路由

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // 设置Gin的模式为release
	Router := gin.New()
	Router.Use(gin.Recovery())

	systemRouter := router.GroupApp.System

	cors_config := SetCors()
	Router.Use(middleware.BlockerMiddleware())
	Router.Use(cors.New(cors_config))
	// 设置404处理器
	Router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	// Swagger 文档路由
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 公开接口组 - 不需要认证
	PublicGroup := Router.Group("/api")
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}

	// 私有接口组 - 需要认证
	PrivateGroup := Router.Group("/api")
	PrivateGroup.Use(middleware.AuthMiddleware())
	PrivateGroup.Use(middleware.OperationLogMiddleware()) // 添加操作日志中间件
	{
		systemRouter.InitAuthBaseRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitRoleRouter(PrivateGroup)
		systemRouter.InitPermissionRouter(PrivateGroup)
		systemRouter.InitOperationLogRouter(PrivateGroup)
	}
	logger.Info("Router Init Ok")
	return Router
}

// setWebStatic 设置静态文件路由（已移除，如需要可重新添加）
func setWebStatic(rootRouter *gin.RouterGroup) {
	// 简单的健康检查接口
	rootRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go Web Admin API",
			"version": "1.0.0",
			"status":  "running",
		})
	})
}

func SetCors() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	return config
}
