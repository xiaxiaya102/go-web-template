package initialization

import (
	"feishuReboot/global"
	"feishuReboot/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func InitServer(router *gin.Engine) server {

	logger.Info("Starting HTTP service at %s", global.System.ServerInfo.Port)

	return &http.Server{
		Addr:    ":" + global.System.ServerInfo.Port,
		Handler: router,
	}
}
