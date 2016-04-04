package waveshape

type Pulse struct {
	length float64
}

func NewPulse(length float64) *Pulse {
	out := &Pulse{
		length: length,
	}
	return out
}

func (p *Pulse) SetLength(length float64) {
	p.length = length
}

func (p Pulse) Sample(pos float64) float64 {
	if pos < p.length {
		return 1.0
	}
	return -1.0
}
