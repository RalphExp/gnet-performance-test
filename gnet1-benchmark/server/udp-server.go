package server

import (
	"github.com/panjf2000/ants/v2"
	"github.com/panjf2000/gnet"

	"test-util/util"
)

type UDPServer struct {
	gnet.EventServer
	pool *ants.Pool
}

func NewUDPServer(poolSize int) *UDPServer {
	return &UDPServer{
		pool: util.CreateAntsPool(poolSize),
	}
}

func (server *UDPServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {

	buffer := make([]byte, len(frame))
	copy(buffer, frame)

	server.pool.Submit(func() {
		// out, _ = util.GenerateRandomBytes(buffer, len(buffer))
		// time.Sleep(time.Millisecond)
		c.SendTo(buffer)
	})
	return nil, gnet.None
}
