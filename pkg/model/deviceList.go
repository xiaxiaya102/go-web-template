package model

// DeviceList 设备状态常量
const (
	DEVICE_OFFLINE = 0 // 设备离线
	DEVICE_ONLINE  = 1 // 设备在线
)

// DeviceList 人脸设备常量
const (
	NOT_FACE_DEVICE = 0 // 非人脸设备
	IS_FACE_DEVICE  = 1 // 人脸设备
)

type DeviceList struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	IPAddress  string `gorm:"column:ip_address" json:"ip_address,omitempty"`   // IP地址
	LastUsedTm string `gorm:"column:last_used_tm" json:"last_used_tm,omitempty"` // 最后使用时间
	Status     int    `gorm:"column:status" json:"status,omitempty"`           // 设备状态：0离线，1在线
	FaceDevice int    `gorm:"column:face_device" json:"face_device,omitempty"` // 是否为人脸设备：0否，1是
}
