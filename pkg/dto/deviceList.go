package dto

type SaveDeviceListRequest struct {
	IPAddress  string `json:"ip_address" binding:"required"` // IP地址
	LastUsedTm string `json:"last_used_tm"`                  // 最后使用时间
	Status     int    `json:"status"`                        // 设备状态
	FaceDevice int    `json:"face_device"`                   // 是否为人脸设备
}

type UpdateDeviceListRequest struct {
	ID         int    `json:"id" binding:"required"`         // 设备ID
	IPAddress  string `json:"ip_address" binding:"required"` // IP地址
	LastUsedTm string `json:"last_used_tm"`                  // 最后使用时间
	Status     int    `json:"status"`                        // 设备状态
	FaceDevice int    `json:"face_device"`                   // 是否为人脸设备
}

type DeviceListPageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type DeviceListPageResponse struct {
	List  []DeviceListResponse `json:"list"`
	Total int64                `json:"total"`
}

type DeviceListResponse struct {
	ID         int    `json:"id"`           // 设备ID
	IPAddress  string `json:"ip_address"`   // IP地址
	LastUsedTm string `json:"last_used_tm"` // 最后使用时间
	Status     int    `json:"status"`       // 设备状态
	FaceDevice int    `json:"face_device"`  // 是否为人脸设备
}

type UpdateDeviceStatusRequest struct {
	ID     int `json:"id" binding:"required"`     // 设备ID
	Status int `json:"status" binding:"required"` // 设备状态
}

type DeviceListQueryRequest struct {
	Status     *int   `json:"status" form:"status"`         // 设备状态
	FaceDevice *int   `json:"face_device" form:"face_device"` // 是否为人脸设备
	IPAddress  string `json:"ip_address" form:"ip_address"`   // IP地址（模糊查询）
	Page       int    `json:"page" form:"page"`
	PageSize   int    `json:"pageSize" form:"pageSize"`
}

type DeviceStatusSummaryResponse struct {
	TotalDevices    int `json:"total_devices"`    // 总设备数
	OnlineDevices   int `json:"online_devices"`   // 在线设备数
	OfflineDevices  int `json:"offline_devices"`  // 离线设备数
	FaceDevices     int `json:"face_devices"`     // 人脸设备数
	NonFaceDevices  int `json:"non_face_devices"` // 非人脸设备数
}

type BatchUpdateDeviceStatusRequest struct {
	DeviceIDs []int `json:"device_ids" binding:"required"` // 设备ID列表
	Status    int   `json:"status" binding:"required"`     // 设备状态
}
