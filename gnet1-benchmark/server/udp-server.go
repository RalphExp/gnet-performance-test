package server

import (
	"fmt"

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

	buffer := util.GetBuffer()
	copy(buffer, frame)

	err := server.pool.Submit(func() {
		// out, _ = util.GenerateRandomBytes(buffer, len(buffer))
		// time.Sleep(time.Millisecond)
		defer util.PutBuffer(buffer)

		if err := c.SendTo(buffer); err != nil {
			fmt.Printf("%v\n", err)
		}
	})

	if err != nil {
		fmt.Printf("submit error: %v\n", err)
	}
	return nil, gnet.None
}
