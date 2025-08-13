package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
	_ "fmt"
)

// SaveAlarmManage 保存报警管理
func SaveAlarmManage(alarmManage *model.AlarmManage) error {
	return global.DB.Create(alarmManage).Error
}

// UpdateAlarmManage 更新报警管理
func UpdateAlarmManage(alarmManage *model.AlarmManage) error {
	return global.DB.Save(alarmManage).Error
}

// UpdateReportStatus 更新报告状态
func UpdateReportStatus(id int, reportStatus int) error {
	return global.DB.Model(&model.AlarmManage{}).Where("id = ?", id).Update("report_status", reportStatus).Error
}

// DeleteAlarmManage 删除报警管理
func DeleteAlarmManage(id int) error {
	return global.DB.Delete(&model.AlarmManage{}, id).Error
}

// GetAlarmManageByID 根据ID获取报警管理
func GetAlarmManageByID(id int) (*model.AlarmManage, error) {
	var alarmManage model.AlarmManage
	err := global.DB.First(&alarmManage, id).Error
	if err != nil {
		return nil, err
	}
	return &alarmManage, nil
}

// GetAlarmManageList 获取报警管理列表
func GetAlarmManageList() ([]model.AlarmManage, error) {
	var alarmManages []model.AlarmManage
	err := global.DB.Find(&alarmManages).Error
	return alarmManages, err
}

// GetAlarmManagePage 分页获取报警管理列表
func GetAlarmManagePage(page, pageSize int) ([]model.AlarmManage, int64, error) {
	var alarmManages []model.AlarmManage
	var total int64

	// 计算总数
	err := global.DB.Model(&model.AlarmManage{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Order("id DESC").Find(&alarmManages).Error
	if err != nil {
		return nil, 0, err
	}

	return alarmManages, total, nil
}

// GetAlarmManageByReportStatus 根据报告状态获取报警列表
func GetAlarmManageByReportStatus(reportStatus int) ([]model.AlarmManage, error) {
	var alarmManages []model.AlarmManage
	err := global.DB.Where("report_status = ?", reportStatus).Find(&alarmManages).Error
	return alarmManages, err
}

// GetAlarmManageByAlarmType 根据报警类型获取报警列表
func GetAlarmManageByAlarmType(alarmType string) ([]model.AlarmManage, error) {
	var alarmManages []model.AlarmManage
	err := global.DB.Where("alarm_type = ?", alarmType).Find(&alarmManages).Error
	return alarmManages, err
}

// GetAlarmManageByMediaID 根据媒体ID获取报警列表
func GetAlarmManageByMediaID(mediaID int) ([]model.AlarmManage, error) {
	var alarmManages []model.AlarmManage
	err := global.DB.Where("media_id = ?", mediaID).Find(&alarmManages).Error
	return alarmManages, err
}

// GetAlarmManageByQuery 根据查询条件获取报警列表
func GetAlarmManageByQuery(alarmType string, reportStatus *int, mediaID *int, startTime, endTime string, page, pageSize int) ([]model.AlarmManage, int64, error) {
	var alarmManages []model.AlarmManage
	var total int64

	query := global.DB.Model(&model.AlarmManage{})

	// 添加查询条件
	if alarmType != "" {
		query = query.Where("alarm_type = ?", alarmType)
	}
	if reportStatus != nil {
		query = query.Where("report_status = ?", *reportStatus)
	}
	if mediaID != nil {
		query = query.Where("media_id = ?", *mediaID)
	}
	if startTime != "" && endTime != "" {
		query = query.Where("tm BETWEEN ? AND ?", startTime, endTime)
	} else if startTime != "" {
		query = query.Where("tm >= ?", startTime)
	} else if endTime != "" {
		query = query.Where("tm <= ?", endTime)
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&alarmManages).Error
	if err != nil {
		return nil, 0, err
	}

	return alarmManages, total, nil
}
