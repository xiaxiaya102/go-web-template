package main

import (
	"feishuReboot/initialization"
	"feishuReboot/logger"
	"time"
)

// @title           Daily News Reboot API
// @version         1.0
// @description     This is a daily news reboot server API documentation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

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
