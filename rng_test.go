/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import (
	"math/rand"
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
	var sum uint64
	for i := 0; i < b.N; i++ {
		sum += Get()
	}
}

func BenchmarkStdGet(b *testing.B) {
	var sum uint64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += mr.Uint64()
	}
}

func BenchmarkExp(b *testing.B) {
	var sum float64
	for i := 0; i < b.N; i++ {
		sum += Exp()
	}
}

func BenchmarkStdExp(b *testing.B) {
	var sum float64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += mr.ExpFloat64()
	}
}

func BenchmarkOne(b *testing.B) {
	var sum float64
	for i := 0; i < b.N; i++ {
		sum += One()
	}
}

func BenchmarkStdOne(b *testing.B) {
	var sum float64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += mr.Float64()
	}
}

func BenchmarkNormal(b *testing.B) {
	var sum1, sum2 float64
	for i := 0; i < b.N; i++ {
		x, y := Normal()
		sum1 += x
		sum2 += y
	}
}

func BenchmarkStdNormal(b *testing.B) {
	var sum float64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += mr.NormFloat64()
	}
}

func BenchmarkTwo(b *testing.B) {
	var sum float64
	for i := 0; i < b.N; i++ {
		sum += Two()
	}
}
