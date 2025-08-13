package v1

import (
	"feishuReboot/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

func init() {
	//ApiGroupApp.SystemApiGroup.FaceAlarmApi.Init()
}
