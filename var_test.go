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
	Large = 10_000_000
	Small = 128
)

func testfn(t *testing.T, fn func() (float64, float64),
	lo, med, hi, imean, istd float64, name string) {

	px, nx, py, ny := 0, 0, 0, 0
	for i := Large; i > 0; i-- {

		x, y := fn()
		if x != x || x <= lo || x >= hi ||
			y != y || y <= lo || y >= hi {
			t.Fatalf("rng.%s: bad output: %f, %f\n", name, x, y)
		}
		if x > med {
			px++
		} else if x < med {
			nx++
		}
		if y > med {
			py++
		} else if y < med {
			ny++
		}
	}

	mx, vx, my, vy, mxy := .0, .0, .0, .0, .0
	for i := Small; i > 0; i-- {
		x, y := fn()
		mx += x
		my += y
		vx += x * x
		vy += y * y
		mxy += x * y
	}
	t.Logf("mean: %+6.4f stdev: %6.4f (ideal)\n", imean, istd)
	mx, sx := stats(t, mx, vx, px, nx)
	my, sy := stats(t, my, vy, py, ny)
	mxy /= Small
	t.Logf("corr: %+6.4f\n", (mxy-mx*my)/(sx*sy))
}

func stats(t *testing.T, mean, vari float64, p, n int) (float64, float64) {
	mean /= Small
	std := math.Sqrt((vari - mean*mean*Small) / (Small - 1)) // from unbiased variance

	t.Logf("mean: %+6.4f stdev: %6.4f hi: %d lo: %d\n", mean, std, p, n)
	return mean, std
}

func TestExp(t *testing.T) {
	testfn(t, func() (float64, float64) {
		return Exp(), Exp()
	}, 0, math.Ln2, 25, 1, 1, "Exp")
}

func TestOne(t *testing.T) {
	testfn(t, func() (float64, float64) {
		return One(), One()
	}, 0, 0.5, 1, 0.5, 0.288675, "One")
}

func TestTwo(t *testing.T) {
	testfn(t, func() (float64, float64) {
		return Two(), Two()
	}, -1, 0, 1, 0, 0.57735, "Two")
}

func TestNormal(t *testing.T) {
	testfn(t, Normal, -6.5, 0, 6.5, 0, 1, "Normal")
}
