// Package easing provides all the implemented easing functions.
// The type easing is implemented in package dsp. You need to call Init
package easing

import (
	"math"
)

// must map [0.0, 1.0] to [0.0, 1.0]
type EasingFun func(float64) float64
type Easing struct {
	In, Out EasingFun
}

var (
	Linear = Easing{
		In:  linearIn,
		Out: linearOut,
	}
	QuadIn = Easing{
		In:  quadInIn,
		Out: quadInOut,
	}
	QuadOut = Easing{
		In:  quadOutIn,
		Out: quadOutOut,
	}
	QuadInOut = Easing{
		In:  quadInOutIn,
		Out: quadInOutOut,
	}
	QuadOutIn = Easing{
		In:  quadOutInIn,
		Out: quadOutInOut,
	}
	ExpIn = Easing{
		In:  expInIn,
		Out: expInOut,
	}
	ExpOut = Easing{
		In:  expOutIn,
		Out: expOutOut,
	}
	ExpInOut = Easing{
		In:  expInOutIn,
		Out: expInOutOut,
	}
	ExpOutIn = Easing{
		In:  expOutInIn,
		Out: expOutInOut,
	}
)

// y = x
func linearIn(x float64) float64 {
	return x
}

// y = -x + 1
func linearOut(x float64) float64 {
	return 1.0 - x
}

// y = x²
func quadInIn(x float64) float64 {
	return x * x
}

// y = -x² - 2x + 1
func quadInOut(x float64) float64 {
	return 1.0 + x*(x-2.0)
}

// y = -x² + 2x
func quadOutIn(x float64) float64 {
	return x * (2.0 - x)
}

// y = -x² + 1
func quadOutOut(x float64) float64 {
	return 1.0 - x*x
}

// y([0 - 0.5[) = 2x²
// y([0.5 - 1]) = 2 * (-x² + 2x) - 1
func quadInOutIn(x float64) float64 {
	if x < 0.5 {
		return 2 * x * x
	}
	return 2.0*x*(2.0-x) - 1.0
}

// y([0 - 0.5[) = 1 - 2x²
// y([0.5 - 1[) = 2 * (x² - 2x) + 2
func quadInOutOut(x float64) float64 {
	if x < 0.5 {
		return 1.0 - 2.0*x*x
	}
	return 2.0 * (x*(x-2.0) + 1.0)
}

// y([0 - 0.5[) = -2 * (x² - x)
// y([0.5 - 1]) = 2 * (x² - x) + 1
func quadOutInIn(x float64) float64 {
	if x < 0.5 {
		return 2.0 * x * (1.0 - x)
	}
	return 2.0*x*(x-1.0) + 1
}

// y([0 - 0.5[) = 2 * (x² - x) + 1
// y([0.5 - 1]) = -2 * (x² - x)
func quadOutInOut(x float64) float64 {
	if x < 0.5 {
		return 2.0*x*(x-1.0) + 1
	}
	return 2.0 * x * (1.0 - x)
}

// y(x) = 1024 ^ (x - 1)
func expInIn(x float64) float64 {
	return math.Pow(1024.0, x-1.0)
}

// y(x) = 1024 ^ (-x)
func expInOut(x float64) float64 {
	return math.Pow(1024.0, -x)
}

// y(x) = 1 - 1024 ^ (-x)
func expOutIn(x float64) float64 {
	return 1.0 - math.Pow(1024.0, -x)
}

// y(x) = 1 - 1024 ^ (x - 1)
func expOutOut(x float64) float64 {
	return 1.0 - math.Pow(1024.0, x-1.0)
}

// y([0 - 0.5[) = (1024 ^ (2 * (x - 0.5))) / 2
// y([0.5 - 1]) = 1 - (1024 ^ (2 * (0.5 - x))) / 2
func expInOutIn(x float64) float64 {
	if x < 0.5 {
		return 0.5 * math.Pow(1048576.0, x-0.5)
	}
	return 1.0 - 0.5*math.Pow(1048576.0, 0.5-x)
}

// y([0 - 0.5[) = 1 - (1024 ^ (2 * (x - 0.5))) / 2
// y([0.5 - 1]) = (1024 ^ (2 * (0.5 - x))) / 2
func expInOutOut(x float64) float64 {
	if x < 0.5 {
		return 1.0 - 0.5*math.Pow(1048576.0, x-0.5)
	}
	return 0.5 * math.Pow(1048576.0, 0.5-x)
}

// y([0 - 0.5[) = 0.5 - 0.5 * 1048576 ^ (-x)
// y([0.5 - 1]) = 0.5 + 0.5 * 1048576 ^ (x-1)
func expOutInIn(x float64) float64 {
	if x < 0.5 {
		return 0.5 - 0.5*math.Pow(1048576.0, -x)
	}
	return 0.5 + 0.5*math.Pow(1048576.0, x-1.0)
}

// y([0 - 0.5[) = 0.5 + 0.5 * 1048576 ^ (-x)
// y([0.5 - 1]) = 0.5 - 0.5 * 1048576 ^ (x-1)
func expOutInOut(x float64) float64 {
	if x < 0.5 {
		return 0.5 + 0.5*math.Pow(1048576.0, -x)
	}
	return 0.5 - 0.5*math.Pow(1048576.0, x-1)
}
