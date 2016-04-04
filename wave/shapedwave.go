package wave

import (
	"github.com/esqilin/godsp/waveshape"

	"errors"
)

var sampleInterval float64

type ShapedWave struct {
	pos, step float64
	shaper    waveshape.WaveShape
}

func Init(sampleInterval_ float64) {
	sampleInterval = sampleInterval_
}

func NewShapedWave(freq float64, waveShape waveshape.WaveShape) (*ShapedWave, error) {
	wav := &ShapedWave{}
	wav.SetFreq(freq)
	err := wav.SetShape(waveShape)
	return wav, err
}

func (wav ShapedWave) GetFreq() float64 {
	return wav.step / sampleInterval
}

func (wav *ShapedWave) SetFreq(freq float64) {
	wav.step = freq * sampleInterval
}

func (wav *ShapedWave) SetShape(waveShape waveshape.WaveShape) error {
	if nil == waveShape {
		return errors.New("call to ShapedWave.SetShape with nil")
	}
	wav.shaper = waveShape
	return nil
}

func (wav *ShapedWave) Sample() float64 {
	if wav.pos >= 1.0 {
		wav.pos -= 1.0
	}
	val := wav.shaper.Sample(wav.pos)
	wav.pos += wav.step
	return val
}
