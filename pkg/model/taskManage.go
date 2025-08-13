package model

// TaskManage 任务状态常量
const (
	STOP_TASKSTATUS           = 0  // 停止状态
	START_TASKSTATUS          = 1  // 启动状态
	DEVICE_OFFLINE_TASKSTATUS = 2  // 设备离线状态
	NOT_AREA_TASKSTATUS       = -1 // 无区域状态
)

type TaskManage struct {
	ID               int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Name             string `gorm:"column:name" json:"name,omitempty"`
	MediaSource      string `gorm:"column:media_source" json:"media_source,omitempty"`
	WorkPlan         string `gorm:"column:work_plan" json:"work_plan,omitempty"`
	ArithmeticConfig string `gorm:"column:arithmetic_config" json:"arithmetic_config,omitempty"`
	TaskStatus       int    `gorm:"column:task_status" json:"task_status,omitempty"`
	MediaID          int    `gorm:"column:media_id" json:"media_id,omitempty"`
	PlanID           int    `gorm:"column:plan_id" json:"plan_id,omitempty"`
	SophonTypes      string `gorm:"column:sophon_types" json:"sophon_types,omitempty"`
	ReportAddr       string `gorm:"column:report_addr" json:"report_addr,omitempty"`
	AlarmInterval    int    `gorm:"column:alarm_interval" json:"alarm_interval,omitempty"`
	AritParams       string `gorm:"column:arit_params" json:"arit_params,omitempty"`
	StatusDesc       string `gorm:"column:status_desc" json:"status_desc,omitempty"`
	UsedIP           string `gorm:"column:used_ip" json:"used_ip,omitempty"`
}
