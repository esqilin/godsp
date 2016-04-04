package sound

import (
	"github.com/esqilin/godsp"
	"github.com/esqilin/godsp/wave"
)

type WaveSound struct {
	dsp.SoundBase
	w wave.Wave
}

func NewWaveSound(w wave.Wave) *WaveSound {
	return &WaveSound{dsp.SoundBase{}, w}
}

func (s WaveSound) Sample() float64 {
	return s.w.Sample()
}
