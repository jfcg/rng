/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import (
	"fmt"
	"math"
	"testing"
)

func rvTest(t *testing.T, fn func() (float64, float64),
	min, median, max, imean, istd float64, name string) {

	rvPrint(imean, istd, HalfN, HalfN, "(ideal)")

	px, nx, py, ny := 0, 0, 0, 0
	mx, vx, my, vy, mxy := .0, .0, .0, .0, .0
	for i := SampleN; i > 0; i-- {

		x, y := fn()
		if !(min <= x && x < max && min <= y && y < max) {
			t.Errorf("rng.%s: unexpected output: %f, %f\n", name, x, y)
		}
		if x > median {
			px++
		} else if x < median {
			nx++
		}
		if y > median {
			py++
		} else if y < median {
			ny++
		}

		mx += x
		my += y
		vx += x * x
		vy += y * y
		mxy += x * y
	}

	if !(QuarterN < nx && QuarterN < px) {
		t.Errorf("rng.%s: unexpected distribution: %d, %d\n", name, nx, px)
	}
	if !(QuarterN < ny && QuarterN < py) {
		t.Errorf("rng.%s: unexpected distribution: %d, %d\n", name, ny, py)
	}

	mxy /= SampleN
	mx, sx := stats(mx, vx, nx, px)
	my, sy := stats(my, vy, ny, py)
	corr := (mxy - mx*my) / (sx * sy)

	if math.Abs(corr) >= 0.5 {
		t.Errorf("rng.%s: unexpected correlation: %f\n", name, corr)
	} else {
		fmt.Printf("    corr: % 6.4f\n", corr)
	}
}

func rvPrint(mean, stdev float64, n, p int, suffix string) {
	fmt.Printf("    mean: % 6.4f stdev: %6.4f lo: %3d hi: %3d %s\n", mean, stdev, n, p, suffix)
}

func stats(mean, vari float64, n, p int) (float64, float64) {
	mean /= SampleN
	std := math.Sqrt((vari - mean*mean*SampleN) / (SampleN - 1)) // from unbiased variance

	rvPrint(mean, std, n, p, "")
	return mean, std
}

func TestExp(t *testing.T) {
	rvTest(t, func() (float64, float64) {
		return Exp(), Exp()
	}, 0, math.Ln2, 25, 1, 1, "Exp")
}

func TestOne(t *testing.T) {
	rvTest(t, func() (float64, float64) {
		return One(), One()
	}, 0, 0.5, 1, 0.5, 0.288675, "One")
}

func TestTwo(t *testing.T) {
	rvTest(t, func() (float64, float64) {
		return Two(), Two()
	}, -1, 0, 1, 0, 0.57735, "Two")
}

func TestNormal(t *testing.T) {
	rvTest(t, Normal, -6.5, 0, 6.5, 0, 1, "Normal")
}
