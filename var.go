/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import (
	"math"
	"unsafe"
)

// One returns a uniformly distributed random number from (0,1)
func One() float64 {
	var i uint64
	// set sign=exponent=0 to get double from [1,2)
	for {
		i = Get() << 12
		if i != 0 { // avoid 1
			break
		}
	}
	i |= 1<<10 - 1
	i = i>>12 ^ i<<52

	f := *(*float64)(unsafe.Pointer(&i))
	return f - 1
}

// Two returns a uniformly distributed random number from (-1,1)
func Two() float64 {
	// set exponent=0 to get double from ±[1,2)
	i := Get() << 11
	i |= 1<<10 - 1
	i = i>>12 ^ i<<52

	f := *(*float64)(unsafe.Pointer(&i))
	if int64(i) < 0 {
		return f + 1
	}
	return f - 1
}

// Exp returns an exponentially distributed random number with unit mean
func Exp() float64 {
	return -math.Log(One())
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
