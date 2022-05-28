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

func TestModn(t *testing.T) {
	for n := uint64(0); n <= 10*Small; n++ {
		for i := Small; i > 0; i-- {
			if r := Modn(n); r+1 > n {
				t.Fatal("rng.Modn: invalid return", r, "for n=", n)
			}
		}
	}
}

const permN = 100

func permTest(t *testing.T) []uint32 {
	ls := Permute(permN)
	if len(ls) != permN {
		t.Fatal("rng.Permute: invalid permutation length!")
	}

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
}

const readN = 255

func readTest(t *testing.T) []byte {

	buf := make([]byte, readN)
	n, err := Read(buf)
	if n != readN || err != nil {
		t.Fatal("rng.Read: bad return!")
	}

	if *(*uint64)(unsafe.Pointer(&buf[0])) == 0 {
		t.Fatal("rng.Read: unlikely zero first 8 bytes!")
	}

	i := len(buf) - 7
	for ; i < len(buf); i++ {
		if buf[i] != 0 {
			break
		}
	}
	if i >= len(buf) {
		t.Fatal("rng.Read: unlikely zero last 7 bytes!")
	}
	return buf
}

func TestRead(t *testing.T) {
	buf1 := readTest(t)
	buf2 := readTest(t)

	if bytes.Equal(buf1, buf2) {
		t.Fatal("rng.Read: unlikely equal buffers!")
	}
}
