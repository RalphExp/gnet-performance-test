package option

import "time"

type Option func(opts *Options)

type Options struct {
	Concurrency  int
	PacketSize   int
	ReadTimeout  time.Duration // Timeout for sending packets
	WriteTimeout time.Duration // Timeout for receiving packets
	Duration     time.Duration // Duration of the test
	Debug        bool
}

func WithConcurrency(concurrency int) Option {
	return func(opts *Options) {
		opts.Concurrency = concurrency
	}
}

func WithPacketSize(packetSize int) Option {
	return func(opts *Options) {
		opts.PacketSize = packetSize
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.WriteTimeout = timeout
	}
}

func WithDuration(duration time.Duration) Option {
	return func(opts *Options) {
		opts.Duration = duration
	}
}

func WithDebug(debug bool) Option {
	return func(opts *Options) {
		opts.Debug = debug
	}
}
