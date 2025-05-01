package server

import (
	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet/v2"

	"test-util/util"
)

type UDPServer struct {
	gnet.BuiltinEventEngine
	pool *ants.Pool
}

func NewUDPServer(poolSize int) *UDPServer {
	return &UDPServer{
		pool: util.CreateAntsPool(poolSize),
	}
}

func (server *UDPServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	frame, _ := c.Next(-1)

	buffer := util.GetBuffer()
	copy(buffer, frame)

	server.pool.Submit(func() {
		defer util.PutBuffer(buffer)
		c.AsyncWrite(buffer, nil)
	})
	return gnet.None
}
