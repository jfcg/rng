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

// Put u into rng
//go:norace
func Put(u uint64) {
	a := state[0]
	b := state[1]
	c := state[2]
	a ^= u

	b ^= 0x5555555555555555
	c ^= 0x3333333333333333
	b = b>>21 ^ b<<43
	c = c<<21 ^ c>>43

	x := b ^ c&^a
	y := c ^ a&^b
	z := a ^ b&^c

	y ^= 0x6666666666666666
	z ^= 0xaaaaaaaaaaaaaaaa
	y = y>>21 ^ y<<43
	z = z<<21 ^ z>>43

	state[0] = y ^ z&^x
	state[1] = z ^ x&^y
	state[2] = x ^ y&^z
}

// Get from rng
//go:norace
func Get() uint64 {
	a := state[0]
	b := state[1]
	c := state[2]

	b ^= 0x5555555555555555
	c ^= 0x3333333333333333
	b = b>>21 ^ b<<43
	c = c<<21 ^ c>>43

	x := b ^ c&^a
	y := c ^ a&^b
	z := a ^ b&^c

	y ^= 0x6666666666666666
	z ^= 0xaaaaaaaaaaaaaaaa
	y = y>>21 ^ y<<43
	z = z<<21 ^ z>>43

	state[0] = y ^ z&^x
	state[1] = z ^ x&^y
	state[2] = x ^ y&^z
	return a
}

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
