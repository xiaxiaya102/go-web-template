package dto

type SaveAlarmCategoryRequest struct {
	Arithmetic string `json:"arithmetic" binding:"required"` // 所属算法
	AlarmType  string `json:"alarm_type" binding:"required"` // 告警类别
	AlarmDesc  string `json:"alarm_desc"`                    // 告警描述
	AudioFile  string `json:"audio_file"`                    // 语音文件
	SophonType int    `json:"sophon_type"`                   // Sophon类型
}

type UpdateAlarmCategoryRequest struct {
	ID         int    `json:"id" binding:"required"`         // 告警分类ID
	Arithmetic string `json:"arithmetic" binding:"required"` // 所属算法
	AlarmType  string `json:"alarm_type" binding:"required"` // 告警类别
	AlarmDesc  string `json:"alarm_desc"`                    // 告警描述
	AudioFile  string `json:"audio_file"`                    // 语音文件
	SophonType int    `json:"sophon_type"`                   // Sophon类型
}

type AlarmCategoryPageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type AlarmCategoryPageResponse struct {
	List  []AlarmCategoryResponse `json:"list"`
	Total int64                   `json:"total"`
}

type AlarmCategoryResponse struct {
	ID         int    `json:"id"`          // 告警分类ID
	Arithmetic string `json:"arithmetic"`  // 所属算法
	AlarmType  string `json:"alarm_type"`  // 告警类别
	AlarmDesc  string `json:"alarm_desc"`  // 告警描述
	AudioFile  string `json:"audio_file"`  // 语音文件
	SophonType int    `json:"sophon_type"` // Sophon类型
}

type AlarmCategoryQueryRequest struct {
	Arithmetic string `json:"arithmetic" form:"arithmetic"`   // 所属算法（模糊查询）
	AlarmType  string `json:"alarm_type" form:"alarm_type"`   // 告警类别（模糊查询）
	SophonType *int   `json:"sophon_type" form:"sophon_type"` // Sophon类型
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
}

type AlarmCategoryStatisticsResponse struct {
	TotalCategories   int                        `json:"total_categories"`    // 总分类数
	ArithmeticStats   []ArithmeticStatItem       `json:"arithmetic_stats"`    // 算法统计
	AlarmTypeStats    []AlarmTypeStatItem        `json:"alarm_type_stats"`    // 告警类型统计
	SophonTypeStats   []SophonTypeStatItem       `json:"sophon_type_stats"`   // Sophon类型统计
	AudioFileStats    AlarmCategoryAudioStats    `json:"audio_file_stats"`    // 语音文件统计
}

type ArithmeticStatItem struct {
	Arithmetic string `json:"arithmetic"` // 算法名称
	Count      int64  `json:"count"`      // 数量
}

type AlarmTypeStatItem struct {
	AlarmType string `json:"alarm_type"` // 告警类型
	Count     int64  `json:"count"`      // 数量
}

type SophonTypeStatItem struct {
	SophonType int   `json:"sophon_type"` // Sophon类型
	Count      int64 `json:"count"`       // 数量
}

type AlarmCategoryAudioStats struct {
	WithAudio    int64 `json:"with_audio"`    // 有语音文件的数量
	WithoutAudio int64 `json:"without_audio"` // 无语音文件的数量
}

type BatchUpdateAlarmCategoryRequest struct {
	Categories []UpdateAlarmCategoryRequest `json:"categories" binding:"required"` // 告警分类列表
}

type AlarmCategoryByArithmeticResponse struct {
	Arithmetic string                    `json:"arithmetic"`  // 算法名称
	Categories []AlarmCategoryResponse   `json:"categories"`  // 告警分类列表
}

type AlarmCategoryByTypeResponse struct {
	AlarmType  string                  `json:"alarm_type"`  // 告警类型
	Categories []AlarmCategoryResponse `json:"categories"`  // 告警分类列表
}

type AlarmCategoryListResponse struct {
	Arithmetics []string `json:"arithmetics"` // 所有算法列表
	AlarmTypes  []string `json:"alarm_types"` // 所有告警类型列表
}
