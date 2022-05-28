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

func BenchmarkModn(b *testing.B) {
	var sum uint64
	n := uint64(b.N)
	for i := n; i > 0; i-- {
		sum += rng.Modn(n)
	}
}

func BenchmarkStdModn(b *testing.B) {
	var sum int64
	n := int64(b.N)
	mr := rand.New(rand.NewSource(n))
	b.ResetTimer()
	for i := n; i > 0; i-- {
		sum += mr.Int63n(n)
	}
}

func BenchmarkAltModn(b *testing.B) {
	var sum uint64
	n := uint64(b.N)
	ar := altr.New(altr.NewSource(n))
	b.ResetTimer()
	for i := n; i > 0; i-- {
		sum += ar.Uint64n(n)
	}
}

const permN = 1000

func BenchmarkPerm(b *testing.B) {
	var sum uint32
	for i := b.N; i > 0; i-- {
		sum += rng.Permute(permN)[0]
	}
}

func BenchmarkStdPerm(b *testing.B) {
	var sum int
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += mr.Perm(permN)[0]
	}
}

func BenchmarkAltPerm(b *testing.B) {
	var sum int
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		sum += ar.Perm(permN)[0]
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

const readN = 255

func BenchmarkRead(b *testing.B) {
	buf := make([]byte, readN)
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_, _ = rng.Read(buf)
	}
}

func BenchmarkStdRead(b *testing.B) {
	buf := make([]byte, readN)
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_, _ = mr.Read(buf)
	}
}

func BenchmarkAltRead(b *testing.B) {
	buf := make([]byte, readN)
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_, _ = ar.Read(buf)
	}
}

func BenchmarkTwo(b *testing.B) {
	var sum float64
	for i := b.N; i > 0; i-- {
		sum += rng.Two()
	}
}
