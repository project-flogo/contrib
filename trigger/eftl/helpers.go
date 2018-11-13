package eftl

import(
	"net"

	opentracing "github.com/opentracing/opentracing-go"
)

// Span is a tracing span
type Span struct {
	opentracing.Span
}


// GetLocalIP gets the public ip address of the system
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0"
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "0.0.0.0"
}