package wave

import "errors"

type FMWave struct {
	lfo        Wave
	carrier    *ShapedWave
	centerFreq float64
	amount     float64
}

func NewFMWave(lfo Wave, carrier *ShapedWave, amount float64) (*FMWave, error) {
	if nil == lfo || nil == carrier {
		return nil, errors.New("call to NewFMWave with nil argument")
	}
	return &FMWave{
		lfo:        lfo,
		carrier:    carrier,
		centerFreq: carrier.GetFreq(),
		amount:     amount,
	}, nil
}

func (wav *FMWave) SetFreq(freq float64) {
	wav.centerFreq = freq
}

func (wav *FMWave) Sample() float64 {
	wav.carrier.SetFreq(wav.lfo.Sample())
	return wav.carrier.Sample()
}
