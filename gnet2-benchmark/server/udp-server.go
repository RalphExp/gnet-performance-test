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

// React fires when a connection sends the server data.
// Call c.Read() or c.ReadN(n) within the parameter:c to read incoming data from client.
// Parameter:out is the return value which is going to be sent back to the client.
func (svr *UDPServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	buf, _ := c.Next(-1)
	_, _ = c.Write(buf)
	return gnet.None
}
