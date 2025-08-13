package repo

import (
	"feishuReboot/global"
	"feishuReboot/pkg/model"
	"strings"
)

// SaveDeviceList 保存设备信息
func SaveDeviceList(deviceList *model.DeviceList) error {
	return global.DB.Create(deviceList).Error
}

// UpdateDeviceList 更新设备信息
func UpdateDeviceList(deviceList *model.DeviceList) error {
	return global.DB.Save(deviceList).Error
}

// UpdateDeviceStatus 更新设备状态
func UpdateDeviceStatus(id int, status int) error {
	return global.DB.Model(&model.DeviceList{}).Where("id = ?", id).Update("status", status).Error
}

// BatchUpdateDeviceStatus 批量更新设备状态
func BatchUpdateDeviceStatus(deviceIDs []int, status int) error {
	return global.DB.Model(&model.DeviceList{}).Where("id IN ?", deviceIDs).Update("status", status).Error
}

// DeleteDeviceList 删除设备信息
func DeleteDeviceList(id int) error {
	return global.DB.Delete(&model.DeviceList{}, id).Error
}

// GetDeviceListByID 根据ID获取设备信息
func GetDeviceListByID(id int) (*model.DeviceList, error) {
	var deviceList model.DeviceList
	err := global.DB.First(&deviceList, id).Error
	if err != nil {
		return nil, err
	}
	return &deviceList, nil
}

// GetDeviceListByIP 根据IP地址获取设备信息
func GetDeviceListByIP(ipAddress string) (*model.DeviceList, error) {
	var deviceList model.DeviceList
	err := global.DB.Where("ip_address = ?", ipAddress).First(&deviceList).Error
	if err != nil {
		return nil, err
	}
	return &deviceList, nil
}

// GetDeviceListAll 获取所有设备列表
func GetDeviceListAll() ([]model.DeviceList, error) {
	var deviceLists []model.DeviceList
	err := global.DB.Find(&deviceLists).Error
	return deviceLists, err
}

// GetDeviceListPage 分页获取设备列表
func GetDeviceListPage(page, pageSize int) ([]model.DeviceList, int64, error) {
	var deviceLists []model.DeviceList
	var total int64

	// 计算总数
	err := global.DB.Model(&model.DeviceList{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = global.DB.Offset(offset).Limit(pageSize).Order("id DESC").Find(&deviceLists).Error
	if err != nil {
		return nil, 0, err
	}

	return deviceLists, total, nil
}

// GetDeviceListByStatus 根据状态获取设备列表
func GetDeviceListByStatus(status int) ([]model.DeviceList, error) {
	var deviceLists []model.DeviceList
	err := global.DB.Where("status = ?", status).Find(&deviceLists).Error
	return deviceLists, err
}

// GetDeviceListByFaceDevice 根据是否为人脸设备获取设备列表
func GetDeviceListByFaceDevice(faceDevice int) ([]model.DeviceList, error) {
	var deviceLists []model.DeviceList
	err := global.DB.Where("face_device = ?", faceDevice).Find(&deviceLists).Error
	return deviceLists, err
}

// GetDeviceListByQuery 根据查询条件获取设备列表
func GetDeviceListByQuery(status *int, faceDevice *int, ipAddress string, page, pageSize int) ([]model.DeviceList, int64, error) {
	var deviceLists []model.DeviceList
	var total int64

	query := global.DB.Model(&model.DeviceList{})

	// 添加查询条件
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if faceDevice != nil {
		query = query.Where("face_device = ?", *faceDevice)
	}
	if ipAddress != "" {
		query = query.Where("ip_address LIKE ?", "%"+ipAddress+"%")
	}

	// 计算总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&deviceLists).Error
	if err != nil {
		return nil, 0, err
	}

	return deviceLists, total, nil
}

// GetDeviceStatusSummary 获取设备状态统计
func GetDeviceStatusSummary() (map[string]int64, error) {
	summary := make(map[string]int64)

	// 总设备数
	err := global.DB.Model(&model.DeviceList{}).Count(summary["total"]).Error
	if err != nil {
		return nil, err
	}

	// 在线设备数
	err = global.DB.Model(&model.DeviceList{}).Where("status = ?", model.DEVICE_ONLINE).Count(summary["online"]).Error
	if err != nil {
		return nil, err
	}

	// 离线设备数
	err = global.DB.Model(&model.DeviceList{}).Where("status = ?", model.DEVICE_OFFLINE).Count(summary["offline"]).Error
	if err != nil {
		return nil, err
	}

	// 人脸设备数
	err = global.DB.Model(&model.DeviceList{}).Where("face_device = ?", model.IS_FACE_DEVICE).Count(summary["face"]).Error
	if err != nil {
		return nil, err
	}

	// 非人脸设备数
	err = global.DB.Model(&model.DeviceList{}).Where("face_device = ?", model.NOT_FACE_DEVICE).Count(summary["non_face"]).Error
	if err != nil {
		return nil, err
	}

	return summary, nil
}

// UpdateDeviceLastUsedTime 更新设备最后使用时间
func UpdateDeviceLastUsedTime(id int, lastUsedTm string) error {
	return global.DB.Model(&model.DeviceList{}).Where("id = ?", id).Update("last_used_tm", lastUsedTm).Error
}

// UpdateDeviceLastUsedTimeByIP 根据IP更新设备最后使用时间
func UpdateDeviceLastUsedTimeByIP(ipAddress, lastUsedTm string) error {
	return global.DB.Model(&model.DeviceList{}).Where("ip_address = ?", ipAddress).Update("last_used_tm", lastUsedTm).Error
}

// CheckDeviceExists 检查设备是否存在
func CheckDeviceExists(ipAddress string) (bool, error) {
	var count int64
	err := global.DB.Model(&model.DeviceList{}).Where("ip_address = ?", ipAddress).Count(&count).Error
	return count > 0, err
}

// GetOnlineDevices 获取在线设备列表
func GetOnlineDevices() ([]model.DeviceList, error) {
	return GetDeviceListByStatus(model.DEVICE_ONLINE)
}

// GetOfflineDevices 获取离线设备列表
func GetOfflineDevices() ([]model.DeviceList, error) {
	return GetDeviceListByStatus(model.DEVICE_OFFLINE)
}

// GetFaceDevices 获取人脸设备列表
func GetFaceDevices() ([]model.DeviceList, error) {
	return GetDeviceListByFaceDevice(model.IS_FACE_DEVICE)
}

// SearchDevicesByIP 根据IP地址搜索设备
func SearchDevicesByIP(ipPattern string) ([]model.DeviceList, error) {
	var deviceLists []model.DeviceList
	searchPattern := "%" + strings.TrimSpace(ipPattern) + "%"
	err := global.DB.Where("ip_address LIKE ?", searchPattern).Find(&deviceLists).Error
	return deviceLists, err
}
