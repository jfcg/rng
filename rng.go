/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Package rng is a compact, fast, [sponge]-based, lockless and hard-to-predict random number
// generator. Not suitable for cryptographic applications because it has 128 bits capacity.
//
// [sponge]: https://en.wikipedia.org/wiki/Sponge_function
package rng

import (
	"os"
	"strings"
	"time"

	"github.com/jfcg/sixb/v2"
)

var state [3]uint64

const (
	// xor masks: equal amount of 1s & 0s with periods: 2, 4, 8, 16, 32, 64
	xm1 = 0x5555555555555555
	xm2 = 0x3333333333333333
	xm3 = 0x3535353535353535
	xm4 = 0x3355335533553355
	xm5 = 0x3333555533335555
	xm6 = 0x3333333355555555

	// rotation amount
	rta = 21 // 64 / 3
)

// mix(u,v) is a permutation (bijection) on state (a,b,c) consisting of
// an affine map & the nonlinear chi map. inlined.
func mix(a, b, c, u, v uint64) (x, y, z uint64) {
	b ^= u
	c ^= v
	b = b>>rta ^ b<<(64-rta) // to right
	c = c<<rta ^ c>>(64-rta) // to left

	return b ^ c&^a, c ^ a&^b, a ^ b&^c
}

// put u into rng
//
//go:norace
//go:nosplit
func put(u uint64) {
	a := state[0]
	b := state[1]
	c := state[2]
	a ^= u

	a, b, c = mix(a, b, c, xm1, xm2)
	a, b, c = mix(a, b, c, xm3, xm4)
	a, b, c = mix(a, b, c, xm5, xm6)

	state[0] = a
	state[1] = b
	state[2] = c
}

// Put u into rng
func Put[T sixb.Integer](u T) {
	put(uint64(u))
}

// Get random 64 bits from rng
//
//go:norace
//go:nosplit
func Get() uint64 {
	a := state[0]
	b := state[1]
	c := state[2]
	r := a

	a, b, c = mix(a, b, c, xm1, xm2)
	a, b, c = mix(a, b, c, xm3, xm4)
	a, b, c = mix(a, b, c, xm5, xm6)

	state[0] = a
	state[1] = b
	state[2] = c
	return r
}

// putStr inserts a pointer & a string into rng
//
//go:nosplit
func putStr(s string) {
	lu := sixb.Integers[uint32](s)
	if len(lu) <= 0 {
		return
	}
	Put(sixb.PtrToInt(&lu[0]))

	for _, u := range lu {
		Put(u)
	}
}

// putList inserts a pointer & sum of strings into rng
//
//go:nosplit
func putList(ls []string) {
	if len(ls) <= 0 {
		return
	}
	Put(sixb.PtrToInt(&ls[0]))

	var sum strings.Builder
	for _, s := range ls {
		sum.WriteString(s)
	}
	putStr(sum.String())
}

//go:nosplit
func init() {
	// insert various data into rng
	now := time.Now()
	zone, off := now.Zone()
	Put(now.UnixNano())
	Put(off)
	Put(os.Getpid())
	Put(os.Getppid())
	Put(sixb.PtrToInt(&now))

	host, _ := os.Hostname()
	wdir, _ := os.Getwd()
	putStr(zone + host + wdir)

	putList(os.Args)
	putList(os.Environ())
}
