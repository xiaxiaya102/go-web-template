package service

import (
	"bufio"
	"feishuReboot/pkg/model"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type NetworkConfigService struct{}

// GetNetworkConfig 获取指定网络接口的配置信息
func (n *NetworkConfigService) GetNetworkConfig(device string) (*model.NetworkConfig, error) {
	config := &model.NetworkConfig{
		Device: device,
	}

	// 获取IP地址、MAC地址等基本信息
	if err := n.getInterfaceInfo(config); err != nil {
		return nil, err
	}

	// 获取网关信息
	if err := n.getGatewayInfo(config); err != nil {
		return nil, err
	}

	// 获取DNS信息
	if err := n.getDNSInfo(config); err != nil {
		return nil, err
	}

	// 获取带宽信息
	if err := n.getBandwidthInfo(config); err != nil {
		return nil, err
	}

	// 判断是否为静态IP或DHCP
	if err := n.getNetworkStatus(config); err != nil {
		return nil, err
	}

	return config, nil
}

// getInterfaceInfo 获取网络接口基本信息
func (n *NetworkConfigService) getInterfaceInfo(config *model.NetworkConfig) error {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, iface := range interfaces {
		if iface.Name == config.Device {
			config.MAC = iface.HardwareAddr.String()

			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}

			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						config.IP = ipnet.IP.String()
						config.Netmask = net.IP(ipnet.Mask).String()
						break
					}
				}
			}
			break
		}
	}

	return nil
}

// getGatewayInfo 获取网关信息
func (n *NetworkConfigService) getGatewayInfo(config *model.NetworkConfig) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("ip", "route", "show", "default")
	case "windows":
		cmd = exec.Command("route", "print", "0.0.0.0")
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	if runtime.GOOS == "linux" {
		// 解析Linux路由信息
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "default") && strings.Contains(line, config.Device) {
				fields := strings.Fields(line)
				if len(fields) >= 3 {
					config.Gateway = fields[2]
					break
				}
			}
		}
	} else if runtime.GOOS == "windows" {
		// 解析Windows路由信息
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "0.0.0.0") && strings.Contains(line, "0.0.0.0") {
				fields := strings.Fields(line)
				if len(fields) >= 3 {
					config.Gateway = fields[2]
					break
				}
			}
		}
	}

	return nil
}

// getDNSInfo 获取DNS信息
func (n *NetworkConfigService) getDNSInfo(config *model.NetworkConfig) error {
	var dnsServers []string

	switch runtime.GOOS {
	case "linux":
		file, err := os.Open("/etc/resolv.conf")
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "nameserver") {
				fields := strings.Fields(line)
				if len(fields) >= 2 {
					dnsServers = append(dnsServers, fields[1])
				}
			}
		}
	case "windows":
		cmd := exec.Command("nslookup", "localhost")
		output, err := cmd.Output()
		if err == nil {
			lines := strings.Split(string(output), "\n")
			for _, line := range lines {
				if strings.Contains(line, "Address:") && !strings.Contains(line, "#") {
					parts := strings.Split(line, ":")
					if len(parts) >= 2 {
						dns := strings.TrimSpace(parts[1])
						if dns != "" && dns != "127.0.0.1" {
							dnsServers = append(dnsServers, dns)
						}
					}
				}
			}
		}
	}

	config.DNS = strings.Join(dnsServers, ",")
	return nil
}

// getBandwidthInfo 获取带宽信息
func (n *NetworkConfigService) getBandwidthInfo(config *model.NetworkConfig) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("ethtool", config.Device)
	case "windows":
		// Windows下获取网络速度比较复杂，这里简化处理
		config.Bandwidth = "Unknown"
		config.BandwidthVarity = false
		return nil
	default:
		config.Bandwidth = "Unknown"
		config.BandwidthVarity = false
		return nil
	}

	output, err := cmd.Output()
	if err != nil {
		// 如果ethtool命令失败，设置默认值
		config.Bandwidth = "Unknown"
		config.BandwidthVarity = false
		return nil
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Speed:") {
			re := regexp.MustCompile(`Speed:\s*(\d+)Mb/s`)
			matches := re.FindStringSubmatch(line)
			if len(matches) >= 2 {
				config.Bandwidth = matches[1] + "Mb/s"
			}
			break
		}
	}

	if config.Bandwidth == "" {
		config.Bandwidth = "Unknown"
	}
	config.BandwidthVarity = false

	return nil
}

// getNetworkStatus 判断网络配置状态（静态或DHCP）
func (n *NetworkConfigService) getNetworkStatus(config *model.NetworkConfig) error {
	// 这里简化处理，实际应该检查网络配置文件
	// Linux: /etc/network/interfaces 或 /etc/netplan/
	// Windows: 通过WMI或注册表查询

	switch runtime.GOOS {
	case "linux":
		// 检查是否有DHCP客户端进程
		cmd := exec.Command("ps", "aux")
		output, err := cmd.Output()
		if err == nil {
			if strings.Contains(string(output), "dhclient") || strings.Contains(string(output), "dhcpcd") {
				config.Status = model.NETWORK_STATUS_DYNAMIC
			} else {
				config.Status = model.NETWORK_STATUS_STATIC
			}
		} else {
			config.Status = model.NETWORK_STATUS_STATIC
		}
	case "windows":
		// Windows下默认设置为动态
		config.Status = model.NETWORK_STATUS_DYNAMIC
	default:
		config.Status = model.NETWORK_STATUS_STATIC
	}

	return nil
}

// GetNetworkInterfaces 获取所有网络接口列表
func (n *NetworkConfigService) GetNetworkInterfaces() ([]model.NetworkConfig, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var configs []model.NetworkConfig
	for _, iface := range interfaces {
		// 跳过回环接口
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		config, err := n.GetNetworkConfig(iface.Name)
		if err != nil {
			continue
		}
		configs = append(configs, *config)
	}

	return configs, nil
}

// SetStaticIP 设置静态IP
func (n *NetworkConfigService) SetStaticIP(device, ip, netmask, gateway, dns string) error {
	switch runtime.GOOS {
	case "linux":
		return n.setLinuxStaticIP(device, ip, netmask, gateway, dns)
	case "windows":
		return n.setWindowsStaticIP(device, ip, netmask, gateway, dns)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// setLinuxStaticIP 在Linux系统上设置静态IP
func (n *NetworkConfigService) setLinuxStaticIP(device, ip, netmask, gateway, dns string) error {
	// 设置IP地址
	cmd := exec.Command("ip", "addr", "flush", "dev", device)
	if err := cmd.Run(); err != nil {
		return err
	}

	// 计算CIDR
	cidr, err := n.netmaskToCIDR(netmask)
	if err != nil {
		return err
	}

	cmd = exec.Command("ip", "addr", "add", ip+"/"+strconv.Itoa(cidr), "dev", device)
	if err := cmd.Run(); err != nil {
		return err
	}

	// 启用接口
	cmd = exec.Command("ip", "link", "set", device, "up")
	if err := cmd.Run(); err != nil {
		return err
	}

	// 设置默认网关
	if gateway != "" {
		cmd = exec.Command("ip", "route", "add", "default", "via", gateway, "dev", device)
		cmd.Run() // 忽略错误，可能已存在
	}

	return nil
}

// setWindowsStaticIP 在Windows系统上设置静态IP
func (n *NetworkConfigService) setWindowsStaticIP(device, ip, netmask, gateway, dns string) error {
	// Windows下设置静态IP需要管理员权限
	cmd := exec.Command("netsh", "interface", "ip", "set", "address", device, "static", ip, netmask, gateway)
	if err := cmd.Run(); err != nil {
		return err
	}

	if dns != "" {
		cmd = exec.Command("netsh", "interface", "ip", "set", "dns", device, "static", dns)
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// SetDHCP 设置DHCP
func (n *NetworkConfigService) SetDHCP(device string) error {
	switch runtime.GOOS {
	case "linux":
		return n.setLinuxDHCP(device)
	case "windows":
		return n.setWindowsDHCP(device)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// setLinuxDHCP 在Linux系统上设置DHCP
func (n *NetworkConfigService) setLinuxDHCP(device string) error {
	// 清除静态IP配置
	cmd := exec.Command("ip", "addr", "flush", "dev", device)
	if err := cmd.Run(); err != nil {
		return err
	}

	// 启动DHCP客户端
	cmd = exec.Command("dhclient", device)
	return cmd.Run()
}

// setWindowsDHCP 在Windows系统上设置DHCP
func (n *NetworkConfigService) setWindowsDHCP(device string) error {
	cmd := exec.Command("netsh", "interface", "ip", "set", "address", device, "dhcp")
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("netsh", "interface", "ip", "set", "dns", device, "dhcp")
	return cmd.Run()
}

// netmaskToCIDR 将子网掩码转换为CIDR格式
func (n *NetworkConfigService) netmaskToCIDR(netmask string) (int, error) {
	ip := net.ParseIP(netmask)
	if ip == nil {
		return 0, fmt.Errorf("invalid netmask: %s", netmask)
	}

	mask := ip.To4()
	if mask == nil {
		return 0, fmt.Errorf("invalid IPv4 netmask: %s", netmask)
	}

	cidr := 0
	for _, b := range mask {
		for b != 0 {
			cidr++
			b <<= 1
		}
	}

	return cidr, nil
}
