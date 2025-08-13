package model

type ModelsInfo struct {
	ID          int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Name        string `gorm:"column:name" json:"name,omitempty"`               // 模型名称
	Algorithm   string `gorm:"column:algorithm" json:"algorithm,omitempty"`     // 算法
	DateTime    string `gorm:"column:date_time" json:"date_time,omitempty"`     // 日期时间
	Description string `gorm:"column:description" json:"description,omitempty"` // 描述
	SophonType  int    `gorm:"column:sophon_type" json:"sophon_type,omitempty"` // Sophon类型
	Classify    string `gorm:"column:classify" json:"classify,omitempty"`       // 分类
	Specialty   int    `gorm:"column:specialty" json:"specialty,omitempty"`     // 专业性
}
