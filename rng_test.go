/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import (
	"bytes"
	"sort"
	"testing"
	"unsafe"
)

func TestGet(t *testing.T) {
	ls := make([]uint64, Large)

	for i := range ls {
		ls[i] = Get()
	}
	sort.Slice(ls, func(i, k int) bool { return ls[i] < ls[k] })

	for i := len(ls) - 1; i > 0; i-- {
		if ls[i] == ls[i-1] {
			t.Fatal("rng.Get: collision!")
		}
	}
}

func testn(t *testing.T, n uint64) {
	for i := Small; i >= 0; i-- {
		if r := Modn(n); r+1 > n {
			t.Fatal("rng.Modn: invalid return", r, "for n=", n)
		}
	}
}

func TestModn(t *testing.T) {
	for n := ^uint64(10 * Small); n != 10*Small; n++ {
		testn(t, n)
	}
	for n := uint64(1<<63 - 10*Small); n != 1<<63+10*Small; n++ {
		testn(t, n)
	}
	for n := maxu/3 + 1; n != maxu/3+10*Small; n++ {
		testn(t, n)
	}
}

const permN = 100

func permTest(t *testing.T) []uint32 {
	ls := make([]uint32, permN)
	Permute(ls)

	i := len(ls) - 1
	for ; i >= 0; i-- {
		if ls[i] != uint32(i) {
			break
		}
	}
	if i < 0 {
		t.Fatal("rng.Permute: unlikely identity permutation!")
	}

	i = len(ls) - 1
	for k := uint32(0); i >= 0; i-- {
		if ls[i] != k {
			break
		}
		k++
	}
	if i < 0 {
		t.Fatal("rng.Permute: unlikely inverse permutation!")
	}
	return ls
}

func permTest2(t *testing.T, ls []uint32) {

	sort.Slice(ls, func(i, k int) bool { return ls[i] < ls[k] })

	for i := len(ls) - 1; i >= 0; i-- {
		if ls[i] != uint32(i) {
			t.Fatal("rng.Permute: not a permutation!")
		}
	}
}

func TestPermute(t *testing.T) {
	ls := permTest(t)
	lu := permTest(t)

	i := len(ls) - 1
	for ; i >= 0; i-- {
		if ls[i] != lu[i] {
			break
		}
	}
	if i < 0 {
		t.Fatal("rng.Permute: unlikely equal permutations!")
	}

	permTest2(t, ls)
	permTest2(t, lu)

	Permute(nil) // should be no-op
	ls = []uint32{3}
	Permute(ls)
	if ls[0] != 0 {
		t.Fatal("rng.Permute: should store single zero")
	}
}

const readN = 255

func readTest(t *testing.T) []byte {

	buf := make([]byte, readN)
	Fill(buf)

	if *(*uint64)(unsafe.Pointer(&buf[0])) == 0 {
		t.Fatal("rng.Fill: unlikely zero first 8 bytes!")
	}

	i := len(buf) - 7
	for ; i < len(buf); i++ {
		if buf[i] != 0 {
			break
		}
	}
	if i >= len(buf) {
		t.Fatal("rng.Fill: unlikely zero last 7 bytes!")
	}
	return buf
}

func TestFill(t *testing.T) {
	buf1 := readTest(t)
	buf2 := readTest(t)

	if bytes.Equal(buf1, buf2) {
		t.Fatal("rng.Fill: unlikely equal buffers!")
	}
}
