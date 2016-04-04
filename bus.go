package dsp

type Bus interface {
	Sound
	AddSound(Sound) float64
}
