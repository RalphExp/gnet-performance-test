package server

import (
	"fmt"
	"net"
	"sync"
	"test-util/util"

	"github.com/panjf2000/ants/v2"
)

type UDPServer struct {
	Addr    net.Addr
	Conn    *net.UDPConn
	Threads int
	pool    *ants.Pool
}

func NewUDPServer(addr string, threads int, poolSize int) (*UDPServer, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen udp: %w", err)
	}

	return &UDPServer{
		Addr:    serverAddr,
		Conn:    conn,
		Threads: threads,
		pool:    util.CreateAntsPool(poolSize),
	}, nil
}

func (s *UDPServer) Serve() error {
	fmt.Printf("UDP Server start: %s\n", s.Addr.String())
	var wg sync.WaitGroup

	for i := 0; i < s.Threads; i++ {
		wg.Add(1)
		buffer := make([]byte, 1024)
		go func() {
			defer wg.Done()
			for {
				n, addr, err := s.Conn.ReadFromUDP(buffer)
				if err != nil {
					fmt.Printf("recvfrom error: %v\n", err)
					continue
				}

				data := make([]byte, n)
				copy(data, buffer[:n])
				s.pool.Submit(func() {
					// out, _ := util.GenerateRandomBytes(data, n)
					_, err = s.Conn.WriteToUDP(data, addr)
					if err != nil {
						fmt.Printf("sendto error: %v\n", err)
					}
				})
			}
		}()
	}
	wg.Wait()
	return nil
}

func (s *UDPServer) Close() error {
	if s.Conn != nil {
		return s.Conn.Close()
	}
	return nil
}
