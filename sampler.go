package dsp

type Sampler interface {
	PlaySound(*Sound)
	Sample() float64
}

// SimpleSampler can play on Sound at a time
type SimpleSampler struct {
	SoundBase
	s Sound
}

func NewSimpleSampler() *SimpleSampler {
	return &SimpleSampler{SoundBase{}, nil}
}

func (s *SimpleSampler) StopSound() {
	s.s.Stop()
}

func (s *SimpleSampler) unsetSound() {
	s.s = nil
}

func (s *SimpleSampler) PlaySound(so Sound) {
	if nil != s.s {
		s.s.Stop()
	}
	s.s = so
	so.OnEnd(s.unsetSound)
}

func (s *SimpleSampler) Stop() {
	s.StopSound()
	s.SoundBase.Stop()
}

func (s *SimpleSampler) Sample() float64 {
	return s.s.Sample()
}

// MultiSampler can play multiple Sounds at the same time
type MultiSampler struct {
	SoundBase
	in map[Sound]struct{}
}

func NewMultiSampler() *MultiSampler {
	return &MultiSampler{SoundBase{}, make(map[Sound]struct{})}
}

func (s *MultiSampler) PlaySound(so Sound) {
	s.in[so] = struct{}{}
	so.OnEnd(func() { delete(s.in, so) })
}

func (s *MultiSampler) Sample() float64 {
	var out float64
	for sa := range s.in {
		out += sa.Sample()
	}
	return out
}
