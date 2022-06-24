/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

const maxu uint64 = 1<<64 - 1

// Modn returns random integer from 0..n-1 for n ≥ 2, or
// returns n-1 for n < 2. This is more uniform than Get() % n.
//go:nosplit
func Modn(n uint64) uint64 {
	k := n - 1
	if n&k == 0 { // n=0 or power of 2 ?
		if n > 1 {
			return Get() & k
		}
		return k
	}

	v := Get()

	if int64(n) < 0 { // n > 2^63 ?
		for v >= n {
			v = Get()
		}
		return v
	}

	// mostly avoid one division
	if v > maxu-n {
		// largest multiple of n < 2^64
		lastn := maxu - maxu%n
		for v >= lastn {
			v = Get()
		}
	}
	return v % n
}

// Permute fills ls with a random permutation of the integers 0..len(ls)-1.
// It does not fill beyond 2^32 integers.
//go:nosplit
func Permute(ls []uint32) {
	n := uint64(len(ls))
	if n == 0 {
		return
	}
	if n > 1<<32 {
		n = 1 << 32
	}

	ls[0] = 0
	for i := uint64(1); i < n; i++ {

		k := Modn(i + 1)
		if k < i {
			ls[i] = ls[k]
		}
		ls[k] = uint32(i)
	}
}
