package cloud

import (
	"fmt"
	"net"
	"strings"
)

// PortAddress turn string port as addr for [net.Listen]
//
// 8080 -> :8080
func PortAddress(port any) string {
	var portStr string
	switch value := port.(type) {
	case string:
		portStr = value
	case int, int64, uint, uint64, uint16, int16:
		portStr = fmt.Sprintf("%d", value)
	default:
		return ""

	}
	if portStr == "" {
		return ""
	}

	if strings.Contains(portStr, ":") {
		return portStr
	}

	return ":" + portStr
}

// PortAvailable check if port is available that next port can be used
func PortAvailable(port uint16) uint16 {
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return port
		}
		go conn.Close()
		port += 1
	}
	//return port
}
