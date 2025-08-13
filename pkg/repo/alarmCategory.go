package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
	"strings"
)

// SaveAlarmCategory 保存告警分类
func SaveAlarmCategory(alarmCategory *model.AlarmCategory) error {
	return global.DB.Create(alarmCategory).Error
}

// UpdateAlarmCategory 更新告警分类
func UpdateAlarmCategory(alarmCategory *model.AlarmCategory) error {
	return global.DB.Save(alarmCategory).Error
}

// BatchUpdateAlarmCategory 批量更新告警分类
func BatchUpdateAlarmCategory(alarmCategories []model.AlarmCategory) error {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, alarmCategory := range alarmCategories {
		if err := tx.Save(&alarmCategory).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// DeleteAlarmCategory 删除告警分类
func DeleteAlarmCategory(id int) error {
	return global.DB.Delete(&model.AlarmCategory{}, id).Error
}

// GetAlarmCategoryByID 根据ID获取告警分类
func GetAlarmCategoryByID(id int) (*model.AlarmCategory, error) {
	var alarmCategory model.AlarmCategory
	err := global.DB.First(&alarmCategory, id).Error
	if err != nil {
		return nil, err
	}
	return &alarmCategory, nil
}

// GetAlarmCategoryAll 获取所有告警分类
func GetAlarmCategoryAll() ([]model.AlarmCategory, error) {
	var alarmCategories []model.AlarmCategory
	err := global.DB.Find(&alarmCategories).Error
	return alarmCategories, err
}

// GetAlarmCategoryPage 分页获取告警分类列表
func GetAlarmCategoryPage(page, pageSize int) ([]model.AlarmCategory, int64, error) {
	var alarmCategories []model.AlarmCategory
	var total int64

	// 计算总数
	err := global.DB.Model(&model.AlarmCategory{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Order("id DESC").Find(&alarmCategories).Error
	if err != nil {
		return nil, 0, err
	}

	return alarmCategories, total, nil
}

// GetAlarmCategoryByArithmetic 根据算法获取告警分类列表
func GetAlarmCategoryByArithmetic(arithmetic string) ([]model.AlarmCategory, error) {
	var alarmCategories []model.AlarmCategory
	err := global.DB.Where("arithmetic = ?", arithmetic).Find(&alarmCategories).Error
	return alarmCategories, err
}

// GetAlarmCategoryByAlarmType 根据告警类型获取告警分类列表
func GetAlarmCategoryByAlarmType(alarmType string) ([]model.AlarmCategory, error) {
	var alarmCategories []model.AlarmCategory
	err := global.DB.Where("alarm_type = ?", alarmType).Find(&alarmCategories).Error
	return alarmCategories, err
}

// GetAlarmCategoryBySophonType 根据Sophon类型获取告警分类列表
func GetAlarmCategoryBySophonType(sophonType int) ([]model.AlarmCategory, error) {
	var alarmCategories []model.AlarmCategory
	err := global.DB.Where("sophon_type = ?", sophonType).Find(&alarmCategories).Error
	return alarmCategories, err
}

// GetAlarmCategoryByQuery 根据查询条件获取告警分类列表
func GetAlarmCategoryByQuery(arithmetic, alarmType string, sophonType *int, page, pageSize int) ([]model.AlarmCategory, int64, error) {
	var alarmCategories []model.AlarmCategory
	var total int64

	query := global.DB.Model(&model.AlarmCategory{})

	// 添加查询条件
	if arithmetic != "" {
		query = query.Where("arithmetic LIKE ?", "%"+arithmetic+"%")
	}
	if alarmType != "" {
		query = query.Where("alarm_type LIKE ?", "%"+alarmType+"%")
	}
	if sophonType != nil {
		query = query.Where("sophon_type = ?", *sophonType)
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&alarmCategories).Error
	if err != nil {
		return nil, 0, err
	}

	return alarmCategories, total, nil
}

// GetAlarmCategoryStatistics 获取告警分类统计
func GetAlarmCategoryStatistics() (map[string]interface{}, error) {
	statistics := make(map[string]interface{})

	// 总分类数
	var totalCount int64
	err := global.DB.Model(&model.AlarmCategory{}).Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	statistics["total_categories"] = totalCount

	// 算法统计
	var arithmeticStats []struct {
		Arithmetic string `json:"arithmetic"`
		Count      int64  `json:"count"`
	}
	err = global.DB.Model(&model.AlarmCategory{}).
		Select("arithmetic, COUNT(*) as count").
		Group("arithmetic").
		Find(&arithmeticStats).Error
	if err != nil {
		return nil, err
	}
	statistics["arithmetic_stats"] = arithmeticStats

	// 告警类型统计
	var alarmTypeStats []struct {
		AlarmType string `json:"alarm_type"`
		Count     int64  `json:"count"`
	}
	err = global.DB.Model(&model.AlarmCategory{}).
		Select("alarm_type, COUNT(*) as count").
		Group("alarm_type").
		Find(&alarmTypeStats).Error
	if err != nil {
		return nil, err
	}
	statistics["alarm_type_stats"] = alarmTypeStats

	// Sophon类型统计
	var sophonTypeStats []struct {
		SophonType int   `json:"sophon_type"`
		Count      int64 `json:"count"`
	}
	err = global.DB.Model(&model.AlarmCategory{}).
		Select("sophon_type, COUNT(*) as count").
		Group("sophon_type").
		Find(&sophonTypeStats).Error
	if err != nil {
		return nil, err
	}
	statistics["sophon_type_stats"] = sophonTypeStats

	// 语音文件统计
	var withAudio, withoutAudio int64
	err = global.DB.Model(&model.AlarmCategory{}).
		Where("audio_file != '' AND audio_file IS NOT NULL").
		Count(&withAudio).Error
	if err != nil {
		return nil, err
	}

	err = global.DB.Model(&model.AlarmCategory{}).
		Where("audio_file = '' OR audio_file IS NULL").
		Count(&withoutAudio).Error
	if err != nil {
		return nil, err
	}

	audioStats := map[string]int64{
		"with_audio":    withAudio,
		"without_audio": withoutAudio,
	}
	statistics["audio_file_stats"] = audioStats

	return statistics, nil
}

// SearchAlarmCategory 搜索告警分类
func SearchAlarmCategory(keyword string) ([]model.AlarmCategory, error) {
	var alarmCategories []model.AlarmCategory
	searchPattern := "%" + strings.TrimSpace(keyword) + "%"
	err := global.DB.Where("arithmetic LIKE ? OR alarm_type LIKE ? OR alarm_desc LIKE ?",
		searchPattern, searchPattern, searchPattern).Find(&alarmCategories).Error
	return alarmCategories, err
}

// CheckAlarmCategoryExists 检查告警分类是否存在
func CheckAlarmCategoryExists(arithmetic, alarmType string) (bool, error) {
	var count int64
	err := global.DB.Model(&model.AlarmCategory{}).
		Where("arithmetic = ? AND alarm_type = ?", arithmetic, alarmType).
		Count(&count).Error
	return count > 0, err
}

// GetAlarmCategoryCount 获取告警分类总数
func GetAlarmCategoryCount() (int64, error) {
	var count int64
	err := global.DB.Model(&model.AlarmCategory{}).Count(&count).Error
	return count, err
}

// GetAlarmCategoryCountByArithmetic 根据算法获取告警分类数量
func GetAlarmCategoryCountByArithmetic(arithmetic string) (int64, error) {
	var count int64
	err := global.DB.Model(&model.AlarmCategory{}).Where("arithmetic = ?", arithmetic).Count(&count).Error
	return count, err
}

// GetAlarmCategoryCountByAlarmType 根据告警类型获取告警分类数量
func GetAlarmCategoryCountByAlarmType(alarmType string) (int64, error) {
	var count int64
	err := global.DB.Model(&model.AlarmCategory{}).Where("alarm_type = ?", alarmType).Count(&count).Error
	return count, err
}

// GetAlarmCategoriesWithAudio 获取有语音文件的告警分类
func GetAlarmCategoriesWithAudio() ([]model.AlarmCategory, error) {
	var alarmCategories []model.AlarmCategory
	err := global.DB.Where("audio_file != '' AND audio_file IS NOT NULL").Find(&alarmCategories).Error
	return alarmCategories, err
}

// GetAlarmCategoriesWithoutAudio 获取无语音文件的告警分类
func GetAlarmCategoriesWithoutAudio() ([]model.AlarmCategory, error) {
	var alarmCategories []model.AlarmCategory
	err := global.DB.Where("audio_file = '' OR audio_file IS NULL").Find(&alarmCategories).Error
	return alarmCategories, err
}

// GetAllArithmetics 获取所有算法列表
func GetAllArithmetics() ([]string, error) {
	var arithmetics []string
	err := global.DB.Model(&model.AlarmCategory{}).Pluck("DISTINCT arithmetic", &arithmetics).Error
	return arithmetics, err
}

// GetAllAlarmTypes 获取所有告警类型列表
func GetAllAlarmTypes() ([]string, error) {
	var alarmTypes []string
	err := global.DB.Model(&model.AlarmCategory{}).Pluck("DISTINCT alarm_type", &alarmTypes).Error
	return alarmTypes, err
}
