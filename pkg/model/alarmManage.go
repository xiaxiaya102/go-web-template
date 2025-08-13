package model

// AlarmManage 报警状态常量
const (
	NOT_REPORT      = -1 // 默认状态，未配置上报地址
	REPORT_FAIL     = 0  // 上报失败
	REPORT_SUCCESS  = 1  // 上报成功
	REPORT_RETRY    = 2  // 等待重试
)

type AlarmManage struct {
	ID           int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	AlarmTask    string `gorm:"column:alarm_task" json:"alarm_task,omitempty"`
	VideoMedia   string `gorm:"column:video_media" json:"video_media,omitempty"`
	MediaURL     string `gorm:"column:media_url" json:"media_url,omitempty"`
	Img          string `gorm:"column:img" json:"img,omitempty"`
	InitImg      string `gorm:"column:init_img" json:"init_img,omitempty"`        // 原图
	AlarmType    string `gorm:"column:alarm_type" json:"alarm_type,omitempty"`
	ReportStatus int    `gorm:"column:report_status" json:"report_status,omitempty"`
	ReportAddr   string `gorm:"column:report_addr" json:"report_addr,omitempty"`
	AlarmVideo   string `gorm:"column:alarm_video" json:"alarm_video,omitempty"`
	AlarmDetail  string `gorm:"column:alarm_detail" json:"alarm_detail,omitempty"`
	Tm           string `gorm:"column:tm" json:"tm,omitempty"`
	MediaID      int    `gorm:"column:media_id" json:"media_id,omitempty"`
}
