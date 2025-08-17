package dto

// PageResult 分页结果
type PageResult struct {
	List     interface{} `json:"list"`     // 数据列表
	Total    int64       `json:"total"`    // 总数
	Page     int         `json:"page"`     // 当前页
	PageSize int         `json:"pageSize"` // 每页大小
}

// BatchDeleteRequest 批量删除请求
type BatchDeleteRequest struct {
	IDs []uint `json:"ids" binding:"required"` // ID列表
}

// ClearLogsRequest 清空日志请求
type ClearLogsRequest struct {
	StartTime string `json:"startTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
	Days      int    `json:"days"`      // 清空天数之前的日志
}

// OperationLogStats 操作日志统计
type OperationLogStats struct {
	TotalCount     int64            `json:"totalCount"`     // 总操作次数
	ModuleStats    []ModuleStats    `json:"moduleStats"`    // 模块统计
	OperationStats []OperationStats `json:"operationStats"` // 操作统计
	UserStats      []UserStats      `json:"userStats"`      // 用户统计
	DailyStats     []DailyStats     `json:"dailyStats"`     // 日期统计
}

// ModuleStats 模块统计
type ModuleStats struct {
	Module string `json:"module"` // 模块名
	Count  int64  `json:"count"`  // 操作次数
}

// OperationStats 操作统计
type OperationStats struct {
	Operation string `json:"operation"` // 操作类型
	Count     int64  `json:"count"`     // 操作次数
}

// UserStats 用户统计
type UserStats struct {
	Username string `json:"username"` // 用户名
	Count    int64  `json:"count"`    // 操作次数
}

// DailyStats 日期统计
type DailyStats struct {
	Date  string `json:"date"`  // 日期
	Count int64  `json:"count"` // 操作次数
}
