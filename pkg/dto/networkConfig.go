package dto

type NetworkConfigResponse struct {
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

type UpdateNetworkConfigRequest struct {
	Device  string `json:"device" binding:"required"` // 网络设备，如"eth0"
	DNS     string `json:"dns"`                       // DNS服务器
	Gateway string `json:"gateway"`                   // 网关地址
	IP      string `json:"ip"`                        // IP地址
	Netmask string `json:"netmask"`                   // 子网掩码
	Status  string `json:"status" binding:"required"` // 网络状态：static/dynamic
}

type NetworkInterfaceListResponse struct {
	Interfaces []NetworkInterfaceInfo `json:"interfaces"`
}

type NetworkInterfaceInfo struct {
	Name      string `json:"name"`       // 接口名称，如"eth0"
	IP        string `json:"ip"`         // IP地址
	MAC       string `json:"mac"`        // MAC地址
	Status    string `json:"status"`     // 接口状态：up/down
	IsDefault bool   `json:"is_default"` // 是否为默认接口
}

type NetworkStatusResponse struct {
	Connected     bool   `json:"connected"`      // 网络连接状态
	InterfaceName string `json:"interface_name"` // 当前使用的网络接口
	Speed         string `json:"speed"`          // 网络速度
	Duplex        string `json:"duplex"`         // 双工模式：full/half
}

type SetStaticIPRequest struct {
	Device  string `json:"device" binding:"required"`  // 网络设备
	IP      string `json:"ip" binding:"required"`      // IP地址
	Netmask string `json:"netmask" binding:"required"` // 子网掩码
	Gateway string `json:"gateway" binding:"required"` // 网关地址
	DNS     string `json:"dns"`                        // DNS服务器
}

type SetDHCPRequest struct {
	Device string `json:"device" binding:"required"` // 网络设备
}
