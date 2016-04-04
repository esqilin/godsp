package waveshape

type moog struct{}

func NewMoog() *moog {
	return &moog{}
}

func (moog) Sample(pos float64) float64 {
	out := pos*2.0 - 1.0
	// das is irgendwas:
	if 0.049 > pos {
		return -0.95 + 0.05*sineTable[int(nSamplesF*(pos*20))]
	} else if 0.05 > pos {
		out += 1.0
	} else if 0.98 < pos {
		return 0.98 + 0.02*sineTable[int(nSamplesF*((1.0-pos)*20))]
	} else if 0.975 < pos {
		out -= 1.0
	}
	return out
}
