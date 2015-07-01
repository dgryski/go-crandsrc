// Package crandsrc provides a source for math/rand via crypto/rand
package crandsrc

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

type Source struct{}

func (s Source) Int63() int64 {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		// :(
		panic("no crypto source")
	}
	return int64(binary.LittleEndian.Uint64(b[:])) & 0x7fffffffffffffff
}

func (r Source) Seed(s int64) {
	// ignored
}

var _ = mrand.Source(Source{})
