package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
	"strings"
)

// SaveParamConf 保存参数配置
func SaveParamConf(paramConf *model.ParamConf) error {
	return global.DB.Create(paramConf).Error
}

// UpdateParamConf 更新参数配置
func UpdateParamConf(paramConf *model.ParamConf) error {
	return global.DB.Save(paramConf).Error
}

// UpdateParamValue 更新参数值
func UpdateParamValue(id int, paramValue string) error {
	return global.DB.Model(&model.ParamConf{}).Where("id = ?", id).Update("param_value", paramValue).Error
}

// BatchUpdateParamConf 批量更新参数配置
func BatchUpdateParamConf(paramConfs []model.ParamConf) error {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, paramConf := range paramConfs {
		if err := tx.Save(&paramConf).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// DeleteParamConf 删除参数配置
func DeleteParamConf(id int) error {
	return global.DB.Delete(&model.ParamConf{}, id).Error
}

// DeleteParamConfByKey 根据配置键删除所有相关配置
func DeleteParamConfByKey(key string) error {
	return global.DB.Where("key = ?", key).Delete(&model.ParamConf{}).Error
}

// GetParamConfByID 根据ID获取参数配置
func GetParamConfByID(id int) (*model.ParamConf, error) {
	var paramConf model.ParamConf
	err := global.DB.First(&paramConf, id).Error
	if err != nil {
		return nil, err
	}
	return &paramConf, nil
}

// GetParamConfByKey 根据配置键获取参数配置列表
func GetParamConfByKey(key string) ([]model.ParamConf, error) {
	var paramConfs []model.ParamConf
	err := global.DB.Where("key = ?", key).Find(&paramConfs).Error
	return paramConfs, err
}

// GetParamConfByKeyAndParam 根据配置键和参数名获取配置
func GetParamConfByKeyAndParam(key, param string) (*model.ParamConf, error) {
	var paramConf model.ParamConf
	err := global.DB.Where("key = ? AND param = ?", key, param).First(&paramConf).Error
	if err != nil {
		return nil, err
	}
	return &paramConf, nil
}

// GetParamConfAll 获取所有参数配置
func GetParamConfAll() ([]model.ParamConf, error) {
	var paramConfs []model.ParamConf
	err := global.DB.Find(&paramConfs).Error
	return paramConfs, err
}

// GetParamConfPage 分页获取参数配置列表
func GetParamConfPage(page, pageSize int) ([]model.ParamConf, int64, error) {
	var paramConfs []model.ParamConf
	var total int64

	// 计算总数
	err := global.DB.Model(&model.ParamConf{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Order("id DESC").Find(&paramConfs).Error
	if err != nil {
		return nil, 0, err
	}

	return paramConfs, total, nil
}

// GetParamConfByQuery 根据查询条件获取参数配置列表
func GetParamConfByQuery(key, param string, page, pageSize int) ([]model.ParamConf, int64, error) {
	var paramConfs []model.ParamConf
	var total int64

	query := global.DB.Model(&model.ParamConf{})

	// 添加查询条件
	if key != "" {
		query = query.Where("key LIKE ?", "%"+key+"%")
	}
	if param != "" {
		query = query.Where("param LIKE ?", "%"+param+"%")
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&paramConfs).Error
	if err != nil {
		return nil, 0, err
	}

	return paramConfs, total, nil
}

// GetAllParamKeys 获取所有配置键
func GetAllParamKeys() ([]string, error) {
	var keys []string
	err := global.DB.Model(&model.ParamConf{}).Pluck("DISTINCT key", &keys).Error
	return keys, err
}

// CheckParamConfExists 检查参数配置是否存在
func CheckParamConfExists(key, param string) (bool, error) {
	var count int64
	err := global.DB.Model(&model.ParamConf{}).Where("key = ? AND param = ?", key, param).Count(&count).Error
	return count > 0, err
}

// GetParamValue 获取参数值
func GetParamValue(key, param string) (string, error) {
	var paramConf model.ParamConf
	err := global.DB.Where("key = ? AND param = ?", key, param).First(&paramConf).Error
	if err != nil {
		return "", err
	}
	return paramConf.ParamValue, nil
}

// SetParamValue 设置参数值（如果不存在则创建，存在则更新）
func SetParamValue(key, param, paramValue string) error {
	var paramConf model.ParamConf
	err := global.DB.Where("key = ? AND param = ?", key, param).First(&paramConf).Error

	if err != nil {
		// 不存在，创建新记录
		newParamConf := &model.ParamConf{
			Key:        key,
			Param:      param,
			ParamValue: paramValue,
		}
		return global.DB.Create(newParamConf).Error
	} else {
		// 存在，更新记录
		return global.DB.Model(&paramConf).Update("param_value", paramValue).Error
	}
}

// SearchParamConf 搜索参数配置
func SearchParamConf(keyword string) ([]model.ParamConf, error) {
	var paramConfs []model.ParamConf
	searchPattern := "%" + strings.TrimSpace(keyword) + "%"
	err := global.DB.Where("key LIKE ? OR param LIKE ? OR param_value LIKE ?",
		searchPattern, searchPattern, searchPattern).Find(&paramConfs).Error
	return paramConfs, err
}

// GetParamConfCount 获取参数配置总数
func GetParamConfCount() (int64, error) {
	var count int64
	err := global.DB.Model(&model.ParamConf{}).Count(&count).Error
	return count, err
}

// GetParamConfCountByKey 根据配置键获取参数配置数量
func GetParamConfCountByKey(key string) (int64, error) {
	var count int64
	err := global.DB.Model(&model.ParamConf{}).Where("key = ?", key).Count(&count).Error
	return count, err
}
