package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type PlanTemplateRouter struct{}

func (p *PlanTemplateRouter) InitPlanTemplateRouter(Router *gin.RouterGroup) {
	planTemplateRouter := Router.Group("planTemplate")
	planTemplateApi := v1.ApiGroupApp.SystemApiGroup.PlanTemplateApi
	{
		planTemplateRouter.POST("savePlanTemplate", planTemplateApi.SavePlanTemplate)
		planTemplateRouter.POST("updatePlanTemplate", planTemplateApi.UpdatePlanTemplate)
		planTemplateRouter.DELETE("delPlanTemplate/:id", planTemplateApi.DelPlanTemplate)
		planTemplateRouter.GET("getPlanTemplate/:id", planTemplateApi.GetPlanTemplate)
		planTemplateRouter.GET("getPlanTemplateList", planTemplateApi.GetPlanTemplateList)
		planTemplateRouter.GET("getPlanTemplatePage", planTemplateApi.GetPlanTemplatePage)
	}
}
