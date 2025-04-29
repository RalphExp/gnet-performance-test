package main

import (
	"gnet2-benchmark/server"

	"github.com/panjf2000/gnet/v2"
)

func main() {
	server := server.NewUDPServer()

	// Placeholder for the UDP server logic
	println("UDP Server is running...")

	gnet.Run(server, "udp://:5000",
		gnet.WithReusePort(true),
		gnet.WithMulticore(true))
}
