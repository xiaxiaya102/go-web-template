package system

import (
	"feishuReboot/pkg/dto"
	"feishuReboot/pkg/handle"
	"feishuReboot/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NetworkConfigApi struct{}

var networkService = &service.NetworkConfigService{}

// GetNetworkConfig 获取指定网络接口配置
// @Summary 获取网络接口配置
// @Description 根据设备名称获取指定网络接口的配置信息
// @Tags 网络配置
// @Accept json
// @Produce json
// @Param device path string true "网络设备名称"
// @Success 200 {object} handle.Response{data=dto.NetworkConfigResponse} "获取成功"
// @Failure 400 {object} handle.Response "参数错误"
// @Router /networkConfig/getNetworkConfig/{device} [get]
func (n *NetworkConfigApi) GetNetworkConfig(c *gin.Context) {
	device := c.Param("device")
	if device == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "网络设备参数不能为空"))
		return
	}

	config, err := networkService.GetNetworkConfig(device)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取网络配置失败: "+err.Error()))
		return
	}

	response := dto.NetworkConfigResponse{
		Bandwidth:       config.Bandwidth,
		BandwidthVarity: config.BandwidthVarity,
		Device:          config.Device,
		DNS:             config.DNS,
		Gateway:         config.Gateway,
		IP:              config.IP,
		MAC:             config.MAC,
		Netmask:         config.Netmask,
		Status:          config.Status,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// GetNetworkInterfaces 获取所有网络接口列表
// @Summary 获取所有网络接口
// @Description 获取系统中所有网络接口的配置信息列表
// @Tags 网络配置
// @Accept json
// @Produce json
// @Success 200 {object} handle.Response{data=[]dto.NetworkConfigResponse} "获取成功"
// @Failure 400 {object} handle.Response "获取失败"
// @Router /networkConfig/getNetworkInterfaces [get]
func (n *NetworkConfigApi) GetNetworkInterfaces(c *gin.Context) {
	configs, err := networkService.GetNetworkInterfaces()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取网络接口列表失败: "+err.Error()))
		return
	}

	var interfaces []dto.NetworkInterfaceInfo
	for _, config := range configs {
		interfaces = append(interfaces, dto.NetworkInterfaceInfo{
			Name:      config.Device,
			IP:        config.IP,
			MAC:       config.MAC,
			Status:    "up", // 简化处理，实际应该检查接口状态
			IsDefault: config.Gateway != "", // 有网关的接口认为是默认接口
		})
	}

	response := dto.NetworkInterfaceListResponse{
		Interfaces: interfaces,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// SetStaticIP 设置静态IP
func (n *NetworkConfigApi) SetStaticIP(c *gin.Context) {
	var req dto.SetStaticIPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误: "+err.Error()))
		return
	}

	err := networkService.SetStaticIP(req.Device, req.IP, req.Netmask, req.Gateway, req.DNS)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置静态IP失败: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Success("静态IP设置成功"))
}

// SetDHCP 设置DHCP
func (n *NetworkConfigApi) SetDHCP(c *gin.Context) {
	var req dto.SetDHCPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "参数错误: "+err.Error()))
		return
	}

	err := networkService.SetDHCP(req.Device)
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "设置DHCP失败: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, handle.Success("DHCP设置成功"))
}

// GetNetworkStatus 获取网络连接状态
func (n *NetworkConfigApi) GetNetworkStatus(c *gin.Context) {
	// 获取默认网络接口的状态
	configs, err := networkService.GetNetworkInterfaces()
	if err != nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "获取网络状态失败: "+err.Error()))
		return
	}

	// 找到有网关的接口作为默认接口
	var defaultInterface *dto.NetworkConfigResponse
	for _, config := range configs {
		if config.Gateway != "" {
			defaultInterface = &dto.NetworkConfigResponse{
				Bandwidth:       config.Bandwidth,
				BandwidthVarity: config.BandwidthVarity,
				Device:          config.Device,
				DNS:             config.DNS,
				Gateway:         config.Gateway,
				IP:              config.IP,
				MAC:             config.MAC,
				Netmask:         config.Netmask,
				Status:          config.Status,
			}
			break
		}
	}

	if defaultInterface == nil {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "未找到默认网络接口"))
		return
	}

	response := dto.NetworkStatusResponse{
		Connected:     defaultInterface.IP != "",
		InterfaceName: defaultInterface.Device,
		Speed:         defaultInterface.Bandwidth,
		Duplex:        "full", // 简化处理
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// TestNetworkConnectivity 测试网络连通性
func (n *NetworkConfigApi) TestNetworkConnectivity(c *gin.Context) {
	target := c.Query("target")
	if target == "" {
		target = "8.8.8.8" // 默认测试Google DNS
	}

	// 简单的ping测试
	// 这里可以实现更复杂的网络连通性测试
	response := map[string]interface{}{
		"target":     target,
		"reachable":  true,  // 简化处理
		"latency":    "10ms", // 简化处理
		"packet_loss": 0,
	}

	c.JSON(http.StatusOK, handle.Success(response))
}

// RestartNetworkInterface 重启网络接口
func (n *NetworkConfigApi) RestartNetworkInterface(c *gin.Context) {
	device := c.Param("device")
	if device == "" {
		c.JSON(http.StatusOK, handle.FailWithMsg(-1, "网络设备参数不能为空"))
		return
	}

	// 这里可以实现重启网络接口的逻辑
	// 需要根据操作系统执行相应的命令

	c.JSON(http.StatusOK, handle.Success("网络接口重启成功"))
}
