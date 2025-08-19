// Copyright (c) 2022-present, Serhat Şevki Dinçer.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package rng

import (
	"math"
	"os"
	"sync/atomic"
	"time"
	"unsafe"

	sb "github.com/jfcg/sixb/v2"
)

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

// mix is a permutation (bijection) on state (a,b,c) made of:
// - affine map(u,v)
// - nonlinear chi map
// - rotation: abc -> bca
func mix(a, b, c, u, v uint64) (x, y, z uint64) { // inlined
	b ^= u
	c ^= v
	b = b>>rta ^ b<<(64-rta)
	c = c<<rta ^ c>>(64-rta) // affine map

	return b ^ c&^a, c ^ a&^b, a ^ b&^c // chi & rotation
}

// Prng is a compact (24 bytes), fast pseudo (deterministic) random number generator.
// Instances with same initialization vectors will produce same results for same call
// sequences, such as:
//
//	var p Prng // Initialize with an arbitrary number of calls to p.Put().
//	p.Put(c1)  // Each distinct initialization vector c
//	p.Put(c2)  // will yield a distinct Prng instance.
//
//	n := p.Get() // generate arbitrary numbers
//	x := p.Two() // etc.
type Prng struct {
	a, b, c uint64
}

// Put u into p.
//
//go:nosplit
func (p *Prng) Put(u uint64) {
	a, b, c := p.a, p.b, p.c
	a ^= u

	a, b, c = mix(a, b, c, xm1, xm2)
	a, b, c = mix(a, b, c, xm3, xm4)
	a, b, c = mix(a, b, c, xm5, xm6)

	p.a, p.b, p.c = a, b, c
}

// Get 64 bits from p. Note that p.Get() % n is not uniform, use [Prng.Modn](n) instead.
//
//go:nosplit
func (p *Prng) Get() uint64 {
	a, b, c := p.a, p.b, p.c
	r := a

	a, b, c = mix(a, b, c, xm1, xm2)
	a, b, c = mix(a, b, c, xm3, xm4)
	a, b, c = mix(a, b, c, xm5, xm6)

	p.a, p.b, p.c = a, b, c
	return r
}

// One returns a uniformly distributed number in interval [0, 1) from p.
func (p *Prng) One() float64 { // inlined

	i := int64(p.Get() >> 11) // [0, 2^53)
	return float64(i) / (1 << 53)
}

// OneR returns a uniformly distributed number in interval (0, 1] from p.
func (p *Prng) OneR() float64 { // inlined

	i := int64(p.Get()>>11) + 1 // (0, 2^53]
	return float64(i) / (1 << 53)
}

// Two returns a uniformly distributed number in interval [-1, 1) from p.
func (p *Prng) Two() float64 { // inlined

	i := int64(p.Get()) >> 10 // [-2^53, 2^53)
	return float64(i) / (1 << 53)
}

// TwoR returns a uniformly distributed number in interval (-1, 1] from p.
func (p *Prng) TwoR() float64 { // inlined

	i := int64(p.Get())>>10 + 1 // (-2^53, 2^53]
	return float64(i) / (1 << 53)
}

// Exp returns an exponentially distributed number (mean=1) from p.
func (p *Prng) Exp() float64 { // inlined
	return -math.Log(p.OneR())
}

// Normal returns two independent & normally distributed
// numbers (mean=0, variance=1) from p.
//
//go:nosplit
func (p *Prng) Normal() (float64, float64) {
	var x, y, k float64
	for !(0 < k && k <= 1) {
		x = p.Two()
		y = p.Two()
		k = x*x + y*y
	}
	k = math.Sqrt(-2 * math.Log(k) / k)
	return k * x, k * y
}

// Modn returns a uniformly selected integer in 0,..,n-1
// from p for n ≥ 2, and returns n-1 for n < 2.
//
//go:nosplit
func (p *Prng) Modn(n uint64) uint64 {
	if k := n - 1; n&k == 0 { // n=0 or power of 2 ?
		if n > 1 {
			return p.Get() & k
		}
		return k
	}

	v := p.Get()

	if int64(n) < 0 { // n > 2^63 ?
		for v >= n {
			v = p.Get()
		}
		return v
	}

	// mostly avoid one division
	if v > ^n {
		// largest multiple of n < 2^64
		lastn := ^((1<<64 - 1) % n)
		for v >= lastn {
			v = p.Get()
		}
	}
	return v % n
}

// Permute fills lu with a uniformly selected permutation of
// integers 0,..,len(lu)-1 from p, up to 2^32 entries.
//
//go:nosplit
func (p *Prng) Permute(lu []uint32) {
	if len(lu) <= 0 {
		return
	}
	n := uint64(len(lu))
	if n > 1<<32 {
		n = 1 << 32
	}

	lu[0] = 0
	for i := uint64(1); i < n; i++ {

		k := p.Modn(i + 1)
		if k < i {
			lu[i] = lu[k]
		}
		lu[k] = uint32(i)
	}
}

// Fill buf with bytes from p.
//
//go:nosplit
func (p *Prng) Fill(buf []byte) {

	lu := sb.Slice[uint64](buf)
	for i := range lu { // fill 8 bytes at a time
		lu[i] = p.Get()
	}

	if r := len(lu) << 3; r < len(buf) { // fill 1 to 7 remaining bytes
		var u [8]byte
		*(*uint64)(unsafe.Pointer(&u)) = p.Get()
		copy(buf[r:], u[:])
	}
}

// Read wraps [Prng.Fill] and always returns len(buf), nil
//
//go:nosplit
func (p *Prng) Read(buf []byte) (n int, err error) {
	p.Fill(buf)
	return len(buf), nil
}

// Reset p.
//
//go:nosplit
func (p *Prng) Reset() {
	p.a, p.b, p.c = 0, 0, 0
}

var randCount uint64

// Randomize p. After this call, results of p's methods
// will not be reproducible until [Prng.Reset]() is called.
//
//go:nosplit
func (p *Prng) Randomize() {
	// insert various data into p to differentiate Prng instances:

	// in the same process
	now := time.Now()
	put(p, now.UnixNano())                  // current time
	put(p, sb.PtrToInt(p))                  // instance address
	put(p, atomic.AddUint64(&randCount, 1)) // unique counter

	// on the same computer
	put(p, os.Getpid())       // process id
	put(p, os.Getppid())      // parent process id
	put(p, sb.PtrToInt(&now)) // address in stack
	putList(p, os.Args)       // cmd line args

	// on different computers
	zone, off := now.Zone()
	host, _ := os.Hostname()
	wdir, _ := os.Getwd()
	put(p, off)
	putStr(p, zone+host+wdir)
	putList(p, os.Environ())
}
