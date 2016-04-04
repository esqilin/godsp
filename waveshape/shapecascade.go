package waveshape

import "errors"

type ShapeModulator interface {
	Sample() float64
}

type ShapeSwitcher struct {
	s1, s2 WaveShape
	mod    ShapeModulator
}

func NewShapeSwitcher(shape1, shape2 WaveShape, modulator ShapeModulator) (*ShapeSwitcher, error) {
	if nil == shape1 {
		return nil, errors.New("shape1 cannot be nil")
	}
	if nil == shape2 {
		return nil, errors.New("shape2 cannot be nil")
	}
	if nil == modulator {
		return nil, errors.New("modulator cannot be nil")
	}
	return &ShapeSwitcher{shape1, shape2, modulator}, nil
}

func (ss ShapeSwitcher) Sample(pos float64) float64 {
	mod := ss.mod.Sample()
	return mod*ss.s1.Sample(pos) + (1.0-mod)*ss.s2.Sample(pos)
}
