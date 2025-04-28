package main

import (
	"flag"
	"test-client/client"
	"test-client/option"
)

func main() {
	// Define command-line flags
	serverAddr := flag.String("server", "localhost:5000", "Server address")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent connections")
	packetSize := flag.Int("packet-size", 1024, "Size of each packet in bytes")
	readTimeout := flag.Duration("read-timeout", 100000000, "Read timeout duration")
	writeTimeout := flag.Duration("write-timeout", 100000000, "Write timeout duration")
	duration := flag.Duration("duration", 1000000000, "Duration of the test")
	debug := flag.Bool("debug", false, "Enable debug mode")

	// Parse command-line flags
	flag.Parse()

	// Print the parsed options
	println("Server Address:", *serverAddr)
	println("Concurrency:", *concurrency)
	println("Packet Size:", *packetSize)
	println("Read Timeout(ms):", *readTimeout/1000000)
	println("Write Timeout(ms):", *writeTimeout/1000000)
	println("Duration (ms):", *duration/1000000)
	println("Debug Mode:", *debug)

	cli := client.NewUDPClient(*serverAddr,
		option.WithConcurrency(*concurrency),
		option.WithPacketSize(*packetSize),
		option.WithReadTimeout(*readTimeout),
		option.WithWriteTimeout(*writeTimeout),
		option.WithDuration(*duration),
		option.WithDebug(*debug),
	)
	cli.Start()
}
