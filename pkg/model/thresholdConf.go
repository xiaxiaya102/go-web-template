package model

type ThresholdConf struct {
	ID         int     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	ParamDesc  string  `gorm:"column:param_desc" json:"param_desc,omitempty"`   // 参数描述
	ParamValue float64 `gorm:"column:param_value" json:"param_value,omitempty"` // 参数值
	SophonType int     `gorm:"column:sophon_type" json:"sophon_type,omitempty"` // Sophon类型
}
