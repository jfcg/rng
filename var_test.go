/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import (
	"math"
	"testing"
)

const (
	Large = 1e8
	Small = 128
)

func testfn(t *testing.T, fn func() float64, lo, med, hi float64, name string) {
	pos, neg := 0, 0
	for i := Large; i > 0; i-- {

		x := fn()
		if x != x || x <= lo || x >= hi {
			t.Fatalf("rng.%s: bad output: %f\n", name, x)
		}
		if x > med {
			pos++
		} else if x < med {
			neg++
		}
	}

	mean, vari := .0, .0
	for i := Small; i > 0; i-- {
		x := fn()
		mean += x
		vari += x * x
	}
	print(t, pos, neg, mean, vari)
}

func TestExp(t *testing.T) {
	testfn(t, Exp, 0, math.Ln2, 9e9, "Exp")
}

func TestOne(t *testing.T) {
	testfn(t, One, 0, 0.5, 1, "One")
}

func TestTwo(t *testing.T) {
	testfn(t, Two, -1, 0, 1, "Two")
}

func print(t *testing.T, p, n int, mean, vari float64) {
	mean /= Small
	vari = (vari - Small*(mean*mean)) / (Small - 1)
	t.Log("hi:", p, "lo:", n)
	t.Logf("mean: %+6.4f variance: %6.4f\n", mean, vari)
}

func TestNormal(t *testing.T) {
	pos, neg := 0, 0
	for i := Large / 2; i > 0; i-- {

		x, y := Normal()
		if x != x || x > 6.3 || x < -6.3 ||
			y != y || y > 6.3 || y < -6.3 {
			t.Fatal("rng.Normal: unusual outputs:", x, y)
		}
		if x > 0 {
			pos++
		} else if x < 0 {
			neg++
		}
		if y > 0 {
			pos++
		} else if y < 0 {
			neg++
		}
	}

	mean, vari := .0, .0
	for i := Small / 2; i > 0; i-- {
		x, y := Normal()
		mean += x
		mean += y
		vari += x * x
		vari += y * y
	}
	print(t, pos, neg, mean, vari)
}
