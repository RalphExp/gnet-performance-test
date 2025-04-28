package server

import (
	"github.com/panjf2000/gnet"
)

type UDPServer struct {
	gnet.EventServer
}

func NewUDPServer() *UDPServer {
	return &UDPServer{}
}

// React fires when a connection sends the server data.
// Call c.Read() or c.ReadN(n) within the parameter:c to read incoming data from client.
// Parameter:out is the return value which is going to be sent back to the client.
func (svr *UDPServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	out = frame[:]
	return out, gnet.None
}
