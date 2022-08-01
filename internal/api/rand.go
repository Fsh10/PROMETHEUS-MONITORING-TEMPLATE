package api


import (
	"math/rand"
	"time"
	"fmt"

)

var (
	// rng is a shared random number generator for the package
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)