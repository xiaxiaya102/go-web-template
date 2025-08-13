package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
)

// SaveTaskManage 保存任务管理
func SaveTaskManage(taskManage *model.TaskManage) error {
	return global.DB.Create(taskManage).Error
}

// UpdateTaskManage 更新任务管理
func UpdateTaskManage(taskManage *model.TaskManage) error {
	return global.DB.Save(taskManage).Error
}

// UpdateTaskStatus 更新任务状态
func UpdateTaskStatus(id int, taskStatus int) error {
	return global.DB.Model(&model.TaskManage{}).Where("id = ?", id).Update("task_status", taskStatus).Error
}

// DeleteTaskManage 删除任务管理
func DeleteTaskManage(id int) error {
	return global.DB.Delete(&model.TaskManage{}, id).Error
}

// GetTaskManageByID 根据ID获取任务管理
func GetTaskManageByID(id int) (*model.TaskManage, error) {
	var taskManage model.TaskManage
	err := global.DB.First(&taskManage, id).Error
	if err != nil {
		return nil, err
	}
	return &taskManage, nil
}

// GetTaskManageList 获取任务管理列表
func GetTaskManageList() ([]model.TaskManage, error) {
	var taskManages []model.TaskManage
	err := global.DB.Find(&taskManages).Error
	return taskManages, err
}

// GetTaskManagePage 分页获取任务管理列表
func GetTaskManagePage(page, pageSize int) ([]model.TaskManage, int64, error) {
	var taskManages []model.TaskManage
	var total int64

	// 计算总数
	err := global.DB.Model(&model.TaskManage{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Find(&taskManages).Error
	if err != nil {
		return nil, 0, err
	}

	return taskManages, total, nil
}

// GetTaskManageByStatus 根据状态获取任务列表
func GetTaskManageByStatus(status int) ([]model.TaskManage, error) {
	var taskManages []model.TaskManage
	err := global.DB.Where("task_status = ?", status).Find(&taskManages).Error
	return taskManages, err
}

// GetTaskManageByMediaID 根据媒体ID获取任务列表
func GetTaskManageByMediaID(mediaID int) ([]model.TaskManage, error) {
	var taskManages []model.TaskManage
	err := global.DB.Where("media_id = ?", mediaID).Find(&taskManages).Error
	return taskManages, err
}

// GetTaskManageByPlanID 根据计划ID获取任务列表
func GetTaskManageByPlanID(planID int) ([]model.TaskManage, error) {
	var taskManages []model.TaskManage
	err := global.DB.Where("plan_id = ?", planID).Find(&taskManages).Error
	return taskManages, err
}
