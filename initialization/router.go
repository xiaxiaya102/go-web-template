package initialization

import (
	"feishuReboot/dist"
	"feishuReboot/logger"
	"feishuReboot/middleware"
	"feishuReboot/router"
	"net/http"
	"strings"

	_ "feishuReboot/docs" // 这里会自动生成
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	{
		systemRouter.InitAuthBaseRouter(PrivateGroup)
		systemRouter.InitMediaConfigRouter(PrivateGroup)
		systemRouter.InitPlanTemplateRouter(PrivateGroup)
		systemRouter.InitTaskManageRouter(PrivateGroup)
		systemRouter.InitAlarmManageRouter(PrivateGroup)
		systemRouter.InitNetworkConfigRouter(PrivateGroup)
		systemRouter.InitDeviceListRouter(PrivateGroup)
		systemRouter.InitParamConfRouter(PrivateGroup)
		systemRouter.InitModelsInfoRouter(PrivateGroup)
		systemRouter.InitAlarmCategoryRouter(PrivateGroup)
		systemRouter.InitThresholdConfRouter(PrivateGroup)
	}
	logger.Info("Router Init Ok")
	return Router
}

func setWebStatic(rootRouter *gin.RouterGroup) {
	// 静态资源路由 - 处理 /static/ 路径
	rootRouter.GET("/static/*filepath", func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "public, max-age=3600")

		// 获取文件路径
		filepath := c.Param("filepath")

		// 设置正确的 MIME 类型
		if strings.HasSuffix(filepath, ".js") {
			c.Writer.Header().Set("Content-Type", "application/javascript")
		} else if strings.HasSuffix(filepath, ".css") {
			c.Writer.Header().Set("Content-Type", "text/css")
		} else if strings.HasSuffix(filepath, ".png") {
			c.Writer.Header().Set("Content-Type", "image/png")
		} else if strings.HasSuffix(filepath, ".jpg") || strings.HasSuffix(filepath, ".jpeg") {
			c.Writer.Header().Set("Content-Type", "image/jpeg")
		}

		staticServer := http.FileServer(http.FS(dist.Assets))
		// 不要移除 /static 前缀，因为 embed 的路径包含 static
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	// 处理 img 路径
	rootRouter.GET("/img/*filepath", func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "public, max-age=3600")

		filepath := c.Param("filepath")
		if strings.HasSuffix(filepath, ".png") {
			c.Writer.Header().Set("Content-Type", "image/png")
		} else if strings.HasSuffix(filepath, ".svg") {
			c.Writer.Header().Set("Content-Type", "image/svg+xml")
		}

		staticServer := http.FileServer(http.FS(dist.Assets))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	// manifest.json
	rootRouter.GET("/manifest.json", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.File("dist/manifest.json")
	})

	// favicon
	rootRouter.GET("/favicon.ico", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "image/x-icon")
		staticServer := http.FileServer(http.FS(dist.Favicon))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	// 根路由 - 必须在最后
	rootRouter.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		staticServer := http.FileServer(http.FS(dist.IndexHtml))
		staticServer.ServeHTTP(c.Writer, c.Request)
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
