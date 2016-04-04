package wave

import "errors"

type AMWave struct {
	lfo     Wave
	carrier Wave
	amount  float64
}

func NewAMWave(lfo, carrier Wave, amount float64) (*AMWave, error) {
	if nil == lfo || nil == carrier {
		return nil, errors.New("call to NewAMWave with nil argument")
	}
	return &AMWave{
		lfo:     lfo,
		carrier: carrier,
		amount:  amount,
	}, nil
}

func (wav *AMWave) SetFreq(freq float64) {
	wav.carrier.SetFreq(freq)
}

func (wav *AMWave) Sample() float64 {
	return wav.amount * wav.lfo.Sample() * wav.carrier.Sample()
}
