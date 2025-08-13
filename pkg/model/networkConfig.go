package model

// NetworkConfig 网络状态常量
const (
	NETWORK_STATUS_STATIC  = "static"  // 静态IP
	NETWORK_STATUS_DYNAMIC = "dynamic" // 动态IP(DHCP)
)

type NetworkConfig struct {
	Bandwidth       string `json:"bandwidth"`        // 带宽，如"100Mb/s"
	BandwidthVarity bool   `json:"bandwidth_varity"` // 带宽变化
	Device          string `json:"device"`           // 网络设备，如"eth0"
	DNS             string `json:"dns"`              // DNS服务器
	Gateway         string `json:"gateway"`          // 网关地址
	IP              string `json:"ip"`               // IP地址
	MAC             string `json:"mac"`              // MAC地址
	Netmask         string `json:"netmask"`          // 子网掩码
	Status          string `json:"status"`           // 网络状态：static/dynamic
}
