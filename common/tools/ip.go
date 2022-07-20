package tools

import (
	"net"
	"net/http"
)

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}
	return remoteAddr
}
