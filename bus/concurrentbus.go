package bus

import (
	"github.com/esqilin/godsp"
)

type ConcurrentBus struct {
	dsp.SoundBase
	cs map[<-chan float64]struct{}
}

func NewConcurrent() *ConcurrentBus {
	return &ConcurrentBus{dsp.SoundBase{}, make(map[<-chan float64]struct{})}
}

func (b *ConcurrentBus) Sample() float64 {
	out := 0.0
	for ch := range b.cs {
		v, more := <-ch
		if !more {
			delete(b.cs, ch)
			continue
		}
		out += v
	}
	return out
}

func (b *ConcurrentBus) PlaySound(s dsp.Sound) {
	ch := make(chan float64, 1)
	b.cs[ch] = struct{}{}
	done := make(chan struct{}, 1)

	s.OnEnd(func() { done <- struct{}{} })

	go sampleSound(s, ch, done)
}

func sampleSound(s dsp.Sound, ch chan<- float64, done <-chan struct{}) {
	for {
		select {
		case <-done:
			close(ch)
			return
		default:
			ch <- s.Sample()
		}
	}
}
