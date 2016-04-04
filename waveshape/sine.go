package waveshape

// TODO: make WaveTable struct: func Sample will always be the same

import (
	"errors"
	"math"
)

func initSineTable() {
	sineTable = make([]float64, nSamples)
	delta := PI2 / nSamplesF
	for i := 0; i < nSamples; i++ {
		sineTable[i] = float64(math.Sin(float64(i) * delta))
	}
}

func NewSine() (*lookupShape, error) {
	if nil == sineTable {
		return nil, errors.New("wave table not initialized (call waveshape.Init first)")
	}
	return NewLookupShape(&sineTable)
}
