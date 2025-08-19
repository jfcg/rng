// Copyright (c) 2022-present, Serhat Şevki Dinçer.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bench

import (
	"math/rand"
	rand2 "math/rand/v2"
	"testing"
	"unsafe"

	"github.com/jfcg/rng"
)

func BenchmarkGet(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		_ = p.Get()
	}
}

func BenchmarkStdGet(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	for range b.N {
		_ = r.Uint64()
	}
}

func BenchmarkPcgGet(b *testing.B) {
	r := rand2.NewPCG(uint64(b.N), uint64(b.N+1))
	for range b.N {
		_ = r.Uint64()
	}
}

var chaSeed [4]uint64

func initCha(n int) *rand2.ChaCha8 {
	chaSeed[0] = uint64(n)
	chaSeed[1] = uint64(n + 1)
	chaSeed[2] = uint64(n + 2)
	chaSeed[3] = uint64(n + 3)
	return rand2.NewChaCha8(*(*[32]byte)(unsafe.Pointer(&chaSeed)))
}

func BenchmarkChaGet(b *testing.B) {
	r := initCha(b.N)
	for range b.N {
		_ = r.Uint64()
	}
}

func BenchmarkModn(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for i := range b.N {
		_ = p.Modn(uint64(i + 1))
	}
}

func BenchmarkStdModn(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	for i := range b.N {
		_ = r.Int63n(int64(i + 1))
	}
}

func BenchmarkPcgModn(b *testing.B) {
	r := rand2.New(rand2.NewPCG(uint64(b.N), uint64(b.N+1)))
	for i := range b.N {
		_ = r.Int64N(int64(i + 1))
	}
}

func BenchmarkChaModn(b *testing.B) {
	r := rand2.New(initCha(b.N))
	for i := range b.N {
		_ = r.Int64N(int64(i + 1))
	}
}

const permN = 100

var permLs [permN]uint32

func BenchmarkPerm(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		p.Permute(permLs[:])
	}
}

func BenchmarkStdPerm(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	for range b.N {
		_ = r.Perm(permN)
	}
}

func BenchmarkPcgPerm(b *testing.B) {
	r := rand2.New(rand2.NewPCG(uint64(b.N), uint64(b.N+1)))
	for range b.N {
		_ = r.Perm(permN)
	}
}

func BenchmarkChaPerm(b *testing.B) {
	r := rand2.New(initCha(b.N))
	for range b.N {
		_ = r.Perm(permN)
	}
}

func BenchmarkExp(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		_ = p.Exp()
	}
}

func BenchmarkStdExp(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	for range b.N {
		_ = r.ExpFloat64()
	}
}

func BenchmarkPcgExp(b *testing.B) {
	r := rand2.New(rand2.NewPCG(uint64(b.N), uint64(b.N+1)))
	for range b.N {
		_ = r.ExpFloat64()
	}
}
func BenchmarkChaExp(b *testing.B) {
	r := rand2.New(initCha(b.N))
	for range b.N {
		_ = r.ExpFloat64()
	}
}

func BenchmarkOne(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		_ = p.One()
	}
}

func BenchmarkStdOne(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	for range b.N {
		_ = r.Float64()
	}
}

func BenchmarkPcgOne(b *testing.B) {
	r := rand2.New(rand2.NewPCG(uint64(b.N), uint64(b.N+1)))
	for range b.N {
		_ = r.Float64()
	}
}
func BenchmarkChaOne(b *testing.B) {
	r := rand2.New(initCha(b.N))
	for range b.N {
		_ = r.Float64()
	}
}

func BenchmarkTwo(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		_ = p.Two()
	}
}

func BenchmarkTri2(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		_ = p.Tri2()
	}
}

func BenchmarkNormal(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		_, _ = p.Normal() // two samples per call
	}
}

func BenchmarkStdNormal(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	for range b.N {
		_ = r.NormFloat64()
	}
}

func BenchmarkPcgNormal(b *testing.B) {
	r := rand2.New(rand2.NewPCG(uint64(b.N), uint64(b.N+1)))
	for range b.N {
		_ = r.NormFloat64()
	}
}
func BenchmarkChaNormal(b *testing.B) {
	r := rand2.New(initCha(b.N))
	for range b.N {
		_ = r.NormFloat64()
	}
}

const readN = 1007

var readLs [readN]byte

func BenchmarkRead(b *testing.B) {
	var p rng.Prng
	p.Put(uint64(b.N))
	p.Put(uint64(b.N + 1))
	for range b.N {
		p.Fill(readLs[:])
	}
}

func BenchmarkStdRead(b *testing.B) {
	r := rand.New(rand.NewSource(int64(b.N)))
	for range b.N {
		_, _ = r.Read(readLs[:])
	}
}

func BenchmarkChaRead(b *testing.B) {
	r := initCha(b.N)
	for range b.N {
		_, _ = r.Read(readLs[:])
	}
}
