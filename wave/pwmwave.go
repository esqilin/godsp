package wave

import (
	"github.com/esqilin/godsp/waveshape"

	"errors"
)

type PWMWave struct {
	pWavShape *waveshape.Pulse
	pWav      *ShapedWave
	lfo       Wave
	amount    float64
}

func NewPWMWave(freq float64, lfo Wave, amount float64) (*PWMWave, error) {
	if nil == lfo {
		return nil, errors.New("call to NewPWMWave with nil argument")
	}
	ws := waveshape.NewPulse(0.0)
	w, _ := NewShapedWave(freq, ws)
	wav := &PWMWave{
		pWavShape: ws,
		pWav:      w,
		lfo:       lfo,
	}
	wav.SetAmount(amount)
	return wav, nil
}

func (wav *PWMWave) SetAmount(amount float64) {
	wav.amount = 0.5 * amount
}

func (wav *PWMWave) SetFreq(freq float64) {
	wav.SetFreq(freq)
}

func (wav *PWMWave) Sampe() float64 {
	wav.pWavShape.SetLength(0.5 + wav.amount*wav.lfo.Sample())
	return wav.pWav.Sample()
}
