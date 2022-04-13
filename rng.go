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

	"github.com/jfcg/sixb"
)

var state [3]uint64

//go:norace
func round(a, b, c uint64) {
	b++
	b = b>>21 ^ b<<43
	c--
	c = c<<21 ^ c>>43

	state[0] = b ^ c&^a
	state[1] = c ^ a&^b
	state[2] = a ^ b&^c
}

// Put x into rng
//go:norace
func Put(x uint64) {
	a := state[0]
	b := state[1]
	c := state[2]
	a ^= x
	round(a, b, c)
}

// Get from rng
//go:norace
func Get() uint64 {
	a := state[0]
	b := state[1]
	c := state[2]
	round(a, b, c)
	return a
}

func putS(s string) {
	lu := sixb.StoU8(s)
	for _, u := range lu {
		Put(u)
	}
}

func putLs(ls []string) {
	for _, s := range ls {
		putS(s)
	}
}

func init() {
	// insert various data into rng
	Put(uint64(os.Getpid()))
	Put(uint64(os.Getppid()))

	str, err := os.Getwd()
	if err == nil {
		putS(str)
	}
	str, err = os.Hostname()
	if err == nil {
		putS(str)
	}
	putLs(os.Args)
	putLs(os.Environ())
}
