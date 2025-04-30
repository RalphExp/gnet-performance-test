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

	txError atomic.Int64 // packet send error
	rxError atomic.Int64 // packet receive error
}

func parseProtoAddr(protoAddr string) (string, int) {

	// Split the protoAddr into host and port
	host := "localhost"
	port := 0

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

func generateRandomBytes(length int, buffer []byte) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLen := len(charset)
	for i := 0; i < length; i++ {
		buffer[i] = charset[rand.Intn(charsetLen)]
	}
}

func NewUDPTestSuite(protoAddr string, options ...option.Option) *UDPTestSuite {
	client := &UDPTestSuite{}
	client.host, client.port = parseProtoAddr(protoAddr)
	client.options = loadOptions(options...)
	return client
}

func (suite *UDPTestSuite) Start() {
	// Placeholder for the UDP client logic
	fmt.Printf("UDP Client is running, starting %d connections...\n", suite.options.Threads)

	if suite.options.Duration == 0 {
		suite.options.Duration = 24 * time.Hour
	}

	var wg sync.WaitGroup
	errChan := make(chan error, suite.options.Threads)

	ctx, cancel := context.WithCancel(context.Background())
	startTime := time.Now()

	for i := 0; i < suite.options.Threads; i++ {
		func(idx int) {
			conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", suite.host, suite.port))
			if err != nil {
				fmt.Printf("Goroutine[%d] connect to server error: %s", idx, err.Error())
				errChan <- err
				return
			}
			wg.Add(2)

			// recv Goroutine
			go func() {
				defer wg.Done()
				rbuf := make([]byte, suite.options.PacketSize)
				for {
					select {
					case <-ctx.Done():
						delay := time.After(2 * time.Second)
						for {
							select {
							case <-delay:
								conn.Close()
								return
							default:
								conn.SetReadDeadline(time.Now().Add(suite.options.ReadTimeout))
								n, err := conn.Read(rbuf)
								if err == nil {
									if n != suite.options.PacketSize {
										// check packet size
										suite.rxError.Add(1)
									} else {
										suite.rx.Add(1)
									}
								} else if !strings.Contains(err.Error(), "i/o timeout") {
									suite.rxError.Add(1)
									if suite.options.Debug {
										fmt.Printf("Goroutine[%d] recv error: %s\n", idx, err.Error())
									}
								}
							}
						}
					default:
						conn.SetReadDeadline(time.Now().Add(suite.options.ReadTimeout))
						n, err := conn.Read(rbuf)
						if err == nil {
							if n != suite.options.PacketSize {
								// check packet size
								suite.rxError.Add(1)
							} else {
								suite.rx.Add(1)
							}
						} else if !strings.Contains(err.Error(), "i/o timeout") {
							suite.rxError.Add(1)
							if suite.options.Debug {
								fmt.Printf("Goroutine[%d] recv error: %s\n", idx, err.Error())
							}
						}
					}
				}
			}()

			// send Goroutine
			go func() {
				defer wg.Done()
				wbuf := make([]byte, suite.options.PacketSize)
				for {
					select {
					case <-ctx.Done():
						return
					default:
						generateRandomBytes(suite.options.PacketSize, wbuf)
						// if suite.options.Debug {
						// 	fmt.Printf("Goroutine[%d] sent buffer: %s\n", idx, string(wbuf))
						// }
						_, err := conn.Write(wbuf)
						if err == nil {
							suite.tx.Add(1)
							time.Sleep(time.Millisecond)
						} else {
							suite.txError.Add(1)
							if suite.options.Debug {
								fmt.Printf("Goroutine[%d] send error: %s\n", idx, err.Error())
							}
						}
					}
				}
			}()
		}(i)
	}

	select {
	case <-time.After(suite.options.Duration):
		cancel()
	case e := <-errChan:
		fmt.Printf("get error from client: %s\n", e.Error())
		cancel()
	}

	wg.Wait()
	endTime := time.Now()
	fmt.Printf("UDP Client finished. Send: %d, Recv: %d, Send Error: %d, Recv Error: %d\n",
		suite.tx.Load(), suite.rx.Load(), suite.txError.Load(), suite.rxError.Load())
	fmt.Printf("Duration: %s\n", endTime.Sub(startTime))
	fmt.Printf("Throughput: %f packets/sec\n", float64(suite.tx.Load())/endTime.Sub(startTime).Seconds())
	fmt.Printf("Latency: %f ms\n", float64(endTime.Sub(startTime).Milliseconds())/float64(suite.tx.Load()))
	fmt.Printf("Packet Loss: %f%%\n", float64(suite.tx.Load()-suite.rx.Load())/float64(suite.tx.Load())*100)
}
