// Copyright (c) 2022-present, Serhat Şevki Dinçer.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package rng provides a compact, fast, [sponge]-based, lockless and hard-to-predict
// global random number generator. It also provides [Prng] for deterministic behavior.
// Neither is suitable for cryptographic applications because of 128 bits capacity.
//
// [sponge]: https://en.wikipedia.org/wiki/Sponge_function
package rng

import (
	"unsafe"

	sb "github.com/jfcg/sixb/v2"
)

var global Prng

//go:nosplit
//go:norace
func init() {
	global.Randomize()
}

// Get 64 bits from global rng. Note that Get() % n is not uniform, use [Modn](n) instead.
//
//go:norace
func Get() uint64 {
	return global.Get()
}

// One returns a uniformly distributed number in interval [0, 1) from global rng.
//
//go:norace
func One() float64 {
	return global.One()
}

// OneR returns a uniformly distributed number in interval (0, 1] from global rng.
//
//go:norace
func OneR() float64 {
	return global.OneR()
}

// Two returns a uniformly distributed number in interval [-1, 1) from global rng.
//
//go:norace
func Two() float64 {
	return global.Two()
}

// TwoR returns a uniformly distributed number in interval (-1, 1] from global rng.
//
//go:norace
func TwoR() float64 {
	return global.TwoR()
}

// Tri1 returns a number from symmetric triangular distribution in interval (0, 1) from global rng.
//
//go:norace
func Tri1() float64 {
	return global.Tri1()
}

// Tri2 returns a number from symmetric triangular distribution in interval (-1, 1) from global rng.
//
//go:norace
func Tri2() float64 {
	return global.Tri2()
}

// Exp returns an exponentially distributed number (mean=1) from global rng.
//
//go:norace
func Exp() float64 {
	return global.Exp()
}

// Normal returns two independent & normally distributed
// numbers (mean=0, variance=1) from global rng.
//
//go:norace
func Normal() (float64, float64) {
	return global.Normal()
}

// Modn returns a uniformly selected integer in 0,..,n-1
// from global rng for n ≥ 2, and returns n-1 for n < 2.
//
//go:norace
func Modn(n uint64) uint64 {
	return global.Modn(n)
}

// Permute fills lu with a uniformly selected permutation of
// integers 0,..,len(lu)-1 from global rng, up to 2^32 entries.
//
//go:norace
func Permute(lu []uint32) {
	global.Permute(lu)
}

// Fill buf with bytes from global rng.
//
//go:norace
func Fill(buf []byte) {
	global.Fill(buf)
}

// put u into p.
func put[T sb.Integer](p *Prng, u T) {
	p.Put(uint64(u))
}

// putStr inserts a pointer and a string into p.
//
//go:nosplit
func putStr(p *Prng, s string) {
	lb := sb.Bytes(s)
	if len(lb) <= 0 {
		return
	}
	put(p, sb.PtrToInt(&lb[0]))

	lu := sb.Slice[uint64](lb)
	for _, u := range lu { // put 8 bytes at a time
		p.Put(u)
	}

	if r := len(lu) << 3; r < len(lb) { // put 1 to 7 remaining bytes
		var u [8]byte
		copy(u[:], lb[r:])
		p.Put(*(*uint64)(unsafe.Pointer(&u)))
	}
}

// putList inserts a pointer and sum of strings into p.
//
//go:nosplit
func putList(p *Prng, ls []string) {
	if len(ls) <= 0 {
		return
	}
	put(p, sb.PtrToInt(&ls[0]))

	sum := sb.Bytes(ls[0])
	for i := 1; i < len(ls); i++ {
		sum = append(sum, ls[i]...)
	}
	putStr(p, sb.String(sum))
}
