package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
)

// SavePlanTemplate 保存计划模板
func SavePlanTemplate(planTemplate *model.PlanTemplate) error {
	return global.DB.Create(planTemplate).Error
}

// UpdatePlanTemplate 更新计划模板
func UpdatePlanTemplate(planTemplate *model.PlanTemplate) error {
	return global.DB.Save(planTemplate).Error
}

// DeletePlanTemplate 删除计划模板
func DeletePlanTemplate(id int) error {
	return global.DB.Delete(&model.PlanTemplate{}, id).Error
}

// GetPlanTemplateByID 根据ID获取计划模板
func GetPlanTemplateByID(id int) (*model.PlanTemplate, error) {
	var planTemplate model.PlanTemplate
	err := global.DB.First(&planTemplate, id).Error
	if err != nil {
		return nil, err
	}
	return &planTemplate, nil
}

// GetPlanTemplateList 获取计划模板列表
func GetPlanTemplateList() ([]model.PlanTemplate, error) {
	var planTemplates []model.PlanTemplate
	err := global.DB.Find(&planTemplates).Error
	return planTemplates, err
}

// GetPlanTemplatePage 分页获取计划模板列表
func GetPlanTemplatePage(page, pageSize int) ([]model.PlanTemplate, int64, error) {
	var planTemplates []model.PlanTemplate
	var total int64

	// 计算总数
	err := global.DB.Model(&model.PlanTemplate{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Find(&planTemplates).Error
	if err != nil {
		return nil, 0, err
	}

	return planTemplates, total, nil
}
