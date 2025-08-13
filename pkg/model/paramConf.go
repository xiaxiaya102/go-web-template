package model

type ParamConf struct {
	ID         int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Key        string `gorm:"column:key" json:"key,omitempty"`           // 配置键
	Param      string `gorm:"column:param" json:"param,omitempty"`       // 参数名称
	ParamValue string `gorm:"column:param_value" json:"param_value,omitempty"` // 参数值
}
