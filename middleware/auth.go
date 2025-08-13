package middleware

import (
	"feishuReboot/pkg/repo"
	"feishuReboot/pkg/service"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_token := service.Token(c.Request)
		if _token != "" {
			user := service.GetUser(_token)
			if user != nil {
				now := time.Now()
				if now.Before(user.ExpireTime) {
					// 如果距离过期时间还有超过10分钟，则延长过期时间
					if user.ExpireTime.Sub(now) < time.Minute*10 {
						user.ExpireTime = now.Add(time.Hour * 2)
						repo.UpdateUser(user)
					}
					c.Next()
					return
				}
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": "无效的token"})
	}
}
