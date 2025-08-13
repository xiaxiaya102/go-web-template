package middleware

import (
	"feishuReboot/global"
	"feishuReboot/pkg/buserr"
	"feishuReboot/pkg/handle"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlockerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.BlockAllRequests {
			c.JSON(http.StatusServiceUnavailable, handle.FailWithMsg(buserr.Upgradeing, "服务器升级中，暂不可用"))
			// c.File("/var/lib/application-web/dist/updating.html")
			c.Abort()
		}
	}
}
