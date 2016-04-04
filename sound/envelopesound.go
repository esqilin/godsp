package sound

import (
	"github.com/esqilin/godsp"
)

type EnvelopeSound struct {
	dsp.SoundBase
	sound    dsp.Sound
	envelope *dsp.Envelope
	velocity float64
}

func NewEnvelopeSound(envelope *dsp.Envelope, sound dsp.Sound, velocity float64) *EnvelopeSound {
	es := &EnvelopeSound{sound: sound, envelope: envelope, velocity: velocity}
	envelope.OnEnd(func() { sound.Stop() })
	sound.OnEnd(func() { es.Stop() })
	return es
}

func (es EnvelopeSound) Sample() float64 {
	return es.velocity * es.sound.Sample() * es.envelope.Sample()
}

func (es *EnvelopeSound) Release() {
	es.envelope.Release()
}
