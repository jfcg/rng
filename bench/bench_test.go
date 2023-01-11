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
	for i := b.N; i > 0; i-- {
		_ = rng.Get()
	}
}

func BenchmarkStdGet(b *testing.B) {
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = mr.Uint64()
	}
}

func BenchmarkAltGet(b *testing.B) {
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = ar.Uint64()
	}
}

func BenchmarkModn(b *testing.B) {
	for i := uint64(b.N); i > 0; i-- {
		_ = rng.Modn(i)
	}
}

func BenchmarkStdModn(b *testing.B) {
	n := int64(b.N)
	mr := rand.New(rand.NewSource(n))
	b.ResetTimer()
	for i := n; i > 0; i-- {
		_ = mr.Int63n(i)
	}
}

func BenchmarkAltModn(b *testing.B) {
	n := uint64(b.N)
	ar := altr.New(altr.NewSource(n))
	b.ResetTimer()
	for i := n; i > 0; i-- {
		_ = ar.Uint64n(i)
	}
}

const permN = 200

func BenchmarkPerm(b *testing.B) {
	ls := make([]uint32, permN)
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		rng.Permute(ls)
	}
}

func BenchmarkStdPerm(b *testing.B) {
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = mr.Perm(permN)
	}
}

func BenchmarkAltPerm(b *testing.B) {
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = ar.Perm(permN)
	}
}

func BenchmarkExp(b *testing.B) {
	for i := b.N; i > 0; i-- {
		_ = rng.Exp()
	}
}

func BenchmarkStdExp(b *testing.B) {
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = mr.ExpFloat64()
	}
}

func BenchmarkAltExp(b *testing.B) {
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = ar.ExpFloat64()
	}
}

func BenchmarkOne(b *testing.B) {
	for i := b.N; i > 0; i-- {
		_ = rng.One()
	}
}

func BenchmarkStdOne(b *testing.B) {
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = mr.Float64()
	}
}

func BenchmarkAltOne(b *testing.B) {
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = ar.Float64()
	}
}

func BenchmarkTwo(b *testing.B) {
	for i := b.N; i > 0; i-- {
		_ = rng.Two()
	}
}

func BenchmarkNormal(b *testing.B) {
	for i := b.N; i > 0; i-- {
		_, _ = rng.Normal()
	}
}

func BenchmarkStdNormal(b *testing.B) {
	mr := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = mr.NormFloat64()
	}
}

func BenchmarkAltNormal(b *testing.B) {
	ar := altr.New(altr.NewSource(uint64(b.N)))
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		_ = ar.NormFloat64()
	}
}

const readN = 255

func BenchmarkFill(b *testing.B) {
	buf := make([]byte, readN)
	b.ResetTimer()
	for i := b.N; i > 0; i-- {
		rng.Fill(buf)
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
