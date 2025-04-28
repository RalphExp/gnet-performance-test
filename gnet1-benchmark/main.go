package main

import (
	"gnet1-benchmark/server"

	"github.com/panjf2000/gnet"
)

func main() {
	server := server.NewUDPServer()

	// Placeholder for the UDP server logic
	println("UDP Server is running...")

	gnet.Serve(server, "udp://:5000",
		gnet.WithReusePort(true),
		gnet.WithMulticore(true))
}
