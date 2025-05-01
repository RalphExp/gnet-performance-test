package util

import (
	"math/rand"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

const (
	antsPoolExpiryDuration = 120 * time.Second
	antsPoolNonblocking    = true
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var bufPool = sync.Pool{
	New: func() any {
		return make([]byte, 1024)
	},
}

func GetBuffer() []byte {
	return bufPool.Get().([]byte)
}

func PutBuffer(b any) {
	bufPool.Put(b)
}

func GenerateRandomBytes(_ []byte, size int) (out []byte, err error) {
	out = make([]byte, size)

	charsetLen := len(charset)
	for i := range size {
		out[i] = charset[rand.Intn(charsetLen)]
	}
	return out, err
}

func CreateAntsPool(poolSize int) *ants.Pool {
	options := ants.Options{
		ExpiryDuration: antsPoolExpiryDuration,
		Nonblocking:    antsPoolNonblocking,
		PreAlloc:       false,
	}

	pool, _ := ants.NewPool(poolSize, ants.WithOptions(options))
	return pool
}
