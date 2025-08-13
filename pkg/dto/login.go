package dto

// LoginRequest 登录请求参数
type LoginRequest struct {
	UserName string `json:"userName" binding:"required" example:"admin"` // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

// LogoutRequest 登出请求参数
type LogoutRequest struct {
	Token string `json:"token" binding:"required" example:"abc123"` // 用户token
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token" example:"abc123"` // 用户token
}
