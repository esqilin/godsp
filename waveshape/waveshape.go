package waveshape

import "math"

const PI2 = 2 * math.Pi

var (
	nSamples  int
	nSamplesF float64
	sineTable []float64
)

type WaveShape interface {
	Sample(float64) float64
}

func Init(maxSamples float64) {
	nSamplesF = maxSamples
	nSamples = int(nSamplesF)
	initSineTable()
}

func MaxSamples() int {
	return nSamples
}
