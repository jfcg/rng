package rng

import "testing"

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

func TestTwo(t *testing.T) {
	testfn(t, Two, -1, 0, 1, "Two")
}

func print(t *testing.T, p, n int, mean, vari float64) {
	mean /= Small
	vari = (vari - Small*(mean*mean)) / (Small - 1)
	t.Log("hi:", p, "lo:", n, "med:", Large-p-n)
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
