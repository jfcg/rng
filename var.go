/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import "math"

// One returns a uniformly distributed random number from half-open interval [0, 1)
func One() float64 {
	i := int64(Get() >> 11) // [0, 2^53)
	return float64(i) / (1 << 53)
}

// Two returns a uniformly distributed random number from half-open interval [-1, 1)
func Two() float64 {
	i := int64(Get()) >> 10 // [-2^53, 2^53)
	return float64(i) / (1 << 53)
}

// Exp returns an exponentially distributed random number with unit mean
//
//go:nosplit
func Exp() float64 {
	var x float64
	for x == 0 {
		x = One()
	}
	return -math.Log(x)
}

// Normal returns two independent & normally distributed
// random numbers with zero mean and unit variance
//
//go:nosplit
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
