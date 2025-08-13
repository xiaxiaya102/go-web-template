package dto

// SaveTaskManageRequest 保存任务管理请求
type SaveTaskManageRequest struct {
	Name             string `json:"name" binding:"required" example:"监控任务1"`
	MediaSource      string `json:"media_source" example:"rtsp://192.168.1.100/stream"`
	WorkPlan         string `json:"work_plan" example:"全天候监控"`
	ArithmeticConfig string `json:"arithmetic_config" example:"人脸识别+行为分析"`
	TaskStatus       int    `json:"task_status" example:"1"`
	MediaID          int    `json:"media_id" example:"1"`
	PlanID           int    `json:"plan_id" example:"1"`
	SophonTypes      string `json:"sophon_types" example:"1,2,3"`
	ReportAddr       string `json:"report_addr" example:"http://192.168.1.200/api/alarm"`
	AlarmInterval    int    `json:"alarm_interval" example:"30"`
	AritParams       string `json:"arit_params" example:"{\"threshold\":0.8}"`
	StatusDesc       string `json:"status_desc" example:"运行中"`
	UsedIP           string `json:"used_ip" example:"192.168.1.100"`
}

type UpdateTaskManageRequest struct {
	ID               int    `json:"id" binding:"required"`
	Name             string `json:"name" binding:"required"`
	MediaSource      string `json:"media_source"`
	WorkPlan         string `json:"work_plan"`
	ArithmeticConfig string `json:"arithmetic_config"`
	TaskStatus       int    `json:"task_status"`
	MediaID          int    `json:"media_id"`
	PlanID           int    `json:"plan_id"`
	SophonTypes      string `json:"sophon_types"`
	ReportAddr       string `json:"report_addr"`
	AlarmInterval    int    `json:"alarm_interval"`
	AritParams       string `json:"arit_params"`
	StatusDesc       string `json:"status_desc"`
	UsedIP           string `json:"used_ip"`
}

type TaskManagePageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type TaskManagePageResponse struct {
	List  []TaskManageResponse `json:"list"`
	Total int64                `json:"total"`
}

type TaskManageResponse struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	MediaSource      string `json:"media_source"`
	WorkPlan         string `json:"work_plan"`
	ArithmeticConfig string `json:"arithmetic_config"`
	TaskStatus       int    `json:"task_status"`
	MediaID          int    `json:"media_id"`
	PlanID           int    `json:"plan_id"`
	SophonTypes      string `json:"sophon_types"`
	ReportAddr       string `json:"report_addr"`
	AlarmInterval    int    `json:"alarm_interval"`
	AritParams       string `json:"arit_params"`
	StatusDesc       string `json:"status_desc"`
	UsedIP           string `json:"used_ip"`
}

type UpdateTaskStatusRequest struct {
	ID         int `json:"id" binding:"required"`
	TaskStatus int `json:"task_status" binding:"required"`
}
