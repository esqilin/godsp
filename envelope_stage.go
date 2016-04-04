package dsp

import (
	"github.com/esqilin/godsp/easing"

	"math"
)

type envStage struct {
	iSample, nSamples, step float64
	easing                  easing.EasingFun
	next                    *envStage
	Next                    *envStage
	sampleFun               func(float64) float64
}

func newEnvStage(duration float64, easingFun easing.EasingFun, next *envStage) *envStage {
	nSamples := duration * sampleRate
	es := &envStage{
		nSamples:  nSamples,
		step:      1.0 / nSamples,
		easing:    easingFun,
		next:      next,
		sampleFun: func(x float64) float64 { return x },
	}
	es.Next = es
	return es
}

func (es *envStage) Sample() float64 {
	out := es.easing(es.iSample * es.step)
	es.iSample += 1.0
	if es.iSample >= es.nSamples {
		es.proceed()
	}
	return es.sampleFun(out)
}

func (es *envStage) proceed() {
	if nil != es.next {
		es.next.reset()
	}
	es.Next = es.next
}

func (es *envStage) reset() {
	es.iSample = 0
	es.Next = es
}

func newAttackStage(duration float64, easing easing.Easing, next *envStage) *envStage {
	return newEnvStage(duration, easing.In, next)
}

// newDecayStage target must be in [0.0, 1.0]
func newDecayStage(duration float64, easing easing.Easing, next *envStage, target float64) *envStage {
	rangeSize := 1.0 - target
	out := newEnvStage(duration, easing.Out, next)
	out.sampleFun = func(x float64) float64 { return target + rangeSize*x }
	return out
}

// newDecayStage level must be in [0.0, 1.0]
func newSustainStage(level float64, next *envStage, noSustain *bool) *envStage {
	out := newEnvStage(math.MaxFloat64, easing.Linear.In, next)
	out.sampleFun = func(float64) float64 {
		if *noSustain {
			out.nSamples = 0.0
		}
		return level
	}
	return out
}

// newDecayStage origin must be in [0.0, 1.0]
func newReleaseStage(duration float64, easing easing.Easing, origin float64) *envStage {
	out := newEnvStage(duration, easing.Out, nil)
	out.sampleFun = func(x float64) float64 { return origin * x }
	return out
}
