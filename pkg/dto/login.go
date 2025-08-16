package dto

// LoginRequest 登录请求参数
type LoginRequest struct {
	UserName string `json:"userName" binding:"required" example:"admin"`  // 用户名
	Password string `json:"password" binding:"required" example:"123456"` // 密码
}

// LogoutRequest 登出请求参数
type LogoutRequest struct {
	Token string `json:"token" binding:"required" example:"abc123"` // 用户token
}

// RefreshTokenRequest 刷新token请求参数
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // 刷新token
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token        string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`         // 访问token
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."` // 刷新token
	ExpiresIn    int64  `json:"expires_in" example:"86400"`                                      // token过期时间（秒）
}
