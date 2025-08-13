package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
	"strings"
)

// SaveThresholdConf 保存阈值配置
func SaveThresholdConf(thresholdConf *model.ThresholdConf) error {
	return global.DB.Create(thresholdConf).Error
}

// UpdateThresholdConf 更新阈值配置
func UpdateThresholdConf(thresholdConf *model.ThresholdConf) error {
	return global.DB.Save(thresholdConf).Error
}

// UpdateThresholdValue 更新阈值参数值
func UpdateThresholdValue(id int, paramValue float64) error {
	return global.DB.Model(&model.ThresholdConf{}).Where("id = ?", id).Update("param_value", paramValue).Error
}

// BatchUpdateThresholdConf 批量更新阈值配置
func BatchUpdateThresholdConf(thresholdConfs []model.ThresholdConf) error {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, thresholdConf := range thresholdConfs {
		if err := tx.Save(&thresholdConf).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// DeleteThresholdConf 删除阈值配置
func DeleteThresholdConf(id int) error {
	return global.DB.Delete(&model.ThresholdConf{}, id).Error
}

// GetThresholdConfByID 根据ID获取阈值配置
func GetThresholdConfByID(id int) (*model.ThresholdConf, error) {
	var thresholdConf model.ThresholdConf
	err := global.DB.First(&thresholdConf, id).Error
	if err != nil {
		return nil, err
	}
	return &thresholdConf, nil
}

// GetThresholdConfAll 获取所有阈值配置
func GetThresholdConfAll() ([]model.ThresholdConf, error) {
	var thresholdConfs []model.ThresholdConf
	err := global.DB.Find(&thresholdConfs).Error
	return thresholdConfs, err
}

// GetThresholdConfPage 分页获取阈值配置列表
func GetThresholdConfPage(page, pageSize int) ([]model.ThresholdConf, int64, error) {
	var thresholdConfs []model.ThresholdConf
	var total int64

	// 计算总数
	err := global.DB.Model(&model.ThresholdConf{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Order("id DESC").Find(&thresholdConfs).Error
	if err != nil {
		return nil, 0, err
	}

	return thresholdConfs, total, nil
}

// GetThresholdConfBySophonType 根据Sophon类型获取阈值配置列表
func GetThresholdConfBySophonType(sophonType int) ([]model.ThresholdConf, error) {
	var thresholdConfs []model.ThresholdConf
	err := global.DB.Where("sophon_type = ?", sophonType).Find(&thresholdConfs).Error
	return thresholdConfs, err
}

// GetThresholdConfByParamDesc 根据参数描述获取阈值配置
func GetThresholdConfByParamDesc(paramDesc string) (*model.ThresholdConf, error) {
	var thresholdConf model.ThresholdConf
	err := global.DB.Where("param_desc = ?", paramDesc).First(&thresholdConf).Error
	if err != nil {
		return nil, err
	}
	return &thresholdConf, nil
}

// GetThresholdConfByValueRange 根据值范围获取阈值配置列表
func GetThresholdConfByValueRange(minValue, maxValue float64) ([]model.ThresholdConf, error) {
	var thresholdConfs []model.ThresholdConf
	err := global.DB.Where("param_value BETWEEN ? AND ?", minValue, maxValue).Find(&thresholdConfs).Error
	return thresholdConfs, err
}

// GetThresholdConfByQuery 根据查询条件获取阈值配置列表
func GetThresholdConfByQuery(paramDesc string, sophonType *int, minValue, maxValue *float64, page, pageSize int) ([]model.ThresholdConf, int64, error) {
	var thresholdConfs []model.ThresholdConf
	var total int64

	query := global.DB.Model(&model.ThresholdConf{})

	// 添加查询条件
	if paramDesc != "" {
		query = query.Where("param_desc LIKE ?", "%"+paramDesc+"%")
	}
	if sophonType != nil {
		query = query.Where("sophon_type = ?", *sophonType)
	}
	if minValue != nil && maxValue != nil {
		query = query.Where("param_value BETWEEN ? AND ?", *minValue, *maxValue)
	} else if minValue != nil {
		query = query.Where("param_value >= ?", *minValue)
	} else if maxValue != nil {
		query = query.Where("param_value <= ?", *maxValue)
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&thresholdConfs).Error
	if err != nil {
		return nil, 0, err
	}

	return thresholdConfs, total, nil
}

// GetThresholdConfStatistics 获取阈值配置统计
func GetThresholdConfStatistics() (map[string]interface{}, error) {
	statistics := make(map[string]interface{})

	// 总阈值数
	var totalCount int64
	err := global.DB.Model(&model.ThresholdConf{}).Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	statistics["total_thresholds"] = totalCount

	// Sophon类型统计
	var sophonTypeStats []struct {
		SophonType int   `json:"sophon_type"`
		Count      int64 `json:"count"`
	}
	err = global.DB.Model(&model.ThresholdConf{}).
		Select("sophon_type, COUNT(*) as count").
		Group("sophon_type").
		Find(&sophonTypeStats).Error
	if err != nil {
		return nil, err
	}
	statistics["sophon_type_stats"] = sophonTypeStats

	// 值范围统计
	var valueRangeStats struct {
		MinValue float64 `json:"min_value"`
		MaxValue float64 `json:"max_value"`
		AvgValue float64 `json:"avg_value"`
	}
	err = global.DB.Model(&model.ThresholdConf{}).
		Select("MIN(param_value) as min_value, MAX(param_value) as max_value, AVG(param_value) as avg_value").
		Scan(&valueRangeStats).Error
	if err != nil {
		return nil, err
	}
	statistics["value_range_stats"] = valueRangeStats

	// 参数描述统计
	var paramDescStats []struct {
		ParamDesc string `json:"param_desc"`
		Count     int64  `json:"count"`
	}
	err = global.DB.Model(&model.ThresholdConf{}).
		Select("param_desc, COUNT(*) as count").
		Group("param_desc").
		Find(&paramDescStats).Error
	if err != nil {
		return nil, err
	}
	statistics["param_desc_stats"] = paramDescStats

	return statistics, nil
}

// SearchThresholdConf 搜索阈值配置
func SearchThresholdConf(keyword string) ([]model.ThresholdConf, error) {
	var thresholdConfs []model.ThresholdConf
	searchPattern := "%" + strings.TrimSpace(keyword) + "%"
	err := global.DB.Where("param_desc LIKE ?", searchPattern).Find(&thresholdConfs).Error
	return thresholdConfs, err
}

// CheckThresholdConfExists 检查阈值配置是否存在
func CheckThresholdConfExists(paramDesc string, sophonType int) (bool, error) {
	var count int64
	err := global.DB.Model(&model.ThresholdConf{}).
		Where("param_desc = ? AND sophon_type = ?", paramDesc, sophonType).
		Count(&count).Error
	return count > 0, err
}

// GetThresholdConfCount 获取阈值配置总数
func GetThresholdConfCount() (int64, error) {
	var count int64
	err := global.DB.Model(&model.ThresholdConf{}).Count(&count).Error
	return count, err
}

// GetThresholdConfCountBySophonType 根据Sophon类型获取阈值配置数量
func GetThresholdConfCountBySophonType(sophonType int) (int64, error) {
	var count int64
	err := global.DB.Model(&model.ThresholdConf{}).Where("sophon_type = ?", sophonType).Count(&count).Error
	return count, err
}

// GetThresholdValue 获取阈值参数值
func GetThresholdValue(paramDesc string, sophonType int) (float64, error) {
	var thresholdConf model.ThresholdConf
	err := global.DB.Where("param_desc = ? AND sophon_type = ?", paramDesc, sophonType).First(&thresholdConf).Error
	if err != nil {
		return 0, err
	}
	return thresholdConf.ParamValue, nil
}

// SetThresholdValue 设置阈值参数值（如果不存在则创建，存在则更新）
func SetThresholdValue(paramDesc string, paramValue float64, sophonType int) error {
	var thresholdConf model.ThresholdConf
	err := global.DB.Where("param_desc = ? AND sophon_type = ?", paramDesc, sophonType).First(&thresholdConf).Error

	if err != nil {
		// 不存在，创建新记录
		newThresholdConf := &model.ThresholdConf{
			ParamDesc:  paramDesc,
			ParamValue: paramValue,
			SophonType: sophonType,
		}
		return global.DB.Create(newThresholdConf).Error
	} else {
		// 存在，更新记录
		return global.DB.Model(&thresholdConf).Update("param_value", paramValue).Error
	}
}

// GetAllParamDescs 获取所有参数描述列表
func GetAllParamDescs() ([]string, error) {
	var paramDescs []string
	err := global.DB.Model(&model.ThresholdConf{}).Pluck("DISTINCT param_desc", &paramDescs).Error
	return paramDescs, err
}

// GetAllSophonTypes 获取所有Sophon类型列表
func GetAllSophonTypes() ([]int, error) {
	var sophonTypes []int
	err := global.DB.Model(&model.ThresholdConf{}).Pluck("DISTINCT sophon_type", &sophonTypes).Error
	return sophonTypes, err
}
