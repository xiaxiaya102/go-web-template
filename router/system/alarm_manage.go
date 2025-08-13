package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type AlarmManageRouter struct{}

func (a *AlarmManageRouter) InitAlarmManageRouter(Router *gin.RouterGroup) {
	alarmManageRouter := Router.Group("alarmManage")
	alarmManageApi := v1.ApiGroupApp.SystemApiGroup.AlarmManageApi
	{
		// 基础CRUD操作
		alarmManageRouter.POST("saveAlarmManage", alarmManageApi.SaveAlarmManage)
		alarmManageRouter.POST("updateAlarmManage", alarmManageApi.UpdateAlarmManage)
		alarmManageRouter.POST("updateReportStatus", alarmManageApi.UpdateReportStatus)
		alarmManageRouter.DELETE("delAlarmManage/:id", alarmManageApi.DelAlarmManage)
		alarmManageRouter.GET("getAlarmManage/:id", alarmManageApi.GetAlarmManage)
		alarmManageRouter.GET("getAlarmManageList", alarmManageApi.GetAlarmManageList)
		alarmManageRouter.GET("getAlarmManagePage", alarmManageApi.GetAlarmManagePage)
		
		// 按条件查询
		alarmManageRouter.GET("getAlarmManageByReportStatus/:status", alarmManageApi.GetAlarmManageByReportStatus)
		alarmManageRouter.GET("getAlarmManageByAlarmType/:type", alarmManageApi.GetAlarmManageByAlarmType)
		alarmManageRouter.GET("getAlarmManageByMediaID/:mediaId", alarmManageApi.GetAlarmManageByMediaID)
		alarmManageRouter.GET("queryAlarmManage", alarmManageApi.QueryAlarmManage)
	}
}
