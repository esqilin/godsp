package waveshape

import "math"

type triangle struct{}

func NewTriangle() *triangle {
	return &triangle{}
}

func (triangle) Sample(pos float64) float64 {
	return 1.0 - 4.0*math.Abs(pos-0.5)
}
