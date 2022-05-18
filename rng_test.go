/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import (
	"sort"
	"testing"
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
