package service

import (
	"feishuReboot/pkg/model"
	"feishuReboot/pkg/repo"
	"net/http"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	acceptLanguage = "Accept-Language"
	authorization  = "Authorization"
	tokenHeader    = "Token"
	contentType    = "Content-Type"
	multipart      = "multipart/form-data"
	Pattern        = "2006-01-02 15:04:05"
)

var tokenCache *cache.Cache

func init() {
	tokenCache = cache.New(2*time.Hour, 5*time.Minute)
}

func Token(request *http.Request) string {
	// 首先尝试从 Authorization 头部获取 Bearer token
	authHeader := request.Header.Get(authorization)
	if authHeader != "" {
		// 检查是否是 Bearer token 格式
		if strings.HasPrefix(authHeader, "Bearer ") {
			return strings.TrimPrefix(authHeader, "Bearer ")
		}
		// 如果不是 Bearer 格式，直接返回整个值（兼容旧格式）
		return authHeader
	}

	// 兼容旧的 Token 头部
	return request.Header.Get(tokenHeader)
}

func GetUser(token string) *model.User {
	user, found := tokenCache.Get(token)
	if found {
		return user.(*model.User)
	} else {
		user, _ := repo.QueryUserWithToken(token)
		return user
	}
}

func SetUser(token string, user *model.User) {
	tokenCache.Set(token, user, 2*time.Hour)
}

func RemoveUser(token string) {
	tokenCache.Delete(token)
}

func IsMultiPartRequest(request *http.Request) bool {
	return strings.Contains(request.Header.Get(contentType), multipart)
}
