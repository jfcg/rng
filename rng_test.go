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

	"github.com/jfcg/sixb/v2"
)

func TestGet(t *testing.T) {
	ls := make([]uint64, 1<<26) // 0.5 GiB
	for i := range ls {
		ls[i] = Get()
	}
	sort.Slice(ls, func(i, k int) bool { return ls[i] < ls[k] })

	for i := len(ls) - 1; i > 0; i-- {
		if ls[i] == ls[i-1] {
			t.Error("rng.Get: unexpected collision!")
		}
	}
}

const (
	SampleN  = 256
	HalfN    = SampleN / 2
	QuarterN = SampleN / 4
	Many     = 40 * SampleN
)

func TestModn(t *testing.T) {
	ranges := []uint64{
		^uint64(Many), Many,
		1<<63 - Many, 1<<63 + Many,
		1<<64/3 + 1, 1<<64/3 + Many}

	for r := len(ranges) - 1; r > 0; r -= 2 {
		for n := ranges[r-1]; n != ranges[r]; n++ {
			for i := Many; i > 0; i-- {
				if k := Modn(n); k+1 > n {
					t.Fatal("rng.Modn: invalid return", k, "for n:", n)
				}
			}
		}
	}
}

const minPerm = 10

func permTest(t *testing.T, n int) []uint32 {
	ls := make([]uint32, n)
	Permute(ls)
	if n < minPerm {
		return ls
	}

	i := len(ls) - 1
	for ; i >= 0; i-- {
		if ls[i] != uint32(i) {
			break
		}
	}
	if i < 0 {
		t.Error("rng.Permute: unlikely identity permutation!")
	}

	i = len(ls) - 1
	for k := uint32(0); i >= 0; i-- {
		if ls[i] != k {
			break
		}
		k++
	}
	if i < 0 {
		t.Error("rng.Permute: unlikely inverse permutation!")
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
	for n := 0; n <= Many; n++ {
		ls := permTest(t, n)
		lu := permTest(t, n)

		if n >= minPerm {
			i := len(ls) - 1
			for ; i >= 0; i-- {
				if ls[i] != lu[i] {
					break
				}
			}
			if i < 0 {
				t.Error("rng.Permute: unlikely equal permutations!")
			}
		}

		permTest2(t, ls)
		permTest2(t, lu)
	}
}

func fillTest(t *testing.T, n int) []byte {
	buf := make([]byte, n)
	buf2 := make([]byte, n)
	p := global

	n, err := p.Read(buf)
	if n != len(buf) || err != nil {
		t.Fatal("rng.Prng.Read: must return len(buf), nil")
	}

	Fill(buf2)
	if !bytes.Equal(buf, buf2) {
		t.Error("rng.Fill: must fill same data on buffers!")
	}

	if n >= 16 {
		lu := sixb.Slice[uint64](buf)
		a := lu[0]
		b := lu[len(lu)-1]
		if a == 0 && b == 0 || ^a == 0 && ^b == 0 {
			t.Error("rng.Fill: unlikely all zeros/ones at start/end!")
		}
	}
	return buf
}

func TestFill(t *testing.T) {
	for n := 0; n <= Many; n++ {
		buf1 := fillTest(t, n)
		buf2 := fillTest(t, n)

		if n >= 8 && bytes.Equal(buf1, buf2) {
			t.Error("rng.Fill: unlikely equal buffers!")
		}
	}
}

func TestReset(t *testing.T) {
	var p Prng
	a, b, c := p.Get(), p.Get(), p.Get()
	p.Reset()
	d, e, f := p.Get(), p.Get(), p.Get()

	if a != d || b != e || c != f {
		t.Fatal("rng.Prng.Reset: expected:", a, b, c, "got:", d, e, f)
	}
}
