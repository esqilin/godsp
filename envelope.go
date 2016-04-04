package dsp

import (
	"github.com/esqilin/godsp/easing"
)

type Envelope struct {
	SoundBase
	noSustain bool
	stages    [4]*envStage
	curStage  *envStage
}

// NewEnvelope returns a pointer to an Envelope struct,
func NewEnvelope(attack, decay, sustain, release float64) *Envelope {
	out := &Envelope{}

	rs := newReleaseStage(release, easing.QuadIn, sustain)
	ss := newSustainStage(sustain, rs, &out.noSustain)
	ds := newDecayStage(decay, easing.QuadIn, ss, sustain)
	as := newAttackStage(attack, easing.QuadOut, ds)
	out.stages = [4]*envStage{as, ds, ss, rs}
	out.curStage = as
	return out
}

func (e *Envelope) Restart() {
	e.curStage = e.stages[0]
	e.curStage.reset()
	e.noSustain = false
}

func (e *Envelope) Sample() float64 {
	if nil == e.curStage {
		e.Stop()
		return 0.0
	}
	out := e.curStage.Sample()
	e.curStage = e.curStage.Next
	return out
}

func (e *Envelope) Release() {
	e.noSustain = true
}

func (e Envelope) IsFinished() bool {
	return nil == e.curStage
}
