package waveshape

import (
	"math/rand"
	"time"
)

type WhiteNoise struct {
	rGen *rand.Rand
}

func NewWhiteNoise() *WhiteNoise {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &WhiteNoise{rGen: r}
}

func (p WhiteNoise) Sample(pos float64) float64 {
	return p.rGen.Float64()
}
