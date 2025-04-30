package main

import (
	"flag"
	"go124-benchmark/server"
	"runtime"
)

func main() {
	threads := flag.Int("threads", runtime.NumCPU(), "Number of server goroutines")
	address := flag.String("address", ":5000", "server listen address")
	flag.Parse()

	udpserver, e := server.NewUDPServer(*address, *threads)
	if e != nil {
		panic(e)
	}
	udpserver.Serve()
}
