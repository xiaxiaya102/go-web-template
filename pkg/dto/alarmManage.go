package dto

type SaveAlarmManageRequest struct {
	AlarmTask    string `json:"alarm_task" binding:"required"`
	VideoMedia   string `json:"video_media"`
	MediaURL     string `json:"media_url"`
	Img          string `json:"img"`
	InitImg      string `json:"init_img"`
	AlarmType    string `json:"alarm_type"`
	ReportStatus int    `json:"report_status"`
	ReportAddr   string `json:"report_addr"`
	AlarmVideo   string `json:"alarm_video"`
	AlarmDetail  string `json:"alarm_detail"`
	Tm           string `json:"tm"`
	MediaID      int    `json:"media_id"`
}

type UpdateAlarmManageRequest struct {
	ID           int    `json:"id" binding:"required"`
	AlarmTask    string `json:"alarm_task" binding:"required"`
	VideoMedia   string `json:"video_media"`
	MediaURL     string `json:"media_url"`
	Img          string `json:"img"`
	InitImg      string `json:"init_img"`
	AlarmType    string `json:"alarm_type"`
	ReportStatus int    `json:"report_status"`
	ReportAddr   string `json:"report_addr"`
	AlarmVideo   string `json:"alarm_video"`
	AlarmDetail  string `json:"alarm_detail"`
	Tm           string `json:"tm"`
	MediaID      int    `json:"media_id"`
}

type AlarmManagePageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type AlarmManagePageResponse struct {
	List  []AlarmManageResponse `json:"list"`
	Total int64                 `json:"total"`
}

type AlarmManageResponse struct {
	ID           int    `json:"id"`
	AlarmTask    string `json:"alarm_task"`
	VideoMedia   string `json:"video_media"`
	MediaURL     string `json:"media_url"`
	Img          string `json:"img"`
	InitImg      string `json:"init_img"`
	AlarmType    string `json:"alarm_type"`
	ReportStatus int    `json:"report_status"`
	ReportAddr   string `json:"report_addr"`
	AlarmVideo   string `json:"alarm_video"`
	AlarmDetail  string `json:"alarm_detail"`
	Tm           string `json:"tm"`
	MediaID      int    `json:"media_id"`
}

type UpdateReportStatusRequest struct {
	ID           int `json:"id" binding:"required"`
	ReportStatus int `json:"report_status" binding:"required"`
}

type AlarmManageQueryRequest struct {
	AlarmType    string `json:"alarm_type" form:"alarm_type"`
	ReportStatus *int   `json:"report_status" form:"report_status"`
	MediaID      *int   `json:"media_id" form:"media_id"`
	StartTime    string `json:"start_time" form:"start_time"`
	EndTime      string `json:"end_time" form:"end_time"`
	Page         int    `json:"page" form:"page"`
	PageSize     int    `json:"pageSize" form:"pageSize"`
}
