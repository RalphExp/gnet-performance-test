package option

import "time"

type Option func(opts *Options)

type Options struct {
	Threads      int
	PacketSize   int
	ReadTimeout  time.Duration // Timeout for sending packets
	WriteTimeout time.Duration // Timeout for receiving packets
	Duration     time.Duration // Duration of the test
	Debug        bool
}

func WithThreads(threads int) Option {
	return func(opts *Options) {
		opts.Threads = threads
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
