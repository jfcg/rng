package rng

import (
	"math"
	"unsafe"
)

// Two returns a uniformly distributed random number from (-1,1)
func Two() float64 {
	// set exponent=0 to get double from +/-[1,2)
	i := Get() << 11
	i |= 1<<10 - 1
	i = i>>12 ^ i<<52

	f := *(*float64)(unsafe.Pointer(&i))
	if int64(i) < 0 {
		return f + 1
	}
	return f - 1
}

// Normal returns two independent & normally distributed
// random numbers with zero mean and unit variance
func Normal() (float64, float64) {
	var x, y, k float64
	for {
		x = Two()
		y = Two()
		k = x*x + y*y
		if 0 < k && k < 1 {
			break
		}
	}
	k = math.Sqrt(-2 * math.Log(k) / k)
	return k * x, k * y
}
