package wave

// TODO: bullshit all this
type BiquadWave struct {
	s1, s2, a1, a2 float64
}

func NewBiquadWave(a [3]float64, b [2]float64) *BiquadWave {
	return &BiquadWave{a[2] + a[1], a[2], a[1], a[2]}
}

func (w *BiquadWave) Sample() float64 {
	y := w.s1
	w.s1 = w.s2 - w.a1*y
	w.s2 = -w.a2 * y
	return y
}

func (w BiquadWave) OnEnd(func()) {
	// empty
}

func (w BiquadWave) Stop() {
	// emtpy
}
