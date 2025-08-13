package system

import (
	"feishuReboot/pkg/dto"
	"feishuReboot/pkg/handle"
	"feishuReboot/pkg/model"
	"feishuReboot/pkg/repo"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DeviceListApi struct{}

// SaveDeviceList 保存设备信息
// @Summary 保存设备信息
// @Description 保存新的设备信息
// @Tags 设备管理
// @Accept json
// @Produce json
// @Param request body dto.SaveDeviceListRequest true "设备信息"
// @Success 200 {object} handle.Response{data=model.DeviceList} "保存成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /deviceList/saveDeviceList [post]
func (d *DeviceListApi) SaveDeviceList(c *gin.Context) {
	var req dto.SaveDeviceListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查IP地址是否已存在
	exists, err := repo.CheckDeviceExists(req.IPAddress)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查设备失败"))
		return
	}
	if exists {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设备IP地址已存在"))
		return
	}

	// 如果没有提供最后使用时间，使用当前时间
	if req.LastUsedTm == "" {
		req.LastUsedTm = time.Now().Format("2006-01-02 15:04:05")
	}

	deviceList := &model.DeviceList{
		IPAddress:  req.IPAddress,
		LastUsedTm: req.LastUsedTm,
		Status:     req.Status,
		FaceDevice: req.FaceDevice,
	}

	if err := repo.SaveDeviceList(deviceList); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "保存失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(deviceList))
}

// UpdateDeviceList 更新设备信息
func (d *DeviceListApi) UpdateDeviceList(c *gin.Context) {
	var req dto.UpdateDeviceListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 检查设备是否存在
	existingDevice, err := repo.GetDeviceListByID(req.ID)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设备不存在"))
		return
	}

	// 如果IP地址发生变化，检查新IP是否已被其他设备使用
	if existingDevice.IPAddress != req.IPAddress {
		exists, err := repo.CheckDeviceExists(req.IPAddress)
		if err != nil {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "检查设备失败"))
			return
		}
		if exists {
			c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设备IP地址已存在"))
			return
		}
	}

	deviceList := &model.DeviceList{
		ID:         req.ID,
		IPAddress:  req.IPAddress,
		LastUsedTm: req.LastUsedTm,
		Status:     req.Status,
		FaceDevice: req.FaceDevice,
	}

	if err := repo.UpdateDeviceList(deviceList); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(deviceList))
}

// UpdateDeviceStatus 更新设备状态
func (d *DeviceListApi) UpdateDeviceStatus(c *gin.Context) {
	var req dto.UpdateDeviceStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.UpdateDeviceStatus(req.ID, req.Status); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "更新状态失败"))
		return
	}

	// 如果设备上线，更新最后使用时间
	if req.Status == model.DEVICE_ONLINE {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		repo.UpdateDeviceLastUsedTime(req.ID, currentTime)
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// BatchUpdateDeviceStatus 批量更新设备状态
func (d *DeviceListApi) BatchUpdateDeviceStatus(c *gin.Context) {
	var req dto.BatchUpdateDeviceStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if len(req.DeviceIDs) == 0 {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设备ID列表不能为空"))
		return
	}

	if err := repo.BatchUpdateDeviceStatus(req.DeviceIDs, req.Status); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "批量更新状态失败"))
		return
	}

	// 如果设备上线，更新最后使用时间
	if req.Status == model.DEVICE_ONLINE {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		for _, deviceID := range req.DeviceIDs {
			repo.UpdateDeviceLastUsedTime(deviceID, currentTime)
		}
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// DelDeviceList 删除设备信息
func (d *DeviceListApi) DelDeviceList(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	if err := repo.DeleteDeviceList(id); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "删除失败"))
		return
	}

	c.JSON(http.StatusOK, handle.Success(nil))
}

// GetDeviceList 根据ID获取设备信息
func (d *DeviceListApi) GetDeviceList(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	deviceList, err := repo.GetDeviceListByID(id)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	response := dto.DeviceListResponse{
		ID:         deviceList.ID,
		IPAddress:  deviceList.IPAddress,
		LastUsedTm: deviceList.LastUsedTm,
		Status:     deviceList.Status,
		FaceDevice: deviceList.FaceDevice,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetDeviceListAll 获取所有设备列表
func (d *DeviceListApi) GetDeviceListAll(c *gin.Context) {
	deviceLists, err := repo.GetDeviceListAll()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetDeviceListPage 分页获取设备列表
func (d *DeviceListApi) GetDeviceListPage(c *gin.Context) {
	var req dto.DeviceListPageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	deviceLists, total, err := repo.GetDeviceListPage(req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	response := dto.DeviceListPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetDeviceListByStatus 根据状态获取设备列表
func (d *DeviceListApi) GetDeviceListByStatus(c *gin.Context) {
	statusStr := c.Param("status")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	deviceLists, err := repo.GetDeviceListByStatus(status)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetOnlineDevices 获取在线设备列表
func (d *DeviceListApi) GetOnlineDevices(c *gin.Context) {
	deviceLists, err := repo.GetOnlineDevices()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetOfflineDevices 获取离线设备列表
func (d *DeviceListApi) GetOfflineDevices(c *gin.Context) {
	deviceLists, err := repo.GetOfflineDevices()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// GetFaceDevices 获取人脸设备列表
func (d *DeviceListApi) GetFaceDevices(c *gin.Context) {
	deviceLists, err := repo.GetFaceDevices()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}

// QueryDeviceList 根据查询条件获取设备列表
func (d *DeviceListApi) QueryDeviceList(c *gin.Context) {
	var req dto.DeviceListQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误"))
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	deviceLists, total, err := repo.GetDeviceListByQuery(
		req.Status,
		req.FaceDevice,
		req.IPAddress,
		req.Page,
		req.PageSize,
	)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	response := dto.DeviceListPageResponse{
		List:  responses,
		Total: total,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetDeviceStatusSummary 获取设备状态统计
func (d *DeviceListApi) GetDeviceStatusSummary(c *gin.Context) {
	summary, err := repo.GetDeviceStatusSummary()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取统计失败"))
		return
	}

	response := dto.DeviceStatusSummaryResponse{
		TotalDevices:   int(summary["total"]),
		OnlineDevices:  int(summary["online"]),
		OfflineDevices: int(summary["offline"]),
		FaceDevices:    int(summary["face"]),
		NonFaceDevices: int(summary["non_face"]),
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// SearchDevicesByIP 根据IP地址搜索设备
func (d *DeviceListApi) SearchDevicesByIP(c *gin.Context) {
	ipPattern := c.Query("ip")
	if ipPattern == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "IP地址参数不能为空"))
		return
	}

	deviceLists, err := repo.SearchDevicesByIP(ipPattern)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "搜索失败"))
		return
	}

	var responses []dto.DeviceListResponse
	for _, deviceList := range deviceLists {
		responses = append(responses, dto.DeviceListResponse{
			ID:         deviceList.ID,
			IPAddress:  deviceList.IPAddress,
			LastUsedTm: deviceList.LastUsedTm,
			Status:     deviceList.Status,
			FaceDevice: deviceList.FaceDevice,
		})
	}

	c.JSON(http.StatusOK, handle.Success(responses))
}
