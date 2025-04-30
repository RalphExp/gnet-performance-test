package main

import (
	"flag"
	"fmt"
	"gnet2-benchmark/server"

	"github.com/panjf2000/gnet/v2"
)

func main() {
	address := flag.String("bind", ":5000", "server listen address")
	poolSize := flag.Int("pool-size", 1024*1024, "ant pool size")
	flag.Parse()
	server := server.NewUDPServer(*poolSize)

	// Placeholder for the UDP server logic
	println("UDP Server is running...")

	gnet.Run(server, fmt.Sprintf("udp://%s", *address),
		gnet.WithReusePort(true),
		gnet.WithMulticore(true))
}
