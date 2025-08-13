package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type TaskManageRouter struct{}

func (t *TaskManageRouter) InitTaskManageRouter(Router *gin.RouterGroup) {
	taskManageRouter := Router.Group("taskManage")
	taskManageApi := v1.ApiGroupApp.SystemApiGroup.TaskManageApi
	{
		// 基础CRUD操作
		taskManageRouter.POST("saveTaskManage", taskManageApi.SaveTaskManage)
		taskManageRouter.POST("updateTaskManage", taskManageApi.UpdateTaskManage)
		taskManageRouter.POST("updateTaskStatus", taskManageApi.UpdateTaskStatus)
		taskManageRouter.DELETE("delTaskManage/:id", taskManageApi.DelTaskManage)
		taskManageRouter.GET("getTaskManage/:id", taskManageApi.GetTaskManage)
		taskManageRouter.GET("getTaskManageList", taskManageApi.GetTaskManageList)
		taskManageRouter.GET("getTaskManagePage", taskManageApi.GetTaskManagePage)
		
		// 按条件查询
		taskManageRouter.GET("getTaskManageByStatus/:status", taskManageApi.GetTaskManageByStatus)
		taskManageRouter.GET("getTaskManageByMediaID/:mediaId", taskManageApi.GetTaskManageByMediaID)
	}
}
