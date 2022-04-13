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

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Get()
	}
}
