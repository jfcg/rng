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

func u2f(i uint64) float64 {
	return *(*float64)(unsafe.Pointer(&i))
}

// One returns a uniformly distributed random number from (0,1)
func One() float64 {
	var i uint64
	// set sign=exponent=0 to get double from [1,2)
	for i == 0 { // avoid 1
		i = Get() << 12
	}
	i |= 1<<10 - 1
	i = i>>12 ^ i<<52

	return u2f(i) - 1
}

// Two returns a uniformly distributed random number from (-1,1)
func Two() float64 {
	// set exponent=0 to get double from ±[1,2)
	i := Get()
	s := i & 1
	i <<= 11
	s <<= 11
	i |= 1<<10 - 1
	s |= 1<<10 - 1
	i = i>>12 ^ i<<52
	s = s>>12 ^ s<<52 // s = sign(i)

	return u2f(i) - u2f(s)
}

// Exp returns an exponentially distributed random number with unit mean
func Exp() float64 {
	return -math.Log(One())
}

// Normal returns two independent & normally distributed
// random numbers with zero mean and unit variance
func Normal() (float64, float64) {
	var x, y, k float64
	for !(0 < k && k < 1) {
		x = Two()
		y = Two()
		k = x*x + y*y
	}
	k = math.Sqrt(-2 * math.Log(k) / k)
	return k * x, k * y
}
