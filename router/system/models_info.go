package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type ModelsInfoRouter struct{}

func (m *ModelsInfoRouter) InitModelsInfoRouter(Router *gin.RouterGroup) {
	modelsInfoRouter := Router.Group("modelsInfo")
	modelsInfoApi := v1.ApiGroupApp.SystemApiGroup.ModelsInfoApi
	{
		// 基础CRUD操作
		modelsInfoRouter.POST("saveModelsInfo", modelsInfoApi.SaveModelsInfo)
		modelsInfoRouter.POST("updateModelsInfo", modelsInfoApi.UpdateModelsInfo)
		modelsInfoRouter.POST("batchUpdateModelsInfo", modelsInfoApi.BatchUpdateModelsInfo)
		modelsInfoRouter.DELETE("delModelsInfo/:id", modelsInfoApi.DelModelsInfo)
		modelsInfoRouter.GET("getModelsInfo/:id", modelsInfoApi.GetModelsInfo)
		modelsInfoRouter.GET("getModelsInfoAll", modelsInfoApi.GetModelsInfoAll)
		modelsInfoRouter.GET("getModelsInfoPage", modelsInfoApi.GetModelsInfoPage)
		
		// 按条件查询
		modelsInfoRouter.GET("getModelsInfoByAlgorithm/:algorithm", modelsInfoApi.GetModelsInfoByAlgorithm)
		modelsInfoRouter.GET("getModelsInfoByClassify/:classify", modelsInfoApi.GetModelsInfoByClassify)
		modelsInfoRouter.GET("getModelsInfoBySophonType/:sophonType", modelsInfoApi.GetModelsInfoBySophonType)
		modelsInfoRouter.GET("queryModelsInfo", modelsInfoApi.QueryModelsInfo)
		modelsInfoRouter.GET("searchModelsInfo", modelsInfoApi.SearchModelsInfo)
		
		// 统计和工具接口
		modelsInfoRouter.GET("getModelsInfoStatistics", modelsInfoApi.GetModelsInfoStatistics)
		modelsInfoRouter.GET("getAllAlgorithms", modelsInfoApi.GetAllAlgorithms)
		modelsInfoRouter.GET("getAllClassifies", modelsInfoApi.GetAllClassifies)
	}
}
