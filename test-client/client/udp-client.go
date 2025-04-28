package client

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"test-client/option"
	"time"
)

type UDPTestSuite struct {
	host    string
	port    int
	options *option.Options

	tx atomic.Int64 // packet sent
	rx atomic.Int64 // packet received
}

func parseProtoAddr(protoAddr string) (string, int) {

	// Split the protoAddr into host and port
	host := "localhost"
	port := 5000

	fmt.Printf("server address:%s\n", protoAddr)
	section := strings.Split(protoAddr, ":")
	if len(section) != 2 {
		panic("invalid address format: " + protoAddr)
	}

	host = section[0]
	port, err := strconv.Atoi(section[1])
	if err != nil {
		panic("invalid address format: " + protoAddr)
	}
	return host, port
}

func loadOptions(options ...option.Option) *option.Options {
	opts := new(option.Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

func NewUDPClient(protoAddr string, options ...option.Option) *UDPTestSuite {
	client := &UDPTestSuite{}
	client.host, client.port = parseProtoAddr(protoAddr)
	client.options = loadOptions(options...)
	return client
}

func generateRandomBytes(length int, buffer []byte) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := len(charset)
	for i := 0; i < length; i++ {
		buffer[i] = charset[rand.Intn(charsetLen)]
	}
}

func (c *UDPTestSuite) Start() {
	// Placeholder for the UDP client logic
	fmt.Printf("UDP Client is running, starting %d connections...\n", c.options.Concurrency)

	if c.options.Duration == 0 {
		c.options.Duration = 24 * time.Hour
	}

	var wg sync.WaitGroup
	errChan := make(chan error, c.options.Concurrency)

	ctx, cancel := context.WithCancel(context.Background())
	startTime := time.Now()

	for i := range c.options.Concurrency {
		wg.Add(1)

		go func(idx int) {
			client, err := net.Dial("udp", fmt.Sprintf("%s:%d", c.host, c.port))
			wbuf := make([]byte, c.options.PacketSize)
			rbuf := make([]byte, c.options.PacketSize) // we don't care about the content
			if err != nil {
				fmt.Printf("Goroutine[%d] connect to server error: %s", idx, err.Error())
				errChan <- err
			}

		loop:
			for {
				select {
				case <-ctx.Done():
					break loop
				default:
					// Send a packet
					generateRandomBytes(c.options.PacketSize, wbuf)

					if c.options.Debug {
						fmt.Printf("buffer: %s\n", string(wbuf))
					}

					client.SetWriteDeadline(time.Now().Add(c.options.WriteTimeout * time.Millisecond))

					_, err = client.Write(wbuf)
					if err == nil {
						c.tx.Add(1)
					} else {
						fmt.Printf("Goroutine[%d] sending packet error: %s\n", idx, err.Error())
						// TODO: some error is considered as a success
						continue
					}

					_, err = client.Read(rbuf) // read the response and drop it
					client.SetReadDeadline(time.Now().Add(c.options.ReadTimeout * time.Millisecond))

					if err == nil {
						c.rx.Add(1)
					} else {
						fmt.Printf("Goroutine[%d] reading packet error: %s\n", idx, err.Error())
						continue
					}
				}
			}
			defer wg.Done()
		}(i)
	}

	select {
	case <-time.After(c.options.Duration):
		cancel()
	case e := <-errChan:
		fmt.Printf("get error from client: %s\n", e.Error())
		cancel()
	}

	wg.Wait()
	endTime := time.Now()
	fmt.Printf("UDP Client finished. Sent: %d, Received: %d\n", c.tx.Load(), c.rx.Load())
	fmt.Printf("Duration: %s\n", endTime.Sub(startTime))
	fmt.Printf("Throughput: %f packets/sec\n", float64(c.tx.Load())/endTime.Sub(startTime).Seconds())
	fmt.Printf("Latency: %f ms\n", float64(endTime.Sub(startTime).Milliseconds())/float64(c.tx.Load()))
	fmt.Printf("Packet Loss: %f%%\n", float64(c.tx.Load()-c.rx.Load())/float64(c.tx.Load())*100)
}
