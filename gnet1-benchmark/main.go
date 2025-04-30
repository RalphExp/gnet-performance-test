package main

import (
	"flag"
	"fmt"
	"gnet1-benchmark/server"

	"github.com/panjf2000/gnet"
)

func main() {
	address := flag.String("bind", ":5000", "server listen address")
	poolSize := flag.Int("pool-size", 1024*1024, "size of ant pool")
	flag.Parse()

	server := server.NewUDPServer(*poolSize)

	// Placeholder for the UDP server logic
	println("UDP Server is running...")

	gnet.Serve(server, fmt.Sprintf("udp://%s", *address),
		gnet.WithReusePort(true),
		gnet.WithMulticore(true))
}
