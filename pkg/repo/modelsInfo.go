package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
	"strings"
)

// SaveModelsInfo 保存模型信息
func SaveModelsInfo(modelsInfo *model.ModelsInfo) error {
	return global.DB.Create(modelsInfo).Error
}

// UpdateModelsInfo 更新模型信息
func UpdateModelsInfo(modelsInfo *model.ModelsInfo) error {
	return global.DB.Save(modelsInfo).Error
}

// BatchUpdateModelsInfo 批量更新模型信息
func BatchUpdateModelsInfo(modelsInfos []model.ModelsInfo) error {
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, modelsInfo := range modelsInfos {
		if err := tx.Save(&modelsInfo).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// DeleteModelsInfo 删除模型信息
func DeleteModelsInfo(id int) error {
	return global.DB.Delete(&model.ModelsInfo{}, id).Error
}

// GetModelsInfoByID 根据ID获取模型信息
func GetModelsInfoByID(id int) (*model.ModelsInfo, error) {
	var modelsInfo model.ModelsInfo
	err := global.DB.First(&modelsInfo, id).Error
	if err != nil {
		return nil, err
	}
	return &modelsInfo, nil
}

// GetModelsInfoByName 根据名称获取模型信息
func GetModelsInfoByName(name string) (*model.ModelsInfo, error) {
	var modelsInfo model.ModelsInfo
	err := global.DB.Where("name = ?", name).First(&modelsInfo).Error
	if err != nil {
		return nil, err
	}
	return &modelsInfo, nil
}

// GetModelsInfoAll 获取所有模型信息
func GetModelsInfoAll() ([]model.ModelsInfo, error) {
	var modelsInfos []model.ModelsInfo
	err := global.DB.Find(&modelsInfos).Error
	return modelsInfos, err
}

// GetModelsInfoPage 分页获取模型信息列表
func GetModelsInfoPage(page, pageSize int) ([]model.ModelsInfo, int64, error) {
	var modelsInfos []model.ModelsInfo
	var total int64

	// 计算总数
	err := global.DB.Model(&model.ModelsInfo{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Order("id DESC").Find(&modelsInfos).Error
	if err != nil {
		return nil, 0, err
	}

	return modelsInfos, total, nil
}

// GetModelsInfoByAlgorithm 根据算法获取模型信息列表
func GetModelsInfoByAlgorithm(algorithm string) ([]model.ModelsInfo, error) {
	var modelsInfos []model.ModelsInfo
	err := global.DB.Where("algorithm = ?", algorithm).Find(&modelsInfos).Error
	return modelsInfos, err
}

// GetModelsInfoByClassify 根据分类获取模型信息列表
func GetModelsInfoByClassify(classify string) ([]model.ModelsInfo, error) {
	var modelsInfos []model.ModelsInfo
	err := global.DB.Where("classify = ?", classify).Find(&modelsInfos).Error
	return modelsInfos, err
}

// GetModelsInfoBySophonType 根据Sophon类型获取模型信息列表
func GetModelsInfoBySophonType(sophonType int) ([]model.ModelsInfo, error) {
	var modelsInfos []model.ModelsInfo
	err := global.DB.Where("sophon_type = ?", sophonType).Find(&modelsInfos).Error
	return modelsInfos, err
}

// GetModelsInfoBySpecialty 根据专业性获取模型信息列表
func GetModelsInfoBySpecialty(specialty int) ([]model.ModelsInfo, error) {
	var modelsInfos []model.ModelsInfo
	err := global.DB.Where("specialty = ?", specialty).Find(&modelsInfos).Error
	return modelsInfos, err
}

// GetModelsInfoByQuery 根据查询条件获取模型信息列表
func GetModelsInfoByQuery(name, algorithm, classify string, sophonType, specialty *int, page, pageSize int) ([]model.ModelsInfo, int64, error) {
	var modelsInfos []model.ModelsInfo
	var total int64

	query := global.DB.Model(&model.ModelsInfo{})

	// 添加查询条件
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if algorithm != "" {
		query = query.Where("algorithm LIKE ?", "%"+algorithm+"%")
	}
	if classify != "" {
		query = query.Where("classify LIKE ?", "%"+classify+"%")
	}
	if sophonType != nil {
		query = query.Where("sophon_type = ?", *sophonType)
	}
	if specialty != nil {
		query = query.Where("specialty = ?", *specialty)
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&modelsInfos).Error
	if err != nil {
		return nil, 0, err
	}

	return modelsInfos, total, nil
}

// GetModelsInfoStatistics 获取模型信息统计
func GetModelsInfoStatistics() (map[string]interface{}, error) {
	statistics := make(map[string]interface{})

	// 总模型数
	var totalCount int64
	err := global.DB.Model(&model.ModelsInfo{}).Count(&totalCount).Error
	if err != nil {
		return nil, err
	}
	statistics["total_models"] = totalCount

	// 算法统计
	var algorithmStats []struct {
		Algorithm string `json:"algorithm"`
		Count     int64  `json:"count"`
	}
	err = global.DB.Model(&model.ModelsInfo{}).
		Select("algorithm, COUNT(*) as count").
		Group("algorithm").
		Find(&algorithmStats).Error
	if err != nil {
		return nil, err
	}
	statistics["algorithm_stats"] = algorithmStats

	// 分类统计
	var classifyStats []struct {
		Classify string `json:"classify"`
		Count    int64  `json:"count"`
	}
	err = global.DB.Model(&model.ModelsInfo{}).
		Select("classify, COUNT(*) as count").
		Group("classify").
		Find(&classifyStats).Error
	if err != nil {
		return nil, err
	}
	statistics["classify_stats"] = classifyStats

	// Sophon类型统计
	var sophonTypeStats []struct {
		SophonType int   `json:"sophon_type"`
		Count      int64 `json:"count"`
	}
	err = global.DB.Model(&model.ModelsInfo{}).
		Select("sophon_type, COUNT(*) as count").
		Group("sophon_type").
		Find(&sophonTypeStats).Error
	if err != nil {
		return nil, err
	}
	statistics["sophon_type_stats"] = sophonTypeStats

	// 专业性统计
	var specialtyStats []struct {
		Specialty int   `json:"specialty"`
		Count     int64 `json:"count"`
	}
	err = global.DB.Model(&model.ModelsInfo{}).
		Select("specialty, COUNT(*) as count").
		Group("specialty").
		Find(&specialtyStats).Error
	if err != nil {
		return nil, err
	}
	statistics["specialty_stats"] = specialtyStats

	return statistics, nil
}

// SearchModelsInfo 搜索模型信息
func SearchModelsInfo(keyword string) ([]model.ModelsInfo, error) {
	var modelsInfos []model.ModelsInfo
	searchPattern := "%" + strings.TrimSpace(keyword) + "%"
	err := global.DB.Where("name LIKE ? OR algorithm LIKE ? OR description LIKE ? OR classify LIKE ?",
		searchPattern, searchPattern, searchPattern, searchPattern).Find(&modelsInfos).Error
	return modelsInfos, err
}

// CheckModelsInfoExists 检查模型信息是否存在
func CheckModelsInfoExists(name string) (bool, error) {
	var count int64
	err := global.DB.Model(&model.ModelsInfo{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

// GetAllAlgorithms 获取所有算法列表
func GetAllAlgorithms() ([]string, error) {
	var algorithms []string
	err := global.DB.Model(&model.ModelsInfo{}).Pluck("DISTINCT algorithm", &algorithms).Error

	return algorithms, err
}

// GetAllClassifies 获取所有分类列表
func GetAllClassifies() ([]string, error) {
	var classifies []string
	err := global.DB.Model(&model.ModelsInfo{}).Pluck("DISTINCT classify", &classifies).Error

	return classifies, err
}

// GetModelsInfoCount 获取模型信息总数
func GetModelsInfoCount() (int64, error) {
	var count int64
	err := global.DB.Model(&model.ModelsInfo{}).Count(&count).Error
	return count, err
}

// GetModelsInfoCountByAlgorithm 根据算法获取模型数量
func GetModelsInfoCountByAlgorithm(algorithm string) (int64, error) {
	var count int64
	err := global.DB.Model(&model.ModelsInfo{}).Where("algorithm = ?", algorithm).Count(&count).Error
	return count, err
}

// GetModelsInfoCountByClassify 根据分类获取模型数量
func GetModelsInfoCountByClassify(classify string) (int64, error) {
	var count int64
	err := global.DB.Model(&model.ModelsInfo{}).Where("classify = ?", classify).Count(&count).Error
	return count, err
}
