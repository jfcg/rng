/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package bench

import (
	"math/rand"
	"testing"

	"github.com/jfcg/rng"
	altr "golang.org/x/exp/rand"
)

func BenchmarkGet(b *testing.B) {
	var sum uint64
	for i := b.N; i > 0; i-- {
		sum += rng.Get()
	}
}

func BenchmarkStdGet(b *testing.B) {
	var sum uint64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += mr.Uint64()
	}
}

func BenchmarkAltGet(b *testing.B) {
	var sum uint64
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += ar.Uint64()
	}
}

func BenchmarkExp(b *testing.B) {
	var sum float64
	for i := b.N; i > 0; i-- {
		sum += rng.Exp()
	}
}

func BenchmarkStdExp(b *testing.B) {
	var sum float64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += mr.ExpFloat64()
	}
}

func BenchmarkAltExp(b *testing.B) {
	var sum float64
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += ar.ExpFloat64()
	}
}

func BenchmarkOne(b *testing.B) {
	var sum float64
	for i := b.N; i > 0; i-- {
		sum += rng.One()
	}
}

func BenchmarkStdOne(b *testing.B) {
	var sum float64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += mr.Float64()
	}
}

func BenchmarkAltOne(b *testing.B) {
	var sum float64
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += ar.Float64()
	}
}

func BenchmarkNormal(b *testing.B) {
	var sum1, sum2 float64
	for i := b.N; i > 0; i-- {
		x, y := rng.Normal()
		sum1 += x
		sum2 += y
	}
}

func BenchmarkStdNormal(b *testing.B) {
	var sum float64
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += mr.NormFloat64()
	}
}

func BenchmarkAltNormal(b *testing.B) {
	var sum float64
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += ar.NormFloat64()
	}
}

func BenchmarkTwo(b *testing.B) {
	var sum float64
	for i := b.N; i > 0; i-- {
		sum += rng.Two()
	}
}
