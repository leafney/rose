package rose

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

// LocalIP gets the first NIC's IP address.
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if nil != err {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("can't get local IP")
}

// LocalMacAddr gets the first NIC's MAC address.
func LocalMacAddr() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, inter := range interfaces {
		address, err := inter.Addrs()
		if err != nil {
			return "", err
		}

		for _, address := range address {
			// check the address type and if it is not a loopback the display it
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return inter.HardwareAddr.String(), nil
				}
			}
		}
	}

	return "", errors.New("can't get local mac")
}

// IsLocalIPAddr 检测 IP 地址字符串是否是内网地址
func IsLocalIPAddr(ip string) bool {
	return IsLocalIP(net.ParseIP(ip))
}

// IsLocalIP 检测 IP 地址是否是内网地址
// 通过直接对比ip段范围效率更高，详见：https://github.com/thinkeridea/go-extend/issues/2
func IsLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	ip := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	var err error
	if ip, _, err = net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// ClientPublicIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !IsLocalIPAddr(ip) {
			return ip
		}
	}

	if ip = strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" && !IsLocalIPAddr(ip) {
		return ip
	}

	if ip = RemoteIP(r); !IsLocalIPAddr(ip) {
		return ip
	}

	return ""
}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	return ip
}

// GetRemoteClientIP 获取客户端IP地址
// 优先获取公网IP（解析 X-Real-IP 和 X-Forwarded-For），无公网IP则获取内网IP
func GetRemoteClientIP(r *http.Request) string {
	// 优先尝试获取公网IP
	ip := ClientPublicIP(r)
	if StrIsEmpty(ip) {
		// 其次尝试获取内网IP
		ip = ClientIP(r)
	}

	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}
