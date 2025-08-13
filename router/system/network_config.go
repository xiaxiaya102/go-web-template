package system

import (
	v1 "feishuReboot/api/v1"
	"github.com/gin-gonic/gin"
)

type NetworkConfigRouter struct{}

func (n *NetworkConfigRouter) InitNetworkConfigRouter(Router *gin.RouterGroup) {
	networkConfigRouter := Router.Group("networkConfig")
	networkConfigApi := v1.ApiGroupApp.SystemApiGroup.NetworkConfigApi
	{
		// 获取网络配置信息
		networkConfigRouter.GET("getNetworkConfig/:device", networkConfigApi.GetNetworkConfig)
		networkConfigRouter.GET("getNetworkInterfaces", networkConfigApi.GetNetworkInterfaces)
		networkConfigRouter.GET("getNetworkStatus", networkConfigApi.GetNetworkStatus)
		
		// 设置网络配置
		networkConfigRouter.POST("setStaticIP", networkConfigApi.SetStaticIP)
		networkConfigRouter.POST("setDHCP", networkConfigApi.SetDHCP)
		
		// 网络工具
		networkConfigRouter.GET("testConnectivity", networkConfigApi.TestNetworkConnectivity)
		networkConfigRouter.POST("restartInterface/:device", networkConfigApi.RestartNetworkInterface)
	}
}
