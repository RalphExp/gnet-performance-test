package main

import (
	"flag"
	"log"

	"go124-benchmark/server"

	"net/http"
	_ "net/http/pprof"
)

func main() {
	address := flag.String("bind", ":5000", "server listen address")
	paddress := flag.String("prof", "localhost:5001", "profile listen address")
	poolSize := flag.Int("pool-size", 1024*1024, "size of ant pool")
	flag.Parse()

	udpserver, e := server.NewUDPServer(*address, *poolSize)
	if e != nil {
		log.Printf("Error creating UDP server: %v\n", e)
		panic(e)
	}

	go func() {
		log.Printf("pprof listening on %s\n", *paddress)
		if err := http.ListenAndServe(*paddress, nil); err != nil {
			log.Fatal(err)
		}
	}()
	udpserver.Serve()
}
