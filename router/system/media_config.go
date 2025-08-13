package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type MediaConfigRouter struct{}

func (m *MediaConfigRouter) InitMediaConfigRouter(Router *gin.RouterGroup) {
	mediaConfigRouter := Router.Group("mainModule")
	mediaConfigApi := v1.ApiGroupApp.SystemApiGroup.MediaConfigApi
	{
		mediaConfigRouter.POST("saveMediaConfig", mediaConfigApi.SaveMediaConfig)
		mediaConfigRouter.POST("updateMediaConfig", mediaConfigApi.UpdateMediaConfig)
		mediaConfigRouter.DELETE("delMediaConfig/:id", mediaConfigApi.DelMediaConfig)
		mediaConfigRouter.GET("getMediaConfig/:id", mediaConfigApi.GetMediaConfig)
		mediaConfigRouter.GET("getMediaConfigPage", mediaConfigApi.GetMediaConfigPage)
	}
}
