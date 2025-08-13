package model

type PlanTemplate struct {
	ID   int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Name string `gorm:"column:name" json:"name,omitempty"`
	Plan string `gorm:"column:plan" json:"plan,omitempty"`
}
