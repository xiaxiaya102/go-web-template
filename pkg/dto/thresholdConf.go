package dto

type SaveThresholdConfRequest struct {
	ParamDesc  string  `json:"param_desc" binding:"required"` // 参数描述
	ParamValue float64 `json:"param_value"`                   // 参数值
	SophonType int     `json:"sophon_type"`                   // Sophon类型
}

type UpdateThresholdConfRequest struct {
	ID         int     `json:"id" binding:"required"`         // 阈值配置ID
	ParamDesc  string  `json:"param_desc" binding:"required"` // 参数描述
	ParamValue float64 `json:"param_value"`                   // 参数值
	SophonType int     `json:"sophon_type"`                   // Sophon类型
}

type ThresholdConfPageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type ThresholdConfPageResponse struct {
	List  []ThresholdConfResponse `json:"list"`
	Total int64                   `json:"total"`
}

type ThresholdConfResponse struct {
	ID         int     `json:"id"`          // 阈值配置ID
	ParamDesc  string  `json:"param_desc"`  // 参数描述
	ParamValue float64 `json:"param_value"` // 参数值
	SophonType int     `json:"sophon_type"` // Sophon类型
}

type ThresholdConfQueryRequest struct {
	ParamDesc  string   `json:"param_desc" form:"param_desc"`   // 参数描述（模糊查询）
	SophonType *int     `json:"sophon_type" form:"sophon_type"` // Sophon类型
	MinValue   *float64 `json:"min_value" form:"min_value"`     // 最小值
	MaxValue   *float64 `json:"max_value" form:"max_value"`     // 最大值
	Page       int      `json:"page" form:"page"`
	PageSize   int      `json:"pageSize" form:"pageSize"`
}

type UpdateThresholdValueRequest struct {
	ID         int     `json:"id" binding:"required"` // 阈值配置ID
	ParamValue float64 `json:"param_value"`           // 参数值
}

type BatchUpdateThresholdConfRequest struct {
	Thresholds []UpdateThresholdConfRequest `json:"thresholds" binding:"required"` // 阈值配置列表
}

type ThresholdConfStatisticsResponse struct {
	TotalThresholds   int                         `json:"total_thresholds"`    // 总阈值数
	SophonTypeStats   []SophonTypeStatItem        `json:"sophon_type_stats"`   // Sophon类型统计
	ValueRangeStats   ThresholdValueRangeStats    `json:"value_range_stats"`   // 值范围统计
	ParamDescStats    []ParamDescStatItem         `json:"param_desc_stats"`    // 参数描述统计
}



type ThresholdValueRangeStats struct {
	MinValue float64 `json:"min_value"` // 最小值
	MaxValue float64 `json:"max_value"` // 最大值
	AvgValue float64 `json:"avg_value"` // 平均值
}

type ParamDescStatItem struct {
	ParamDesc string `json:"param_desc"` // 参数描述
	Count     int64  `json:"count"`      // 数量
}

type ThresholdConfBySophonTypeResponse struct {
	SophonType int                       `json:"sophon_type"` // Sophon类型
	Thresholds []ThresholdConfResponse   `json:"thresholds"`  // 阈值配置列表
}

type ThresholdConfByRangeResponse struct {
	MinValue   float64                 `json:"min_value"`   // 最小值
	MaxValue   float64                 `json:"max_value"`   // 最大值
	Thresholds []ThresholdConfResponse `json:"thresholds"`  // 阈值配置列表
}

type ThresholdConfListResponse struct {
	ParamDescs  []string `json:"param_descs"`  // 所有参数描述列表
	SophonTypes []int    `json:"sophon_types"` // 所有Sophon类型列表
}
