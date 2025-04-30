package server

import (
	"fmt"
	"net"
	"sync"
)

type UDPServer struct {
	Addr    net.Addr
	Conn    *net.UDPConn
	Threads int
}

func NewUDPServer(addr string, threads int) (*UDPServer, error) {
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
				} else {
					// fmt.Printf("recvfrom %s: %d bytes\n", addr.String(), n)
				}

				data := buffer[:n]
				_, err = s.Conn.WriteToUDP(data, addr)
				if err != nil {
					fmt.Printf("sendto error: %v\n", err)
					continue
				}
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
