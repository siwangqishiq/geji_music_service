package util

import (
	"fmt"
	"geji/config"
	"net"
)

func Is[T any](value bool, left T, right T) T {
	if value {
		return left
	} else {
		return right
	}
}

// AddAll 将 src 的所有元素追加到 dst，返回新切片
func AddAll[T any](dst, src []T) []T {
	return append(dst, src...)
}

func GetLocalIP() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, i := range interfaces {
		if i.Flags&net.FlagUp == 0 {
			continue
		}

		// 排除回环
		if i.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {

			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP

			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String()
		}
	}

	return ""
}

func GetUrlFromLocalPath(localPath string) string {
	var host string = config.PRD_IP
	if config.IS_DEBUG {
		host = GetLocalIP()
	}

	return fmt.Sprintf("http://%s:%d/media/%s", host, config.Port, localPath)
}
