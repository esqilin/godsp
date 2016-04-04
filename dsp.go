package dsp

import (
	"github.com/esqilin/godsp/easing"
	"github.com/esqilin/godsp/wave"
	"github.com/esqilin/godsp/waveshape"
)

const (
	OSC_MIN_FREQUENCY = 0.1
)

type Easing easing.Easing
type EasingFun easing.EasingFun

var (
	sampleRate, sampleInterval float64 // interval in seconds
	Sine                       waveshape.WaveShape
	Square                     waveshape.WaveShape
	Sawtooth                   waveshape.WaveShape
	Triangle                   waveshape.WaveShape
	StdTuning                  Tuning
)

func Init(sampleRate_ uint32) {
	sampleRate = float64(sampleRate_)
	sampleInterval = 1.0 / sampleRate
	wave.Init(sampleInterval)
	maxSamples := sampleRate / OSC_MIN_FREQUENCY
	waveshape.Init(maxSamples)

	var err error
	Sine, err = waveshape.NewSine()
	if nil != err {
		panic(err.Error())
	}
	Square = waveshape.NewSquare()
	Triangle = waveshape.NewTriangle()
	Sawtooth = waveshape.NewSawtooth()

	StdTuning = NewEqualTemperedTuning(440.0)
}
