package main

import (
	"flag"
	"fmt"
	"gnet1-benchmark/server"
	"log"

	"github.com/panjf2000/gnet"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	address := flag.String("bind", ":5000", "server listen address")
	paddress := flag.String("pprof", "localhost:5001", "profile listen address")
	poolSize := flag.Int("pool-size", 1024*1024, "size of ant pool")
	flag.Parse()

	server := server.NewUDPServer(*poolSize)

	// Placeholder for the UDP server logic
	log.Printf("UDP Server is running on %s\n", *address)

	go func() {
		log.Printf("pprof listening on %s\n", *paddress)
		if err := http.ListenAndServe(*paddress, nil); err != nil {
			log.Fatal(err)
		}
	}()

	gnet.Serve(server, fmt.Sprintf("udp://%s", *address),
		gnet.WithReusePort(true),
		gnet.WithMulticore(true))
}
