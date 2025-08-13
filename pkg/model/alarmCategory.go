package model

type AlarmCategory struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Arithmetic string `gorm:"column:arithmetic" json:"arithmetic,omitempty"`     // 所属算法
	AlarmType  string `gorm:"column:alarm_type" json:"alarm_type,omitempty"`    // 告警类别
	AlarmDesc  string `gorm:"column:alarm_desc" json:"alarm_desc,omitempty"`    // 告警描述
	AudioFile  string `gorm:"column:audio_file" json:"audio_file,omitempty"`    // 语音文件
	SophonType int    `gorm:"column:sophon_type" json:"sophon_type,omitempty"`  // Sophon类型
}
