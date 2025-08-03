// Copyright (c) 2022-present, Serhat Şevki Dinçer.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bench

import (
	"math/rand"
	rand2 "math/rand/v2"
	"testing"

	"github.com/jfcg/rng"
)

func BenchmarkGet(b *testing.B) {
	var p rng.Prng
	for range b.N {
		_ = p.Get()
	}
}

func BenchmarkStdGet(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for range b.N {
		_ = r.Uint64()
	}
}

func BenchmarkPcgGet(b *testing.B) {
	r := rand2.NewPCG(1, 2)
	b.ResetTimer()
	for range b.N {
		_ = r.Uint64()
	}
}

var chaSeed = [32]byte{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func BenchmarkChaGet(b *testing.B) {
	r := rand2.NewChaCha8(chaSeed)
	b.ResetTimer()
	for range b.N {
		_ = r.Uint64()
	}
}

func BenchmarkModn(b *testing.B) {
	var p rng.Prng
	for i := range b.N {
		_ = p.Modn(uint64(i + 1))
	}
}

func BenchmarkStdModn(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for i := range b.N {
		_ = r.Int63n(int64(i + 1))
	}
}

func BenchmarkPcgModn(b *testing.B) {
	r := rand2.New(rand2.NewPCG(1, 2))
	b.ResetTimer()
	for i := range b.N {
		_ = r.Int64N(int64(i + 1))
	}
}

func BenchmarkChaModn(b *testing.B) {
	r := rand2.New(rand2.NewChaCha8(chaSeed))
	b.ResetTimer()
	for i := range b.N {
		_ = r.Int64N(int64(i + 1))
	}
}

const permN = 100

func BenchmarkPerm(b *testing.B) {
	var p rng.Prng
	lu := make([]uint32, permN)
	b.ResetTimer()
	for range b.N {
		p.Permute(lu)
	}
}

func BenchmarkStdPerm(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for range b.N {
		_ = r.Perm(permN)
	}
}

func BenchmarkPcgPerm(b *testing.B) {
	r := rand2.New(rand2.NewPCG(1, 2))
	b.ResetTimer()
	for range b.N {
		_ = r.Perm(permN)
	}
}

func BenchmarkChaPerm(b *testing.B) {
	r := rand2.New(rand2.NewChaCha8(chaSeed))
	b.ResetTimer()
	for range b.N {
		_ = r.Perm(permN)
	}
}

func BenchmarkExp(b *testing.B) {
	var p rng.Prng
	for range b.N {
		_ = p.Exp()
	}
}

func BenchmarkStdExp(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for range b.N {
		_ = r.ExpFloat64()
	}
}

func BenchmarkPcgExp(b *testing.B) {
	r := rand2.New(rand2.NewPCG(1, 2))
	b.ResetTimer()
	for range b.N {
		_ = r.ExpFloat64()
	}
}
func BenchmarkChaExp(b *testing.B) {
	r := rand2.New(rand2.NewChaCha8(chaSeed))
	b.ResetTimer()
	for range b.N {
		_ = r.ExpFloat64()
	}
}

func BenchmarkOne(b *testing.B) {
	var p rng.Prng
	for range b.N {
		_ = p.One()
	}
}

func BenchmarkStdOne(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for range b.N {
		_ = r.Float64()
	}
}

func BenchmarkPcgOne(b *testing.B) {
	r := rand2.New(rand2.NewPCG(1, 2))
	b.ResetTimer()
	for range b.N {
		_ = r.Float64()
	}
}
func BenchmarkChaOne(b *testing.B) {
	r := rand2.New(rand2.NewChaCha8(chaSeed))
	b.ResetTimer()
	for range b.N {
		_ = r.Float64()
	}
}

func BenchmarkTwo(b *testing.B) {
	var p rng.Prng
	for range b.N {
		_ = p.Two()
	}
}

func BenchmarkNormal(b *testing.B) {
	var p rng.Prng
	for range b.N {
		_, _ = p.Normal()
	}
}

func BenchmarkStdNormal(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for range b.N {
		_ = r.NormFloat64()
	}
}

func BenchmarkPcgNormal(b *testing.B) {
	r := rand2.New(rand2.NewPCG(1, 2))
	b.ResetTimer()
	for range b.N {
		_ = r.NormFloat64()
	}
}
func BenchmarkChaNormal(b *testing.B) {
	r := rand2.New(rand2.NewChaCha8(chaSeed))
	b.ResetTimer()
	for range b.N {
		_ = r.NormFloat64()
	}
}

const readN = 500

func BenchmarkRead(b *testing.B) {
	var p rng.Prng
	buf := make([]byte, readN)
	b.ResetTimer()
	for range b.N {
		p.Fill(buf)
	}
}

func BenchmarkStdRead(b *testing.B) {
	buf := make([]byte, readN)
	r := rand.New(rand.NewSource(int64(b.N)))
	b.ResetTimer()
	for range b.N {
		_, _ = r.Read(buf)
	}
}

func BenchmarkChaRead(b *testing.B) {
	buf := make([]byte, readN)
	r := rand2.NewChaCha8(chaSeed)
	b.ResetTimer()
	for range b.N {
		_, _ = r.Read(buf)
	}
}
