package main

import (
	"flag"
	"fmt"
	"gnet1-benchmark/server"

	"github.com/panjf2000/gnet"
)

func main() {
	address := flag.String("address", ":5000", "server listen address")
	flag.Parse()

	server := server.NewUDPServer()

	// Placeholder for the UDP server logic
	println("UDP Server is running...")

	gnet.Serve(server, fmt.Sprintf("udp://%s", *address),
		gnet.WithReusePort(true),
		gnet.WithMulticore(true))
}
