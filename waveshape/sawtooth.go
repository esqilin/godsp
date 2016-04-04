package waveshape

type sawtooth struct{}

func NewSawtooth() *sawtooth {
	return &sawtooth{}
}

func (sawtooth) Sample(pos float64) float64 {
	return pos*2.0 - 1.0
}
