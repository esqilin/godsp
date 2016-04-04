package wave

type Wave interface {
	Sample() float64
	SetFreq(float64)
}
