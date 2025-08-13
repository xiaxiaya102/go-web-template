package dto

// Result 通用响应结构
type Result struct {
	ErrCode int         `json:"errCode" example:"0"`                   // 响应码，0表示成功
	ErrMsg  string      `json:"errMsg" example:"ok"`                   // 响应消息
	Result  interface{} `json:"result,omitempty" swaggertype:"object"` // 响应数据
}

// Response 通用响应结构别名，用于 Swagger 文档
type Response = Result
