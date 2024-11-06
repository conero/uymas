package cloud

import (
	"fmt"
	"gitee.com/conero/uymas/v2/rock"
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
}

// PortAsWeb set port as web address, portArgs support parameter [2]string, []string, [2]any, []any
//
// PortAsWeb(8080) -> http://localhost:8080
//
// PortAsWeb([2]string{8080, 10.10.16.241}) -> http://10.10.16.241:8080
//
// PortAsWeb([]string{8080, 10.10.16.241}) -> http://10.10.16.241:8080
//
// PortAsWeb([2]any{8080, 10.10.16.241}) -> http://10.10.16.241:8080
//
// PortAsWeb([]any{8080, 10.10.16.241}) -> http://10.10.16.241:8080
func PortAsWeb(portArgs any, isHttpsArgs ...bool) string {
	host := "localhost"
	var vPort = portArgs
	switch vPortArgs := portArgs.(type) {
	case [2]string:
		vPort = vPortArgs[0]
		host = vPortArgs[1]
	case []string:
		vPort = rock.ParamIndex(1, "", vPortArgs...)
		host = rock.ParamIndex(2, "localhost", vPortArgs...)
	case [2]any:
		vPort = vPortArgs[0]
		host = fmt.Sprintf("%v", vPortArgs[1])
	case []any:
		vPort = rock.ParamIndex(1, "", vPortArgs...)
		host = fmt.Sprintf("%v", rock.ParamIndex(2, "localhost", vPortArgs...))
	}
	portStr := PortAddress(vPort)
	if portStr == "" {
		return ""
	}
	isHttps := rock.Param(false, isHttpsArgs...)
	idx := strings.Index(portStr, ":")
	if idx == 0 && isHttps {
		if portStr == ":443" {
			portStr = ""
		}
		return "https://" + host + portStr
	}
	if idx == 0 {
		if portStr == ":80" {
			portStr = ""
		}
		return "http://" + host + portStr
	}
	return portStr
}
