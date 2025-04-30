module gnet1-benchmark

go 1.24

require (
	github.com/panjf2000/gnet v1.4.4
	test-util v0.0.0-00010101000000-000000000000
)

require github.com/panjf2000/ants/v2 v2.4.4 // indirect

require (
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
)

replace test-util => ../test-util
