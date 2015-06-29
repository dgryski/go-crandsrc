package crandsrc

import (
	"math/rand"
	"time"
)

func ExampleRand() {

	var r *rand.Rand

	// Instead of this
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Use this...
	r = rand.New(Source{})

	// and then call this...
	r.Int63()
}
