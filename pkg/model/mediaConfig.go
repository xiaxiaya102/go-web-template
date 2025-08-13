package model

// MediaConfig 状态常量
const (
	EXCEPTION           = -1 // 异常
	NORMAL              = 0  // 正常
	CONNECT_ING         = 1  // 正在连接
	INVALID_IP          = 2  // 无效的ip
	INVALID_DATA        = 3  // 无权访问该资源
	NOT_AUTH            = 4  // 无权访问该资源
	ERROR_USER_PASSWORD = 5  // 无效的用户名或密码
	CONNECT_FAIL        = 6  // 连接视频失败
	FILE_NOT_FIND       = 7  // 设备视频文件没找到
)

// MediaConfig 类型常量
const (
	LOCAL_TYPE   = 0 // 设备内视频
	RTSP_TYPE    = 1 // rtsp
	GB28181_TYPE = 2 // gb28181
	RTMP_TYPE    = 3 // rtmp
)

type MediaConfig struct {
	ID        int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Name      string `gorm:"column:name" json:"name,omitempty"`
	Url       string `gorm:"column:url" json:"url,omitempty"`
	Status    string `gorm:"column:status" json:"status,omitempty"`
	Code      string `gorm:"column:code" json:"code,omitempty"`
	Type      string `gorm:"column:type" json:"type,omitempty"`
	MediaDesc string `gorm:"column:media_desc" json:"media_desc,omitempty"`
}
