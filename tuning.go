package dsp

import "math"

type Tuning [128]float64

func NewEqualTemperedTuning(a4freq float64) Tuning {
	fs := Tuning{}
	base := math.Pow(2.0, 1.0/12.0)
	for i := 0.0; i < 128.0; i += 1.0 {
		fs[int(i)] = a4freq * math.Pow(base, i-69.0)
	}
	return fs
}
