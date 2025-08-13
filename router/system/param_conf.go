package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type ParamConfRouter struct{}

func (p *ParamConfRouter) InitParamConfRouter(Router *gin.RouterGroup) {
	paramConfRouter := Router.Group("paramConf")
	paramConfApi := v1.ApiGroupApp.SystemApiGroup.ParamConfApi
	{
		// 基础CRUD操作
		paramConfRouter.POST("saveParamConf", paramConfApi.SaveParamConf)
		paramConfRouter.POST("updateParamConf", paramConfApi.UpdateParamConf)
		paramConfRouter.POST("updateParamValue", paramConfApi.UpdateParamValue)
		paramConfRouter.POST("batchUpdateParamConf", paramConfApi.BatchUpdateParamConf)
		paramConfRouter.DELETE("delParamConf/:id", paramConfApi.DelParamConf)
		paramConfRouter.DELETE("delParamConfByKey/:key", paramConfApi.DelParamConfByKey)
		paramConfRouter.GET("getParamConf/:id", paramConfApi.GetParamConf)
		paramConfRouter.GET("getParamConfAll", paramConfApi.GetParamConfAll)
		paramConfRouter.GET("getParamConfPage", paramConfApi.GetParamConfPage)
		
		// 按条件查询
		paramConfRouter.GET("getParamConfByKey/:key", paramConfApi.GetParamConfByKey)
		paramConfRouter.GET("queryParamConf", paramConfApi.QueryParamConf)
		paramConfRouter.GET("searchParamConf", paramConfApi.SearchParamConf)
		
		// 参数值操作
		paramConfRouter.GET("getParamValue", paramConfApi.GetParamValue)
		paramConfRouter.POST("setParamValue", paramConfApi.SetParamValue)
		
		// 工具接口
		paramConfRouter.GET("getAllParamKeys", paramConfApi.GetAllParamKeys)
	}
}
