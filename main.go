package main

import (
	"go-web-template/initialization"
	"go-web-template/logger"
	"time"
)

// @title           Go Web Admin API
// @version         1.0
// @description     Go Web后台管理框架API文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8089
// @BasePath  /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	//初始化参数
	initialization.InitBase()
	logger.Info("初始化完成")

	Router := initialization.Routers()
	s := initialization.InitServer(Router)

	time.Sleep(10 * time.Microsecond)

	err := s.ListenAndServe()
	if err != nil {
		logger.Error("An error occurred starting HTTP listener %v", err)
	}

}
