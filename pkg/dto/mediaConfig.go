package dto

// SaveMediaConfigRequest 保存媒体配置请求
type SaveMediaConfigRequest struct {
	Name      string `json:"name" binding:"required" example:"摄像头1"`      // 媒体名称
	Url       string `json:"url" binding:"required" example:"rtsp://..."`  // 媒体URL
	Status    string `json:"status" example:"active"`                      // 状态
	Code      string `json:"code" example:"CAM001"`                        // 编码
	Type      string `json:"type" example:"camera"`                        // 类型
	MediaDesc string `json:"media_desc" example:"前门摄像头"`                   // 媒体描述
}

// UpdateMediaConfigRequest 更新媒体配置请求
type UpdateMediaConfigRequest struct {
	ID        int    `json:"id" binding:"required" example:"1"`            // 媒体配置ID
	Name      string `json:"name" binding:"required" example:"摄像头1"`      // 媒体名称
	Url       string `json:"url" binding:"required" example:"rtsp://..."`  // 媒体URL
	Status    string `json:"status" example:"active"`                      // 状态
	Code      string `json:"code" example:"CAM001"`                        // 编码
	Type      string `json:"type" example:"camera"`                        // 类型
	MediaDesc string `json:"media_desc" example:"前门摄像头"`                   // 媒体描述
}

type MediaConfigPageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type MediaConfigPageResponse struct {
	List  []MediaConfigResponse `json:"list"`
	Total int64                 `json:"total"`
}

type MediaConfigResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Status    string `json:"status"`
	Code      string `json:"code"`
	Type      string `json:"type"`
	MediaDesc string `json:"media_desc"`
}
