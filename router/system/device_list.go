package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type DeviceListRouter struct{}

func (d *DeviceListRouter) InitDeviceListRouter(Router *gin.RouterGroup) {
	deviceListRouter := Router.Group("deviceList")
	deviceListApi := v1.ApiGroupApp.SystemApiGroup.DeviceListApi
	{
		// 基础CRUD操作
		deviceListRouter.POST("saveDeviceList", deviceListApi.SaveDeviceList)
		deviceListRouter.POST("updateDeviceList", deviceListApi.UpdateDeviceList)
		deviceListRouter.POST("updateDeviceStatus", deviceListApi.UpdateDeviceStatus)
		deviceListRouter.POST("batchUpdateDeviceStatus", deviceListApi.BatchUpdateDeviceStatus)
		deviceListRouter.DELETE("delDeviceList/:id", deviceListApi.DelDeviceList)
		deviceListRouter.GET("getDeviceList/:id", deviceListApi.GetDeviceList)
		deviceListRouter.GET("getDeviceListAll", deviceListApi.GetDeviceListAll)
		deviceListRouter.GET("getDeviceListPage", deviceListApi.GetDeviceListPage)
		
		// 按条件查询
		deviceListRouter.GET("getDeviceListByStatus/:status", deviceListApi.GetDeviceListByStatus)
		deviceListRouter.GET("getOnlineDevices", deviceListApi.GetOnlineDevices)
		deviceListRouter.GET("getOfflineDevices", deviceListApi.GetOfflineDevices)
		deviceListRouter.GET("getFaceDevices", deviceListApi.GetFaceDevices)
		deviceListRouter.GET("queryDeviceList", deviceListApi.QueryDeviceList)
		deviceListRouter.GET("searchDevicesByIP", deviceListApi.SearchDevicesByIP)
		
		// 统计信息
		deviceListRouter.GET("getDeviceStatusSummary", deviceListApi.GetDeviceStatusSummary)
	}
}
