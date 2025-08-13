package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type AlarmCategoryRouter struct{}

func (a *AlarmCategoryRouter) InitAlarmCategoryRouter(Router *gin.RouterGroup) {
	alarmCategoryRouter := Router.Group("alarmCategory")
	alarmCategoryApi := v1.ApiGroupApp.SystemApiGroup.AlarmCategoryApi
	{
		// 基础CRUD操作
		alarmCategoryRouter.POST("saveAlarmCategory", alarmCategoryApi.SaveAlarmCategory)
		alarmCategoryRouter.POST("updateAlarmCategory", alarmCategoryApi.UpdateAlarmCategory)
		alarmCategoryRouter.POST("batchUpdateAlarmCategory", alarmCategoryApi.BatchUpdateAlarmCategory)
		alarmCategoryRouter.DELETE("delAlarmCategory/:id", alarmCategoryApi.DelAlarmCategory)
		alarmCategoryRouter.GET("getAlarmCategory/:id", alarmCategoryApi.GetAlarmCategory)
		alarmCategoryRouter.GET("getAlarmCategoryAll", alarmCategoryApi.GetAlarmCategoryAll)
		alarmCategoryRouter.GET("getAlarmCategoryPage", alarmCategoryApi.GetAlarmCategoryPage)
		
		// 按条件查询
		alarmCategoryRouter.GET("getAlarmCategoryByArithmetic/:arithmetic", alarmCategoryApi.GetAlarmCategoryByArithmetic)
		alarmCategoryRouter.GET("getAlarmCategoryByAlarmType/:alarmType", alarmCategoryApi.GetAlarmCategoryByAlarmType)
		alarmCategoryRouter.GET("getAlarmCategoryBySophonType/:sophonType", alarmCategoryApi.GetAlarmCategoryBySophonType)
		alarmCategoryRouter.GET("queryAlarmCategory", alarmCategoryApi.QueryAlarmCategory)
		alarmCategoryRouter.GET("searchAlarmCategory", alarmCategoryApi.SearchAlarmCategory)
		
		// 语音文件相关
		alarmCategoryRouter.GET("getAlarmCategoriesWithAudio", alarmCategoryApi.GetAlarmCategoriesWithAudio)
		alarmCategoryRouter.GET("getAlarmCategoriesWithoutAudio", alarmCategoryApi.GetAlarmCategoriesWithoutAudio)
		
		// 统计和工具接口
		alarmCategoryRouter.GET("getAlarmCategoryStatistics", alarmCategoryApi.GetAlarmCategoryStatistics)
		alarmCategoryRouter.GET("getAlarmCategoryLists", alarmCategoryApi.GetAlarmCategoryLists)
	}
}
