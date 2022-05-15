/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

// Package rng implements a cheap, fast and hard-to-predict random number generator
// without locks as a feature. Not suitable for cryptographic applications.
package rng

import (
	"os"
	"time"
	"unsafe"

	"github.com/jfcg/sixb"
)

var state [3]uint64

const (
	// xor masks: equal amount of 1s & 0s with periods: 2, 4, 8, 16
	xm1 = 0x5555555555555555
	xm2 = 0x3333333333333333
	xm3 = 0x3535353535353535
	xm4 = 0x3355335533553355

	// rotatiom amount
	rta = 21 // 64 / 3
)

// mix(u,v) is a permutation of state (a,b,c) consisting of
// an affine map & the nonlinear chi map. inlined.
func mix(a, b, c, u, v uint64) (x, y, z uint64) {
	b ^= u
	c ^= v
	b = b>>rta ^ b<<(64-rta) // to right
	c = c<<rta ^ c>>(64-rta) // to left

	return b ^ c&^a, c ^ a&^b, a ^ b&^c
}

// Put u into rng
//go:norace
//go:nosplit
func Put(u uint64) {
	a := state[0]
	b := state[1]
	c := state[2]
	a ^= u

	a, b, c = mix(a, b, c, xm1, xm2)
	a, b, c = mix(a, b, c, xm3, xm4)

	state[0] = a
	state[1] = b
	state[2] = c
}

// Get random 64 bits from rng
//go:norace
//go:nosplit
func Get() uint64 {
	a := state[0]
	b := state[1]
	c := state[2]
	r := a

	a, b, c = mix(a, b, c, xm1, xm2)
	a, b, c = mix(a, b, c, xm3, xm4)

	state[0] = a
	state[1] = b
	state[2] = c
	return r
}

// putS inserts a pointer & a string into rng
func putS(s string) {
	lu := sixb.StoU4(s)
	if len(lu) <= 0 {
		return
	}
	Put(uint64(uintptr(unsafe.Pointer(&lu[0]))))
	for _, u := range lu {
		Put(uint64(u))
	}
}

// putLs inserts a pointer & sum of strings into rng
func putLs(ls []string) {
	if len(ls) <= 0 {
		return
	}
	Put(uint64(uintptr(unsafe.Pointer(&ls[0]))))
	sum := ""
	for _, s := range ls {
		sum += s
	}
	putS(sum)
}

func init() {
	// insert various data into rng
	now := time.Now()
	zone, off := now.Zone()
	Put(uint64(now.UnixNano()))
	Put(uint64(off))
	Put(uint64(os.Getpid()))
	Put(uint64(os.Getppid()))
	Put(uint64(uintptr(unsafe.Pointer(&now))))

	host, _ := os.Hostname()
	wdir, _ := os.Getwd()
	putS(zone + host + wdir)

	putLs(os.Args)
	putLs(os.Environ())
}
