package initialization

import (
	"go-web-template/global"
	"go-web-template/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func InitServer(router *gin.Engine) server {

	logger.Info("Starting HTTP service at %s", global.System.ServerInfo.Port)
	// 接口文档地址
	logger.Info("Swagger doc at http://localhost:%s/swagger/index.html", global.System.ServerInfo.Port)

	return &http.Server{
		Addr:    ":" + global.System.ServerInfo.Port,
		Handler: router,
	}
}
