package main

import (
	"flag"
	"fmt"
	"go124-benchmark/server"
	"runtime"
)

func main() {
	threads := flag.Int("threads", runtime.NumCPU(), "Number of server goroutines")
	address := flag.String("bind", ":5000", "server listen address")
	poolSize := flag.Int("pool-size", 1024*1024, "size of ant pool")
	flag.Parse()

	udpserver, e := server.NewUDPServer(*address, *threads, *poolSize)
	if e != nil {
		fmt.Printf("Error creating UDP server: %v\n", e)
		panic(e)
	}
	udpserver.Serve()
}
