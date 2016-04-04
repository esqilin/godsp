package dsp

import (
	"fmt"
	"os"
)

type Sound interface {
	OnEnd(func())
	Sample() float64
	Stop()
	IsFinished() bool
	Release()
}

// SoundBase provides the OnEnd and Stop functions to remove itself from its
// Sampler
type SoundBase struct {
	onEnd      func()
	isFinished bool
}

func (s *SoundBase) OnEnd(callback func()) {
	if nil != s.onEnd {
		fmt.Fprintln(os.Stderr, "WARNING: Sound.OnEnd called twice; potential memory leak")
	}
	s.onEnd = callback
}

func (s *SoundBase) Stop() {
	if nil != s.onEnd {
		s.onEnd()
		s.onEnd = nil
	}
	s.isFinished = true
}

func (s SoundBase) IsFinished() bool {
	return s.isFinished
}

func (s *SoundBase) Release() {
	s.Stop()
}
