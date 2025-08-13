package dto

// SaveModelsInfoRequest 保存模型信息请求
type SaveModelsInfoRequest struct {
	Name        string `json:"name" binding:"required" example:"人脸识别模型"`        // 模型名称
	Algorithm   string `json:"algorithm" binding:"required" example:"face_detection"`   // 算法
	DateTime    string `json:"date_time" example:"2024-01-01 12:00:00"`                      // 日期时间
	Description string `json:"description" example:"用于人脸识别的深度学习模型"`                    // 描述
	SophonType  int    `json:"sophon_type" example:"1"`                    // Sophon类型
	Classify    string `json:"classify" example:"人脸识别"`                       // 分类
	Specialty   int    `json:"specialty" example:"1"`                      // 专业性
}

type UpdateModelsInfoRequest struct {
	ID          int    `json:"id" binding:"required"`          // 模型ID
	Name        string `json:"name" binding:"required"`        // 模型名称
	Algorithm   string `json:"algorithm" binding:"required"`   // 算法
	DateTime    string `json:"date_time"`                      // 日期时间
	Description string `json:"description"`                    // 描述
	SophonType  int    `json:"sophon_type"`                    // Sophon类型
	Classify    string `json:"classify"`                       // 分类
	Specialty   int    `json:"specialty"`                      // 专业性
}

type ModelsInfoPageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type ModelsInfoPageResponse struct {
	List  []ModelsInfoResponse `json:"list"`
	Total int64                `json:"total"`
}

type ModelsInfoResponse struct {
	ID          int    `json:"id"`          // 模型ID
	Name        string `json:"name"`        // 模型名称
	Algorithm   string `json:"algorithm"`   // 算法
	DateTime    string `json:"date_time"`   // 日期时间
	Description string `json:"description"` // 描述
	SophonType  int    `json:"sophon_type"` // Sophon类型
	Classify    string `json:"classify"`    // 分类
	Specialty   int    `json:"specialty"`   // 专业性
}

type ModelsInfoQueryRequest struct {
	Name       string `json:"name" form:"name"`             // 模型名称（模糊查询）
	Algorithm  string `json:"algorithm" form:"algorithm"`   // 算法（模糊查询）
	Classify   string `json:"classify" form:"classify"`     // 分类（模糊查询）
	SophonType *int   `json:"sophon_type" form:"sophon_type"` // Sophon类型
	Specialty  *int   `json:"specialty" form:"specialty"`   // 专业性
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
}

type ModelsInfoStatisticsResponse struct {
	TotalModels      int                    `json:"total_models"`       // 总模型数
	AlgorithmStats   []AlgorithmStatItem    `json:"algorithm_stats"`    // 算法统计
	ClassifyStats    []ClassifyStatItem     `json:"classify_stats"`     // 分类统计
	SophonTypeStats  []SophonTypeStatItem   `json:"sophon_type_stats"`  // Sophon类型统计
	SpecialtyStats   []SpecialtyStatItem    `json:"specialty_stats"`    // 专业性统计
}

type AlgorithmStatItem struct {
	Algorithm string `json:"algorithm"` // 算法名称
	Count     int64  `json:"count"`     // 数量
}

type ClassifyStatItem struct {
	Classify string `json:"classify"` // 分类名称
	Count    int64  `json:"count"`    // 数量
}



type SpecialtyStatItem struct {
	Specialty int   `json:"specialty"` // 专业性
	Count     int64 `json:"count"`     // 数量
}

type BatchUpdateModelsInfoRequest struct {
	Models []UpdateModelsInfoRequest `json:"models" binding:"required"` // 模型列表
}

type ModelsInfoByAlgorithmResponse struct {
	Algorithm string                `json:"algorithm"` // 算法名称
	Models    []ModelsInfoResponse  `json:"models"`    // 模型列表
}

type ModelsInfoByClassifyResponse struct {
	Classify string               `json:"classify"` // 分类名称
	Models   []ModelsInfoResponse `json:"models"`   // 模型列表
}
