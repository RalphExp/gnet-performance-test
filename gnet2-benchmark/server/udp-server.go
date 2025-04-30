package server

import (
	"github.com/panjf2000/gnet/v2"
)

type UDPServer struct {
	gnet.BuiltinEventEngine
}

func NewUDPServer() *UDPServer {
	return &UDPServer{}
}

func (svr *UDPServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	buf, _ := c.Next(-1)
	_, _ = c.Write(buf)
	return gnet.None
}
