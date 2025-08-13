package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
)

// SaveMediaConfig 保存媒体配置
func SaveMediaConfig(mediaConfig *model.MediaConfig) error {
	return global.DB.Create(mediaConfig).Error
}

// UpdateMediaConfig 更新媒体配置
func UpdateMediaConfig(mediaConfig *model.MediaConfig) error {
	return global.DB.Save(mediaConfig).Error
}

// DeleteMediaConfig 删除媒体配置
func DeleteMediaConfig(id int) error {
	return global.DB.Delete(&model.MediaConfig{}, id).Error
}

// GetMediaConfigByID 根据ID获取媒体配置
func GetMediaConfigByID(id int) (*model.MediaConfig, error) {
	var mediaConfig model.MediaConfig
	err := global.DB.First(&mediaConfig, id).Error
	if err != nil {
		return nil, err
	}
	return &mediaConfig, nil
}

// GetAllMediaConfigs 获取所有媒体配置
func GetAllMediaConfigs() ([]model.MediaConfig, error) {
	var mediaConfigs []model.MediaConfig
	err := global.DB.Find(&mediaConfigs).Error
	return mediaConfigs, err
}

// GetMediaConfigPage 分页获取媒体配置
func GetMediaConfigPage(page, pageSize int) ([]model.MediaConfig, int64, error) {
	var mediaConfigs []model.MediaConfig
	var total int64

	// 获取总数
	global.DB.Model(&model.MediaConfig{}).Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := global.DB.Offset(offset).Limit(pageSize).Find(&mediaConfigs).Error

	return mediaConfigs, total, err
}

// GetMediaConfigByName 根据名称获取媒体配置
func GetMediaConfigByName(name string) (*model.MediaConfig, error) {
	var mediaConfig model.MediaConfig
	err := global.DB.Where("name = ?", name).First(&mediaConfig).Error
	if err != nil {
		return nil, err
	}
	return &mediaConfig, nil
}
