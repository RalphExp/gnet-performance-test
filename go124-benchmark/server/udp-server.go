package server

import (
	"fmt"
	"net"
	"test-util/util"

	"github.com/panjf2000/ants/v2"
)

type UDPServer struct {
	Addr    net.Addr
	Conn    *net.UDPConn
	Threads int
	pool    *ants.Pool
}

func NewUDPServer(addr string, poolSize int) (*UDPServer, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen udp: %w", err)
	}

	return &UDPServer{
		Addr: serverAddr,
		Conn: conn,
		pool: util.CreateAntsPool(poolSize),
	}, nil
}

func (s *UDPServer) Serve() error {
	fmt.Printf("UDP Server start: %s\n", s.Addr.String())
	buffer := make([]byte, 1024)

	for {
		n, addr, err := s.Conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("recvfrom error: %v\n", err)
			continue
		}

		data := util.GetBuffer()
		copy(data, buffer[:n])

		s.pool.Submit(func() {
			defer util.PutBuffer(data)
			_, err = s.Conn.WriteToUDP(data, addr)
			if err != nil {
				fmt.Printf("sendto error: %v\n", err)
			}
		})
	}
	return nil
}

func (s *UDPServer) Close() error {
	if s.Conn != nil {
		return s.Conn.Close()
	}
	return nil
}
