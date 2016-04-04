package sound

import (
	"github.com/esqilin/godsp"
)

type BufferedSound struct {
	dsp.SoundBase
	s    dsp.Sound
	ch   <-chan float64
	done chan<- struct{}
}

func NewBufferedSound(s dsp.Sound, bufferSize int) *BufferedSound {
	ch := make(chan float64, bufferSize)
	done := make(chan struct{}, 1)
	s.OnEnd(func() {
		done <- struct{}{}
	})
	go buffer(s, ch, done)
	return &BufferedSound{dsp.SoundBase{}, s, ch, done}
}

func (s BufferedSound) Sample() float64 {
	out, more := <-s.ch
	if !more {
		s.Stop()
	}
	return out
}

func buffer(s dsp.Sound, ch chan<- float64, done <-chan struct{}) {
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
