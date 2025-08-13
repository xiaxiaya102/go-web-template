package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type ThresholdConfRouter struct{}

func (t *ThresholdConfRouter) InitThresholdConfRouter(Router *gin.RouterGroup) {
	thresholdConfRouter := Router.Group("thresholdConf")
	thresholdConfApi := v1.ApiGroupApp.SystemApiGroup.ThresholdConfApi
	{
		// 基础CRUD操作
		thresholdConfRouter.POST("saveThresholdConf", thresholdConfApi.SaveThresholdConf)
		thresholdConfRouter.POST("updateThresholdConf", thresholdConfApi.UpdateThresholdConf)
		thresholdConfRouter.POST("updateThresholdValue", thresholdConfApi.UpdateThresholdValue)
		thresholdConfRouter.POST("batchUpdateThresholdConf", thresholdConfApi.BatchUpdateThresholdConf)
		thresholdConfRouter.DELETE("delThresholdConf/:id", thresholdConfApi.DelThresholdConf)
		thresholdConfRouter.GET("getThresholdConf/:id", thresholdConfApi.GetThresholdConf)
		thresholdConfRouter.GET("getThresholdConfAll", thresholdConfApi.GetThresholdConfAll)
		thresholdConfRouter.GET("getThresholdConfPage", thresholdConfApi.GetThresholdConfPage)
		
		// 按条件查询
		thresholdConfRouter.GET("getThresholdConfBySophonType/:sophonType", thresholdConfApi.GetThresholdConfBySophonType)
		thresholdConfRouter.GET("queryThresholdConf", thresholdConfApi.QueryThresholdConf)
		thresholdConfRouter.GET("searchThresholdConf", thresholdConfApi.SearchThresholdConf)
		
		// 阈值操作
		thresholdConfRouter.GET("getThresholdValue", thresholdConfApi.GetThresholdValue)
		thresholdConfRouter.POST("setThresholdValue", thresholdConfApi.SetThresholdValue)
		
		// 统计和工具接口
		thresholdConfRouter.GET("getThresholdConfStatistics", thresholdConfApi.GetThresholdConfStatistics)
		thresholdConfRouter.GET("getThresholdConfLists", thresholdConfApi.GetThresholdConfLists)
	}
}
