// Package csrandsrc provides a source for math/rand via crypto/rand
package csrandsrc

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

type Rand struct{}

func (r Rand) Int63() int64 {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		// :(
		panic("no crypto source")
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}

func (r Rand) Seed(s int64) {
	// ignored
}

var _ = mrand.Source(Rand{})
